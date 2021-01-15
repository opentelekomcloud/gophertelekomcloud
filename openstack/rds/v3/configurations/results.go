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

type Configuration struct {
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
	// Configuration Parameters
	Parameters []Parameter `json:"configuration_parameters"`
}

type Parameter struct {
	// Parameter Name
	Name string `json:"name"`
	// Parameter value
	Value string `json:"value"`
	// Whether a restart is required
	RestartRequired bool `json:"restart_required"`
	// Whether the parameter is read-only
	ReadOnly bool `json:"readonly"`
	// Parameter value range
	ValueRange string `json:"value_range"`
	// Parameter type
	Type string `json:"type"`
	// Parameter description
	Description string `json:"description"`
}

type ApplyConfiguration struct {
	// Specifies the parameter template ID.
	ConfigurationID string `json:"configuration_id"`
	// Specifies the parameter template name.
	ConfigurationName string `json:"configuration_name"`
	// Specifies the result of applying the parameter template.
	ApplyResults []ApplyConfigurationResult `json:"apply_results"`
	// Specifies whether each parameter template is applied to DB instances successfully.
	Success bool `json:"success"`
}

type ApplyConfigurationResult struct {
	// Indicates the DB instance ID.
	InstanceID string `json:"instance_id"`
	// Indicates the DB instance name.
	InstanceName string `json:"instance_name"`
	// Indicates whether a reboot is required.
	RestartRequired bool `json:"restart_required"`
	// Indicates whether each parameter template is applied to DB instances successfully.
	Success bool `json:"success"`
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

// Extract is a function that accepts a result and extracts a list of configurations.
func (r ListResult) Extract() ([]Configuration, error) {
	var a struct {
		Configurations []Configuration `json:"configurations"`
	}
	err := r.Result.ExtractInto(&a)
	return a.Configurations, err
}

// ListResult represents the result of a list operation. Call its Extract
// method to interpret it as a list of Configurations.
type ListResult struct {
	golangsdk.Result
}

// Extract is a function that accepts a result and extracts an apply configuration result.
func (r ApplyResult) Extract() (*ApplyConfiguration, error) {
	var response ApplyConfiguration
	err := r.ExtractInto(&response)
	return &response, err
}

// ApplyResult represents the result of a apply operation. Call its Extract
// method to interpret it.
type ApplyResult struct {
	golangsdk.Result
}
