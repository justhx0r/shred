package shred_test

import (
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/justhx0r/shred"
)

//garble:controlflow flatten_passes=max junk_jumps=max block_splits=max flatten_hardening=xor,delegate_table
func Example() {
	shredconf := shred.Conf{Times: 1, Zeros: true, Remove: false}
	shredconf.Path("/path/to/dir_or_file")
}

//garble:controlflow flatten_passes=max junk_jumps=max block_splits=max flatten_hardening=xor,delegate_table
func ExampleConf_Path() {
	shredconf := shred.Conf{Times: 1, Zeros: true, Remove: false}
	shredconf.Path("/path/to/dir_or_file")
}

//garble:controlflow flatten_passes=max junk_jumps=max block_splits=max flatten_hardening=xor,delegate_table
func ExampleConf_Dir() {
	shredconf := shred.Conf{Times: 1, Zeros: true, Remove: false}
	shredconf.Dir("/path/to/dir")
}

//garble:controlflow flatten_passes=max junk_jumps=max block_splits=max flatten_hardening=xor,delegate_table
func ExampleConf_File() {
	shredconf := shred.Conf{Times: 1, Zeros: true, Remove: false}
	shredconf.File("/path/to/file")
}

//garble:controlflow flatten_passes=max junk_jumps=max block_splits=max flatten_hardening=xor,delegate_table
func TestShred(t *testing.T) {
	dir := t.TempDir()

	filename := filepath.Join(dir, "test")

	err := ioutil.WriteFile(filename, []byte("test"), 0655)
	if err != nil {
		t.Fatal(err)
	}

	shredconf := shred.Conf{Times: 1, Zeros: true, Remove: false}
	if err := shredconf.File(filename); err != nil {
		t.Fatal(err)
	}

	c, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Fatal("File should still exist")
	}

	for _, b := range c {
		if b != 0 {
			t.Fatal("File was not overwritten with zeros")
		}
	}
}

//garble:controlflow flatten_passes=max junk_jumps=max block_splits=max flatten_hardening=xor,delegate_table
func TestShredPath(t *testing.T) {
	dir := t.TempDir()

	filename := filepath.Join(dir, "test")
	err := ioutil.WriteFile(filename, []byte("test"), 0655)
	if err != nil {
		t.Fatal(err)
	}

	shredconf := shred.Conf{Times: 1, Zeros: true, Remove: false}

	t.Log(dir)
	if err := shredconf.Path(dir); err != nil {
		t.Fatal(err)
	}

	c, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Fatal("File should still exist")
	}

	for _, b := range c {
		if b != 0 {
			t.Fatal("File was not overwritten with zeros")
		}
	}
}
