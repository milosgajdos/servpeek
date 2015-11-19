// Package command allows to run external commands
package command

import (
	"bufio"
	"io"
	"os/exec"
	"sync"
)

// Command is an external command with arguments
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

// Run executes command and returns Out object that can be used
// to collect the results of the command run
func (c *Command) Run() *Out {
	cmd := exec.Command(c.Cmd, c.Args...)
	cmdStdout, err := cmd.StdoutPipe()
	res := &Out{
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

// RunCombined executes Command and returns combined stdout and stderr as a string
// The difference betweent RunCombined and Run is that Run returns a stream of lines
// ie. stream of strings. RunCombined returns full Stdout/Stderr output in one string
func (c *Command) RunCombined() (string, error) {
	cmd := exec.Command(c.Cmd, c.Args...)
	out, err := cmd.CombinedOutput()
	return string(out), err
}

// Out contains all the information about the result of executed
// Command. It provides a simple API to interact with the result
type Out struct {
	closed bool
	reader io.ReadCloser
	lines  chan string
	mu     sync.Mutex
	line   string
	err    error
}

// Next returns next line from Command output or false if there is none
func (o *Out) Next() (ok bool) {
	o.line, ok = <-o.lines
	return !o.closed && o.err == nil && ok
}

// Text returns a single line of command output
func (o *Out) Text() string {
	return o.line
}

// Err returns the last encountered error of the executed command
func (o *Out) Err() error {
	o.mu.Lock()
	defer o.mu.Unlock()
	return o.err
}

// Close closes standard output reader
func (o *Out) Close() error {
	o.mu.Lock()
	defer o.mu.Unlock()
	if o.closed {
		return nil
	}
	o.closed = true
	return o.reader.Close()
}
