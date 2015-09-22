// package command wraps exec.Command to provide a convenient
// way to execute a command and receive its stdout output line by line
package command

import (
	"bufio"
	"io"
	"os/exec"
	"sync"
)

// Run executes a command with arbitrary number of arguments passed in
// as the function's parameter and returns combined stdout and stderr
func RunCombined(command string, args ...string) (string, error) {
	cmd := exec.Command(command, args...)
	out, err := cmd.CombinedOutput()
	output := string(out)
	if err != nil {
		return output, err
	}
	return output, nil
}

// Command provides an interface to command and its arguments
type Command struct {
	Cmd  string
	Args []string
}

// NewCommand returns pointer to Command
func NewCommand(cmd string, args ...string) *Command {
	return &Command{
		Cmd:  cmd,
		Args: args,
	}
}

// Run executes a command and returns Result object that can be used
// to collect the results of the run command
func (c *Command) Run() *Result {
	cmd := exec.Command(c.Cmd, c.Args...)
	cmdStdout, err := cmd.StdoutPipe()
	res := &Result{
		lines:  make(chan string, 1),
		reader: cmdStdout,
		err:    err,
	}
	if err != nil {
		close(res.lines)
		return res
	}

	go func() {
		defer close(res.lines)
		if err := cmd.Start(); err != nil {
			res.mu.Lock()
			defer res.mu.Unlock()
			res.err = err
			return
		}

		scanner := bufio.NewScanner(res.reader)
		for scanner.Scan() {
			res.lines <- scanner.Text()
		}

		if err := scanner.Err(); err != nil {
			res.mu.Lock()
			defer res.mu.Unlock()
			res.err = err
			return
		}

		if err := cmd.Wait(); err != nil {
			res.mu.Lock()
			defer res.mu.Unlock()
			res.err = err
		}
	}()

	return res
}

// Result contains all the information about the result of the executed
// Command. It provides a simple API to interact with the result
type Result struct {
	closed bool
	reader io.ReadCloser
	lines  chan string
	mu     sync.Mutex
	line   string
	err    error
}

// Next returns next line from result output or false if there is none
func (r *Result) Next() (ok bool) {
	r.line, ok = <-r.lines
	return !r.closed && r.err == nil && ok
}

// Returns a single line of streamed command output
func (r *Result) Text() string {
	return r.line
}

// Err returns the last encountered error of the executed command
func (r *Result) Err() error {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.err
}

// Close closes Result's standard output reader
func (r *Result) Close() error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.closed {
		return nil
	}
	r.closed = true
	return r.reader.Close()
}
