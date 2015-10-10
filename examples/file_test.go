package examples

import (
	"testing"

	"github.com/milosgajdos83/servpeek/resource"
	"github.com/milosgajdos83/servpeek/resource/file"
)

func Test_File(t *testing.T) {
	f := &resource.File{
		Path: "/etc/hosts",
	}

	if ok, err := file.IsRegular(f); err != nil {
		t.Errorf("Error: %s", err)
		if !ok {
			t.Errorf("%s is not a regular file", f)
		}
	}

	owner := "root"
	group := "wheel"
	md5 := "YOUR_MD5SUM_HERE"
	if ok, err := file.IsOwnedBy(f, owner); err != nil {
		t.Errorf("Error: %s", err)
		if !ok {
			t.Errorf("%s not owned by %s", f, owner)
		}
	}

	if ok, err := file.IsGrupedInto(f, group); err != nil {
		t.Errorf("Error: %s", err)
		if !ok {
			t.Errorf("%s not grouped into %s", f, group)
		}
	}

	if ok, err := file.Md5(f, md5); err != nil {
		t.Errorf("Error: %s", err)
		if !ok {
			t.Errorf("%s incorrect MD5 sum of file: %s", f, md5)
		}
	}
}
