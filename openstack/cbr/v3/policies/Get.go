package policies

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func Get(client *golangsdk.ServiceClient, id string) (*Policy, error) {
	raw, err := client.Get(client.ServiceURL("policies", id), nil, nil)
	if err != nil {
		return nil, err
	}

	var res Policy
	err = extract.IntoStructPtr(raw.Body, &res, "policy")
	return &res, err
}
