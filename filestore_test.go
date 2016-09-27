package filestore

import (
	"testing"
)

func TestNewFile(t *testing.T) {
	test_string := "lol"
	fs := New("/tmp/lol")
	name, f := fs.NewFile()
	f.Write([]byte(test_string))
	f.Close()

	f, err := fs.GetFile(name)
	defer f.Close()
	if err != nil {
		t.FailNow()
	}
	contents := make([]byte, 4)
	l, err := f.Read(contents)

	if l != 3 {
		t.Fail()
	}

	if err != nil {
		t.FailNow()
	}

	if string(contents[:l]) != test_string {
		t.Fail()
	}
}
