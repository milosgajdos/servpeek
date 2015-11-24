// +build windows

package utils

import (
	"fmt"
	"runtime"
)

func roleToID(role string, name string) (uint64, error) {
	return 0, fmt.Errorf("RoleToID not implemented on %s/%s", runtime.GOOS, runtime.GOARCH)
}
