package link

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type CreateOpts struct {
	Links []Link `json:"links" required:"true"`
}

type CreateQuery struct {
	// When the parameter is set to true, the API only validates whether the parameters are correctly configured, but does not create any link.
	Validate bool `json:"validate,omitempty"`
}

type Link struct {
	// User who created the link
	LinkConfigValues *ConfigValues `json:"link-config-values" required:"true"`

	// User who created the link.
	CreationUser string `json:"creation-user,omitempty"`

	// Link name.
	Name string `json:"name" required:"true"`

	// Link ID.
	ID int `json:"id,omitempty"`

	// Time when the link was created.
	CreationDate int64 `json:"creation-date,omitempty"`

	// Connector name associated with the link.
	ConnectorName string `json:"connector-name" required:"true"`

	// Time when the link was updated.
	UpdateDate int64 `json:"update-date,omitempty"`

	// Whether to activate the link.
	Enabled bool `json:"enabled"`

	// Update user who updated the link.
	UpdateUser string `json:"update_user,omitempty"`
}

type ConfigValues struct {
	// Array of configuration objects
	Configs []Config `json:"configs" required:"true"`

	// Extended configurations (optional)
	ExtendedConfigs *ExtendedConfigs `json:"extended-configs,omitempty"`

	// Validators (optional)
	Validators []string `json:"validators,omitempty"`
}

type Config struct {
	// Inputs is a list of Input. Each element in the list is in name,value format. For details, see the descriptions of inputs parameters.
	// In the from-config-values data structure, the value of this parameter varies with the source link type.
	// For details, see section "Source Job Parameters" in the Cloud Data Migration User Guide.
	// In the to-cofig-values data structure, the value of this parameter varies with the destination link type.
	// For details, see section "Destination Job Parameters" in the Cloud Data Migration User Guide.
	// For details about the inputs parameter in the driver-config-values data structure, see the job parameter descriptions.
	Inputs []*Input `json:"inputs" required:"true"`

	// Name is a configuration name. The value is fromJobConfig for a source job, toJobConfig for a destination job, and linkConfig for a link.
	Name string `json:"name" required:"true"`

	// Configuration ID (might not be returned)
	ID int `json:"id,omitempty"`

	// Configuration type (might not be returned)
	Type string `json:"type,omitempty"`
}

type Input struct {
	// Parameter name
	Name string `json:"name" required:"true"`

	// Parameter value
	Value string `json:"value" required:"true"`

	// Value type (might not be returned)
	Type string `json:"type,omitempty"`
}

type ExtendedConfigs struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

// Create is used to create a link.
// Send request POST /v1.1/{project_id}/clusters/{cluster_id}/cdm/link
func Create(client *golangsdk.ServiceClient, clusterId string, opts CreateOpts, q *CreateQuery) (*CreateResp, error) {
	url, err := golangsdk.NewURLBuilder().WithEndpoints(clustersEndpoint, clusterId, cdmEndpoint, linkEndpoint).WithQueryParams(q).Build()
	if err != nil {
		return nil, err
	}

	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL(url.String()), b, nil, &golangsdk.RequestOpts{OkCodes: []int{200}})
	if err != nil {
		return nil, err
	}

	var res CreateResp
	err = extract.Into(raw.Body, &res)
	return &res, err
}

// CreateResp holds job running information
type CreateResp struct {
	// Name is a link name.
	Name string `json:"name"`
	// ValidationResult is an array of ValidationResult objects. If a link fails to be created, the failure cause is returned.
	// If a link is successfully created, an empty list is returned.
	ValidationResult []LinkValidationResult `json:"validation-result"`
}

type LinkValidationResult struct {
	// LinkConfig is a validation result of link creation or update.
	LinkConfig []*ValidationLinkConfig `json:"linkConfig,omitempty"`
}

type ValidationLinkConfig struct {
	// Message is an error message.
	Message string `json:"message,omitempty"`
	// Status can be ERROR,WARNING.
	Status string `json:"status,omitempty"`
}
