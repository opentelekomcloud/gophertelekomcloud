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

// DeleteResult represents the result of a Delete operation. Call its
// ExtractErr method to determine if the request succeeded or failed.
type DeleteResult struct {
	golangsdk.ErrResult
}
