package ping

// pingResp represents the ping response.
//
// It has a single field, Timestamp, which is a string representing the timestamp.
type pingResp struct {
	Timestamp string `json:"timestamp"`
}
