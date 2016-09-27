package filestore

import (
	"math/rand"
	"time"
	"os"
	"path/filepath"
)

var hexCharset = []byte("0123456789abcdef")
var randomGenerator = rand.New(rand.NewSource(time.Now().UnixNano()))

func randHex(length int) string {
	b := make([]byte, length)
	for i := 0; i < length; i++ {
		b[i] = hexCharset[randomGenerator.Int() % 16]
	}
	return string(b[:length])
}

func fileExists(fpath string) bool {
	_, err := os.Stat(fpath)
	return err == nil
}

type FileStore struct {
	path string
}

func New(path string) FileStore {
	return FileStore{path}
}

func (fs *FileStore) fullPath(fname string) (fpath string) {
	fpath = filepath.Join(fs.path, fname)
	return
}

func (fs *FileStore) getAvailableFilename() (fname string, fpath string) {
	for {
		fname = randHex(32)
		fpath = fs.fullPath(fname)
		if !fileExists(fpath) {
			return
		}
	}
}

func (fs *FileStore) NewFile() (fname string, f *os.File) {
	fname, fpath := fs.getAvailableFilename()
	f, err := os.Create(fpath)
	if err != nil {
		panic(err)
	}
	return
}

func (fs *FileStore) GetFile(fname string) (f *os.File, err error) {
	fpath := fs.fullPath(fname)
	f, err = os.Open(fpath)
	return
}