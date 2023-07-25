package protectedinstances

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type Instance struct {
	// Instance ID
	ID string `json:"id"`
	// Instance Name
	Name string `json:"name"`
	// Instance Description
	Description string `json:"description"`
	// Protection Group ID
	GroupID string `json:"server_group_id"`
	// Instance Status
	Status string `json:"status"`
	// Instance Progress
	Progress int `json:"progress"`
	// Source Server
	SourceServer string `json:"source_server"`
	// Target Server
	TargetServer string `json:"target_server"`
	// Instance CreatedAt time
	CreatedAt string `json:"created_at"`
	// Instance UpdatedAt time
	UpdatedAt string `json:"updated_at"`
	// Production site AZ of the protection group containing the protected instance.
	PriorityStation string `json:"priority_station"`
	// Attachment
	Attachment []Attachment `json:"attachment"`
	// Tags list
	Tags []Tags `json:"tags"`
	// Metadata
	Metadata map[string]string `json:"metadata"`
}

type Attachment struct {
	// Replication ID
	Replication string `json:"replication"`
	// Device Name
	Device string `json:"device"`
}

type commonResult struct {
	golangsdk.Result
}

// UpdateResult represents the result of a update operation. Call its Extract
// method to interpret it as a Instance.
type UpdateResult struct {
	commonResult
}

// Extract is a function that accepts a result and extracts a instance.
func (r commonResult) Extract() (*Instance, error) {
	response := new(Instance)
	err := r.ExtractIntoStructPtr(response, "protected_instance")
	return response, err
}

// GetResult represents the result of a get operation. Call its Extract
// method to interpret it as a Instance.
type GetResult struct {
	commonResult
}

// InstancePage is a struct which can do the page function
type InstancePage struct {
	pagination.SinglePageBase
}

// IsEmpty determines whether or not a InstancePage is empty.
func (r InstancePage) IsEmpty() (bool, error) {
	instances, err := ExtractInstances(r)
	return len(instances) == 0, err
}

// ExtractInstances interprets the results of a single page from
// a List() API call, producing a slice of []Instance structures.
func ExtractInstances(r pagination.Page) ([]Instance, error) {
	var s []Instance

	err := extract.IntoSlicePtr((r.(InstancePage)).Body, &s, "protected_instances")
	return s, err
}
