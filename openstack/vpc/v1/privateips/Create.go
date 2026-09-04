package privateips

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type CreateOpts struct {
	PrivateIPs []PrivateIPRequest `json:"privateips" required:"true"`
}

type PrivateIPRequest struct {
	SubnetID  string `json:"subnet_id" required:"true"`
	IPAddress string `json:"ip_address,omitempty"`
}

func Create(client *golangsdk.ServiceClient, opts CreateOpts) ([]PrivateIP, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL(client.ProjectID, "privateips"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res struct {
		PrivateIPs []PrivateIP `json:"privateips"`
	}
	err = extract.Into(raw.Body, &res)
	return res.PrivateIPs, err
}
