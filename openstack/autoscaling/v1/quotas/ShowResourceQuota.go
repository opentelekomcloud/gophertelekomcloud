package quotas

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func ShowResourceQuota(client *golangsdk.ServiceClient) (*AllQuotas, error) {
	// GET https://{Endpoint}/autoscaling-api/v1/{project_id}/quotas
	raw, err := client.Get(client.ServiceURL("quotas"), nil, nil)
	if err != nil {
		return nil, err
	}

	var res AllQuotas
	err = extract.IntoStructPtr(raw.Body, &res, "quotas")
	return &res, err
}

type AllQuotas struct {
	Resources []AllResources `json:"resources,omitempty"`
}

type AllResources struct {
	// Specifies the quota type.
	// scaling_Group: AS group quota
	// scaling_Config: AS configuration quota
	// scaling_Policy: AS policy quota
	// scaling_Instance: instance quota
	// bandwidth_scaling_policy: bandwidth scaling policy quota
	Type string `json:"type,omitempty"`
	// Specifies the used amount of the quota.
	// When type is set to scaling_Policy or scaling_Instance,
	// this parameter is reserved, and the system returns -1 as the parameter value.
	// You can query the used quota of AS policies and AS instances in a specified AS group.
	Used int32 `json:"used,omitempty"`
	// Specifies the total quota.
	Quota int32 `json:"quota,omitempty"`
	// Specifies the quota upper limit.
	Max int32 `json:"max,omitempty"`
	// Specifies the quota lower limit.
	Min int32 `json:"min,omitempty"`
}
