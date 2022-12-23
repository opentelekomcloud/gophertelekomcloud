package configurations

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack"
)

// GetForInstance retrieves Configuration applied to particular RDS instance
// configuration ID and Name will be empty
func GetForInstance(c *golangsdk.ServiceClient, instanceID string) (*Configuration, error) {
	// GET https://{Endpoint}/v3/{project_id}/instances/{instance_id}/configurations
	raw, err := c.Get(c.ServiceURL("instances", instanceID, "configurations"), nil, openstack.StdRequestOpts())
	if err != nil {
		return nil, err
	}

	var res Configuration
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type Parameter struct {
	// Indicates the parameter name.
	Name string `json:"name"`
	// Indicates the parameter value.
	Value string `json:"value"`
	// Indicates whether a reboot is required.
	RestartRequired bool `json:"restart_required"`
	// Indicates whether the parameter is read-only.
	ReadOnly bool `json:"readonly"`
	// Indicates the parameter value range. If the type is Integer, the value is 0 or 1. If the type is Boolean, the value is true or false.
	ValueRange string `json:"value_range"`
	// Indicates the parameter type, which can be integer, string, boolean, list, or float.
	Type string `json:"type"`
	// Indicates the parameter description.
	Description string `json:"description"`
}
