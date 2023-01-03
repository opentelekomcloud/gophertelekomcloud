package domains

import (
	"fmt"

	"github.com/opentelekomcloud/gophertelekomcloud"
)

type GetOpts struct {
	// Organization name
	Namespace string `json:"-" required:"true"`
	// Image repository name
	Repository string `json:"-" required:"true"`
	// Name of the account used for image sharing.
	AccessDomain string `json:"access_domain"`
}

func Delete(client *golangsdk.ServiceClient, opts GetOpts) (err error) {
	// DELETE /v2/manage/namespaces/{namespace}/repositories/{repository}/access-domains/{access_domain}
	url := fmt.Sprintf("%s/%s", client.ServiceURL("manage", "namespaces", opts.Namespace, "repos", opts.Repository, "access-domains"), opts.AccessDomain)
	_, err = client.Delete(url, nil)
	return
}
