package main

import (
	"regexp"

	"github.com/milosgajdos83/servpeek/file"
)

// CheckFiles checks various file properties.
// It returns error if any of the checked properties fail to be satisfied.
func CheckFiles() error {
	f := file.NewFile("/etc/hosts")
	if err := file.IsRegular(f); err != nil {
		return err
	}

	owner := "root"
	group := "wheel"
	md5 := "YOUR_MD5SUM_HERE"

	if err := file.IsOwnedBy(f, owner); err != nil {
		return err
	}

	if err := file.IsGrupedInto(f, group); err != nil {
		return err
	}

	if err := file.MD5Equal(f, md5); err != nil {
		return err
	}

	content := regexp.MustCompile(`localhost`)
	if err := file.Contains(f, content); err != nil {
		return err
	}

	return nil
}
