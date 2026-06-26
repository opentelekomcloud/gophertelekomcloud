package projects

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type ListOptsBuilder interface {
	ToProjectListQuery() (string, error)
}

type ListOpts struct {
	ID      string `q:"id"`
	Limit   int    `q:"limit"`
	Name    string `q:"name"`
	Offset  int    `q:"offset"`
	SortDir string `q:"sort_dir"`
	SortKey string `q:"sort_key"`
	Status  int    `q:"status"`
}

func (opts ListOpts) ToProjectListQuery() (string, error) {
	q, err := golangsdk.BuildQueryString(opts)
	if err != nil {
		return "", err
	}
	return q.String(), nil
}

func List(client *golangsdk.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := rootURL(client)
	if opts != nil {
		query, err := opts.ToProjectListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}

	return pagination.NewPager(client, url, func(r pagination.PageResult) pagination.Page {
		return ProjectPage{pagination.SinglePageBase(r)}
	})
}

type CreateOptsBuilder interface {
	ToProjectCreateMap() (map[string]interface{}, error)
}

type CreateOpts struct {
	Name        string `json:"name" required:"true"`
	Description string `json:"description,omitempty"`
}

func (opts CreateOpts) ToProjectCreateMap() (map[string]interface{}, error) {
	b, err := golangsdk.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}
	return b, nil
}

func Create(client *golangsdk.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToProjectCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(rootURL(client), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200, 201},
	})
	return
}

func Get(client *golangsdk.ServiceClient, id string) (r GetResult) {
	_, r.Err = client.Get(resourceURL(client, id), &r.Body, nil)
	return
}

type UpdateOptsBuilder interface {
	ToProjectUpdateMap() (map[string]interface{}, error)
}

type UpdateOpts struct {
	Name        string `json:"name" required:"true"`
	Description string `json:"description,omitempty"`
	Type        string `json:"type,omitempty"`
}

func (opts UpdateOpts) ToProjectUpdateMap() (map[string]interface{}, error) {
	b, err := golangsdk.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}
	return b, nil
}

func Update(client *golangsdk.ServiceClient, id string, opts UpdateOptsBuilder) (r UpdateResult) {
	b, err := opts.ToProjectUpdateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Put(resourceURL(client, id), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

type ActionOptsBuilder interface {
	ToProjectActionMap() (map[string]interface{}, error)
}

type ActionOpts struct {
	Action string `json:"action" required:"true"`
}

func (opts ActionOpts) ToProjectActionMap() (map[string]interface{}, error) {
	b, err := golangsdk.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}
	return b, nil
}

type ActionResult struct {
	golangsdk.ErrResult
}

func Action(client *golangsdk.ServiceClient, id string, opts ActionOptsBuilder) (r ActionResult) {
	b, err := opts.ToProjectActionMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(resourceURL(client, id)+"/action", b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200, 204},
	})
	return
}
