package metadata

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

func Update(client *golangsdk.ServiceClient, snapshotId string, opts map[string]string) (map[string]string, error) {
	b, err := build.RequestBody(opts, "metadata")
	if err != nil {
		return nil, err
	}

	// PUT /v3/{project_id}/snapshots/{snapshot_id}/metadata
	raw, err := client.Put(client.ServiceURL("snapshots", snapshotId, "metadata"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return extraMetadata(err, raw)
}
