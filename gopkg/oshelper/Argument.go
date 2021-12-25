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
	"strconv"
	"strings"

	"gitlab.com/zoralab/cerigo/strings/strhelper"
)

const (
	base10        = 10
	bit64         = 64
	exampleTextEN = "EXAMPLES"
)

// Argument is the data structure for an user's control flag.
//
// It is safe to create using the conventional `&struct{}` method.
type Argument struct {
	// Name is the element strictly for internal referencing use only.
	//
	// This is for selecting or identifying the Argument without needing
	// to iterate over the Label slice.
	//
	// This field is **MANDATORY**.
	Name string

	// ValueLabel is the changing the help's value label during printout.
	//
	// This is mainly facilitaed for supporting i18n purposes. The default
	// is `VALUE`.
	//
	// Example:
	//   1. If set to "number", the help printout would be:
	//        "-h, --help, help [ NUMBER ]"
	//   2. If set to empty (""), the help printout would be:
	//        "-h, --help, help"
	//   3. If set to "名字", the help printout would be:
	//        "-h, --help, help [ 名字 ]"
	//
	// This field is **MANDATORY**.
	ValueLabel string

	// Value is the variable pointer for ArgParser to insert results into.
	//
	// The data type for this field **MUST** be a supported data types
	// (see "SUPPORTED DATA TYPES" section). Otherwise, the value shall be
	// discarded.
	//
	// This field **ONLY** accepts pointer of your data type's variable.
	//
	// This field is **MANDATORY**.
	Value interface{}

	// Help is this Argument's main descriptions for manual printout.
	//
	// This should be single paragraph holding the abstract about this
	// Argument. It shall be printed to the right.
	//
	// This package uses a word-wrap algorithm to style the output. Please
	// avoid multi-paragraphs (are you mad?) if possible.
	//
	// This field is **MANDATORY**.
	Help string

	// ExampleLabel is the string altering "EXAMPLES" text in HelpExamples.
	//
	// It is offered explicitly for facilitatng i18n. By default, it uses
	// English word "Examples".
	//
	// This field is optional.
	ExampleLabel string

	value string

	// Label is a list of the Argument's "flags" (e.g. -h, --help, help)
	//
	// An Argument can offers multiple "flags" for a single purpose. The
	// minimum is a single flag.
	//
	// This field is **MANDATORY**.
	Label []string

	// HelpExamples are the examples for using this Argument.
	//
	// This is mainly for printing help output assistances. You may supply
	// as many as you want (sensibly). Argument does automatic numbering
	// so no manual efforts are needed.
	//
	// If the list is empty (`[]string{}`) or `nil`, it is treated as no
	// example available.
	//
	// Example cases:
	//
	// CASE 1 - If set to: `[]string{}` or `nil`:
	//   -a, --append [ VALUE ]     to append a round number into the
	//                              inputs. This can be a very long value
	//                              and can be a very long description.
	//
	// CASE 2 - If set to: `[]string{"$ ./program -a 123",
	// "$ ./program --append 123", "$ ./program -append '123'", }`:
	//   -a, --append [ VALUE ]     to append a round number into the
	//                              inputs. This can be a very long value
	//                              and can be a very long description.
	//                              EXAMPLES
	//                                1) $ ./program -a 123
	//                                2) $ ./program --append 123
	//                                3) $ ./program -append '123'
	//
	// This field is optional.
	HelpExamples []string

	// DisableHelp is to instruct ArgumentParser to discard help printout.
	//
	// This is useful for controlling printout of this Argument. The default
	// is off (`false`).
	//
	// This field is optional.
	DisableHelp bool
}

func (me *Argument) setValue(value string) {
	var i int

	// trim whitespace
	s := strings.TrimSpace(value)

	// remove leading quote
	i = 0
	if s[i] == '"' || s[i] == '\'' {
		s = s[1:]
	}

	// remove tailing quote
	i = len(s) - 1
	if s[i] == '"' || s[i] == '\'' {
		s = s[:i]
	}

	// save value
	me.value = s
}

func (me *Argument) convert() {
	if me.value == "" {
		return
	}

	switch p := me.Value.(type) {
	case *uint:
		value, _ := strconv.ParseUint(me.value, base10, bit64)
		*p = uint(value)
	case *uint8:
		value, _ := strconv.ParseUint(me.value, base10, bit64)
		*p = uint8(value)
	case *uint16:
		value, _ := strconv.ParseUint(me.value, base10, bit64)
		*p = uint16(value)
	case *uint32:
		value, _ := strconv.ParseUint(me.value, base10, bit64)
		*p = uint32(value)
	case *uint64:
		*p, _ = strconv.ParseUint(me.value, base10, bit64)
	case *int:
		value, _ := strconv.ParseInt(me.value, base10, bit64)
		*p = int(value)
	case *int8:
		value, _ := strconv.ParseInt(me.value, base10, bit64)
		*p = int8(value)
	case *int16:
		value, _ := strconv.ParseInt(me.value, base10, bit64)
		*p = int16(value)
	case *int32:
		value, _ := strconv.ParseInt(me.value, base10, bit64)
		*p = int32(value)
	case *int64:
		*p, _ = strconv.ParseInt(me.value, base10, bit64)
	case *float32:
		value, _ := strconv.ParseFloat(me.value, bit64)
		*p = float32(value)
	case *float64:
		*p, _ = strconv.ParseFloat(me.value, bit64)
	case *string:
		*p = me.value
	case *bool:
		*p, _ = strconv.ParseBool(me.value)
	case *[]byte:
		*p = []byte(me.value)
	case **[]byte:
		value := []byte(me.value)
		*p = &value
	}
}

func (me *Argument) help(width uint, descWidth uint, withEx bool) (out string) {
	var l, s, title, format string
	var hl []string
	var i int

	if len(me.Label) == 0 || me.DisableHelp {
		return out
	}

	// 1. prepare
	format = fmt.Sprintf("%s%dv%s\n", "%-", width, "%v")
	l = me.labelStatement()

	// 2. format help statement
	hl = me.wordWrap(me.Help, descWidth)
	for i, s = range hl {
		if i == 0 {
			out = fmt.Sprintf(format, l, s)
			continue
		}

		out += fmt.Sprintf(format, "", s)
	}

	// 3. format example section
	title = me.ExampleLabel
	if title == "" {
		title = exampleTextEN
	}

	if len(me.HelpExamples) != 0 && withEx {
		hl = []string{title}
		i = 1

		for _, s = range me.HelpExamples {
			if s == "" {
				continue
			}

			hl = append(hl, fmt.Sprintf("  %d) %s", i, s))
			i++
		}

		hl = me.indentWrap(hl, width)
		for i, s = range hl {
			if i != 0 {
				out += newLine
			}

			out += s
		}

		out += newLine
	}

	return out
}

func (me *Argument) cleanupLabel() {
	sl := []string{}

	for _, s := range me.Label {
		if s != "" {
			sl = append(sl, s)
		}
	}

	me.Label = sl
}

func (me *Argument) labelStatement() (out string) {
	me.cleanupLabel()
	out = strings.Join(me.Label, ", ")

	if me.ValueLabel != "" {
		out += fmt.Sprintf(" [ %s ]", me.ValueLabel)
	}

	return out
}

func (me *Argument) wordWrap(in string, flagWidth uint) []string {
	s := &strhelper.Styler{}
	return s.WordWrap(in, flagWidth)
}

func (me *Argument) indentWrap(in []string, flagWidth uint) (out []string) {
	s := &strhelper.Styler{}
	return s.Indent(in, strhelper.CharSpaceIndent, flagWidth)
}
