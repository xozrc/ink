package fileutil

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestExist(t *testing.T) {
	f, err := ioutil.TempFile(os.TempDir(), "fileutil")
	if err != nil {
		t.Fatal(err)
	}
	f.Close()

	if g := Exist(f.Name()); g != true {
		t.Errorf("exist = %v, want true", g)
	}

	os.Remove(f.Name())
	if g := Exist(f.Name()); g != false {
		t.Errorf("exist = %v, want false", g)
	}
}
