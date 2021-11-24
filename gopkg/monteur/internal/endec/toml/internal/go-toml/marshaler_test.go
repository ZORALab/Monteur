package toml_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"github.com/pelletier/go-toml/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

//nolint:funlen
func TestMarshal(t *testing.T) {
	someInt := 42

	type structInline struct {
		A interface{} `inline:"true"`
	}

	examples := []struct {
		desc     string
		v        interface{}
		expected string
		err      bool
	}{
		{
			desc: "simple map and string",
			v: map[string]string{
				"hello": "world",
			},
			expected: "hello = 'world'",
		},
		{
			desc: "map with new line in key",
			v: map[string]string{
				"hel\nlo": "world",
			},
			err: true,
		},
		{
			desc: `map with " in key`,
			v: map[string]string{
				`hel"lo`: "world",
			},
			expected: `'hel"lo' = 'world'`,
		},
		{
			desc: "map in map and string",
			v: map[string]map[string]string{
				"table": {
					"hello": "world",
				},
			},
			expected: `
[table]
hello = 'world'`,
		},
		{
			desc: "map in map in map and string",
			v: map[string]map[string]map[string]string{
				"this": {
					"is": {
						"a": "test",
					},
				},
			},
			expected: `
[this]
[this.is]
a = 'test'`,
		},
		{
			desc: "map in map in map and string with values",
			v: map[string]interface{}{
				"this": map[string]interface{}{
					"is": map[string]string{
						"a": "test",
					},
					"also": "that",
				},
			},
			expected: `
[this]
also = 'that'
[this.is]
a = 'test'`,
		},
		{
			desc: "simple string array",
			v: map[string][]string{
				"array": {"one", "two", "three"},
			},
			expected: `array = ['one', 'two', 'three']`,
		},
		{
			desc:     "empty string array",
			v:        map[string][]string{},
			expected: ``,
		},
		{
			desc:     "map",
			v:        map[string][]string{},
			expected: ``,
		},
		{
			desc: "nested string arrays",
			v: map[string][][]string{
				"array": {{"one", "two"}, {"three"}},
			},
			expected: `array = [['one', 'two'], ['three']]`,
		},
		{
			desc: "mixed strings and nested string arrays",
			v: map[string][]interface{}{
				"array": {"a string", []string{"one", "two"}, "last"},
			},
			expected: `array = ['a string', ['one', 'two'], 'last']`,
		},
		{
			desc: "array of maps",
			v: map[string][]map[string]string{
				"top": {
					{"map1.1": "v1.1"},
					{"map2.1": "v2.1"},
				},
			},
			expected: `
[[top]]
'map1.1' = 'v1.1'
[[top]]
'map2.1' = 'v2.1'
`,
		},
		{
			desc: "map with two keys",
			v: map[string]string{
				"key1": "value1",
				"key2": "value2",
			},
			expected: `
key1 = 'value1'
key2 = 'value2'`,
		},
		{
			desc: "simple struct",
			v: struct {
				A string
			}{
				A: "foo",
			},
			expected: `A = 'foo'`,
		},
		{
			desc: "one level of structs within structs",
			v: struct {
				A interface{}
			}{
				A: struct {
					K1 string
					K2 string
				}{
					K1: "v1",
					K2: "v2",
				},
			},
			expected: `
[A]
K1 = 'v1'
K2 = 'v2'
`,
		},
		{
			desc: "structs in array with interfaces",
			v: map[string]interface{}{
				"root": map[string]interface{}{
					"nested": []interface{}{
						map[string]interface{}{"name": "Bob"},
						map[string]interface{}{"name": "Alice"},
					},
				},
			},
			expected: `
[root]
[[root.nested]]
name = 'Bob'
[[root.nested]]
name = 'Alice'
`,
		},
		{
			desc: "string escapes",
			v: map[string]interface{}{
				"a": `'"\`,
			},
			expected: `a = "'\"\\"`,
		},
		{
			desc: "string utf8 low",
			v: map[string]interface{}{
				"a": "'Ę",
			},
			expected: `a = "'Ę"`,
		},
		{
			desc: "string utf8 low 2",
			v: map[string]interface{}{
				"a": "'\u10A85",
			},
			expected: "a = \"'\u10A85\"",
		},
		{
			desc: "string utf8 low 2",
			v: map[string]interface{}{
				"a": "'\u10A85",
			},
			expected: "a = \"'\u10A85\"",
		},
		{
			desc: "emoji",
			v: map[string]interface{}{
				"a": "'😀",
			},
			expected: "a = \"'😀\"",
		},
		{
			desc: "control char",
			v: map[string]interface{}{
				"a": "'\u001A",
			},
			expected: `a = "'\u001A"`,
		},
		{
			desc: "multi-line string",
			v: map[string]interface{}{
				"a": "hello\nworld",
			},
			expected: `a = "hello\nworld"`,
		},
		{
			desc: "multi-line forced",
			v: struct {
				A string `multiline:"true"`
			}{
				A: "hello\nworld",
			},
			expected: `A = """
hello
world"""`,
		},
		{
			desc: "inline field",
			v: struct {
				A map[string]string `inline:"true"`
				B map[string]string
			}{
				A: map[string]string{
					"isinline": "yes",
				},
				B: map[string]string{
					"isinline": "no",
				},
			},
			expected: `
A = {isinline = 'yes'}
[B]
isinline = 'no'
`,
		},
		{
			desc: "mutiline array int",
			v: struct {
				A []int `multiline:"true"`
				B []int
			}{
				A: []int{1, 2, 3, 4},
				B: []int{1, 2, 3, 4},
			},
			expected: `
A = [
  1,
  2,
  3,
  4
]
B = [1, 2, 3, 4]
`,
		},
		{
			desc: "mutiline array in array",
			v: struct {
				A [][]int `multiline:"true"`
			}{
				A: [][]int{{1, 2}, {3, 4}},
			},
			expected: `
A = [
  [1, 2],
  [3, 4]
]
`,
		},
		{
			desc: "nil interface not supported at root",
			v:    nil,
			err:  true,
		},
		{
			desc: "nil interface not supported in slice",
			v: map[string]interface{}{
				"a": []interface{}{"a", nil, 2},
			},
			err: true,
		},
		{
			desc: "nil pointer in slice uses zero value",
			v: struct {
				A []*int
			}{
				A: []*int{nil},
			},
			expected: `A = [0]`,
		},
		{
			desc: "nil pointer in slice uses zero value",
			v: struct {
				A []*int
			}{
				A: []*int{nil},
			},
			expected: `A = [0]`,
		},
		{
			desc: "pointer in slice",
			v: struct {
				A []*int
			}{
				A: []*int{&someInt},
			},
			expected: `A = [42]`,
		},
		{
			desc: "inline table in inline table",
			v: structInline{
				A: structInline{
					A: structInline{
						A: "hello",
					},
				},
			},
			expected: `A = {A = {A = 'hello'}}`,
		},
		{
			desc: "empty slice in map",
			v: map[string][]string{
				"a": {},
			},
			expected: `a = []`,
		},
		{
			desc: "map in slice",
			v: map[string][]map[string]string{
				"a": {{"hello": "world"}},
			},
			expected: `
[[a]]
hello = 'world'`,
		},
		{
			desc: "newline in map in slice",
			v: map[string][]map[string]string{
				"a\n": {{"hello": "world"}},
			},
			err: true,
		},
		{
			desc: "newline in map in slice",
			v: map[string][]map[string]*customTextMarshaler{
				"a": {{"hello": &customTextMarshaler{1}}},
			},
			err: true,
		},
		{
			desc: "empty slice of empty struct",
			v: struct {
				A []struct{}
			}{
				A: []struct{}{},
			},
			expected: `A = []`,
		},
		{
			desc: "nil field is ignored",
			v: struct {
				A interface{}
			}{
				A: nil,
			},
			expected: ``,
		},
		{
			desc: "private fields are ignored",
			v: struct {
				Public  string
				private string
			}{
				Public:  "shown",
				private: "hidden",
			},
			expected: `Public = 'shown'`,
		},
		{
			desc: "fields tagged - are ignored",
			v: struct {
				Public  string `toml:"-"`
				private string
			}{
				Public: "hidden",
			},
			expected: ``,
		},
		{
			desc: "nil value in map is ignored",
			v: map[string]interface{}{
				"A": nil,
			},
			expected: ``,
		},
		{
			desc: "new line in table key",
			v: map[string]interface{}{
				"hello\nworld": 42,
			},
			err: true,
		},
		{
			desc: "new line in parent of nested table key",
			v: map[string]interface{}{
				"hello\nworld": map[string]interface{}{
					"inner": 42,
				},
			},
			err: true,
		},
		{
			desc: "new line in nested table key",
			v: map[string]interface{}{
				"parent": map[string]interface{}{
					"in\ner": map[string]interface{}{
						"foo": 42,
					},
				},
			},
			err: true,
		},
		{
			desc: "invalid map key",
			v:    map[int]interface{}{},
			err:  true,
		},
		{
			desc: "unhandled type",
			v: struct {
				A chan int
			}{
				A: make(chan int),
			},
			err: true,
		},
		{
			desc: "numbers",
			v: struct {
				A float32
				B uint64
				C uint32
				D uint16
				E uint8
				F uint
				G int64
				H int32
				I int16
				J int8
				K int
			}{
				A: 1.1,
				B: 42,
				C: 42,
				D: 42,
				E: 42,
				F: 42,
				G: 42,
				H: 42,
				I: 42,
				J: 42,
				K: 42,
			},
			expected: `
A = 1.1
B = 42
C = 42
D = 42
E = 42
F = 42
G = 42
H = 42
I = 42
J = 42
K = 42`,
		},
	}

	for _, e := range examples {
		e := e
		t.Run(e.desc, func(t *testing.T) {
			b, err := toml.Marshal(e.v)
			if e.err {
				require.Error(t, err)

				return
			}

			require.NoError(t, err)
			equalStringsIgnoreNewlines(t, e.expected, string(b))

			// make sure the output is always valid TOML
			defaultMap := map[string]interface{}{}
			err = toml.Unmarshal(b, &defaultMap)
			require.NoError(t, err)

			testWithAllFlags(t, func(t *testing.T, flags int) {
				t.Helper()

				var buf bytes.Buffer
				enc := toml.NewEncoder(&buf)
				setFlags(enc, flags)

				err := enc.Encode(e.v)
				require.NoError(t, err)

				inlineMap := map[string]interface{}{}
				err = toml.Unmarshal(buf.Bytes(), &inlineMap)
				require.NoError(t, err)

				require.Equal(t, defaultMap, inlineMap)
			})
		})
	}
}

type flagsSetters []struct {
	name string
	f    func(enc *toml.Encoder, flag bool)
}

var allFlags = flagsSetters{
	{"arrays-multiline", (*toml.Encoder).SetArraysMultiline},
	{"tables-inline", (*toml.Encoder).SetTablesInline},
	{"indent-tables", (*toml.Encoder).SetIndentTables},
}

func setFlags(enc *toml.Encoder, flags int) {
	for i := 0; i < len(allFlags); i++ {
		enabled := flags&1 > 0
		allFlags[i].f(enc, enabled)
	}
}

func testWithAllFlags(t *testing.T, testfn func(t *testing.T, flags int)) {
	t.Helper()
	testWithFlags(t, 0, allFlags, testfn)
}

func testWithFlags(t *testing.T, flags int, setters flagsSetters, testfn func(t *testing.T, flags int)) {
	t.Helper()

	if len(setters) == 0 {
		testfn(t, flags)

		return
	}

	s := setters[0]

	for _, enabled := range []bool{false, true} {
		name := fmt.Sprintf("%s=%t", s.name, enabled)
		newFlags := flags << 1

		if enabled {
			newFlags++
		}

		t.Run(name, func(t *testing.T) {
			testWithFlags(t, newFlags, setters[1:], testfn)
		})
	}
}

func equalStringsIgnoreNewlines(t *testing.T, expected string, actual string) {
	t.Helper()
	cutset := "\n"
	assert.Equal(t, strings.Trim(expected, cutset), strings.Trim(actual, cutset))
}

//nolint:funlen
func TestMarshalIndentTables(t *testing.T) {
	examples := []struct {
		desc     string
		v        interface{}
		expected string
	}{
		{
			desc: "one kv",
			v: map[string]interface{}{
				"foo": "bar",
			},
			expected: `foo = 'bar'`,
		},
		{
			desc: "one level table",
			v: map[string]map[string]string{
				"foo": {
					"one": "value1",
					"two": "value2",
				},
			},
			expected: `
[foo]
  one = 'value1'
  two = 'value2'
`,
		},
		{
			desc: "two levels table",
			v: map[string]interface{}{
				"root": "value0",
				"level1": map[string]interface{}{
					"one": "value1",
					"level2": map[string]interface{}{
						"two": "value2",
					},
				},
			},
			expected: `
root = 'value0'
[level1]
  one = 'value1'
  [level1.level2]
    two = 'value2'
`,
		},
	}

	for _, e := range examples {
		e := e
		t.Run(e.desc, func(t *testing.T) {
			var buf strings.Builder
			enc := toml.NewEncoder(&buf)
			enc.SetIndentTables(true)
			err := enc.Encode(e.v)
			require.NoError(t, err)
			equalStringsIgnoreNewlines(t, e.expected, buf.String())
		})
	}
}

type customTextMarshaler struct {
	value int64
}

func (c *customTextMarshaler) MarshalText() ([]byte, error) {
	if c.value == 1 {
		return nil, fmt.Errorf("cannot represent 1 because this is a silly test")
	}
	return []byte(fmt.Sprintf("::%d", c.value)), nil
}

func TestMarshalTextMarshaler_NoRoot(t *testing.T) {
	c := customTextMarshaler{}
	_, err := toml.Marshal(&c)
	require.Error(t, err)
}

func TestMarshalTextMarshaler_Error(t *testing.T) {
	m := map[string]interface{}{"a": &customTextMarshaler{value: 1}}
	_, err := toml.Marshal(m)
	require.Error(t, err)
}

func TestMarshalTextMarshaler_ErrorInline(t *testing.T) {
	type s struct {
		A map[string]interface{} `inline:"true"`
	}

	d := s{
		A: map[string]interface{}{"a": &customTextMarshaler{value: 1}},
	}

	_, err := toml.Marshal(d)
	require.Error(t, err)
}

func TestMarshalTextMarshaler(t *testing.T) {
	m := map[string]interface{}{"a": &customTextMarshaler{value: 2}}
	r, err := toml.Marshal(m)
	require.NoError(t, err)
	equalStringsIgnoreNewlines(t, "a = '::2'", string(r))
}

type brokenWriter struct{}

func (b *brokenWriter) Write([]byte) (int, error) {
	return 0, fmt.Errorf("dead")
}

func TestEncodeToBrokenWriter(t *testing.T) {
	w := brokenWriter{}
	enc := toml.NewEncoder(&w)
	err := enc.Encode(map[string]string{"hello": "world"})
	require.Error(t, err)
}

func TestEncoderSetIndentSymbol(t *testing.T) {
	var w strings.Builder
	enc := toml.NewEncoder(&w)
	enc.SetIndentTables(true)
	enc.SetIndentSymbol(">>>")
	err := enc.Encode(map[string]map[string]string{"parent": {"hello": "world"}})
	require.NoError(t, err)
	expected := `
[parent]
>>>hello = 'world'`
	equalStringsIgnoreNewlines(t, expected, w.String())
}

func TestIssue436(t *testing.T) {
	data := []byte(`{"a": [ { "b": { "c": "d" } } ]}`)

	var v interface{}
	err := json.Unmarshal(data, &v)
	require.NoError(t, err)

	var buf bytes.Buffer
	err = toml.NewEncoder(&buf).Encode(v)
	require.NoError(t, err)

	expected := `
[[a]]
[a.b]
c = 'd'
`
	equalStringsIgnoreNewlines(t, expected, buf.String())
}

func TestIssue424(t *testing.T) {
	type Message1 struct {
		Text string
	}

	type Message2 struct {
		Text string `multiline:"true"`
	}

	msg1 := Message1{"Hello\\World"}
	msg2 := Message2{"Hello\\World"}

	toml1, err := toml.Marshal(msg1)
	require.NoError(t, err)

	toml2, err := toml.Marshal(msg2)
	require.NoError(t, err)

	msg1parsed := Message1{}
	err = toml.Unmarshal(toml1, &msg1parsed)
	require.NoError(t, err)
	require.Equal(t, msg1, msg1parsed)

	msg2parsed := Message2{}
	err = toml.Unmarshal(toml2, &msg2parsed)
	require.NoError(t, err)
	require.Equal(t, msg2, msg2parsed)
}

func TestIssue567(t *testing.T) {
	var m map[string]interface{}
	err := toml.Unmarshal([]byte("A = 12:08:05"), &m)
	require.NoError(t, err)
	require.IsType(t, m["A"], toml.LocalTime{})
}

func TestIssue590(t *testing.T) {
	type CustomType int
	var cfg struct {
		Option CustomType `toml:"option"`
	}
	err := toml.Unmarshal([]byte("option = 42"), &cfg)
	require.NoError(t, err)
}

func ExampleMarshal() {
	type MyConfig struct {
		Version int
		Name    string
		Tags    []string
	}

	cfg := MyConfig{
		Version: 2,
		Name:    "go-toml",
		Tags:    []string{"go", "toml"},
	}

	b, err := toml.Marshal(cfg)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))

	// Output:
	// Version = 2
	// Name = 'go-toml'
	// Tags = ['go', 'toml']
}