package quotas

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func ShowPolicyAndInstanceQuota(client *golangsdk.ServiceClient, scalingGroupId string) (*AllQuotas, error) {
	// GET /autoscaling-api/v1/{project_id}/quotas/{scaling_group_id}
	raw, err := client.Get(client.ServiceURL("quotas", scalingGroupId), nil, nil)
	if err != nil {
		return nil, err
	}

	var res AllQuotas
	err = extract.IntoStructPtr(raw.Body, &res, "quotas")
	return &res, err
}
