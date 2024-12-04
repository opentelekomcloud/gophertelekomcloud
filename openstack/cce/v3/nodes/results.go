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

// Job Structure
type Job struct {
	// API type, fixed value "Job"
	Kind string `json:"kind"`
	// API version, fixed value "v3"
	Apiversion string `json:"apiVersion"`
	// Node metadata
	Metadata JobMetadata `json:"metadata"`
	// Node detailed parameters
	Spec JobSpec `json:"spec"`
	// Node status information
	Status JobStatus `json:"status"`
}

type JobMetadata struct {
	// ID of the job
	ID string `json:"uid"`
}

type JobSpec struct {
	// Type of job
	Type string `json:"type"`
	// ID of the cluster where the job is located
	ClusterID string `json:"clusterUID"`
	// ID of the IaaS resource for the job operation
	ResourceID string `json:"resourceID"`
	// The name of the IaaS resource for the job operation
	ResourceName string `json:"resourceName"`
	// List of child jobs
	SubJobs []Job `json:"subJobs"`
	// ID of the parent job
	OwnerJob string `json:"ownerJob"`
}

type JobStatus struct {
	// Job status
	Phase string `json:"phase"`
	// The reason why the job becomes the current state
	Reason string `json:"reason"`
	// The job becomes the current state details
	Message string `json:"message"`
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
