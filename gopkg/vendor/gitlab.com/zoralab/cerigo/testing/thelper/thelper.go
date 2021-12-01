package thelper

import (
	"fmt"
	"math/big"
	"strings"
	"testing"
	"time"
)

const (
	// AllType is the Flush configuration input that tells Flush(...)
	// to clean both Log and Failed containers.
	AllType = uint(1)

	// LogType is the Flush configuration input that tells Flush(...)
	// to clean only the Log containers.
	LogType = uint(2)

	// FailedType is the Flush configuration input that tells Flush(...)
	// to clean only the Failed containers.
	FailedType = uint(3)
)

// TController is the interface that allows THelper to report out all
// the messages.
//
// It takes 2 basic inputs:
//   1. Errorf - printing by failed cases
//   2. Logf   - printing by passed cases
type TController interface {
	Errorf(format string, args ...interface{})
	Logf(format string, args ...interface{})
}

// THelper is a test helper structure offering various assertions, generator
// and some calculator functions for simplifying the testing.
//
// This structure has 3 public elements:
//   1. Controller   - holding TController interface like * testing.T
//   2. FailKeyword  - the "FAIL" announcement that allows you to quickly
//                     search the failing cases one by one. It can be unique
//                     to you.
//   3. QuietMode    - instructs the helper to only print failed cases,
//                     leaving passed cases quiet.
//
// This structure has private elements so use NewTHelper(...) function to
// create one instead of using the conventional structre{} method.
type THelper struct {
	Controller  TController
	FailKeyword string
	QuietMode   bool
	log         []string
	failed      []string
}

// NewTHelper is to create a properly defined THelper object.
//
// It takes 1 input:
//   1. t    - your t *testing.T test case instructor
//
// By default, the public interfacing elements are configured in the following
// manners. They are customizable after the THelper object is created and
// initialized.
//   1. Controller   - the given *testing.T object
//   2. FailKeyword  - "TESTFAIL"
//   3. QuietMode    - false
func NewTHelper(t *testing.T) *THelper {
	return &THelper{
		Controller:  t,
		FailKeyword: "TESTFAIL",
		log:         []string{},
		failed:      []string{},
	}
}

// Errorf is to print an error message reporting to the T testing controller.
//
// It works similarly with testing.T.Errorf(...) function from testing package
// with a few differences:
//   1. messages are queued, only printed out when THelper.Conclude() (a.k.a.
//      concluding the test) is made.
func (h *THelper) Errorf(format string, a ...interface{}) {
	var s string

	if h.FailKeyword != "" {
		newFormat := "%v: " + format

		la := append([]interface{}{h.FailKeyword}, a...)
		s = fmt.Sprintf(newFormat, la...)
	} else {
		s = fmt.Sprintf(format, a...)
	}

	h.failed = append(h.failed, s)
}

// Logf is to print a log message reporting to the T testing controller.
//
// It works similarly with testing.T.Logf(...) function from testing package
// with a few differences:
//   1. messages are queued, only printed out when THelper.Conclude() (a.k.a.
//      concluding the test) is made.
func (h *THelper) Logf(format string, a ...interface{}) {
	s := fmt.Sprintf(format, a...)
	h.log = append(h.log, s)
}

// LogScenario is to log a compatible Scenario object holding the test case
// meta information.
//
// It utilizes THelper.Logf(...) to printout the the scenario elements in a
// standardized way. If a test suite defined a Scenario derivative types, the
// Scenario type conversion should happen in that test suite itself.
//
// LogScenario facilitates data map parameter that allows tester to print
// out any data alongside the logging process. The string label of the list
// is the field name while the value can be anything since it is a interface{}.
//
// The data value access is unsafe (without using any mutexes). Hence, it is
// the test suite responsibility to ensure the data concurrency safety before
// feeding into LogScenario data field.
//
// If data map list is not available (nil), LogScenario will stop printing
// after switches.
func (h *THelper) LogScenario(s Scenario, data map[string]interface{}) {
	h.Logf("CASE\n%v\n", s.UID)
	h.Logf("TEST TYPE\n%v\n", s.TestType)
	h.Logf("DESCRIPTION\n%v\n", strings.TrimSpace(s.Description))

	// process switches
	h.Logf("SWITCHES")

	for k, v := range s.Switches {
		h.Logf("%-20v = %v", k, v)
	}

	h.Logf("\n")

	// process data
	if len(data) == 0 {
		return
	}

	h.Logf("GOT")

	for k, v := range data {
		h.Logf("%-20v = %#v", k, v)
	}

	h.Logf("\n")
}

// Flush is to reset all message containers in a THelper structure to initial
// conditions for next run.
//
// Flush is made available in case of manual interventions.
func (h *THelper) Flush(t uint) {
	switch t {
	case LogType:
		h.log = []string{}
	case FailedType:
		h.failed = []string{}
	case AllType:
		h.log = []string{}
		h.failed = []string{}
	}
}

// Conclude processes the message containers and pretty print them out under
// a single output message.
//
// Conclude uses the internal TController interface to perform the printing.
// The outcome can be either:
//   1. using the Errorf(...) when there is any logged error/failed messages
//   2. using the Logf(...) for normal passed cases' messages
//
// If QuietMode is true for the #2, Conclude() will not print anything.
//
// Upon successful output, the THelper structure automatically runs
// Flush(AllType) function, setting itself clean for next test run.
func (h *THelper) Conclude() {
	var log, err string

	if h.Controller == nil {
		goto endConclusion
	}

	// process error messages
	err = strings.Join(h.failed, "\n")

	// process log messages
	log = strings.Join(h.log, "\n")

endConclusion:
	switch {
	case h.QuietMode && err == "":
	case h.QuietMode && err != "":
		fallthrough
	case log != "" && err != "":
		h.Controller.Errorf("\n%v\n%v\n\n\n", err, log)
	case log != "":
		h.Controller.Logf("\n%v\n\n\n", log)
	case err != "":
		h.Controller.Errorf("\n%v\n\n\n", err)
	}
	h.Flush(AllType)
}

// ExpectExists is an assertion to check a given object is nil or exist.
//
// It takes 3 inputs:
//   1. label       - for error logging purposes, naming the asserted object.
//   2. subject     - the test subject being asserted.
//   3. expected    - the expectation for the subject existence. True means it
//                    is expected to exist.
//
// This function automatically logs the error message into the THelper if the
// test subject is behaving outside of expectation.
//
// It returns:
//   1. 0     - passed assertion
//   2. non-0 - failed assertion
func (h *THelper) ExpectExists(label string,
	subject interface{},
	expected bool) int {
	exist := false

	switch x := subject.(type) {
	case bool:
		if x {
			exist = true
		}
	default:
		if x != nil {
			exist = true
		}
	}

	if (expected && exist) || (!expected && !exist) {
		return 0
	}

	s := label
	if s == "" {
		s = "object"
	}

	h.Errorf("%s - expect exist=%v, got=%v", s, expected, exist)

	return 1
}

// ExpectError is an assertion to compare an error object existence.
//
// It takes 2 inputs:
//   1. subject     - the error test subject
//   2. expected    - the expectation for the subject existence.
//
// If the test subject is behaving outside of the given expectation, this
// function automatically logs the error message into the THelper.
//
// It returns:
//   1. 0     - passed assertion
//   2. non-0 - failed assertion
func (h *THelper) ExpectError(subject error, expected bool) int {
	return h.ExpectExists("error", subject, expected)
}

// SnapTime is a function to snap the current time.
//
// It returns:
//   1. Time        - time.Now() object.
func (h *THelper) SnapTime() *time.Time {
	s := time.Now()
	return &s
}

// CalculateDuration is a function to calculate time duration from the given
// start and stop time.
//
// It calculates by taking the stop time subtracting the start time. Hence, if
// the 2 inputs are misplaced, the return value should be negative.
//
// It takes 2 inputs:
//   1. start    - start time
//   2. stop     - stop time
//
// It returns:
//   1. duration - time duration
func (h *THelper) CalculateDuration(start *time.Time,
	stop *time.Time) time.Duration {
	if start == nil || stop == nil {
		return 0
	}

	return stop.Sub(*start)
}

// CalculateTimeLimits is to calculate the 2 time limits (min and max) for an
// expected stop time.

// It first calculates the expected stop time which is by adding the expected
// duration time (duration) to the given start time (start).
//
// Then, It calculates:
//   1. the minimum limit by subtracting the stop time from the limit range
//      (range).
//   2. the maximum limit by adding the limit range (range) into the stop time.
//
// Providing a nil to the start time will cause the function to do nothing.
// returning nil for both minimum and maximum values.
//
// It takes 2 inputs:
//   1. start       - the reference time
//   2. duration    - the expected duration for the stop time in nanoseconds
//   3. ranges      - the symmetric time range between the stop time in
//                    nanoseconds.
//
// It returns:
//   1. minimum     - the timestamp at the stop time's minimum limit
//   2. maximum     - the timestamp at the stop time's maximum limit
func (h *THelper) CalculateTimeLimits(start *time.Time,
	duration int64,
	ranges int64) (minimum *time.Time, maximum *time.Time) {
	if start == nil {
		return nil, nil
	}

	stop := start.Add(time.Duration(duration))
	l := stop.Add(time.Duration(-ranges))
	t := stop.Add(time.Duration(ranges))

	return &l, &t
}

// ExpectInTime is an assertion to check a given time is within a given set of
// time limits.
//
// It takes 1 main input and 2 optional inputs. They are:
//   1. subject      - the given test subject
//   2. minimum      - the minimum time limit (optional)
//   3. maximum      - the maximum time limit (optional)
//
// Depending on the availability of minimum and maximum, ExpectInTime carries
// out the assertion differently:
//   1. If no limit is available, it always return true since the test subject
//      is always in range.
//   2. If only maximum is available, it will check the test subject is always
//      below tha maximum limit. It will assert failed if it went over.
//   3. If only minimum is available, it will check the test subject is always
//      above the minimum limit. It will assert failed if it went under.
//   4. If both maximum and minimum are available, it will check the test
//      subject falls within range. It will assert failed if it went outside.
//   5. If test subject is not available, it will assert an error.
//
// ExpectInTime automatically logs the error message when meeting a failed
// assertion. If everything is good, ExpectInTime does nothing.
//
// It returns:
//   1. 0      - assertion is a passed case
//   2. not 0  - assertion is a failed case
func (h *THelper) ExpectInTime(subject *time.Time,
	minimum *time.Time,
	maximum *time.Time) int {
	switch {
	case subject == nil:
		h.Errorf("bad assertion - missing time subject")
	case maximum == nil && minimum == nil:
		return 0
	case maximum != nil && minimum == nil:
		return h.expectWithinMaxTimeLimit(subject, maximum)
	case maximum == nil && minimum != nil:
		return h.expectWithinMinTimeLimit(subject, minimum)
	case maximum != nil && minimum != nil:
		return h.expectWithinTimeRange(subject, minimum, maximum)
	}

	return 1
}

func (h *THelper) expectWithinTimeRange(subject *time.Time,
	minimum *time.Time,
	maximum *time.Time) int {
	if maximum.Sub(*subject) > 0 && subject.Sub(*minimum) > 0 {
		return 0
	}

	h.Errorf("out of range: %v | %v | %v", *minimum, *subject, *maximum)

	return 1
}

func (h *THelper) expectWithinMinTimeLimit(subject *time.Time,
	minimum *time.Time) int {
	if subject.Sub(*minimum) > 0 {
		return 0
	}

	h.Errorf("over below limit: %v below %v", *subject, *minimum)

	return 1
}

func (h *THelper) expectWithinMaxTimeLimit(subject *time.Time,
	maximum *time.Time) int {
	if maximum.Sub(*subject) > 0 {
		return 0
	}

	h.Errorf("over maximum limit: %v over %v", *subject, *maximum)

	return 1
}

//nolint:dupl
// ExpectSameBytesSlices is an assertion to check 2 given bytes slices are the
// same.
//
// It checks for nil, length, and lastly byte by byte before returning a
// passing result. It automatically log the error message when encountering a
// failed assertion.
//
// If a and b bytes slices are nil, ExpectSameBytesSlices passes the assertion
// immediately since nil is always the same as nil.
//
// It takes 4 inputs:
//   1. labelA    - label for first slice (A) reporting purpose
//   2. a         - the first byte slice (A)
//   3. labelB    - label for second slice (B) reporting purpose
//   4. b         - the second byte slice (B)
//
// It returns:
//   1. 0         - passed assertion
//   2. 1         - A is nil but B is not nil
//   3. 2         - A is not nil but B is nil
//   4. 3         - both A and B has different length
//   4. 4         - incorrect data matching during byte-to-byte comparison
func (h *THelper) ExpectSameBytesSlices(labelA string,
	a *[]byte,
	labelB string,
	b *[]byte) int {
	switch {
	case a == nil && b == nil:
		return 0
	case a == nil && b != nil:
		h.Errorf("%s is nil while %s is not nil",
			labelA,
			labelB,
		)

		return 1
	case a != nil && b == nil:
		h.Errorf("%s is nil while %s is not nil",
			labelB,
			labelA,
		)

		return 2
	}

	la := len(*a)
	lb := len(*b)

	if la != lb {
		h.Errorf("%s and %s has incorrect length: %v|%v",
			labelA,
			labelB,
			la,
			lb,
		)

		return 3
	}

	for i := 0; i < la; i++ {
		if (*a)[i] != (*b)[i] {
			h.Errorf("%s and %s has incorrect data: %v|%v",
				labelA,
				labelB,
				*a,
				*b,
			)

			return 4
		}
	}

	return 0
}

// ExpectUIDCorrectness is to check the current index against the test case
// uid.
//
// If there is a mismatch, it automatically log a failed message.
//
// It takes 3 inputs:
//   1. index         - the current index number of the test cases (for loop)
//   2. uid           - the uid of the test case from the table driven testing
//   3. startFromZero - to state whether the uid starts from 0.
//
// It returns:
//   1. 0             - successful assertion
//   2. non-zero      - failed assertion
func (h *THelper) ExpectUIDCorrectness(index int,
	uid int,
	startFromZero bool) (verdict int) {
	i := index

	if !startFromZero {
		i++
	}

	if i == uid {
		return 0
	}

	h.Errorf("uid mistmatched: index=%v uid=%v", i, uid)

	return 1
}

//nolint:dupl
// ExpectSameStringSlices is an assertion to check 2 given string slices'
// contents are the same.
//
// It checks for nil, length, and lastly line by line before returning a
// passing result. It automatically logs the error message when encountering a
// failed assertion.
//
// If a and b bytes slices are nil, ExpectSameStringSlices(...) passes the
// assertion immediately since nil is always the same as nil.
//
// It takes 4 inputs:
//   1. labelA    - label for first slice (A) reporting purpose
//   2. a         - the first string slice (A)
//   3. labelB    - label for second slice (B) reporting purpose
//   4. b         - the second string slice (B)
//
// It returns:
//   1. 0         - passed assertion
//   2. 1         - A is nil but B is not nil
//   3. 2         - A is not nil but B is nil
//   4. 3         - both A and B has different length
//   4. 4         - incorrect data matching during line-by-line comparison
func (h *THelper) ExpectSameStringSlices(labelA string,
	a *[]string,
	labelB string,
	b *[]string) int {
	switch {
	case a == nil && b == nil:
		return 0
	case a == nil && b != nil:
		h.Errorf("%s is nil while %s is not nil",
			labelA,
			labelB,
		)

		return 1
	case a != nil && b == nil:
		h.Errorf("%s is nil while %s is not nil",
			labelB,
			labelA,
		)

		return 2
	}

	la := len(*a)
	lb := len(*b)

	if la != lb {
		h.Errorf("%s and %s has incorrect length: %v|%v",
			labelA,
			labelB,
			la,
			lb,
		)

		return 3
	}

	for i := 0; i < la; i++ {
		if (*a)[i] != (*b)[i] {
			h.Errorf("%s and %s has incorrect string line: %v|%v",
				labelA,
				labelB,
				*a,
				*b,
			)

			return 4
		}
	}

	return 0
}

// ExpectSameStrings is an assertion to compare 2 given strings are having
// the same contents.
//
// If the assertion fails, this function automatically logs an error message
// when using the 2 labels.
//
// It takes 4 inputs:
//   1. labelA    - label for first (a) string
//   2. a         - the first string
//   3. labelB    - label for second (b) string
//   4. b         - the second string
//
// It returns:
//   1. 0     - passed assertion
//   2. 1     - failed assertion
func (h *THelper) ExpectSameStrings(labelA string,
	subjectA string,
	labelB string,
	subjectB string) int {
	if subjectA != subjectB {
		h.Errorf("%s is not the same string as %s: %s|%s",
			labelA,
			labelB,
			subjectA,
			subjectB)

		return 1
	}

	return 0
}

// ExpectStringHasKeywords is an assertion to check a given sub-string is
// inside the string subject.
//
// If the assertion fails (such as missing substring), this function
// automatically logs an error message using the 2 labels.
//
// It takes 4 inputs:
//   1. labelA         - label for the subject
//   2. subject        - the subject string
//   3. labelB         - label for the substring
//   4. substring      - the sub-string
//
// It returns:
//   1. 0              - passed assertion
//   2. 1              - failed assertion
func (h *THelper) ExpectStringHasKeywords(labelA string,
	subject string,
	labelB string,
	substring string) int {
	if strings.Contains(subject, substring) {
		return 0
	}

	h.Errorf("%s does not have substring %s: %v|%v",
		labelA,
		labelB,
		subject,
		substring)

	return 1
}

// ExpectSameBool is an assertion to compare 2 given bools are having the same
// contents.
//
// If the assertion fails (such as the string pair are not the same), this
// function automatically logs an error message using the 2 labels.
//
// It takes 4 inputs:
//   1. labelA         - label for first (a) bool
//   2. a              - the first bool
//   3. labelB         - label for second (b) bool
//   4. b              - the second bool
//
// It returns:
//   1. 0              - passed assertion
//   2. 1              - failed assertion
func (h *THelper) ExpectSameBool(labelA string,
	subjectA bool,
	labelB string,
	subjectB bool) int {
	if subjectB == subjectA {
		return 0
	}

	h.Errorf("%s is not the same bool as %s: %v|%v",
		labelA,
		labelB,
		subjectA,
		subjectB)

	return 1
}

// ExpectSameFloat32 is an assertion to compare 2 given float32 are same.
//
// If the assertion fails (such as the string pair are not the same), this
// function automatically logs an error message using the 2 labels.
//
// It takes 4 inputs:
//   1. labelA         - label for first (a) bool
//   2. a              - the first float32
//   3. labelB         - label for second (b) bool
//   4. b              - the second float32
//
// It returns:
//   1. 0              - passed assertion
//   2. 1              - failed assertion
func (h *THelper) ExpectSameFloat32(labelA string,
	subjectA float32,
	labelB string,
	subjectB float32) int {
	if subjectB == subjectA {
		return 0
	}

	h.Errorf("%s is not the same float32 as %s: %v|%v",
		labelA,
		labelB,
		subjectA,
		subjectB)

	return -1
}

// ExpectSameFloat64 is an assertion to compare 2 given float64 are same.
//
// If the assertion fails (such as the string pair are not the same), this
// function automatically logs an error message using the 2 labels.
//
// It takes 4 inputs:
//   1. labelA         - label for first (a) bool
//   2. a              - the first float64
//   3. labelB         - label for second (b) bool
//   4. b              - the second float64
//
// It returns:
//   1. 0              - passed assertion
//   2. 1              - failed assertion
func (h *THelper) ExpectSameFloat64(labelA string,
	subjectA float64,
	labelB string,
	subjectB float64) int {
	if subjectB == subjectA {
		return 0
	}

	h.Errorf("%s is not the same float64 as %s: %v|%v",
		labelA,
		labelB,
		subjectA,
		subjectB)

	return -1
}

// ExpectSameBigFloat is an assertion to compare 2 math/big.Float are same.
//
// If the assertion fails (such as the string pair are not the same), this
// function automatically logs an error message using the 2 labels.
//
// It takes 4 inputs:
//   1. labelA         - label for first (a) bool
//   2. a              - the first *big.Float
//   3. labelB         - label for second (b) bool
//   4. b              - the second *big.Float
//
// It returns:
//   1. 0              - passed assertion (including a == b == nil)
//   2. 1              - failed assertion
func (h *THelper) ExpectSameBigFloat(labelA string,
	subjectA *big.Float,
	labelB string,
	subjectB *big.Float) int {
	switch {
	case subjectA == nil && subjectB != nil:
		h.Errorf("%s is nil while %s is not: %v|%v",
			labelA,
			labelB,
			subjectA,
			subjectB)

		return -1
	case subjectA != nil && subjectB == nil:
		h.Errorf("%s is nil while %s is not: %v|%v",
			labelB,
			labelA,
			subjectB,
			subjectA)

		return -1
	case subjectA == nil && subjectB == nil:
		return 0
	}

	if subjectA.Cmp(subjectB) == 0 {
		return 0
	}

	h.Errorf("%s is not the same *big.Float as %s: %v|%v",
		labelA,
		labelB,
		subjectA,
		subjectB)

	return -1
}
