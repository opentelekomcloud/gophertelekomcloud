package groups

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func Get(client *golangsdk.ServiceClient, id string) (*Group, error) {
	raw, err := client.Get(client.ServiceURL("scaling_group", id), nil, nil)
	if err != nil {
		return nil, err
	}

	var res Group
	err = extract.IntoStructPtr(raw.Body, &res, "scaling_group")
	return &res, err
}
