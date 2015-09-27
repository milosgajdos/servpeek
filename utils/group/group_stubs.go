// +build !cgo,!windows,!plan9 android

package group

import (
	"fmt"
	"runtime"
)

func init() {
	implemented = false
}

func lookupGroup(groupname string) (*Group, error) {
	return nil, fmt.Errorf("user: LookupGroup not implemented on %s/%s", runtime.GOOS, runtime.GOARCH)
}

func lookupGroupId(int) (*Group, error) {
	return nil, fmt.Errorf("user: LookupGroupId not implemented on %s/%s", runtime.GOOS, runtime.GOARCH)
}
