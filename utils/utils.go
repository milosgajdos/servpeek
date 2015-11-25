// Package utils provides useful utility functions
package utils

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"hash"
	"io"

	"github.com/milosgajdos83/servpeek/utils/command"
)

// BuildCmd builds Command from cmd name and arguments
func BuildCmd(cmd string, args ...string) *command.Command {
	return command.NewCommand(cmd, args...)

}

// RoleToID converts username/groupname to their numeric id representations: uid/gid
// Returns error if the role is not supported, usernamd/groupname have not been found
func RoleToID(role string, name string) (uint64, error) {
	return roleToID(role, name)
}

// HashSum calculates hash sum of requested hash type from data stored in Reader r
// It returns hash sum in Hex encoded string. HashSum uses io.Copy to copy the underlying
// data into hash.Hash Writer and returns error only if io.Copy fails with error
func HashSum(hashType string, r io.Reader) (string, error) {
	var hasher hash.Hash

	switch hashType {
	case "md5":
		hasher = md5.New()
	case "sha256":
		hasher = sha256.New()
	default:
		return "", fmt.Errorf("Unsupported hash type requested: %s", hashType)
	}
	if _, err := io.Copy(hasher, r); err != nil {
		return "", err
	}
	return hex.EncodeToString(hasher.Sum(nil)), nil
}
