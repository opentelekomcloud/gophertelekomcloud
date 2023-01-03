package repositories

import (
	"fmt"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

func Delete(client *golangsdk.ServiceClient, organization, repository string) (r DeleteResult) {
	_, r.Err = client.Delete(client.ServiceURL("manage", "namespaces", organization, "repos", repository), nil)
	return
}

type ListOptsBuilder interface {
	ToRepositoryListQuery() (string, error)
}

type ListOpts struct {
	// Organization (namespace) name
	Organization string `q:"namespace"`
	// Image repository name
	Name string `q:"name"`
	// Image repository category.
	Category string `q:"category"`

	// Sorting by column.
	// You can set this parameter to `name`, `updated_time`, and `tag_count`.
	// The parameters OrderColumn and OrderType should always be used together.
	OrderColumn string `q:"order_column"`
	// Sorting type.
	// You can set this parameter to `desc` (descending sort) and `asc` (ascending sort).
	OrderType string `q:"order_type"`

	Offset *int `q:"offset,omitempty"` // offset 0 is a valid value
	Limit  int  `q:"limit,omitempty"`
}

const defaultLimit = 25

func (opts ListOpts) ToRepositoryListQuery() (string, error) {
	if opts.Limit == 0 && opts.Offset != nil {
		opts.Limit = defaultLimit
	}
	if opts.Limit != 0 && opts.Offset == nil {
		return "", fmt.Errorf("offset has to be defined if the limit is set")
	}
	if (opts.OrderColumn != "" && opts.OrderType == "") || (opts.OrderColumn == "" && opts.OrderType != "") {
		return "", fmt.Errorf("`OrderColumn` and `OrderType` should always be used together")
	}
	q, err := golangsdk.BuildQueryString(opts)
	if err != nil {
		return "", err
	}
	return q.String(), nil
}

func List(client *golangsdk.ServiceClient, opts ListOptsBuilder) (p pagination.Pager) {
	url := client.ServiceURL("manage", "repos")
	if opts != nil {
		q, err := opts.ToRepositoryListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += q
	}

	return pagination.NewPager(client, url, func(r pagination.PageResult) pagination.Page {
		return RepositoryPage{pagination.OffsetPageBase{PageResult: r}}
	})
}

func Get(client *golangsdk.ServiceClient, organization, repository string) (r GetResult) {
	_, r.Err = client.Get(client.ServiceURL("manage", "namespaces", organization, "repos", repository), &r.Body, nil)
	return
}

type UpdateOptsBuilder interface {
	ToRepositoryUpdateMap() (map[string]interface{}, error)
}

type UpdateOpts struct {
	Category    string `json:"category,omitempty"`
	Description string `json:"description,omitempty"`
	IsPublic    bool   `json:"is_public"` // this is mandatory field, so no pointers
}

func (opts UpdateOpts) ToRepositoryUpdateMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "")
}

func Update(client *golangsdk.ServiceClient, organization, repository string, opts UpdateOptsBuilder) (r UpdateResult) {
	b, err := opts.ToRepositoryUpdateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Patch(client.ServiceURL("manage", "namespaces", organization, "repos", repository), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{201},
	})
	return
}
