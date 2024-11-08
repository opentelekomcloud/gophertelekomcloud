package lifecycle

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// Get an instance with detailed information by id
func Get(client *golangsdk.ServiceClient, id string) (*Instance, error) {
	raw, err := client.Get(client.ServiceURL("instances", id), nil, nil)
	if err != nil {
		return nil, err
	}

	var res Instance
	err = extract.Into(raw.Body, &res)
	return &res, err
}
