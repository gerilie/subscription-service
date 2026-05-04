package test

// Test represents a generic test case with input arguments and expected output.
//
// A is the type of input arguments.
// W is the type of the expected result.
type Test[A, W any] struct {
	Name  string
	Args  A
	Want  W
	Error error
}

// TestNoExpected represents a generic test case with input arguments
// where no expected output value is asserted.
//
// A is the type of input arguments.
type TestNoExpected[A any] struct {
	Name  string
	Args  A
	Error error
}
