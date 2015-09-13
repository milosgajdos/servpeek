// package command wraps exec.CombinedOutput to provide a convenient
// way to execute a command and return stdout and stderr
package command

import (
	"os/exec"
	"strings"
)

// Build builds up a command from a list of strings and returns
// its string interpretation
func Build(args ...string) string {
	return strings.Join(args, " ")
}

// Run executes a command with arbitrary number of arguments passed in
// as the function's parameter and returns combined stdour and stderr
func Run(command string, args ...string) (output string, err error) {
	cmd := exec.Command(command, args...)
	out, err := cmd.CombinedOutput()
	output = string(out)
	if err != nil {
		return output, err
	}
	return output, nil
}
