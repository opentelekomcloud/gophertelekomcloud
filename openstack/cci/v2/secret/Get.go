package secret

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func Get(client *golangsdk.ServiceClient, nameSpace, name string) (*SecretResp, error) {
	raw, err := client.Get(client.ServiceURL("namespaces", nameSpace, "secrets", name), nil, &golangsdk.RequestOpts{
		OkCodes:  []int{200},
		JSONBody: nil,
	})
	if err != nil {
		return nil, err
	}

	var res SecretResp
	return &res, extract.Into(raw.Body, &res)
}
