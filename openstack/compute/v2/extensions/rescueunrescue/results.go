package rescueunrescue

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type commonResult struct {
	golangsdk.Result
}

// RescueResult is the response from a Rescue operation. Call its Extract
// method to retrieve adminPass for a rescued server.
type RescueResult struct {
	commonResult
}

// UnrescueResult is the response from an UnRescue operation. Call its ExtractErr
// method to determine if the call succeeded or failed.
type UnrescueResult struct {
	golangsdk.ErrResult
}

// Extract interprets any RescueResult as an AdminPass, if possible.
func (raw RescueResult) Extract() (string, error) {
	var res struct {
		AdminPass string `json:"adminPass"`
	}
	err = extract.Into(raw.Body, &res)
	return &res, err
}
