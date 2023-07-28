package pauseunpause

// PauseResult is the response from a Pause operation. Call its ExtractErr
// method to determine if the request succeeded or failed.
type PauseResult struct {
	Err error
}

// UnpauseResult is the response from an Unpause operation. Call its ExtractErr
// method to determine if the request succeeded or failed.
type UnpauseResult struct {
	Err error
}
