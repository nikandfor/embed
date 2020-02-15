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

func EncodeFile(d []byte) (cont []byte) {
	if len(d) == 0 {
		return
	}

	z := snappy.Encode(nil, d)
	e := base64.StdEncoding.EncodeToString(z)

	return []byte(e)
}

func TestFile(t *testing.T) {
	var f File

	SetFile(&f, false, EncodeFile(nil))
	assert.Len(t, f.Data(), 0)

	b := make([]byte, 100)
	_, err := f.Reader().Read(b)
	assert.Error(t, err)

	SetFile(&f, false, EncodeFile([]byte("content")))
	assert.Equal(t, []byte("content"), f.Data())

	n, err := f.Reader().Read(b)
	assert.NoError(t, err)
	assert.Equal(t, []byte("content"), b[:n])

	SetFile(&f, true, []byte("content"))
	assert.Equal(t, []byte("content"), f.Data())

	n, err = f.Reader().Read(b)
	assert.NoError(t, err)
	assert.Equal(t, []byte("content"), b[:n])
}

func TestBadFile(t *testing.T) {
	var f File

	SetFile(&f, false, append(EncodeFile([]byte("content")), "123"...))

	assert.Panics(t, func() { f.Data() })

	SetFile(&f, false, []byte(base64.StdEncoding.EncodeToString([]byte("content"))))

	assert.Panics(t, func() { f.Data() })
}

func TestFSFile(t *testing.T) {
	tm := time.Now()
	tm2 := time.Now()

	var fs FS

	var _ http.FileSystem = fs

	AddFile(&fs, "file_path", 7, tm, os.FileMode(0600), false, nil, EncodeFile([]byte("content")))
	AddFile(&fs, "file", 6, tm2, os.FileMode(0641), false, nil, EncodeFile([]byte("valval")))
	AddFile(&fs, "empty", 6, tm2, os.FileMode(0641), false, nil, nil)

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
	f, err = fs.Open("/empty")
	assert.NoError(t, err)

	//
	err = f.Close()
	assert.NoError(t, err)

	_, err = f.Seek(1, io.SeekStart)
	assert.Equal(t, ErrClosed, err)

	_, err = f.Read(b)
	assert.Equal(t, ErrClosed, err)

	_, err = f.(io.ReaderAt).ReadAt(b, 2)
	assert.Equal(t, ErrClosed, err)

	//
	_, err = fs.Open("nonexisted")
	assert.True(t, os.IsNotExist(err))
}

func TestFSFileNoc(t *testing.T) {
	tm := time.Now()
	tm2 := time.Now()

	var fs FS
	NotCompressed(&fs, true)

	var _ http.FileSystem = fs

	AddFile(&fs, "file_path", 7, tm, os.FileMode(0600), false, nil, []byte("content"))
	AddFile(&fs, "file", 6, tm2, os.FileMode(0641), false, nil, []byte("valval"))

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
	_, err = fs.Open("nonexisted")
	assert.True(t, os.IsNotExist(err))
}

func TestFSDir(t *testing.T) {
	tm := time.Now()
	tm2 := time.Now()
	tm3 := time.Now()

	var fs FS
	fs.Index = true

	AddFile(&fs, "dir", 4096, tm, os.FileMode(020000000775), true, []string{"a", "b"}, nil)
	AddFile(&fs, "dir/a", 7, tm2, os.FileMode(0600), false, nil, EncodeFile([]byte("content")))
	AddFile(&fs, "dir/b", 6, tm3, os.FileMode(0641), false, nil, EncodeFile([]byte("valval")))

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

	AddFile(&fs, "file1", 7, tm, os.FileMode(0600), false, nil, append(EncodeFile([]byte("content")), "123"...))
	AddFile(&fs, "file2", 7, tm, os.FileMode(0600), false, nil, []byte(base64.StdEncoding.EncodeToString([]byte("content"))))

	_, err := fs.Open("/file1")
	assert.Error(t, err)

	_, err = fs.Open("/file2")
	assert.Error(t, err)
}

func TestFSData(t *testing.T) {
	tm := time.Now()

	var fs FS

	AddFile(&fs, ".", 0, tm, 0, true, []string{"file1", "file2"}, nil)

	_, err := fs.Data("/")
	assert.Equal(t, ErrNoContent, err)

	AddFile(&fs, "file1", 7, tm, os.FileMode(0600), false, nil, EncodeFile([]byte("content")))

	data, err := fs.Data("/file1")
	assert.NoError(t, err)
	assert.Equal(t, []byte("content"), data)

	NotCompressed(&fs, true)
	AddFile(&fs, "file2", 7, tm, os.FileMode(0600), false, nil, []byte("content2"))

	data, err = fs.Data("/file2")
	assert.NoError(t, err)
	assert.Equal(t, []byte("content2"), data)

	_, err = fs.Data("/nofile")
	assert.True(t, os.IsNotExist(err))
}

func TestCoverMustTime(t *testing.T) {
	tm := MustTime(time.Parse(TimeFormat, TimeFormat))
	assert.Equal(t, TimeFormat, tm.Format(TimeFormat))

	assert.Panics(t, func() { MustTime(tm, io.EOF) })
}

func BenchmarkFile(b *testing.B) {
	b.ReportAllocs()

	var f File
	SetFile(&f, false, EncodeFile([]byte("file content")))

	var d []byte

	for i := 0; i < b.N; i++ {
		d = f.Data()
	}

	_ = d
}

func BenchmarkNocFile(b *testing.B) {
	b.ReportAllocs()

	var f File
	SetFile(&f, true, []byte("file content"))

	var d []byte

	for i := 0; i < b.N; i++ {
		d = f.Data()
	}

	_ = d
}

func BenchmarkFS(b *testing.B) {
	b.ReportAllocs()

	var fs FS
	AddFile(&fs, "name", 0, time.Now(), 0, false, nil, EncodeFile([]byte("file content")))

	d := make([]byte, 100)

	for i := 0; i < b.N; i++ {
		f, err := fs.Open("name")
		if err != nil {
			panic(err)
		}
		_, err = f.Read(d)
		if err != nil {
			panic(err)
		}
	}

	_ = d
}

func BenchmarkFSNoCache(b *testing.B) {
	b.ReportAllocs()

	var fs FS
	fs.NoCache = true

	AddFile(&fs, "name", 0, time.Now(), 0, false, nil, EncodeFile([]byte("file content")))

	d := make([]byte, 100)

	for i := 0; i < b.N; i++ {
		f, err := fs.Open("name")
		if err != nil {
			panic(err)
		}
		_, err = f.Read(d)
		if err != nil {
			panic(err)
		}
	}

	_ = d
}

func BenchmarkFSNoCompress(b *testing.B) {
	b.ReportAllocs()

	var fs FS
	NotCompressed(&fs, true)

	AddFile(&fs, "name", 0, time.Now(), 0, false, nil, EncodeFile([]byte("file content")))

	d := make([]byte, 100)

	for i := 0; i < b.N; i++ {
		f, err := fs.Open("name")
		if err != nil {
			panic(err)
		}
		_, err = f.Read(d)
		if err != nil {
			panic(err)
		}
	}

	_ = d
}

func BenchmarkFSDate(b *testing.B) {
	b.ReportAllocs()

	var fs FS
	AddFile(&fs, "name", 0, time.Now(), 0, false, nil, EncodeFile([]byte("file content")))

	var err error
	var d []byte

	for i := 0; i < b.N; i++ {
		d, err = fs.Data("name")
		if err != nil {
			panic(err)
		}
	}

	_ = d
}
