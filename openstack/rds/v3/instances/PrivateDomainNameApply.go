package instances

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ApplyForPrivateDomainOpts struct {
	InstanceId string `json:"-"`
	DnsType    string `json:"dns_type" required:"true"`
}

// This function is used to bind a private domain name to a specified DB instance.
func ApplyForPrivateDomain(client *golangsdk.ServiceClient, opts ApplyForPrivateDomainOpts) (*string, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// POST /v3/{project_id}/instances/{instance_id}/create-dns
	raw, err := client.Post(client.ServiceURL("instances", opts.InstanceId, "create-dns"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200, 202},
	})
	if err != nil {
		return nil, err
	}

	var res JobId
	err = extract.Into(raw.Body, &res)
	return &res.JobId, err
}
