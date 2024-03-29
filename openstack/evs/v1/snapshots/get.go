package snapshots

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func Get(client *golangsdk.ServiceClient, id string) (*Snapshot, error) {
	raw, err := client.Get(client.ServiceURL("snapshots", id), nil, nil)
	if err != nil {
		return nil, err
	}

	var res Snapshot
	err = extract.IntoStructPtr(raw.Body, &res, "snapshot")
	return &res, err
}
