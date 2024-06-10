package instance

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func Get(client *golangsdk.ServiceClient, instanceID string) (*RouterInstanceResp, error) {
	raw, err := client.Get(client.ServiceURL("enterprise-router", "instances", instanceID), nil, nil)
	if err != nil {
		return nil, err
	}

	var res RouterInstanceResp
	err = extract.Into(raw.Body, &res)
	return &res, err
}
