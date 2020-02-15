package main

import (
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
	"text/template"
	"time"

	"github.com/golang/snappy"
	"github.com/nikandfor/cli"
	"github.com/nikandfor/embed"
	"github.com/nikandfor/tlog"
	"github.com/pkg/errors"
)

type file struct {
	Path    string
	Size    int64
	ModTime time.Time
	Files   []string
	Content string
	Mode    os.FileMode
	IsDir   bool
}

func main() {
	cli.App = cli.Command{
		Name:   "embed",
		Usage:  "[OPTIONS] {<exclude_path>}",
		Action: embedFS,
		Flags: []*cli.Flag{
			cli.NewFlag("src,s", "static", "folder of file to embed"),
			cli.NewFlag("dst,d", "static/embedded.go", "output file"),
			cli.NewFlag("package,pkg,p", "static", "package name"),
			cli.NewFlag("var,v", "", "variable name to assign value to"),
			cli.NewFlag("skip-hidden,H", false, "skip hidden files in src"),
			cli.NewFlag("no-compress,noc,C", false, "do not compress and encode files"),
			cli.HelpFlag,
		},
	}

	cli.RunAndExit(os.Args)
}

func embedFS(c *cli.Command) error {
	stat, err := os.Stat(c.String("src"))
	if err != nil {
		return err
	}
	if !stat.IsDir() {
		return embedFile(c)
	}

	var files []*file
	dirs := map[string]*file{}

	err = filepath.Walk(c.String("src"), func(p string, info os.FileInfo, err error) error {
		p = path.Clean(p)

		if p == path.Clean(c.String("dst")) {
			tlog.Printf("SKIP dst:    %v", p)
			return nil
		}
		for _, a := range c.Args {
			if p == filepath.Clean(a) {
				tlog.Printf("SKIP arg:    %v", p)
				return nil
			}
		}

		realPath := p
		if pp := c.String("src"); pp != "" && pp != "." {
			pp = path.Clean(pp)
			if strings.HasPrefix(p, pp) {
				p = p[len(pp):]
			} else {
				return fmt.Errorf("path %v does not have prefix %v", p, pp)
			}
		}
		for strings.HasPrefix(p, "../") {
			p = p[3:]
		}
		if p == "" || p == ".." {
			p = "."
		}
		if p[0] == '/' {
			p = p[1:]
		}

		if c.Bool("skip-hidden") {
			n := path.Base(p)
			if n != "." && n[0] == '.' {
				tlog.Printf("SKIP hidden: %v", realPath)
				if info.IsDir() {
					return filepath.SkipDir
				}
				return nil
			}
		}

		typ := "file"
		if info.IsDir() {
			typ = "dir"
		}
		tlog.Printf("add   %-5v  %6v kb  %v", typ, info.Size()/1024, p)

		f := &file{
			Path:    p,
			Size:    info.Size(),
			Mode:    info.Mode(),
			ModTime: info.ModTime(),
			IsDir:   info.IsDir(),
		}

		files = append(files, f)

		if info.IsDir() {
			dirs[p] = f
		} else {
			data, err := ioutil.ReadFile(realPath)
			if err != nil {
				return errors.Wrap(err, "read src file")
			}

			if c.Bool("noc") {
				f.Content = strconv.Quote(string(data))
			} else {
				if len(data) != 0 {
					z := snappy.Encode(nil, data)
					a := base64.StdEncoding.EncodeToString(z)

					f.Content = a
				}
			}
		}

		if d, fn := path.Split(p); fn != "." {
			if d == "" {
				d = "."
			}
			if l := len(d) - 1; d[l] == '/' {
				d = d[:l]
			}
			tlog.V("dirs").Printf("dir  %q file %q  dirs: %v", d, fn, dirs)
			dirs[d].Files = append(dirs[d].Files, fn)
		}

		return nil
	})
	if err != nil {
		return errors.Wrap(err, "walk src")
	}

	var w io.Writer = os.Stdout
	if q := c.String("dst"); q != "-" {
		f, err := os.Create(q)
		if err != nil {
			return errors.Wrap(err, "create dst")
		}
		defer f.Close()

		w = f
	}

	t := template.New("")
	t.Funcs(template.FuncMap{
		"sprintf": fmt.Sprintf,
	})
	t = template.Must(t.Parse(`// Generated by {{ .import }}/cmd/embed. DO NOT EDIT.

package {{ .package }}

import (
	"time"

	"{{ .import }}"
)

func init() {
{{- $var := .var }}
{{- $noc := .noc }}
{{- if $noc }}
	embed.NotCompressed(&{{ $var }}, true)
{{- end }}
{{- range .files }}
	embed.AddFile(&{{ $var }},
		` + "`{{ .Path }}`," + `
		{{ .Size }},
		embed.MustTime(time.Parse(embed.TimeFormat, "{{ .ModTime }}")),
		{{ .Mode | sprintf "0%o" }},
		{{ .IsDir }},
		{{ if .IsDir -}}
		[]string{
			{{- range $i, $e := .Files }}{{ if $i }}, {{ end }}` + "`{{ . }}`" + `{{ end -}}
		},
		{{ else -}}
		nil,
		{{ end -}}
		{{ if $noc }}{{ if .IsDir }}nil{{ else }}[]byte({{ .Content }}){{ end }},{{ else -}}
		[]byte(` + "`{{ .Content }}`)," + `{{ end }}
	)
{{- end }}
}
`))

	tlog.Printf("type: %v", reflect.TypeOf(embed.File{}).PkgPath())

	err = t.Execute(w, map[string]interface{}{
		"import":  reflect.TypeOf(embed.File{}).PkgPath(),
		"package": c.String("package"),
		"var":     c.String("var"),
		"files":   files,
		"noc":     c.Bool("noc"),
	})
	if err != nil {
		return errors.Wrap(err, "template")
	}

	return nil
}

func embedFile(c *cli.Command) error {
	data, err := ioutil.ReadFile(c.String("src"))
	if err != nil {
		return errors.Wrap(err, "read src")
	}

	var cont string
	if c.Bool("noc") {
		if len(data) == 0 {
			cont = "nil"
		} else {
			cont = strconv.Quote(string(data))
		}
	} else {
		if len(data) != 0 {
			z := snappy.Encode(nil, data)
			cont = base64.StdEncoding.EncodeToString(z)
		}
	}

	var w io.Writer = os.Stdout
	if q := c.String("dst"); q != "-" {
		f, err := os.Create(q)
		if err != nil {
			return errors.Wrap(err, "create dst")
		}
		defer f.Close()

		w = f
	}

	t := template.New("")
	t.Funcs(template.FuncMap{
		"q": strconv.Quote,
	})
	t = template.Must(t.Parse(`// Generated by {{ .import }}/cmd/embed. DO NOT EDIT.

package {{ .package }}

import "{{ .import }}"

func init() {
{{- if .noc }}
	embed.SetFile(&{{ .var }}, true, []byte({{ .content }}))
{{- else }}
	embed.SetFile(&{{ .var }}, false, []byte(` + "`{{ .content }}`))" + `
{{- end }}
}
`))

	err = t.Execute(w, map[string]interface{}{
		"import":  reflect.TypeOf(embed.File{}).PkgPath(),
		"package": c.String("package"),
		"var":     c.String("var"),
		"noc":     c.Bool("noc"),
		"content": cont,
	})
	if err != nil {
		return errors.Wrap(err, "template")
	}

	return nil
}
