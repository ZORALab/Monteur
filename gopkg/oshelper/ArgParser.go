// Copyright 2021 ZORALab Enterprise (hello@zoralab.com)
// Copyright 2021 "Holloway" Chew, Kean Ho (hollowaykeanho@gmail.com)
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package oshelper

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"gitlab.com/zoralab/cerigo/strings/strhelper"
)

const (
	separatorWidth     = uint(5)
	argumentLabel      = "arguments"
	usageLabel         = "usages"
	emptyLineSeparator = "\n\n\n"
	lineCharacter      = "─"
)

// ArgParser is the comamnd line arguments parser data structure.
//
// This structure requires internal initialization so please use
// `NewArgParser()` function to created one over the risky `struct{}` method.
type ArgParser struct {
	flags map[string]*Argument

	// Version is the version number of the program.
	//
	// It was meant for printing out the program version during a help
	// printout.
	//
	// This field only accepts the following data types:
	//   1. `string`
	//   2. `int`
	Version interface{}

	// Examples are the overall examples for using the program.
	//
	// It is meant for users to quickly learn how to use via examples
	// observations.
	//
	// ArgParser does automatic numbering so you do not need to do it
	// manually. If the list is empty, it is treated as no examples.
	// Example:
	//
	// CASE 1 - If set to: `[]string{}` or `nil`:
	//   MY PROGRAM (0.2.0)
	//   ──────────────────
	//   维护喝色黄色而后恶化而这是话语维护喝色黄色而后恶化而这是话而这是话
	//   语维护喝色 黄色而后 恶化而 这是话语 维护喝黄色而后恶化而这 是话语维
	//   护喝色黄 色而
	//
	//   ARGUMENTS
	//   ─────────
	//   ...
	//
	// CASE 2 - If set to: `[]string{"", "$ ./program --help",
	// "$ ./program --list", }`:
	//   MY PROGRAM (0.2.0)
	//   ──────────────────
	//   维护喝色黄色而后恶化而这是话语维护喝色黄色而后恶化而这是话而这是话
	//   语维护喝色 黄色而后 恶化而 这是话语 维护喝黄色而后恶化而这 是话语维
	//   护喝色黄 色而
	//
	//
	//   USAGES
	//   ──────
	//   1)  $ ./program --help
	//   2)  $ ./program --list
	//
	//
	//   ARGUMENTS
	//   ─────────
	//   ...
	Examples []string

	// Name is the name of the program.
	//
	// This is used for help printout.
	//
	// This field is **MANDATORY**.
	Name string

	// Description is the short description of the program.
	//
	// This is use for help printout. Please limit the content to 1 short
	// paragraph. Unicode is supported.
	//
	// This field is **MANDATORY**.
	Description string

	// ArgumentLabel is to alter "ARGUMENT" title in help printout.
	//
	// It is offered for facilitating i18n. The default is "ARGUMENT".
	ArgumentLabel string

	// UsageLabel is to alter "USAGE" title in help printout.
	//
	// It is offered for facilitating i18n. The default is "USAGES".
	UsageLabel string

	args []string

	// ShowExamples decides to print Arguments' examples in help printout.
	//
	// It is the master switch to turn all on/off. The default is off
	// (`false`).
	ShowExamples bool
}

// NewArgParser creates and initializes the ArgParser ready for use.
func NewArgParser() *ArgParser {
	return &ArgParser{
		Name:  os.Args[0],
		flags: map[string]*Argument{},
		args:  os.Args,
	}
}

// Add takes a Argument and registers into the ArgParser.
//
// It returns:
//   1. nil      - registered successfully.
//   2. err      - error occurred (e.g. missing important data)
//   3. panic    - critical error requires upfront attention (e.g. conflicted
//                 labels)
func (me *ArgParser) Add(f *Argument) error {
	if f == nil {
		return fmt.Errorf("missing input argument")
	}

	f.cleanupLabel()

	for _, label := range f.Label {
		x := me.flags[label]
		if x != nil {
			err := fmt.Errorf("%s --> %s",
				"found conflicting argument",
				label)
			panic(err)
		}

		me.flags[label] = f
	}

	return nil
}

// Parse starts the argument parsing process.
func (me *ArgParser) Parse() {
	me.parseArgs()
	me.convertArguments()
}

func (me *ArgParser) parseArgs() {
	if me.args == nil || len(me.args) == 0 {
		me.args = os.Args
	}

	var oldLabel string

	for i, arg := range me.args {
		if i == 0 {
			continue
		}

		label, value, hasTail := me.analyzeArg(arg, oldLabel)
		f := me.flags[label]

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

func (me *ArgParser) analyzeArg(arg string,
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

func (me *ArgParser) convertArguments() {
	for _, flag := range me.flags {
		flag.convert()
	}
}

// PrintHelp generates the help statement for the given Arguments.
//
// Its output string is readily available for printout.
//
// It returns:
//   1. string     - the help statement
func (me *ArgParser) PrintHelp() (out string) {
	me.preparePrintHelp()
	out += me.printHelpHeaders()
	out += me.printArgumentsHelp()

	return out
}

func (me *ArgParser) preparePrintHelp() {
	if me.ArgumentLabel == "" {
		me.ArgumentLabel = argumentLabel
	}

	if me.UsageLabel == "" {
		me.UsageLabel = usageLabel
	}
}

func (me *ArgParser) printHelpHeaders() (out string) {
	var s *strhelper.Styler
	var sl []string
	var i int
	var l, title string
	var first bool

	// 1. print title
	title = me.Name

	// 2. print version if available
	if me.Version != nil {
		title += fmt.Sprintf(" (%v)", me.Version)
	}

	out += me.printHeader(title)

	// 3. print description
	if me.Description != "" {
		_, max := me.termSize()
		s = &strhelper.Styler{}
		sl = s.WordWrap(me.Description, max)
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
	if len(me.Examples) != 0 {
		out += me.printHeader(me.UsageLabel)
		first = true
		i = 1

		for _, l = range me.Examples {
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
	out += me.printHeader(me.ArgumentLabel)

	return out
}

func (me *ArgParser) printHeader(keyword string) (out string) {
	out = strings.ToUpper(keyword) + newLine
	out += strings.Repeat(lineCharacter, len(out)-1) + newLine

	return out
}

func (me *ArgParser) printArgumentsHelp() (out string) {
	// 1. get all flags
	list := map[*Argument]bool{}
	for _, v := range me.flags {
		list[v] = true
	}

	// 2. sort flags alphabetically
	sortedList := map[string]*Argument{}
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
	argWidth, contentWidth := me.calculateWidth(flagWidth)
	first := true

	for _, s := range labelList {
		f := sortedList[s]
		me.processArgumentHelp(f, &out, argWidth, contentWidth, &first)
	}

	for _, s := range dashedLabelList {
		f := sortedList[s]
		me.processArgumentHelp(f, &out, argWidth, contentWidth, &first)
	}

	return out
}

func (me *ArgParser) processArgumentHelp(f *Argument, out *string,
	argWidth uint, contentWidth uint, isFirst *bool) {
	if f.DisableHelp {
		return
	}

	if !*isFirst && me.ShowExamples {
		*out += newLine
	}

	*isFirst = false
	*out += f.help(argWidth, contentWidth, me.ShowExamples)
}

func (me *ArgParser) calculateWidth(width int) (arg uint, content uint) {
	_, max := me.termSize()
	arg = uint(width) + separatorWidth
	content = max - arg

	return arg, content
}

func (me *ArgParser) termSize() (row uint, column uint) {
	return TermSize()
}
