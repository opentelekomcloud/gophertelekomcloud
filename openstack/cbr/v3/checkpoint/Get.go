package checkpoint

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func Get(client *golangsdk.ServiceClient, id string) (*Checkpoint, error) {
	raw, err := client.Get(client.ServiceURL("checkpoints", id), nil, nil)
	if err != nil {
		return nil, err
	}

	var res Checkpoint
	err = extract.IntoStructPtr(raw.Body, &res, "checkpoint")
	return &res, err
}
