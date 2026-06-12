package test

// Expected represents a generic test case with input arguments and expected output.
//
// A is the type of input arguments.
// W is the type of the expected result.
type Expected[A, E any] struct {
	Name     string
	Args     A
	Expected E
	Error    error
	IsError  bool
}

// NoExpected represents a generic test case with input arguments
// where no expected output value is asserted.
//
// A is the type of input arguments.
type NoExpected[A any] struct {
	Name    string
	Args    A
	Error   error
	IsError bool
}

// Request represents an HTTP request.
type Request[B any] struct {
	Code   int
	Method string
	Path   string
	Body   B
}

// Response represents a generic HTTP response with a status code, body, and content type.
//
// B is the type of the response body.
type Response[B any] struct {
	Code        int
	Body        B
	ContentType string
}
