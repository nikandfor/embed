package embed

import (
	"bytes"
	"encoding/base64"
	"errors"
	"io"
	"net/http"
	"os"
	"path"
	"sync"
	"time"

	"github.com/golang/snappy"
)

//go:generate go run ./cmd/embed --pkg tests --skip-hidden --src . --dst tests/dir.go tests/dir.go tests/dir2.go --var FS
//go:generate go run ./cmd/embed --pkg tests --skip-hidden --src . --dst tests/dir2.go --var fs --noc tests/dir.go tests/dir2.go
//go:generate go run ./cmd/embed --pkg tests --skip-hidden --src go.mod --dst tests/file.go --var file
//go:generate go run ./cmd/embed --pkg tests --skip-hidden --src go.mod --dst tests/noc.go --var noc --noc
//go:generate go run ./cmd/embed --pkg tests --skip-hidden --src empty --dst tests/empty.go --var Empty

type (
	// File embeds file content in executable.
	File struct {
		enc []byte
		noc bool
	}

	file struct {
		path    string
		size    int64
		modTime time.Time
		mode    os.FileMode
		isDir   bool

		files   []string
		content []byte

		once    sync.Once
		decoded []byte
	}

	// FS embeds files content in executable.
	FS struct {
		m          map[string]*file
		nocompress bool
		Index      bool // Allow index over directories
		NoCache    bool // Do not cache decoded files in memory
	}

	fsfile struct {
		r bytes.Reader
		*file
		fs FS
		d  int
	}
)

// TimeFormat used by generator
var TimeFormat = "2006-01-02 15:04:05.999999999 -0700 MST"

var (
	ErrClosed     = errors.New("closed file")
	ErrOutOfRange = errors.New("out of range")
	ErrNoContent  = errors.New("no content")
)

// SetFile used by generator.
func SetFile(f *File, noc bool, enc []byte) {
	f.enc = enc
	f.noc = noc
}

// AddFile used by generator.
func AddFile(fs *FS, fpath string, size int64, mod time.Time, mode os.FileMode, isDir bool, files []string, enc []byte) {
	if fs.m == nil {
		fs.m = make(map[string]*file)
	}
	fs.m[fpath] = &file{
		path:    fpath,
		size:    size,
		modTime: mod,
		mode:    mode,
		isDir:   isDir,
		files:   files,
		content: enc,
	}
}

// NotCompressed used by generator.
func NotCompressed(fs *FS, b bool) { fs.nocompress = b }

// Data decodes and returns file content or nil if empty.
// If file is not encoded (flag was provided to generator), content is not copied.
// It means less allocs but also it means you can't modify resulting slice content.
// You may slice it like data[10:100].
func (f File) Data() []byte {
	if f.noc {
		return f.enc
	}
	if f.enc == nil {
		return nil
	}

	z := make([]byte, base64.StdEncoding.DecodedLen(len(f.enc)))
	n, err := base64.StdEncoding.Decode(z, f.enc)
	if err != nil {
		panic(err)
	}

	z = z[:n]

	r, err := snappy.Decode(nil, z)
	if err != nil {
		panic(err)
	}

	return r
}

// Reader decodes content creates reader of it.
func (f File) Reader() io.Reader {
	d := f.Data()
	return bytes.NewReader(d)
}

// Open opens file from embedded fs.
// If file is not encoded (flag was provided to generator), content is not copied.
// It means less allocs but also it means you can't modify resulting slice content.
// You may slice it like data[10:100].
func (fs FS) Open(p string) (_ http.File, err error) {
	if p != "" && p[0] == '/' {
		p = p[1:]
	}
	if p == "" {
		p = "."
	}
	f, ok := fs.m[p]
	if !ok {
		return nil, os.ErrNotExist
	}

	ff := &fsfile{file: f, fs: fs}

	if !f.isDir {
		var dec []byte
		dec, err = fs.decode(f)
		if err != nil {
			return
		}
		ff.r.Reset(dec)
	}

	return ff, nil
}

func (fs FS) Data(p string) ([]byte, error) {
	if p != "" && p[0] == '/' {
		p = p[1:]
	}
	if p == "" {
		p = "."
	}
	f, ok := fs.m[p]
	if !ok {
		return nil, os.ErrNotExist
	}

	if f.isDir {
		return nil, ErrNoContent
	}

	return fs.decode(f)
}

func (fs FS) decode(f *file) (d []byte, err error) {
	if fs.nocompress {
		return f.content, nil
	}
	if fs.NoCache {
		return fs.decodeFile(f)
	}

	f.once.Do(func() {
		f.decoded, err = fs.decodeFile(f)
	})
	if err != nil {
		return
	}

	return f.decoded, nil
}

func (fs FS) decodeFile(f *file) (d []byte, err error) {
	if f.content == nil {
		return
	}

	z := make([]byte, base64.StdEncoding.DecodedLen(len(f.content)))
	n, err := base64.StdEncoding.Decode(z, f.content)
	if err != nil {
		return
	}

	z = z[:n]

	d, err = snappy.Decode(nil, z)
	if err != nil {
		return
	}

	return
}

func (f *fsfile) Close() error { f.d = -1; return nil }

func (f *fsfile) Seek(off int64, whence int) (int64, error) {
	if f.d == -1 {
		return 0, ErrClosed
	}
	return f.r.Seek(off, whence)
}

func (f *fsfile) Read(p []byte) (n int, err error) {
	if f.d == -1 {
		return 0, ErrClosed
	}
	return f.r.Read(p)
}

func (f *fsfile) ReadAt(p []byte, off int64) (n int, err error) {
	if f.d == -1 {
		return 0, ErrClosed
	}
	return f.r.ReadAt(p, off)
}

func (f *fsfile) Stat() (os.FileInfo, error) {
	return f.file, nil
}

func (f *fsfile) Readdir(n int) ([]os.FileInfo, error) {
	if !f.fs.Index {
		return nil, nil
	}

	var res []os.FileInfo
	for f.d < len(f.files) && len(res) < n {
		res = append(res, f.fs.m[path.Join(f.path, f.files[f.d])])
		f.d++
	}

	if f.d == len(f.files) {
		return res, io.EOF
	}

	return res, nil
}

func (f *file) Name() string       { return path.Base(f.path) }
func (f *file) Size() int64        { return f.size }
func (f *file) Mode() os.FileMode  { return f.mode }
func (f *file) ModTime() time.Time { return f.modTime }
func (f *file) IsDir() bool        { return f.isDir }
func (f *file) Sys() interface{}   { return nil }

// MustTime used by generator.
func MustTime(t time.Time, err error) time.Time {
	if err != nil {
		panic(err)
	}
	return t
}
