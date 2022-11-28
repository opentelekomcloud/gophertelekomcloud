package configurations

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
)

type ConfigurationCreate struct {
	// Configuration ID
	ID string `json:"id"`
	// Configuration Name
	Name string `json:"name"`
	// Database version Name
	DatastoreVersionName string `json:"datastore_version_name"`
	// Database Name
	DatastoreName string `json:"datastore_name"`
	// Configuration Description
	Description string `json:"description"`
	// Indicates the creation time in the following format: yyyy-MM-ddTHH:mm:ssZ.
	Created string `json:"created"`
	// Indicates the update time in the following format: yyyy-MM-ddTHH:mm:ssZ.
	Updated string `json:"updated"`
}

func (r CreateResult) ExtractInto(v interface{}) error {
	return r.Result.ExtractIntoStructPtr(v, "configuration")
}

// Extract is a function that accepts a result and extracts a configuration.
func (r CreateResult) Extract() (*ConfigurationCreate, error) {
	var response ConfigurationCreate
	err := r.ExtractInto(&response)
	return &response, err
}

// CreateResult represents the result of a create operation. Call its Extract
// method to interpret it as a Configuration.
type CreateResult struct {
	golangsdk.Result
}

// UpdateResult represents the result of a update operation.
type UpdateResult struct {
	golangsdk.ErrResult
}

// Extract is a function that accepts a result and extracts a configuration.
func (r GetResult) Extract() (*Configuration, error) {
	var response Configuration
	err := r.ExtractInto(&response)
	return &response, err
}

// GetResult represents the result of a get operation. Call its Extract
// method to interpret it as a Configuration.
type GetResult struct {
	golangsdk.Result
}

// DeleteResult represents the result of a delete operation. Call its ExtractErr
// method to determine if the request succeeded or failed.
type DeleteResult struct {
	golangsdk.ErrResult
}
