package validation

// Errors represents a map of field names to their corresponding validation error messages.
type Errors map[string]string

// Resp represents a response containing validation errors grouped by fields.
//
// Fields holds a map where keys are field names and values are error messages.
type Resp struct {
	Fields Errors `json:"fields"`
}
