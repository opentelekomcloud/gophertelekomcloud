package snapshots

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func Get(client *golangsdk.ServiceClient, id string) (*Snapshot, error) {
	raw, err := client.Get(client.ServiceURL("cloudsnapshots", id), nil, nil)
	if err != nil {
		return nil, err
	}

	var res snapshotResponse
	err = extract.Into(raw.Body, &res)
	return &res.Snapshot, err
}
