package instances

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ModifyPrivateDomainNameOpts struct {
	InstanceId string `json:"-"`
	DnsName    string `json:"dns_name" required:"true"`
}

func ModifyPrivateDomainName(client *golangsdk.ServiceClient, opts ModifyPrivateDomainNameOpts) (*string, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// PUT /v3/{project_id}/instances/{instance_id}/modify-dns
	raw, err := client.Put(client.ServiceURL("instances", opts.InstanceId, "modify-dns"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{202},
	})
	if err != nil {
		return nil, err
	}

	var res JobId
	err = extract.Into(raw.Body, &res)
	return &res.JobId, err
}
