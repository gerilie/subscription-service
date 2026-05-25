package ping

// Resp represents the ping response.
//
// It has a single field, Timestamp, which is a string representing the timestamp.
type Resp struct {
	Timestamp string `json:"timestamp"`
}
