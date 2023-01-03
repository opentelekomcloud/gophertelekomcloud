package domains

import (
	"fmt"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

func Get(client *golangsdk.ServiceClient, org, repo, domain string) (r GetResult) {
	_, r.Err = client.Get(fmt.Sprintf("%s/%s", client.ServiceURL("manage", "namespaces", org, "repos", repo, "access-domains"), domain), &r.Body, nil)
	return
}

func List(client *golangsdk.ServiceClient, org, repo string) (p pagination.Pager) {
	return pagination.NewPager(client, client.ServiceURL("manage", "namespaces", org, "repos", repo, "access-domains"), func(r pagination.PageResult) pagination.Page {
		return AccessDomainPage{SinglePageBase: pagination.SinglePageBase(r)}
	})
}
