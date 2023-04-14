package certificates

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
)

type commonResult struct {
	golangsdk.Result
}

func (r commonResult) Extract() (*Certificate, error) {
	s := new(Certificate)
	err := r.ExtractIntoStructPtr(s, "certificate")
	if err != nil {
		return nil, err
	}
	return s, nil
}

type CreateResult struct {
	commonResult
}

// GetResult represents the result of a get operation.
type GetResult struct {
	commonResult
}

type UpdateResult struct {
	commonResult
}

// DeleteResult represents the result of a delete operation.
type DeleteResult struct {
	golangsdk.ErrResult
}
