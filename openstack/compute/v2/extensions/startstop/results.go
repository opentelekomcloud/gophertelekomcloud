package startstop

// StartResult is the response from a Start operation. Call its ExtractErr
// method to determine if the request succeeded or failed.
type StartResult struct {
	Err error
}

// StopResult is the response from Stop operation. Call its ExtractErr
// method to determine if the request succeeded or failed.
type StopResult struct {
	Err error
}
