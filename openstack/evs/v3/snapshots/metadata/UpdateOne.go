package metadata

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

func UpdateOne(client *golangsdk.ServiceClient, volumeId string, key string, opts map[string]string) (map[string]string, error) {
	b, err := build.RequestBody(opts, "meta")
	if err != nil {
		return nil, err
	}

	// PUT /v3/{project_id}/volumes/{volume_id}/metadata/{key}
	raw, err := client.Put(client.ServiceURL("volumes", volumeId, "metadata", key), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return extraMeta(err, raw)
}