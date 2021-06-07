package organizations

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type CreateOptsBuilder interface {
	ToNamespaceCreateMap() (map[string]interface{}, error)
}

type CreateOpts struct {
	Namespace string `json:"namespace"`
}

func (opts CreateOpts) ToNamespaceCreateMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "")
}

func Create(client *golangsdk.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToNamespaceCreateMap()
	if err != nil {
		r.Err = err
		return
	}

	_, r.Err = client.Post(listURL(client), &b, r.Body, nil)
	return
}

func Delete(client *golangsdk.ServiceClient, id string) (r DeleteResult) {
	_, r.Err = client.Delete(organizationURL(client, id), nil)
	return
}

type ListOptsBuilder interface {
	ToNamespaceListQuery() (string, error)
}

type ListOpts struct {
	Namespace string `q:"namespace"`
}

func (opts ListOpts) ToNamespaceListQuery() (string, error) {
	q, err := golangsdk.BuildQueryString(opts)
	if err != nil {
		return "", err
	}
	return q.String(), err
}

func List(client *golangsdk.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := listURL(client)
	if opts != nil {
		q, err := opts.ToNamespaceListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += q
	}
	return pagination.NewPager(client, url, func(r pagination.PageResult) pagination.Page {
		return OrganizationPage{pagination.SinglePageBase(r)}
	})
}

func Get(client *golangsdk.ServiceClient, id string) (r GetResult) {
	_, r.Err = client.Get(organizationURL(client, id), &r.Body, nil)
	return
}
