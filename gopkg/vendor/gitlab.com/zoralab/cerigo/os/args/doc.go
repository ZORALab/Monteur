// Package args is an extended command line interface processor from the
// standard library's flag package. It uses the standard library's "os.Args"
// as the primary input.
//
//
// DIFFERENCES
//
// 1. the Manager object is completely discardable after argument parsing.
//
// 2. Help (-h, --help) is not run by default, allowing one to customize
// accordingly.
//
// 3. Support flag with multiple argument labels like "-h", "--help", "help".
//
//
// SUPPORTED DATA TYPES
//
// args package has limited data type conversion which is strictly towards
// the basic types. It does not use Go's reflect package for code clarity
// reason. They are:
//
//   1. int, int8, int16, int32, int64
//   2. uint, uint8, uint16, uint32, uint64
//   3. float32, float64
//   4. bool
//   5. []byte, *[]byte
//   6. string
//
// If you need to support some custom definition type, you need to convert it
// on your own. Any unsupported data types will be ignored in the conversion
// steps.
//
// **Important Note**:
// You need to supply these data types' variable pointer into Flag's Value
// element. Example, to support variable a that is int type, you fill in the
// Flag's Value as &a instead. Here is an example:
//   var a int
//   f := Flag {
//     ...
//     Value: &a,
//     ...
//   }
//
// UNSUPPORTED DATA TYPES
//
// These are the data types that were confirmed not to be supported to keep
// the package sane to use. They are:
//
//   1. any forms of array
//   2. any forms of maps
//
// *Rationale*:
// It is hard to track all permutations and combinations without using reflect
// package. Also, each time when array or maps occurs, it is use-case specific
// and can't be generalized without complicated algorithms.
//
// Hence, for these cases, user is advised to parse the entire list as string
// and then perform the splitting on your side.
package args
