package domains

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type CreateOpts struct {
	// Organization name
	Namespace string `json:"-" required:"true"`
	// Image repository name
	Repository string `json:"-" required:"true"`
	// Name of the account used for image sharing.
	AccessDomain string `json:"access_domain"`
	// Currently, only the read permission is supported.
	Permit string `json:"permit"`
	// Valid until (UTC). If the sharing is permanent, the value is forever. Otherwise, the sharing is valid until 00:00:00 of the next day.
	Deadline string `json:"deadline"`
	// Description
	Description string `json:"description"`
}

func Create(client *golangsdk.ServiceClient, opts CreateOpts) (err error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return
	}

	// POST /v2/manage/namespaces/{namespace}/repositories/{repository}/access-domains
	url := client.ServiceURL("manage", "namespaces", opts.Namespace, "repos", opts.Repository, "access-domains")
	_, err = client.Post(url, b, nil, nil)
	return
}
