package snapshots

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type RollbackOpts struct {
	VolumeID string `json:"volume_id" required:"true"`
	Name     string `json:"name,omitempty"`
}

func Rollback(client *golangsdk.ServiceClient, snapshotID string, opts RollbackOpts) (*RollbackInfo, error) {
	b, err := build.RequestBody(opts, "rollback")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL("cloudsnapshots", snapshotID, "rollback"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{202},
	})
	if err != nil {
		return nil, err
	}

	var res RollbackResponse
	err = extract.Into(raw.Body, &res)
	return &res.Rollback, err
}

type RollbackResponse struct {
	Rollback RollbackInfo `json:"rollback"`
}

type RollbackInfo struct {
	VolumeID string `json:"volume_id"`
}
