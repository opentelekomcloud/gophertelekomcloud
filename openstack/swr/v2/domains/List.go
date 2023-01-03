package domains

import (
	"fmt"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

func List(client *golangsdk.ServiceClient, org, repo string) (p pagination.Pager) {
	// GET /v2/manage/namespaces/{namespace}/repositories/{repository}/access-domains/{access_domain}
	url := fmt.Sprintf("%s/%s", client.ServiceURL("manage", "namespaces", opts.Namespace, "repos", opts.Repository, "access-domains"), opts.AccessDomain)
	raw, err := client.Get(url, nil, nil)
}
