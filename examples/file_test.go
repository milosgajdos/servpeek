package examples

import (
	"regexp"
	"testing"

	"github.com/milosgajdos83/servpeek/file"
)

func TestFile(t *testing.T) {
	f := file.NewFile("/etc/hosts")

	if err := file.IsRegular(f); err != nil {
		t.Errorf("Error: %s", err)
	}

	owner := "root"
	group := "wheel"
	md5 := "YOUR_MD5SUM_HERE"

	if err := file.IsOwnedBy(f, owner); err != nil {
		t.Errorf("Error: %s", err)
	}

	if err := file.IsGrupedInto(f, group); err != nil {
		t.Errorf("Error: %s", err)
	}

	if err := file.MD5Equal(f, md5); err != nil {
		t.Errorf("Error: %s", err)
	}

	content := regexp.MustCompile(`localhost`)
	if err := file.Contains(f, content); err != nil {
		t.Errorf("Error: %s", err)
	}
}
