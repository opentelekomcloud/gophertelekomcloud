package protocols

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

func List(client *golangsdk.ServiceClient, provider string) pagination.Pager {
	pager := pagination.Pager{
		Client:     client,
		InitialURL: listURL(client, provider),
		CreatePage: func(r pagination.PageResult) pagination.Page {
			return ProtocolPage{LinkedPageBase: pagination.LinkedPageBase{PageResult: r}}
		},
	}
	pager.Headers = map[string]string{
		"Content-Type": "application/json",
	}
	return pager
}

type CreateOptsBuilder interface {
	ToProtocolCreateMap() (map[string]interface{}, error)
}

type CreateOpts struct {
	MappingID string `json:"mapping_id"`
}

func (opts CreateOpts) ToProtocolCreateMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "protocol")
}

func Create(client *golangsdk.ServiceClient, provider, protocol string, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToProtocolCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	url := singleURL(client, provider, protocol)
	_, r.Err = client.Put(url, b, &r.Body, nil)
	return
}

func Get(client *golangsdk.ServiceClient, provider, protocol string) (r GetResult) {
	_, r.Err = client.Get(singleURL(client, provider, protocol), &r.Body, &golangsdk.RequestOpts{
		MoreHeaders: map[string]string{"Content-Type": "application/json"},
	})
	return
}

type UpdateOptsBuilder interface {
	ToProtocolUpdateMap() (map[string]interface{}, error)
}

type UpdateOpts struct {
	MappingID string `json:"mapping_id"`
}

func (opts UpdateOpts) ToProtocolUpdateMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "protocol")
}

func Update(client *golangsdk.ServiceClient, provider, protocol string, opts UpdateOptsBuilder) (r UpdateResult) {
	b, err := opts.ToProtocolUpdateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Patch(singleURL(client, provider, protocol), b, &r.Body, nil)
	return
}

func Delete(client *golangsdk.ServiceClient, provider, protocol string) (r DeleteResult) {
	_, r.Err = client.Delete(singleURL(client, provider, protocol), nil)
	return
}
