package pools

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

// PoolPage is the page returned by a pager when traversing over a
// collection of pools.
type PoolPage struct {
	pagination.PageWithInfo
}

// IsEmpty checks whether a PoolPage struct is empty.
func (r PoolPage) IsEmpty() (bool, error) {
	is, err := ExtractPools(r)
	return len(is) == 0, err
}

// ExtractPools accepts a Page struct, specifically a PoolPage struct,
// and extracts the elements into a slice of Pool structs. In other words,
// a generic collection is mapped into a relevant slice.
func ExtractPools(r pagination.Page) ([]Pool, error) {
	var s []Pool
	err := (r.(PoolPage)).ExtractIntoSlicePtr(&s, "pools")
	if err != nil {
		return nil, err
	}
	return s, nil
}

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
