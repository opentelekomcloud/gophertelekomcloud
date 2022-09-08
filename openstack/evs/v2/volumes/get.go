package volumes

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func Get(client *golangsdk.ServiceClient, id string) (*Volume, error) {
	raw, err := client.Get(client.ServiceURL("volumes", id), nil, nil)
	if err != nil {
		return nil, err
	}

	var res Volume
	err = extract.Into(raw.Body, &res)
	return &res, err
}
