package domains

import (
	"fmt"

	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type UpdateOpts struct {
	// Organization name
	Namespace string `json:"-" required:"true"`
	// Image repository name
	Repository string `json:"-" required:"true"`
	// Name of the account used for image sharing
	AccessDomain string `json:"-" required:"true"`
	// Currently, only the read permission is supported.
	Permit string `json:"permit"`
	// Valid until (UTC). If the sharing is permanent, the value is forever. Otherwise, the sharing is valid until 00:00:00 of the next day.
	Deadline string `json:"deadline"`
	// Description. This parameter is left blank by default.
	Description string `json:"description,omitempty"`
}

func Update(client *golangsdk.ServiceClient, opts UpdateOpts) (err error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return
	}

	// PATCH /v2/manage/namespaces/{namespace}/repositories/{repository}/access-domains/{access_domain}
	url := fmt.Sprintf("%s/%s", client.ServiceURL("manage", "namespaces", opts.Namespace, "repos", opts.Repository, "access-domains"), opts.AccessDomain)
	_, err = client.Patch(url, b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{201},
	})
	return
}
