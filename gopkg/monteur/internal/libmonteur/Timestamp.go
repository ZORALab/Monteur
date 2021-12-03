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

package libmonteur

// Timestamp is the time and date data for parsing from a config file.
//
// By default, the `String()` function generates a timestamp with the following
// format: `2021-12-03T10:41:56+08:00` where symbols like `-`; `T`; `:` except
// Zone (`08:00`); and `+` are automatically filled. Any missing value shall be
// replaced by `X` of its placement. To utilize this default format, the
// supplied data shall only be in digits except Zone. The example above would
// be:
//   1. `Year` = `"2021"`
//   2. `Month` = `"12"`
//   3. `Day` = `"03"`
//   4. `Hour` = `"10"`
//   5. `Minute` = `"41"`
//   6. `Second` = `"56"`
//   7. `Zone` = `"08:00`
//
// If your data is outside of the above format (say `Month` = `"Dec"`), you
// will need to build your own timestamp `String()` function and can only use
// the data structure as a structured data schema.
type Timestamp struct {
	// Day is the day of the date
	Day string

	// Month is the month of the date
	Month string

	// Year is the year of the date
	Year string

	// Hour is the hour of the time
	Hour string

	// Minute is the minute of the time
	Minute string

	// Second is the second of the time
	Second string

	// Zone is the zone of the time
	Zone string
}

func (me *Timestamp) String() (s string) {
	s = me.StringYear()
	s += "-"
	s += me.StringMonth()
	s += "-"
	s += me.StringDay()
	s += "T"
	s += me.StringHour()
	s += ":"
	s += me.StringMinute()
	s += ":"
	s += me.StringSecond()
	s += "+"
	s += me.StringZone()

	return s
}

func (me *Timestamp) StringZone() string {
	if me.Zone != "" {
		return me.Zone
	}

	return "XX"
}

func (me *Timestamp) StringSecond() string {
	if me.Second != "" {
		return me.Second
	}

	return "XX"
}

func (me *Timestamp) StringMinute() string {
	if me.Minute != "" {
		return me.Minute
	}

	return "XX"
}

func (me *Timestamp) StringHour() string {
	if me.Hour != "" {
		return me.Hour
	}

	return "XX"
}

func (me *Timestamp) StringDay() string {
	if me.Day != "" {
		return me.Day
	}

	return "XX"
}

func (me *Timestamp) StringMonth() string {
	if me.Month != "" {
		return me.Month
	}

	return "XX"
}

func (me *Timestamp) StringYear() string {
	if me.Year != "" {
		return me.Year
	}

	return "XXXX"
}
