// Package term is a control unit to interface with command line interface.
//
// It uses multiple standard library packages like `os`, `os/exec`, and etc. to
// build itself and extend the original `os/exec` terminal.
//
// ## Differences between `os/exec`
//
// Here are some of the differences between Cerigo `os/term` package with the
// the standard package `os/exec`:
//
// 1. Unified command line into single line rather than split arguments.
//
// 2. Formulate Terminal object in Go codes, allowing addition of individual
//    Commands.
//
// 3. Independent command management.
//
// 4. Extend `os/exec` asynchonous implementation.
//
// 5. Do not support `Stdin`. Use properly formatted command as a replacement
//    before feeding it into the terminal or start the execution. Example:
//    `$ echo y | sudo apt-get install gimp`
package term
