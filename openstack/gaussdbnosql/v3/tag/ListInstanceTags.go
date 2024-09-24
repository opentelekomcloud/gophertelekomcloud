package tag

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func ListInstanceTags(client *golangsdk.ServiceClient, instanceId string) ([]InstanceTagResult, error) {
	// GET https://{Endpoint}/v3/{project_id}/instances/{instance_id}/tags
	raw, err := client.Get(client.ServiceURL("instances", instanceId, "tags"), nil, nil)
	if err != nil {
		return nil, err
	}

	var res []InstanceTagResult
	err = extract.IntoSlicePtr(raw.Body, &res, "tags")
	return res, err
}
