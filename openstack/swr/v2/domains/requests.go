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

type UpdateOptsBuilder interface {
	ToAccessDomainUpdateMap() (map[string]interface{}, error)
}

// UpdateOpts used for update operations
// For argument details see CreateOpts
type UpdateOpts struct {
	Permit      string `json:"permit"`
	Deadline    string `json:"deadline"`
	Description string `json:"description,omitempty"`
}

func (opts UpdateOpts) ToAccessDomainUpdateMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "")
}

func Update(client *golangsdk.ServiceClient, org, repo, domain string, opts UpdateOptsBuilder) (r UpdateResult) {
	b, err := opts.ToAccessDomainUpdateMap()
	if err != nil {
		r.Err = err
		return
	}

	_, r.Err = client.Patch(fmt.Sprintf("%s/%s", client.ServiceURL("manage", "namespaces", org, "repos", repo, "access-domains"), domain), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{201},
	})
	return
}

func List(client *golangsdk.ServiceClient, org, repo string) (p pagination.Pager) {
	return pagination.NewPager(client, client.ServiceURL("manage", "namespaces", org, "repos", repo, "access-domains"), func(r pagination.PageResult) pagination.Page {
		return AccessDomainPage{SinglePageBase: pagination.SinglePageBase(r)}
	})
}
