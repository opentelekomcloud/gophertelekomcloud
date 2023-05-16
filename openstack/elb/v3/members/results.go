package members

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

// MemberPage is the page returned by a pager when traversing over a
// collection of Members in a Pool.
type MemberPage struct {
	pagination.PageWithInfo
}

func (p MemberPage) IsEmpty() (bool, error) {
	l, err := ExtractMembers(p)
	if err != nil {
		return false, err
	}
	return len(l) == 0, nil
}

// ExtractMembers accepts a Page struct, specifically a MemberPage struct,
// and extracts the elements into a slice of Members structs. In other words,
// a generic collection is mapped into a relevant slice.
func ExtractMembers(r pagination.Page) ([]Member, error) {
	var s []Member
	err := (r.(MemberPage)).ExtractIntoSlicePtr(&s, "members")
	if err != nil {
		return nil, err
	}
	return s, err
}

type commonResult struct {
	golangsdk.Result
}

// Extract is a function that accepts a result and extracts a Member.
func (r commonResult) Extract() (*Member, error) {
	s := new(Member)
	err := r.ExtractIntoStructPtr(s, "member")
	if err != nil {
		return nil, err
	}
	return s, nil
}

// CreateResult represents the result of a Create operation.
// Call its Extract method to interpret it as a Member.
type CreateResult struct {
	commonResult
}

// GetResult represents the result of a Get operation.
// Call its Extract method to interpret it as a Member.
type GetResult struct {
	commonResult
}

// UpdateResult represents the result of an Update operation.
// Call its Extract method to interpret it as a Member.
type UpdateResult struct {
	commonResult
}

// DeleteResult represents the result of a Delete operation.
// Call its ExtractErr method to determine if the request succeeded or failed.
type DeleteResult struct {
	golangsdk.ErrResult
}
