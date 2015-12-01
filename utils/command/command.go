// Package command allows to run external commands
package command

import (
	"bufio"
	"io"
	"os/exec"
	"sync"
)

// Commander provides an interface to execute commands
type Commander interface {
	// Run executes a command and returns Outer
	// which provides an interface to interact with
	// output of the executed command
	Run() Outer
	// RunCombined runs command and returns combined output
	// of both stdout and stderr
	RunCombined() (string, error)
	// AppendArgs allows to append arbitrary number of
	// extra command arguments
	AppendArgs(...string)
}

// Command is an external command with arguments
// Command implements Commander interface
type Command struct {
	// cli command
	Cmd string
	// cli command arguments
	Args []string
}

// NewCommand returns pointer to Command
func NewCommand(cmd string, args ...string) *Command {
	return &Command{
		Cmd:  cmd,
		Args: args,
	}
}

// Run executes Command and returns Outer that can be used
// to collect and analyse the output of the executed command
func (c *Command) Run() Outer {
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

// AppendArgs allows to append arbitrary number of arguments to the
// underlying command that will be executed by Run
func (c *Command) AppendArgs(args ...string) {
	c.Args = append(c.Args, args...)
}

// RunCombined executes Command and returns combined Stdout and Stderr as a string.
// The difference betweent RunCombined and Run is that Run returns a Stdout stream
// ie. stream of line strings. RunCombined returns combined Stdout/Stderr output in one string
func (c *Command) RunCombined() (string, error) {
	cmd := exec.Command(c.Cmd, c.Args...)
	out, err := cmd.CombinedOutput()
	return string(out), err
}

// Outer interface defines an interface to interact with Command output
type Outer interface {
	// Next advances through Command output iteration per line
	Next() bool
	// Text returns a single line of Command output
	Text() string
	// Err returns the last encountered error
	Err() error
	// Close closes standard output of Command output
	Close() error
}

// Out implements Outer interface to provide a simple API
// to interact with executed Command output
type Out struct {
	closed bool
	reader io.ReadCloser
	lines  chan string
	mu     sync.Mutex
	line   string
	err    error
}

// Next forwards the Command Stdout iteration to next line.
// It returns true if there is any output to be processed.
// It returns false all all of Command output has been processed or
// an error occurred during Command execution
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

// Close closes Stdout of the executed command
func (o *Out) Close() error {
	o.mu.Lock()
	defer o.mu.Unlock()
	if o.closed {
		return nil
	}
	o.closed = true
	return o.reader.Close()
}
