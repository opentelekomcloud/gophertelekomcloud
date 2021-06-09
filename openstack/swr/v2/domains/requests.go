package domains

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type CreateOptsBuilder interface {
	ToAccessDomainCreateMap() (map[string]interface{}, error)
}

type CreateOpts struct {
	AccessDomain string `json:"access_domain"`
	// Currently, only the `read` permission is supported.
	Permit string `json:"permit"`
	// End date of image sharing (UTC). When the value is set to `forever`,
	// the image will be permanently available for the domain.
	// The validity period is calculated by day. The shared images expire at 00:00:00 on the day after the end date.
	Deadline    string `json:"deadline"`
	Description string `json:"description"`
}

func (opts CreateOpts) ToAccessDomainCreateMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "")
}

func Create(client *golangsdk.ServiceClient, org, repo string, opts CreateOptsBuilder) (r CreateResult) {
	url := listURL(client, org, repo)
	b, err := opts.ToAccessDomainCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(url, b, &r.Body, nil)
	return
}

func Delete(client *golangsdk.ServiceClient, org, repo, domain string) (r DeleteResult) {
	_, r.Err = client.Delete(singleURL(client, org, repo, domain), nil)
	return
}

func Get(client *golangsdk.ServiceClient, org, repo, domain string) (r GetResult) {
	_, r.Err = client.Get(singleURL(client, org, repo, domain), &r.Body, nil)
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

	_, r.Err = client.Patch(singleURL(client, org, repo, domain), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{201},
	})
	return
}

func List(client *golangsdk.ServiceClient, org, repo string) (p pagination.Pager) {
	return pagination.NewPager(client, listURL(client, org, repo), func(r pagination.PageResult) pagination.Page {
		return AccessDomainPage{SinglePageBase: pagination.SinglePageBase(r)}
	})
}
