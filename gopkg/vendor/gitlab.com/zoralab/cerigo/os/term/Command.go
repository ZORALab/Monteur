package term

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
	"sync"
	"time"
)

const (
	// NoTerminal is the terminalType for No terminal interpreter
	NoTerminal = uint(0)

	// BASHTerminal is the terminalType for BASH terminal interpreter
	BASHTerminal = uint(1)

	// SHTerminal is the terminalType for SHELL terminal interpreter
	SHTerminal = uint(2)

	// DOSTerminal is the terminalType for MSDOS terminal interpreter
	DOSTerminal = uint(3)
)

const (
	defaultTimeout = 60 * time.Second
)

// Command structure is an object that holds a single os.exec command operation.
//
// It has a `Timeout` for setting up a bail-out sequences in case of a given dud
// command. The `Timeout` is strictly nanoseconds. If 0 is provided, it carries
// out the default timeout which is `60` seconds.
//
// You cannot alter the command directly in this structure. For that, you should
// use the `Terminal.Update(...)` function.
type Command struct {
	Timeout uint64

	cmd          *exec.Cmd
	command      string
	terminalType uint
	timeout      time.Duration
	err          error
	stdout       []byte
	stderr       []byte
	sync         *sync.RWMutex
}

func (c *Command) reset() {
	switch c.terminalType {
	case BASHTerminal:
		c.cmd = exec.Command("bash", "-c", c.command) // #nosec
	case SHTerminal:
		c.cmd = exec.Command("sh", "-c", c.command) // #nosec
	case DOSTerminal:
		fallthrough
	default:
		arguments := strings.Split(c.command, " ")
		c.cmd = exec.Command(arguments[0], arguments[1:]...) // #nosec
	}

	switch {
	case c.Timeout > 0:
		c.timeout = time.Duration(c.Timeout)
	default:
		c.timeout = defaultTimeout
	}

	c.cmd.Stdout = &bytes.Buffer{}
	c.cmd.Stderr = &bytes.Buffer{}
	c.sync = &sync.RWMutex{}
}

func (c *Command) setError(err error) {
	c.sync.Lock()
	defer c.sync.Unlock()

	if c.err == nil && err != nil {
		c.err = err
	}
}

func (c *Command) setOutput() {
	c.sync.Lock()
	defer c.sync.Unlock()
	c.stdout = c.cmd.Stdout.(*bytes.Buffer).Bytes()
	c.stderr = c.cmd.Stderr.(*bytes.Buffer).Bytes()
}

// ReadError returns the error value found from the command execution.
//
// This function relies on its mutex synchonization to avoid data races.
func (c *Command) ReadError() (err error) {
	c.sync.Lock()
	defer c.sync.Unlock()

	return c.err
}

// ReadOutput returns `STDOUT` and `STDERR` outputs returned from the execution.
//
// This function relies on its mutex synchonization to avoid data races.
func (c *Command) ReadOutput() (stdout []byte, stderr []byte) {
	c.sync.Lock()
	defer c.sync.Unlock()

	return c.stdout, c.stderr
}

// Run executes the command synchonously.
//
// Under the hood, `Run()` runs `Start()` and `Wait()` in sequence.
func (c *Command) Run() {
	c.Start()
	c.Wait()
}

// Wait runs the waiting and post data process after Start().
func (c *Command) Wait() {
	done := make(chan error, 1)

	go func() {
		done <- c.cmd.Wait()
	}()

	select {
	case <-time.After(c.timeout):
		err := c.cmd.Process.Kill()
		err = fmt.Errorf("timeout with error: %v", err)
		c.setError(err)
	case err := <-done:
		c.setError(err)
		c.setOutput()
	}
}

// Start executes the command asynchonously.
func (c *Command) Start() {
	c.reset()
	err := c.cmd.Start()
	c.setError(err)
}
