package args

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"gitlab.com/zoralab/cerigo/os/term"
	"gitlab.com/zoralab/cerigo/strings/strhelper"
)

const (
	separatorWidth     = uint(5)
	argumentLabel      = "arguments"
	usageLabel         = "usages"
	emptyLineSeparator = "\n\n\n"
	lineCharacter      = "─"
)

// Manager is the core arguments structure that holds all the flags and
// processes them accordingly.
//
// This structure contains important private data elements so use NewManager()
// function to create one instead of the conventional structure{} method.
// Without doing so can cause unforseenable fatal error or panics.
//
// ELEMENTS SPECIFICATIONS
//
// 1. Name
//
// Name is the program name where the field is used to print the overall help
// output in PrintHelp(...) function.
//
// This is compulsory to fill in.
//
//
// 2. Description
//
// Description is a short explanation (limit it to 1 short paragraph)
// explaining the program. It is used to print the overall help output in
// PrintHelp(...) function. Unicode is supported.
//
// This is compulsory to fill in.
//
//
// 3. Version
//
// Version is the version number of the program usually in the form of string
// or number types used to print the overall help output in PrintHelp(...)
// function.
//
// This is optional to fill in.
//
//
// 4. Examples
//
// This is the "how-to" examples for using the program. It is a string slice
// offering the program "usage", allowing users to quickly learn to use.
// Manager does automatic numbering so you do not need to do it manually.
// If the list is empty, it is treated as no examples. Example:
//     CASE 1 - If set to:
//                   []string{}     OR    nil
//     The argument help output would be:
//               MY PROGRAM (0.2.0)
//               ──────────────────
//               维护喝色黄色而后恶化而这是话语维护喝色黄色而后恶化而这是话
//               而这是话 语维护喝色 黄色而后 恶化而 这是话语 维护喝
//               黄色而后恶化而这 是话语维护喝色黄 色而
//
//               ARGUMENTS
//               ─────────
//               ...
//
//     CASE 2 - If set to:
//                   []string{"",
//                            "$ ./program --help",
//                            "$ ./program --list",
//                   }
//     The argument help output would be:
//               MY PROGRAM (0.2.0)
//               ──────────────────
//               维护喝色黄色而后恶化而这是话语维护喝色黄色而后恶化而这是话
//               而这是话 语维护喝色 黄色而后 恶化而 这是话语 维护喝
//               黄色而后恶化而这 是话语维护喝色黄 色而
//
//
//               USAGES
//               ──────
//               1)  $ ./program --help
//               2)  $ ./program --list
//
//
//               ARGUMENTS
//               ─────────
//               ...
//
// This is optional to fill in.
//
//
// 5. ShowFlagExamples
//
// This is the master switch to show all flag respective examples. By default,
// it is false (off).
//
// This is optional to fill in.
//
//
// 6. ArgumentLabel
//
// This is the string that alters the ARGUMENT title in the help printout. It
// is offered explicitly to facilitate i18n. By default, it uses English word
// "ARGUMENTS".
//
// This is optional to fill in.
//
//
// 7. UsageLabel
//
// This is the string that alters the USAGE title in the help printout. It is
// offered explicitly to facilitate i18n. By default, it uses English word
// "USAGES".
type Manager struct {
	// program metadata
	Name             string
	Description      string
	Version          interface{}
	Examples         []string
	ShowFlagExamples bool

	// label texts for i18n
	ArgumentLabel string
	UsageLabel    string

	// private elements
	flags     map[string]*Flag
	args      []string
	maxColumn uint
}

// NewManager creates and initializes the Manager structure ready for use.
func NewManager() *Manager {
	return &Manager{
		Name:  os.Args[0],
		flags: map[string]*Flag{},
		args:  os.Args,
	}
}

// Add takes a Flag object and registers into the manager.
//
// It returns:
//   1. nil      - successful register
//   2. err      - error occurred (e.g. missing flag)
//   3. panic    - critical error requires upfront attention (e.g. conflicted
//                 labels)
func (m *Manager) Add(f *Flag) error {
	if f == nil {
		return fmt.Errorf(ErrorMissingFlag)
	}

	f.cleanupLabel()

	for _, label := range f.Label {
		x := m.flags[label]
		if x != nil {
			err := fmt.Errorf("%s --> %s",
				ErrorConflictedFlag,
				label)
			panic(err)
		}

		m.flags[label] = f
	}

	return nil
}

// Parse starts the argument parsing process.
func (m *Manager) Parse() {
	m.parseArgs()
	m.convertFlags()
}

func (m *Manager) parseArgs() {
	if m.args == nil || len(m.args) == 0 {
		m.args = os.Args
	}

	var oldLabel string

	for i, arg := range m.args {
		if i == 0 {
			continue
		}

		label, value, hasTail := m.analyzeArg(arg, oldLabel)
		f := m.flags[label]

		if f == nil {
			return
		}

		f.setValue(value)

		oldLabel = ""
		if hasTail {
			oldLabel = label
		}
	}
}

func (m *Manager) analyzeArg(arg string,
	oldLabel string) (label string, value string, hasTail bool) {
	if strings.Contains(arg, "=") {
		// type: (path=usr/bin, -path=/usr/bin, -p=/usr/bin)
		//       (path="core=ice", path='core=ice', path=core=ice)
		ret := strings.SplitN(arg, "=", 2)
		return ret[0], ret[1], false
	}

	if arg[:1] == "-" {
		// type: (-h, --help)
		return arg, arg, true
	}

	if oldLabel != "" {
		// type: (/usr/bin [tailing value])
		return oldLabel, arg, false
	}

	// type: standalone (install, help, version)
	return arg, arg, false
}

func (m *Manager) convertFlags() {
	for _, flag := range m.flags {
		flag.convert()
	}
}

// PrintHelp creates the help menu based on the given flags and manager's data.
// It will output the final formatted string with help contents ready for
// printout.
//
// It returns:
//   1. string     - the help statement
func (m *Manager) PrintHelp() (out string) {
	m.preparePrintHelp()
	out += m.printHelpHeaders()
	out += m.printFlagsHelp()

	return out
}

func (m *Manager) preparePrintHelp() {
	if m.ArgumentLabel == "" {
		m.ArgumentLabel = argumentLabel
	}

	if m.UsageLabel == "" {
		m.UsageLabel = usageLabel
	}
}

func (m *Manager) printHelpHeaders() (out string) {
	var s *strhelper.Styler
	var sl []string
	var i int
	var l, title string
	var first bool

	// 1. print title
	title = m.Name

	// 2. print version if available
	if m.Version != nil {
		title += fmt.Sprintf(" (%v)", m.Version)
	}

	out += m.printHeader(title)

	// 3. print description
	if m.Description != "" {
		_, max := m.termSize()
		s = &strhelper.Styler{}
		sl = s.WordWrap(m.Description, max)
		first = true

		for _, l = range sl {
			if !first {
				out += newLine
			}

			first = false
			out += l
		}

		out += emptyLineSeparator
	}

	// 4. print example usages
	if len(m.Examples) != 0 {
		out += m.printHeader(m.UsageLabel)
		first = true
		i = 1

		for _, l = range m.Examples {
			if l == "" {
				continue
			}

			if !first {
				out += newLine
			}

			first = false
			out += fmt.Sprintf("%d)  %s", i, l)
			i++
		}

		out += emptyLineSeparator
	}

	// 5. Argument Titles
	out += m.printHeader(m.ArgumentLabel)

	return out
}

func (m *Manager) printHeader(keyword string) (out string) {
	out = strings.ToUpper(keyword) + newLine
	out += strings.Repeat(lineCharacter, len(out)-1) + newLine

	return out
}

func (m *Manager) printFlagsHelp() (out string) {
	// 1. get all flags
	list := map[*Flag]bool{}
	for _, v := range m.flags {
		list[v] = true
	}

	// 2. sort flags alphabetically
	sortedList := map[string]*Flag{}
	labelList := []string{}
	dashedLabelList := []string{}
	flagWidth := 0

	for k := range list {
		s := k.labelStatement()
		if s == "" {
			continue
		}

		sLength := len(s)

		if s[:1] == "-" {
			dashedLabelList = append(dashedLabelList, s)
		} else {
			labelList = append(labelList, s)
		}

		sortedList[s] = k

		if sLength > flagWidth {
			flagWidth = sLength
		}
	}

	sort.Strings(labelList)
	sort.Strings(dashedLabelList)

	// 3. generate each flags help statements
	argWidth, contentWidth := m.calculateWidth(flagWidth)
	first := true

	for _, s := range labelList {
		f := sortedList[s]
		m.processFlagHelp(f, &out, argWidth, contentWidth, &first)
	}

	for _, s := range dashedLabelList {
		f := sortedList[s]
		m.processFlagHelp(f, &out, argWidth, contentWidth, &first)
	}

	return out
}

func (m *Manager) processFlagHelp(f *Flag,
	out *string,
	argWidth uint,
	contentWidth uint,
	isFirst *bool) {
	if f.DisableHelp {
		return
	}

	if !*isFirst && m.ShowFlagExamples {
		*out += newLine
	}

	*isFirst = false
	*out += f.help(argWidth, contentWidth, m.ShowFlagExamples)
}

func (m *Manager) calculateWidth(flagWidth int) (arg uint, content uint) {
	_, max := m.termSize()
	arg = uint(flagWidth) + separatorWidth
	content = max - arg

	return arg, content
}

//nolint:unparam
func (m *Manager) termSize() (row uint, column uint) {
	t := term.NewTerminal(term.NoTerminal)
	row, column = t.Size()

	if m.maxColumn != 0 {
		column = m.maxColumn
	}

	return row, column
}
