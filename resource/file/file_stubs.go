// +build windows

package file

import (
	"fmt"
	"runtime"

	"github.com/milosgajdos83/servpeek/resource"
)

func isOwnedBy(f resource.Filer, username string) error {
	return fmt.Errorf("IsOwnedBy not implemented on %s/%s", runtime.GOOS, runtime.GOARCH)
}

func isGrupedInto(f resource.Filer, groupname string) error {
	return fmt.Errorf("IsGrupedInto not implemented on %s/%s", runtime.GOOS, runtime.GOARCH)
}
