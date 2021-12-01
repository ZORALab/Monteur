package thelper

// Scenario is a structure holding the standard simulation scenario data
// aligning to table-driven, simulation test approach.
//
// It is designed to cater simulation test approach (using switches to toggle
// test environment preparations). The idea is that the test suite has a
// generator functions set that creates a simulation test environment for
// testing.

// Upon tested, the test suite uses its own assertion to determine the test
// results. Scenario relies heavily on the following standard test processes:
//   1. generate all the scenarios in the list (table driven approach).
//   2. loop through each scenarios:
//     2.1. check the TestType is meant for the test script (algorithm)
//          2.1.1. skip the test case if not compatible using 'continue'
//     2.2. prepare simulation environments
//          2.2.1. prepare happy path parameters
//          2.2.2. alters the parameters accordingly based on Switches.
//          2.2.3. generates BOTH the test inputs and the expected output.
//     2.3. run the test
//          2.3.1. execute the test based on the generated test inputs.
//          2.3.2. capture the test outputs.
//     2.4. assert the result using the generated expected output and test
//          output.
//     2.5. log every details for easier assertion statement.
//   3. repeat step 1 for another test script (algorithm)
//
// To configure the simulation environment, the test suite uses the
// map[string]bool Switches field. It allows the tester to dynamically
// configure the simulation preparation functions, case by case basis.
//
// Also, Scenario facilitates isolation for future test algorithm
// developments for large scale unit testing.
//
type Scenario struct {
	UID         int
	TestType    string
	Description string
	Switches    map[string]bool
}
