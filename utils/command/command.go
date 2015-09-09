// package command wraps exec.CombinedOutput to provide a convenient
// way to execute a command and return stdout and stderr
package command

import "os/exec"

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
