package security_policy

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func Get(client *golangsdk.ServiceClient, id string) (*SecurityPolicy, error) {
	raw, err := client.Get(client.ServiceURL("security-policies", id), nil, &golangsdk.RequestOpts{OkCodes: []int{200}})
	if err != nil {
		return nil, err
	}

	var res SecurityPolicy
	err = extract.Into(raw.Body, &res)
	return &res, err
}
