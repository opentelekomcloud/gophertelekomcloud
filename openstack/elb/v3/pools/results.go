package pools

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
)

type commonResult struct {
	golangsdk.Result
}

// Extract is a function that accepts a result and extracts a pool.
func (r commonResult) Extract() (*Pool, error) {
	s := new(Pool)
	err := r.ExtractIntoStructPtr(s, "pool")
	if err != nil {
		return nil, err
	}
	return s, nil
}

// CreateResult represents the result of a Create operation. Call its Extract
// method to interpret the result as a Pool.
type CreateResult struct {
	commonResult
}

// GetResult represents the result of a Get operation. Call its Extract
// method to interpret the result as a Pool.
type GetResult struct {
	commonResult
}

// UpdateResult represents the result of an Update operation. Call its Extract
// method to interpret the result as a Pool.
type UpdateResult struct {
	commonResult
}

// DeleteResult represents the result of a Delete operation. Call its
// ExtractErr method to determine if the request succeeded or failed.
type DeleteResult struct {
	golangsdk.ErrResult
}
