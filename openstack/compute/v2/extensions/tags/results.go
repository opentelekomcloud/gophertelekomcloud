package tags

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type commonResult struct {
	golangsdk.Result
}

// Tags model
type Tags struct {
	// Tags is a list of any tags. Tags are arbitrarily defined strings
	// attached to a resource.
	Tags []string `json:"tags"`
}

// Extract interprets any commonResult as a Tags.
func (raw commonResult) Extract() (*Tags, error) {
	var res Tags
	err = extract.Into(raw, &res)
	return &res, err
}

// CreateResult represents the result of a Create operation
type CreateResult struct {
	commonResult
}

// GetResult represents the result of a Get operation
type GetResult struct {
	commonResult
}

// DeleteResult model
type DeleteResult struct {
	golangsdk.ErrResult
}
