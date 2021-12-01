package strhelper

import (
	"strings"
)

const (
	spaceLength = 1
)

const (
	// CharTabIndent is the tab character for indent usage
	CharTabIndent = "\t"

	// CharSpaceIndent is the space character for indent usage
	CharSpaceIndent = " "

	// CharLegacyNewLine is the CR character ("\r") for newline usage
	CharLegacyNewLine = "\r"

	// CharUNIXNewLine is the LF character ("\n") for newline usage,
	// commonly used in modern UNIX systems.
	CharUNIXNewLine = "\n"

	// CharWindowsNewLine is the CRLF character("\r\n") for newline usage,
	// commonly used in modern Windows systems.
	CharWindowsNewLine = "\r\n"
)

// Styler is the data structure for styling a given string. It is safe to
// create using the conventional structure{} method.
type Styler struct {
}

// ToUnix is to convert a set of given paragraphs (regardless input os) to use
// Unix newline convention: LF ("\n").
func (s *Styler) ToUnix(paragraphs string) (out string) {
	out = strings.Replace(paragraphs,
		CharWindowsNewLine,
		CharUNIXNewLine, -1)
	out = strings.Replace(out,
		CharLegacyNewLine,
		CharUNIXNewLine,
		-1)

	return out
}

// ToWindows is to convert a set of given paragraphs (regardless input os) to
// use Windows new line convention: CRLF ("\r\n").
func (s *Styler) ToWindows(paragraphs string) (out string) {
	// convert to unix first to avoid replacing any leftover CRLF into
	// LFLF by accident.
	out = s.ToUnix(paragraphs)
	out = strings.Replace(out,
		CharUNIXNewLine,
		CharWindowsNewLine,
		-1)

	return out
}

// WordWrap wraps a given paragraph into characters-width bounded set of
// multi-line strings. If multiple paragraphs are available under a single
// string input, they will be merged into a single paragraph.
//
// When the limit is below a word length, this function returns a minimum of
// 1 word regardless of its word length.

// If the limit is 0, this function returns nil.
//
// It returns:
//   1. data          - successful word-wrap in string slice
//   2. nil           - incomprehensible limit for wrapping
func (s *Styler) WordWrap(paragraph string, limit uint) (out []string) {
	if limit == 0 {
		return nil
	}

	if paragraph == "" {
		return []string{}
	}

	list := strings.Fields(paragraph)
	out = []string{}
	col := 0
	colMAX := int(limit)
	line := ""

	for _, word := range list {
		wordLength := len(word)
		col1 := col + wordLength + spaceLength

		switch {
		case col1 >= colMAX:
			line = strings.TrimSpace(line)
			if line != "" {
				out = append(out, line)
			}

			line = word
			col = wordLength
		case col1 < colMAX:
			col = col1

			if line != "" {
				line += " "
			}

			line += word
		}
	}

	// append the last line
	line = strings.TrimSpace(line)
	out = append(out, line)

	return out
}

// ContentWrap performs multiple paragraphs word wrapping against a
// characters-width boundary limit, producing a slice with multi-line strings
// with designed newline convention to separate each paragraphs.
//
// By default, ContentWrap uses the LF ("\n") newLine characters. This is true
// when newLineCharacter is empty ("") or defined to use LF
// (strhelper.CharUNIXNewLine). If you need other variant like CRLF ("\r\n"),
// you can supply strhelper.CharWindowsNewLine to newLineCharacter parameter.
// If the newLineCharacter has unrecognizable string character(s), this
// function returns nil.
//
// When the limit is below a word length, this function returns a minimum of
// 1 word regardless of its word length.
//
// If the limit is 0, this function returns nil.
//
// It returns:
//   1. data          - successful word-wrap in string slice
//   2. nil           - incomprehensible input(s) for wrapping
func (s *Styler) ContentWrap(paragraphs string,
	limit uint,
	newLineCharacter string) (out []string) {
	var newLine string

	// 1. check inputs
	if limit == 0 {
		return nil
	}

	if paragraphs == "" {
		return []string{}
	}

	switch newLineCharacter {
	case CharUNIXNewLine:
		newLine = CharUNIXNewLine
	case CharWindowsNewLine:
		newLine = CharWindowsNewLine
	case CharLegacyNewLine:
		newLine = CharLegacyNewLine
	default:
		return nil
	}

	// 2. convert to UNIX convention and then split the paragraphs
	p := s.ToUnix(paragraphs)
	list := strings.Split(p, CharUNIXNewLine)

	// 3. wrap each paragraphs
	out = []string{}

	for _, paragraph := range list {
		l := s.WordWrap(paragraph, limit)
		if len(l) == 0 {
			continue
		}

		if len(out) != 0 {
			out = append(out, newLine)
		}

		out = append(out, l...)
		out = append(out, newLine)
	}

	return out
}

// Indent is shifting a paragraph list towards left by the defined size
// (indentSize) using the given indent character (indentChar). This function
// creates the indent by repeatedly generate the given indent character n
// times, where n is the defined size before proceeding to next line.
//
// By default, the character is space (strhelper.CharSpaceIndent). If tab
// character (strhelper.CharTabIndent) is provided, IndentLeft will use tab
// instead. If you provide anything other than those 2, this function returns
// nil.
//
// If the given paragraph is having 0 length or nil, it returns the given
// paragraph as output.
//
// If the given indentSize is 0, however, it returns nil instead.
//
// It returns:
//   1.  data     - successful indent
//   2.  nil      - indent size is 0 or given paragraph is 0 or nil
func (s *Styler) Indent(paragraph []string,
	indentChar string,
	indentSize uint) (out []string) {
	var indent string

	// 1. check variables
	if len(paragraph) == 0 {
		return paragraph
	}

	if indentSize == 0 {
		return nil
	}

	switch indentChar {
	case CharSpaceIndent:
		indent = CharSpaceIndent
	case CharTabIndent:
		indent = CharTabIndent
	default:
		return nil
	}

	// 2. prepare
	out = []string{}
	indent = strings.Repeat(indent, int(indentSize))

	// 3. apply styling
	for _, line := range paragraph {
		switch {
		case line == CharUNIXNewLine:
			out = append(out, line)
		case line == CharWindowsNewLine:
			out = append(out, line)
		case line == CharLegacyNewLine:
			out = append(out, line)
		default:
			out = append(out, indent+line)
		}
	}

	return out
}
