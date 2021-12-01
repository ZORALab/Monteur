// thelper is a test helper package complimenting the standard "testing"
// package.
//
// OBJECTIVES
//
// Its primary goal is to extend the "testing" package by:
//   1. facilitating large scale unit testing (>1000 test cases)
//   2. respecting test codes like the source codes
//   3. simplify any common assertions.
//
// PROBLEMS
//
// The biggest problem for testing in Go is that no-one mentions about large
// scale testing and continuous test development. Most of the tutorial guides
// only mention table-driven approach and shows the basic of it.
//
// In reality, package does grow incrementally and as it grows larger, the
// existing test approaches becomes a nightmare to maintain.
//
// APPROACHES
//
// thelper is still using table-driven test approach like any other gophers.
// The differences are:
//   1. how thelper organizes the source codes and segments of test codes
//   2. use simulation test approach instead of direct testing approach
//   3. the use of switches
//
// This way, it does not limits tester's flexibility in testing while still
// able some systematic test implementations across various packages.
package thelper
