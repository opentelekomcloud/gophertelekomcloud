package suspendresume

// SuspendResult is the response from a Suspend operation. Call its
// ExtractErr method to determine if the request succeeded or failed.
type SuspendResult struct {
	Err error
}

// UnsuspendResult is the response from an Unsuspend operation. Call
// its ExtractErr method to determine if the request succeeded or failed.
type UnsuspendResult struct {
	Err error
}
