package embed

import (
	"encoding/base64"
	"io"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/golang/snappy"
	"github.com/stretchr/testify/assert"
)

func EncodeFile(d []byte) (cont string) {
	if len(d) == 0 {
		return
	}

	z := snappy.Encode(nil, d)
	cont = base64.StdEncoding.EncodeToString(z)

	return
}

func TestFile(t *testing.T) {
	var f File

	SetFile(&f, false, EncodeFile(nil))
	assert.Len(t, f.Data(), 0)

	b := make([]byte, 100)
	n, err := f.Reader().Read(b)
	assert.Error(t, err)

	SetFile(&f, false, EncodeFile([]byte("content")))
	assert.Equal(t, []byte("content"), f.Data())

	n, err = f.Reader().Read(b)
	assert.NoError(t, err)
	assert.Equal(t, []byte("content"), b[:n])

	SetFile(&f, true, "content")
	assert.Equal(t, []byte("content"), f.Data())

	n, err = f.Reader().Read(b)
	assert.NoError(t, err)
	assert.Equal(t, []byte("content"), b[:n])
}

func TestBadFile(t *testing.T) {
	var f File

	SetFile(&f, false, EncodeFile([]byte("content"))+"123")

	assert.Panics(t, func() { f.Data() })

	SetFile(&f, false, base64.StdEncoding.EncodeToString([]byte("content")))

	assert.Panics(t, func() { f.Data() })
}

func TestFSFile(t *testing.T) {
	tm := time.Now()
	tm2 := time.Now()

	var fs FS

	var _ http.FileSystem = fs

	AddFile(&fs, "file_path", 7, tm, 0600, false, nil, EncodeFile([]byte("content")))
	AddFile(&fs, "file", 6, tm2, 0641, false, nil, EncodeFile([]byte("valval")))

	//
	f, err := fs.Open("/file_path")
	assert.NoError(t, err)

	b := make([]byte, 100)
	n, err := f.Read(b)
	assert.NoError(t, err)
	assert.Equal(t, []byte("content"), b[:n])

	s, err := f.Stat()
	assert.NoError(t, err)
	assert.Equal(t, "file_path", s.Name())
	assert.Equal(t, int64(7), s.Size())
	assert.Equal(t, tm, s.ModTime())
	assert.Equal(t, os.FileMode(0600), s.Mode())
	assert.Equal(t, false, s.IsDir())
	assert.Equal(t, nil, s.Sys())

	//
	fs.NoCache = true

	f, err = fs.Open("/file")
	assert.NoError(t, err)

	p, err := f.Seek(1, io.SeekStart)
	assert.NoError(t, err)
	assert.Equal(t, int64(1), p)

	n, err = f.Read(b[:3])
	assert.NoError(t, err)
	assert.Equal(t, []byte("alv"), b[:n])

	n, err = f.(io.ReaderAt).ReadAt(b[:3], 2)
	assert.NoError(t, err)
	assert.Equal(t, []byte("lva"), b[:n])

	s, err = f.Stat()
	assert.NoError(t, err)
	assert.Equal(t, "file", s.Name())
	assert.Equal(t, int64(6), s.Size())
	assert.Equal(t, tm2, s.ModTime())
	assert.Equal(t, os.FileMode(0641), s.Mode())
	assert.Equal(t, false, s.IsDir())

	//
	err = f.Close()
	assert.NoError(t, err)

	_, err = f.Seek(1, io.SeekStart)
	assert.Equal(t, err, ErrClosed)

	_, err = f.Read(b)
	assert.Equal(t, err, ErrClosed)

	_, err = f.(io.ReaderAt).ReadAt(b, 2)
	assert.Equal(t, err, ErrClosed)

	//
	f, err = fs.Open("nonexisted")
	assert.True(t, os.IsNotExist(err))
}

func TestFSFileNoc(t *testing.T) {
	tm := time.Now()
	tm2 := time.Now()

	var fs FS
	NotCompressed(&fs, true)

	var _ http.FileSystem = fs

	AddFile(&fs, "file_path", 7, tm, 0600, false, nil, "content")
	AddFile(&fs, "file", 6, tm2, 0641, false, nil, "valval")

	//
	f, err := fs.Open("/file_path")
	assert.NoError(t, err)

	b := make([]byte, 100)
	n, err := f.Read(b)
	assert.NoError(t, err)
	assert.Equal(t, []byte("content"), b[:n])

	s, err := f.Stat()
	assert.NoError(t, err)
	assert.Equal(t, "file_path", s.Name())
	assert.Equal(t, int64(7), s.Size())
	assert.Equal(t, tm, s.ModTime())
	assert.Equal(t, os.FileMode(0600), s.Mode())
	assert.Equal(t, false, s.IsDir())
	assert.Equal(t, nil, s.Sys())

	//
	fs.NoCache = true

	f, err = fs.Open("/file")
	assert.NoError(t, err)

	p, err := f.Seek(1, io.SeekStart)
	assert.NoError(t, err)
	assert.Equal(t, int64(1), p)

	n, err = f.Read(b[:3])
	assert.NoError(t, err)
	assert.Equal(t, []byte("alv"), b[:n])

	n, err = f.(io.ReaderAt).ReadAt(b[:3], 2)
	assert.NoError(t, err)
	assert.Equal(t, []byte("lva"), b[:n])

	s, err = f.Stat()
	assert.NoError(t, err)
	assert.Equal(t, "file", s.Name())
	assert.Equal(t, int64(6), s.Size())
	assert.Equal(t, tm2, s.ModTime())
	assert.Equal(t, os.FileMode(0641), s.Mode())
	assert.Equal(t, false, s.IsDir())

	//
	err = f.Close()
	assert.NoError(t, err)

	_, err = f.Seek(1, io.SeekStart)
	assert.Equal(t, err, ErrClosed)

	_, err = f.Read(b)
	assert.Equal(t, err, ErrClosed)

	_, err = f.(io.ReaderAt).ReadAt(b, 2)
	assert.Equal(t, err, ErrClosed)

	//
	f, err = fs.Open("nonexisted")
	assert.True(t, os.IsNotExist(err))
}

func TestFSDir(t *testing.T) {
	tm := time.Now()
	tm2 := time.Now()
	tm3 := time.Now()

	var fs FS
	fs.Index = true

	AddFile(&fs, "dir", 4096, tm, 020000000775, true, []string{"a", "b"}, "")
	AddFile(&fs, "dir/a", 7, tm2, 0600, false, nil, EncodeFile([]byte("content")))
	AddFile(&fs, "dir/b", 6, tm3, 0641, false, nil, EncodeFile([]byte("valval")))

	d, err := fs.Open("/dir")
	assert.NoError(t, err)

	s, err := d.Stat()
	assert.NoError(t, err)

	assert.Equal(t, "dir", s.Name())
	assert.Equal(t, true, s.IsDir())

	//
	ff1, err := d.Readdir(1)
	assert.NoError(t, err)

	if assert.Len(t, ff1, 1) {
		assert.Equal(t, "a", ff1[0].Name())
	}

	ff2, err := d.Readdir(3)
	assert.Equal(t, io.EOF, err)

	if assert.Len(t, ff2, 1) {
		assert.Equal(t, "b", ff2[0].Name())
	}

	//
	fs.Index = false

	d, err = fs.Open("/dir")
	assert.NoError(t, err)

	ff1, err = d.Readdir(1)
	assert.NoError(t, err)
	assert.Len(t, ff1, 0)

	_, err = fs.Open("")
	assert.True(t, os.IsNotExist(err))
}

func TestFSBadFile(t *testing.T) {
	tm := time.Now()

	var fs FS

	AddFile(&fs, "file1", 7, tm, 0600, false, nil, EncodeFile([]byte("content"))+"123")
	AddFile(&fs, "file2", 7, tm, 0600, false, nil, base64.StdEncoding.EncodeToString([]byte("content")))

	_, err := fs.Open("/file1")
	assert.Error(t, err)

	_, err = fs.Open("/file2")
	assert.Error(t, err)
}

func TestCoverMustTime(t *testing.T) {
	tm := MustTime(time.Parse(TimeFormat, TimeFormat))
	assert.Equal(t, TimeFormat, tm.Format(TimeFormat))

	assert.Panics(t, func() { MustTime(tm, io.EOF) })
}
