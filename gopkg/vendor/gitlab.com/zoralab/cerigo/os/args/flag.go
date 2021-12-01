package args

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

// Flag is the data structure for each of the argument types. It is safe to
// create using the conventional structure{} method. Flag needs the following
// minimum elements to work properly.
//
//   1. Name
//   2. Label
//   3. Value
//   4. Help
//
// Once done, you can use the Manager's Add(...) function to register the
// Flag object. If you encounter panics, usually that means there is a critical
// issue (e.g. conflicted Label) requires your attention.
//
// ELEMENTS SPECIFICATIONS
//
// 1. Name
//
// Name is the element strictly for internal reference only. This is for
// code developers to reference or to search a particular flag based on
// the name without going through the Label slices.
//
// This is compulsory to fill in.
//
//
// 2. Label
//
// These are the argument "flags" (e.g. -h, --help, help) saved in a string
// slice. A flag is only specific to one purpose but has the ability to offer
// different styles of inputs.
//
// This is compulsory to fill in.
//
//
// 3. ValueLabel
//
// This is for labeling the value part of the value-taking arguments
// (e.g. --add [ VALUE ]). It only affects printing help statement and not
// the parsing algorithms. Example:
//   1. If set to "number", the help printout would be:
//        "-h, --help, help [ NUMBER ]"
//   2. If set to empty (""), the help printout would be:
//        "-h, --help, help"
//   3. If set to "名字", the help printout would be:
//        "-h, --help, help [ 名字 ]"
//
// This is optional to fill in.
//
//
// 4. Value
//
// This is the variable pointer field for Flag to parse the results into. The
// value for this field MUST be a supported data types (see SUPPORTED DATA
// TYPES section) variable pointer. Otherwise, the parsed value will be
// discarded.
//
// Passing a full variable (not a pointer) will not work for this field.
//
// This is compulsory to fill in.
//
//
// 5. Help
//
// The argument help description. This single paragraph string holds the
// short explainations about the flag / arguments (usually printed to the
// right).
//
// This package uses a word-wrap algorithm to style the output. Hence, if you
// pass a multi-paragraphs (are you mad?) string into this field, it will be
// wrapped into a single paragraph instead.
//
// This is compulsory to fill in.
//
//
// 6. Help Examples
//
// This is for printing examples of use for the flag / arguments. It is meant
// for printing help statement. It is a string slice so you can supply as
// many as you want. Flag does automatic numbering so you do not need to
// do it manually. If the list is empty, it is treated as no examples.
// Example:
//
//     CASE 1 - If set to:
//                       []string{}   OR    nil
//     The argument help output would be:
//           -a, --append [ VALUE ]     to append a round number into the
//                                      inputs. This can be a very long value
//                                      and can be a very long description.
//
//     CASE 2 - If set to:
//                       []string{"$ ./program -a 123",
//                                "$ ./program --append 123",
//                                "$ ./program -append '123'",
//                       }
//     The argument help output would be:
//           -a, --append [ VALUE ]     to append a round number into the
//                                      inputs. This can be a very long value
//                                      and can be a very long description.
//                                      EXAMPLES
//                                        1) $ ./program -a 123
//                                        2) $ ./program --append 123
//                                        3) $ ./program -append '123'
//
// This is optional to fill in.
//
//
// 7. DisableHelp
//
// This is a switch to tell the manager to discard the printout for this flag
// while processng the arguments help output. The default is false (off).
//
// This is optional to fill in.
//
//
// 8. ExampleLabel
//
// This is the string that alters the "EXAMPLES" text for the HelpExamples.
// It is offered explicitly to facilitate i18n. By default, it uses English
// word "Examples".
//
// This is optional to fill in.
type Flag struct {
	Name         string
	Label        []string
	ValueLabel   string
	Value        interface{}
	Help         string
	HelpExamples []string
	DisableHelp  bool
	ExampleLabel string

	value string
}

func (f *Flag) setValue(value string) {
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
	f.value = s
}

//nolint:gocyclo
func (f *Flag) convert() {
	if f.value == "" {
		return
	}

	switch p := f.Value.(type) {
	case *uint:
		value, _ := strconv.ParseUint(f.value, base10, bit64)
		*p = uint(value)
	case *uint8:
		value, _ := strconv.ParseUint(f.value, base10, bit64)
		*p = uint8(value)
	case *uint16:
		value, _ := strconv.ParseUint(f.value, base10, bit64)
		*p = uint16(value)
	case *uint32:
		value, _ := strconv.ParseUint(f.value, base10, bit64)
		*p = uint32(value)
	case *uint64:
		*p, _ = strconv.ParseUint(f.value, base10, bit64)
	case *int:
		value, _ := strconv.ParseInt(f.value, base10, bit64)
		*p = int(value)
	case *int8:
		value, _ := strconv.ParseInt(f.value, base10, bit64)
		*p = int8(value)
	case *int16:
		value, _ := strconv.ParseInt(f.value, base10, bit64)
		*p = int16(value)
	case *int32:
		value, _ := strconv.ParseInt(f.value, base10, bit64)
		*p = int32(value)
	case *int64:
		*p, _ = strconv.ParseInt(f.value, base10, bit64)
	case *float32:
		value, _ := strconv.ParseFloat(f.value, bit64)
		*p = float32(value)
	case *float64:
		*p, _ = strconv.ParseFloat(f.value, bit64)
	case *string:
		*p = f.value
	case *bool:
		*p, _ = strconv.ParseBool(f.value)
	case *[]byte:
		*p = []byte(f.value)
	case **[]byte:
		value := []byte(f.value)
		*p = &value
	}
}

func (f *Flag) help(flagWidth uint,
	descriptionWidth uint,
	withExamples bool) (out string) {
	var l, s, title, format string
	var hl []string
	var i int

	if len(f.Label) == 0 || f.DisableHelp {
		return out
	}

	// 1. prepare
	format = fmt.Sprintf("%s%dv%s\n", "%-", flagWidth, "%v")
	l = f.labelStatement()

	// 2. format help statement
	hl = f.wordWrap(f.Help, descriptionWidth)
	for i, s = range hl {
		if i == 0 {
			out = fmt.Sprintf(format, l, s)
			continue
		}

		out += fmt.Sprintf(format, "", s)
	}

	// 3. format example section
	title = f.ExampleLabel
	if title == "" {
		title = exampleTextEN
	}

	if len(f.HelpExamples) != 0 && withExamples {
		hl = []string{title}
		i = 1

		for _, s = range f.HelpExamples {
			if s == "" {
				continue
			}

			hl = append(hl, fmt.Sprintf("  %d) %s", i, s))
			i++
		}

		hl = f.indentWrap(hl, flagWidth)
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

func (f *Flag) cleanupLabel() {
	sl := []string{}

	for _, s := range f.Label {
		if s != "" {
			sl = append(sl, s)
		}
	}

	f.Label = sl
}

func (f *Flag) labelStatement() (out string) {
	f.cleanupLabel()
	out = strings.Join(f.Label, ", ")

	if f.ValueLabel != "" {
		out += fmt.Sprintf(" [ %s ]", f.ValueLabel)
	}

	return out
}

func (f *Flag) wordWrap(in string, flagWidth uint) []string {
	s := &strhelper.Styler{}
	return s.WordWrap(in, flagWidth)
}

func (f *Flag) indentWrap(in []string, flagWidth uint) (out []string) {
	s := &strhelper.Styler{}
	return s.Indent(in, strhelper.CharSpaceIndent, flagWidth)
}
