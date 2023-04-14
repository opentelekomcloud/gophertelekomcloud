package listeners

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type commonResult struct {
	golangsdk.Result
}

// Extract is a function that accepts a result and extracts a listener.
func (r commonResult) Extract() (*Listener, error) {
	s := new(Listener)
	err := r.ExtractIntoStructPtr(s, "listener")
	if err != nil {
		return nil, err
	}
	return s, nil
}

// CreateResult represents the result of a create operation. Call its Extract
// method to interpret it as a Listener.
type CreateResult struct {
	commonResult
}

// GetResult represents the result of a get operation. Call its Extract
// method to interpret it as a Listener.
type GetResult struct {
	commonResult
}

// UpdateResult represents the result of an update operation. Call its Extract
// method to interpret it as a Listener.
type UpdateResult struct {
	commonResult
}

// DeleteResult represents the result of a delete operation. Call its
// ExtractErr method to determine if the request succeeded or failed.
type DeleteResult struct {
	golangsdk.ErrResult
}

type ListenerPage struct {
	pagination.PageWithInfo
}

func (p ListenerPage) IsEmpty() (bool, error) {
	l, err := ExtractListeners(p)
	if err != nil {
		return false, err
	}
	return len(l) == 0, nil
}

func ExtractListeners(r pagination.Page) ([]Listener, error) {
	var s []Listener
	err := (r.(ListenerPage)).ExtractIntoSlicePtr(&s, "listeners")
	if err != nil {
		return nil, err
	}
	return s, nil
}
