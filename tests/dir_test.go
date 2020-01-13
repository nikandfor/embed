package tests

import (
	"io"
	"io/ioutil"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDirOpen(t *testing.T) {
	fs1 := FS(false)

	f, err := fs1.Open("tests/dir_test.go")
	assert.NoError(t, err)

	s, err := f.Stat()
	assert.NoError(t, err)

	assert.Equal(t, "dir_test.go", s.Name())

	err = f.Close()
	assert.NoError(t, err)

	//
	f, err = fs1.Open("go.mod")
	assert.NoError(t, err)

	s, err = f.Stat()
	assert.NoError(t, err)

	assert.Equal(t, "go.mod", s.Name())
	assert.Equal(t, int64(324), s.Size())
	assert.True(t, s.Mode()&0644 == 0644)
	assert.True(t, time.Now().After(s.ModTime()))
	assert.Equal(t, false, s.IsDir())
	assert.Nil(t, s.Sys())

	data, err := ioutil.ReadAll(f)
	assert.NoError(t, err)

	fcont := `module github.com/nikandfor/embed

go 1.13

require (
	github.com/golang/snappy v0.0.1
	github.com/nikandfor/cli v0.0.0-20191209221237-ad6c97f5afac
	github.com/nikandfor/json v0.0.0-20191226181838-acc5fc24730f // indirect
	github.com/nikandfor/tlog v0.3.0
	github.com/pkg/errors v0.8.1
	github.com/stretchr/testify v1.3.0
)
`
	assert.Equal(t, fcont, string(data))

	off, err := f.Seek(0, io.SeekCurrent)
	assert.NoError(t, err)
	assert.Equal(t, int64(len(fcont)), off)

	n, err := f.(io.ReaderAt).ReadAt(data[:40], 10)
	assert.NoError(t, err)
	assert.Equal(t, 40, n)

	err = f.Close()
	assert.NoError(t, err)

	_, err = f.Seek(0, io.SeekStart)
	assert.Equal(t, ErrClosed, err)

	_, err = f.Read(nil)
	assert.Equal(t, ErrClosed, err)

	_, err = f.(io.ReaderAt).ReadAt(nil, 0)
	assert.Equal(t, ErrClosed, err)
}

func TestDirReaddir(t *testing.T) {
	fs := SecondFS(true)

	f, err := fs.Open("")
	assert.NoError(t, err)

	s, err := f.Stat()
	assert.NoError(t, err)

	assert.True(t, s.IsDir())
	assert.Equal(t, ".", s.Name())

	ff, err := f.Readdir(4)
	assert.NoError(t, err)
	if !assert.Len(t, ff, 4) {
		return
	}

	var list []string
	for _, f := range ff {
		list = append(list, f.Name())
	}

	ff, err = f.Readdir(2)
	assert.NoError(t, err)
	if !assert.Len(t, ff, 2) {
		return
	}

	for _, f := range ff {
		list = append(list, f.Name())
	}

	ff, err = f.Readdir(2)
	assert.Equal(t, io.EOF, err)
	if !assert.Len(t, ff, 1) {
		return
	}

	for _, f := range ff {
		list = append(list, f.Name())
	}

	assert.ElementsMatch(t, list, []string{"LICENSE", "README.md", "empty", "go.mod", "go.sum", "main.go", "tests"})

	//
	f2, err := fs.Open("/")
	assert.NoError(t, err)

	s2, err := f2.Stat()
	assert.NoError(t, err)

	assert.True(t, s == s2)

	//
	_, err = fs.Open("not-exists")
	assert.True(t, os.IsNotExist(err))
}

func TestDirReaddirNo(t *testing.T) {
	fs := SecondFS(false)

	f, err := fs.Open("")
	assert.NoError(t, err)

	s, err := f.Stat()
	assert.NoError(t, err)

	assert.True(t, s.IsDir())
	assert.Equal(t, ".", s.Name())

	ff, err := f.Readdir(2)
	assert.NoError(t, err)
	assert.Nil(t, ff)
}
