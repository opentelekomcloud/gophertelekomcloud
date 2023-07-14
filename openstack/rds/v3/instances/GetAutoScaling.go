package instances

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func GetAutoScaling(client *golangsdk.ServiceClient, id string) (*ScalingOpts, error) {
	// PUT https://{Endpoint}/v3/{project_id}/instances/{instance_id}/disk-auto-expansion
	raw, err := client.Get(client.ServiceURL("instances", id, "disk-auto-expansion"), nil, nil)
	if err != nil {
		return nil, err
	}
	var res ScalingOpts
	err = extract.IntoStructPtr(raw.Body, &res, "backup")
	return &res, err
}
