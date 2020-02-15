[![Documentation](https://godoc.org/github.com/nikandfor/embed?status.svg)](http://godoc.org/github.com/nikandfor/embed)
[![Build Status](https://travis-ci.com/nikandfor/embed.svg?branch=master)](https://travis-ci.com/nikandfor/embed)
[![codecov](https://codecov.io/gh/nikandfor/embed/branch/master/graph/badge.svg)](https://codecov.io/gh/nikandfor/embed)
[![GolangCI](https://golangci.com/badges/github.com/nikandfor/embed.svg)](https://golangci.com/r/github.com/nikandfor/embed)
[![Go Report Card](https://goreportcard.com/badge/github.com/nikandfor/embed)](https://goreportcard.com/report/github.com/nikandfor/embed)

# embed

embed files or folders into go executable and get `http.FileSystem`. Multiple files of FSs could be embedded in the same or different packages.

# usage

## File

Create file with a variables of type `embed.File`.
```go
// mypkg/vars.go
package mypkg

var (
    someTemplate embed.File
    Config embed.File
)
```

Generate files content.
```bash
embed --pkg mypkg --src my.tmpl --dst mypkg/tmpl.go --var someTemplate
embed --pkg mypkg --src cfg.json --dst mypkg/config.go --var Config
```

Use.
```go
data := someTemplate.Data()
t, err := template.New("").Parse(string(data))

dec, err := json.NewDecoder(Config.Reader()) // file is decoded and bytes.NewReader(data) is returned.
```

## Folder

Create file with a variable of type `embed.FS`.
```go
// static/var.go
package static

var FS = embed.FS{Index: true} // allow reading directories
```

Generate content.
```bash
embed --pkg static --src front/dist --dst static/embedded.go --var FS --skip-hidden front/dist/not_needed.txt front/dist/any_number_of_excludes.html
```

Use.
```go
http.Handle("/", static.FS)
```

## .gitignore

It's recommended to add generated file names to `.gitignore`. It's intended to separate variables and generated content. It allows to have code compiling even without generated files.
