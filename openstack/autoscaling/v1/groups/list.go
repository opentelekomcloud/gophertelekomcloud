package groups

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ListOpts struct {
	// Specifies the AS group name.
	// Supports fuzzy search.
	Name string `q:"scaling_group_name"`
	// Specifies the AS configuration ID, which can be obtained using the API for querying AS configurations.
	ConfigurationID string `q:"scaling_configuration_id"`
	// Specifies the AS group status. The options are as follows:
	// INSERVICE: indicates that the AS group is functional.
	// PAUSED: indicates that the AS group is paused.
	// ERROR: indicates that the AS group malfunctions.
	// DELETING: indicates that the AS group is being deleted.
	Status string `q:"scaling_group_status"`
	// Specifies the start line number. The default value is 0. The minimum value is 0, and there is no limit on the maximum value.
	StartNumber int `q:"start_number"`
	// Specifies the number of query records. The default value is 20. The value range is 0 to 100.
	Limit int `q:"limit"`
	// Specifies the enterprise project ID. If all_granted_eps is transferred,
	// this API will query AS groups in the enterprise projects that you have permissions to.
	EnterpriseProjectID string `q:"enterprise_project_id"`
}

func List(client *golangsdk.ServiceClient, opts ListOpts) (*ListScalingGroupsResponse, error) {
	var opts2 interface{} = opts
	q, err := build.QueryString(opts2)
	if err != nil {
		return nil, err
	}

	raw, err := client.Get(client.ServiceURL("scaling_group")+q.String(), nil, nil)
	if err != nil {
		return nil, err
	}

	var res ListScalingGroupsResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type ListScalingGroupsResponse struct {
	TotalNumber   int32   `json:"total_number,omitempty"`
	StartNumber   int32   `json:"start_number,omitempty"`
	Limit         int32   `json:"limit,omitempty"`
	ScalingGroups []Group `json:"scaling_groups,omitempty"`
}
