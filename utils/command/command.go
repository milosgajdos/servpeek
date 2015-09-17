// package command wraps exec.CombinedOutput to provide a convenient
// way to execute a command and return stdout and stderr
package command

import (
	"bufio"
	"io"
	"os/exec"
	"sync"
)

type Command struct {
	*exec.Cmd
}

func NewCommand(cmd string, args ...string) *Command {
	return &Command{exec.Command(cmd, args...)}
}

// Run executes a command with arbitrary number of arguments passed in
// as the function's parameter and returns combined stdout and stderr
func RunCombined(command string, args ...string) (string, error) {
	cmd := NewCommand(command, args...)
	out, err := cmd.CombinedOutput()
	output := string(out)
	if err != nil {
		return output, err
	}
	return output, nil
}

func (c *Command) Run() *Result {
	cmdStdout, err := c.StdoutPipe()
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
		if err := c.Start(); err != nil {
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

		if err := c.Wait(); err != nil {
			res.mu.Lock()
			defer res.mu.Unlock()
			res.err = err
		}
	}()

	return res
}

type Result struct {
	closed bool
	reader io.ReadCloser
	lines  chan string
	mu     sync.Mutex
	line   string
	err    error
}

func (r *Result) Next() (ok bool) {
	r.line, ok = <-r.lines
	return !r.closed && r.err == nil && ok
}

func (r *Result) Text() string {
	return r.line
}

func (r *Result) Err() error {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.err
}

func (r *Result) Close() error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.closed {
		return nil
	}
	r.closed = true
	return r.reader.Close()
}
