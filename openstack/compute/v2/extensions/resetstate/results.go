package resetstate

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
)

// ResetResult is the response of a ResetState operation. Call its ExtractErr
// method to determine if the request suceeded or failed.
type ResetResult struct {
	golangsdk.ErrResult
}
