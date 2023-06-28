package monitors

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type commonResult struct {
	golangsdk.Result
}

// Extract is a function that accepts a result and extracts a Monitor.
func (r commonResult) Extract() (*Monitor, error) {
	var res Monitor
	err := extract.IntoStructPtr(res, "healthmonitor")
	return &res, err
}

// CreateResult represents the result of a create operation. Call its Extract
// method to interpret it as a Monitor.
type CreateResult struct {
	commonResult
}

// GetResult represents the result of a get operation. Call its Extract
// method to interpret it as a Monitor.
type GetResult struct {
	commonResult
}

// UpdateResult represents the result of an update operation. Call its Extract
// method to interpret it as a Monitor.
type UpdateResult struct {
	commonResult
}

// DeleteResult represents the result of a delete operation. Call its
// ExtractErr method to determine if the result succeeded or failed.
type DeleteResult struct {
	golangsdk.ErrResult
}
