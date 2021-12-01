package term

import (
	"fmt"
	"io"
	"os"
)

const (
	infoTag    = "[ INFO    ] "
	warningTag = "[ WARNING ] "
	errorTag   = "[ ERROR   ] "
	debugTag   = "[ DEBUG   ] "
)

// Terminal is the Cerigo's Go terminal that holds various user's custom
// commands.
//
// There are private variables in the structure so use the `NewTerminal(...)`
// function to create your terminal object.
type Terminal struct {
	commands     map[string]*Command
	terminalType uint
	stderr       io.Writer
	stdout       io.Writer
	root         bool

	row    uint
	column uint
}

// NewTerminal creates the terminal object. It accepts terminalPrefix parameter.
//
// The `terminalPrefix` is to facilitate terminal interpretations. This setting
// is persisted inside the Terminal settings.
//
// The default is `NoTerminal`.
func NewTerminal(terminalPrefix uint) *Terminal {
	return &Terminal{
		commands:     make(map[string]*Command),
		terminalType: terminalPrefix,
		stderr:       os.Stderr,
		stdout:       os.Stdout,
		root:         os.Getuid() <= 0,
	}
}

// Add is to create the given command string into the terminal with its timeout.
//
// If the timeout is `0`, the generated Command object falls back to default
// timeout which is `60` seconds. This is for handling dud commands that holds
// the control forever.
//
// If the given label is found to be associated with an existing Command,
// `Add(...)` will return an error.
func (t *Terminal) Add(label string, command string, timeout uint64) error {
	if label == "" {
		return fmt.Errorf("empty label")
	}

	if _, ok := t.commands[label]; ok {
		return fmt.Errorf("command already exists: %v", label)
	}

	t.commands[label] = &Command{
		command:      command,
		terminalType: t.terminalType,
		Timeout:      timeout,
	}

	return nil
}

// Delete removes an existing command in the terminal regardlessly.
func (t *Terminal) Delete(label string) {
	delete(t.commands, label)
}

// Execute is to execute a given command string synchonously and without
// persisting inside the terminal object.
//
// This is for quick execution without needing to build up your own terminal.
//
// It takes inputs of:
//   1. `command` - which is your command line.
//   2. `timeout` - timeout to kill the command. Set to 0 for default 2 minutes.
//                  the unit is nanoseconds.
//
// It returns:
//   1. `stdout`, `stderr`, `nil`   - successfully ran command with no error.
//   2. `stdout`, `stderr`, `err`   - successfully ran command but with error.
func (t *Terminal) Execute(command string,
	timeout uint64) (stdout []byte, stderr []byte, err error) {
	c := &Command{
		command:      command,
		terminalType: t.terminalType,
	}
	c.Timeout = timeout
	c.Run()
	err = c.ReadError()
	stdout, stderr = c.ReadOutput()

	return stdout, stderr, err
}

// Get seeks the registered command from the terminal and returns it to you.
//
// If the command is not found, it will return `nil`.
func (t *Terminal) Get(label string) *Command {
	return t.commands[label]
}

// IsRoot is to check whether the terminal has root access.
func (t *Terminal) IsRoot() bool {
	return t.root
}

// PrintStatus is to output a given input string to the STDERR output channel.
//
// It uses `fmt.Printf(...)` as its primary printout with suppressed error.
// Hence, to use `PrintStatus(...)`, you just need to state the `statusID`
// first and use then use it like `fmt.Printf(...)`.
//
// PrintStatus accepts `statusID` based on the constant values in this package.
// By default, the status is `NoTagStatus`.
func (t *Terminal) PrintStatus(statusID uint, format string, a ...interface{}) {
	tag := ""

	switch statusID {
	case InfoStatus:
		tag = infoTag
	case WarningStatus:
		tag = warningTag
	case ErrorStatus:
		tag = errorTag
	case DebugStatus:
		tag = debugTag
	}

	if t.stderr == nil {
		t.stderr = os.Stderr
	}

	_, _ = fmt.Fprintf(t.stderr, tag+format, a...)
}

// Printf is to output the given input formatting string to the STDOUT channel.
//
// It uses fmt.Printf(...) as its primary printout with suppressed error. Hence,
// you must ensure your values are correct before passing into `Printf`.
func (t *Terminal) Printf(format string, a ...interface{}) {
	if t.stdout == nil {
		t.stdout = os.Stdout
	}

	_, _ = fmt.Fprintf(t.stdout, format, a...)
}

// Run is to execute a registered command both asynchonously and synchonously.
//
// To different them, state the `mustWait` intention. If `Run(...)` is set to
// `mustWait`, the execution becomes synchonous and will wait for completions
// before proceeding to next line of Go code.
//
// It returns:
//   1. `*Command`, `nil`     - Command object after successful run.
//   2. `*Command`, `err`     - error encountered.
func (t *Terminal) Run(label string, mustWait bool) (c *Command, err error) {
	c = t.commands[label]
	if c == nil {
		return nil, fmt.Errorf("missing command: %v", label)
	}

	c.Start()

	if mustWait {
		c.Wait()
	}

	return c, c.ReadError()
}

// Update is to update the existing Command object with new Command and its
// timeout.
//
// If the `Command` object is missing (e.g. bad label), `Update(...)` returns an
// `error` instead.
func (t *Terminal) Update(label string, command string, timeout uint64) error {
	if label == "" {
		return fmt.Errorf("empty label")
	}

	c := t.commands[label]
	if c == nil {
		return fmt.Errorf("command not found: %v", label)
	}

	c.command = command
	c.Timeout = timeout

	return nil
}
