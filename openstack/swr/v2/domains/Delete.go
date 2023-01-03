package domains

import (
	"fmt"

	"github.com/opentelekomcloud/gophertelekomcloud"
)

func Delete(client *golangsdk.ServiceClient, org, repo, domain string) (err error) {
	// DELETE /v2/manage/namespaces/{namespace}/repositories/{repository}/access-domains/{access_domain}
	_, err = client.Delete(fmt.Sprintf("%s/%s", client.ServiceURL("manage", "namespaces", org, "repos", repo, "access-domains"), domain), nil)
	return
}
