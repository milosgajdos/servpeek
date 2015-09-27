package examples

import (
	"testing"

	"github.com/milosgajdos83/servpeek/resource/file"
)

func Test_File(t *testing.T) {
	f, err := file.NewFile("/etc/hosts")
	if err != nil {
		t.Fatalf("Could not read file %s: %s", f, err)
	}

	if !f.IsRegular() {
		t.Errorf("File %s is not a regular file", f)
	}

	owner := "root"
	group := "wheel"
	md5 := "SOME_MD5SUM_YOU"
	if ok, err := f.IsOwnedBy(owner); err != nil {
		t.Errorf("Error: %s", err)
		if !ok {
			t.Errorf("File %s not owned by %s", f, owner)
		}
	}

	if ok, err := f.IsGrupedInto(group); err != nil {
		t.Errorf("Error: %s", err)
		if !ok {
			t.Errorf("File %s not grouped into %s", f, group)
		}
	}

	if ok, err := f.Md5(md5); err != nil {
		t.Errorf("Error: %s", err)
		if !ok {
			t.Errorf("Incorrect MD5 sum of file %s: %s", f, md5)
		}
	}
}
