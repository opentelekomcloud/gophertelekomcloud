package nodes

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
)

// ListNode describes the Node Structure of cluster
type ListNode struct {
	// API type, fixed value "List"
	Kind string `json:"kind"`
	// API version, fixed value "v3"
	Apiversion string `json:"apiVersion"`
	// all Clusters
	Nodes []Nodes `json:"items"`
}

type commonResult struct {
	golangsdk.Result
}

// Extract is a function that accepts a result and extracts a node.
func (r commonResult) Extract() (*Nodes, error) {
	var s Nodes
	err := r.ExtractInto(&s)
	return &s, err
}

// ExtractNode is a function that accepts a ListOpts struct, which allows you to filter and sort
// the returned collection for greater efficiency.
func (r commonResult) ExtractNode() ([]Nodes, error) {
	var s ListNode
	err := r.ExtractInto(&s)
	if err != nil {
		return nil, err
	}
	return s.Nodes, nil
}

// ExtractJob is a function that accepts a result and extracts a job.
func (r commonResult) ExtractJob() (*Job, error) {
	var s Job
	err := r.ExtractInto(&s)
	return &s, err
}

// ListResult represents the result of a list operation. Call its ExtractNode
// method to interpret it as a Nodes.
type ListResult struct {
	commonResult
}

// CreateResult represents the result of a create operation. Call its Extract
// method to interpret it as a Node.
type CreateResult struct {
	commonResult
}

// GetResult represents the result of a get operation. Call its Extract
// method to interpret it as a Node.
type GetResult struct {
	commonResult
}

// UpdateResult represents the result of an update operation. Call its Extract
// method to interpret it as a Node.
type UpdateResult struct {
	commonResult
}

// DeleteResult represents the result of a delete operation. Call its ExtractErr
// method to determine if the request succeeded or failed.
type DeleteResult struct {
	golangsdk.ErrResult
}
