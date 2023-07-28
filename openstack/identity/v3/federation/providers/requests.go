package providers

import (
	"fmt"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type CreateOptsBuilder interface {
	Id() string
	ToProviderCreateMap() (map[string]interface{}, error)
}

type CreateOpts struct {
	ID          string `json:"-"`
	Description string `json:"description,omitempty"`
	Enabled     bool   `json:"enabled,omitempty"`
}

func (opts CreateOpts) ToProviderCreateMap() (map[string]interface{}, error) {
	return build.RequestBodyMap(opts, "identity_provider")
}

func (opts CreateOpts) Id() string {
	return opts.ID
}

func Create(client *golangsdk.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	if opts.Id() == "" {
		r.Err = fmt.Errorf("missing identity provider ID")
		return
	}
	b, err := opts.ToProviderCreateMap()
	if err != nil {
		r.Err = fmt.Errorf("error building provider create body: %s", err)
	}
	_, r.Err = client.Put(providerURL(client, opts.Id()), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{201},
	})
	return
}

func Get(client *golangsdk.ServiceClient, id string) (r GetResult) {
	_, r.Err = client.Get(providerURL(client, id), &r.Body, nil)
	return
}

func List(client *golangsdk.ServiceClient) pagination.Pager {
	return pagination.Pager{
		Client:     client,
		InitialURL: listURL(client),
		CreatePage: func(r pagination.PageResult) pagination.Page {
			return ProviderPage{pagination.LinkedPageBase{PageResult: r}}
		},
	}
}

type UpdateOptsBuilder interface {
	ToUpdateOptsMap() (map[string]interface{}, error)
}

type UpdateOpts struct {
	Description string `json:"description,omitempty"`
	Enabled     *bool  `json:"enabled,omitempty"`
}

func (opts UpdateOpts) ToUpdateOptsMap() (map[string]interface{}, error) {
	return build.RequestBodyMap(opts, "identity_provider")
}

func Update(client *golangsdk.ServiceClient, id string, opts UpdateOptsBuilder) (r UpdateResult) {
	b, err := opts.ToUpdateOptsMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Patch(providerURL(client, id), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

func Delete(client *golangsdk.ServiceClient, id string) (r DeleteResult) {
	_, r.Err = client.Delete(providerURL(client, id), nil)
	return
}
