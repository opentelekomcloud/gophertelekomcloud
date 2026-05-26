package snapshots

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type UpdateOpts struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

func Update(client *golangsdk.ServiceClient, id string, opts UpdateOpts) (*Snapshot, error) {
	b, err := build.RequestBody(opts, "snapshot")
	if err != nil {
		return nil, err
	}

	raw, err := client.Put(client.ServiceURL("cloudsnapshots", id), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res snapshotResponse
	err = extract.Into(raw.Body, &res)
	return &res.Snapshot, err
}
