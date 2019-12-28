[![Documentation](https://godoc.org/github.com/nikandfor/embed?status.svg)](http://godoc.org/github.com/nikandfor/embed)
[![Build Status](https://travis-ci.com/nikandfor/embed.svg?branch=master)](https://travis-ci.com/nikandfor/embed)
[![codecov](https://codecov.io/gh/nikandfor/embed/branch/master/graph/badge.svg)](https://codecov.io/gh/nikandfor/embed)
[![GolangCI](https://golangci.com/badges/github.com/nikandfor/embed.svg)](https://golangci.com/r/github.com/nikandfor/embed)
[![Go Report Card](https://goreportcard.com/badge/github.com/nikandfor/embed)](https://goreportcard.com/report/github.com/nikandfor/embed)

# embed
embed files or folders into go executable and get `http.FileSystem`

# usage

file
```bash
embed --pkg mypkg --src my.tmpl --dst tmpl.go --name MyTemplate
```
```go
data := MyTemplateReadAll()
t, err := template.New("").Parse(string(data))
```

folder
```bash
embed --pkg static --src build/dist --prefix build/dist --dst static/static.go --skip-hidden
```
```go
allowDirIndex := true
fs := static.FS(allowDirIndex)
http.Handle("/", http.FileServer(fs))
```
