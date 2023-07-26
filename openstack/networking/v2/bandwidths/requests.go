package bandwidths

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type CreateOptsBuilder interface {
	ToBandwidthCreateMap() (map[string]any, error)
}

type CreateOpts struct {
	Name string `json:"name" required:"true"`
	Size int    `json:"size" required:"true"`
}

func (opts CreateOpts) ToBandwidthCreateMap() (map[string]any, error) {
	return golangsdk.BuildRequestBody(opts, "bandwidth")
}

func Create(client *golangsdk.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToBandwidthCreateMap()
	if err != nil {
		r.Err = err
		return
	}

	_, r.Err = client.Post(rootURL(client), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200, 201},
	})
	return
}

type UpdateOptsBuilder interface {
	ToBandwidthUpdateMap() (map[string]any, error)
}

type UpdateOpts struct {
	Name string `json:"name,omitempty"`
	Size int    `json:"size,omitempty"`
}

func (opts UpdateOpts) ToBandwidthUpdateMap() (map[string]any, error) {
	return golangsdk.BuildRequestBody(opts, "bandwidth")
}

func Update(client *golangsdk.ServiceClient, bandwidthID string, opts UpdateOptsBuilder) (r CreateResult) {
	b, err := opts.ToBandwidthUpdateMap()
	if err != nil {
		r.Err = err
		return
	}

	_, r.Err = client.Put(resourceURL(client, bandwidthID), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

func Get(client *golangsdk.ServiceClient, bandwidthID string) (r GetResult) {
	_, r.Err = client.Get(resourceURL(client, bandwidthID), &r.Body, nil)
	return
}

func List(client *golangsdk.ServiceClient) pagination.Pager {
	return pagination.Pager{
		Client:     client,
		InitialURL: rootURL(client),
		CreatePage: func(r pagination.PageResult) pagination.Page {
			return BandwidthPage{SinglePageBase: pagination.SinglePageBase{PageResult: r}}
		},
	}
}

func Delete(client *golangsdk.ServiceClient, bandwidthID string) (r DeleteResult) {
	_, r.Err = client.Delete(resourceURL(client, bandwidthID), nil)
	return
}

type InsertOptsBuilder interface {
	ToBandwidthInsertMap() (map[string]any, error)
}

type InsertOpts struct {
	PublicIpInfo []PublicIpInfoInsertOpts `json:"publicip_info" required:"true"`
}

type PublicIpInfoInsertOpts struct {
	PublicIpID   string `json:"publicip_id" required:"true"`
	PublicIpType string `json:"publicip_type,omitempty"`
}

func (opts InsertOpts) ToBandwidthInsertMap() (map[string]any, error) {
	return golangsdk.BuildRequestBody(opts, "bandwidth")
}

func Insert(client *golangsdk.ServiceClient, bandwidthID string, opts InsertOptsBuilder) (r CreateResult) {
	b, err := opts.ToBandwidthInsertMap()
	if err != nil {
		r.Err = err
		return
	}

	_, r.Err = client.Post(insertURL(client, bandwidthID), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200, 201},
	})
	return
}

type RemoveOptsBuilder interface {
	ToBandwidthRemoveMap() (map[string]any, error)
}

type RemoveOpts struct {
	ChargeMode   string           `json:"charge_mode" required:"true"`
	Size         int              `json:"size" required:"true"`
	PublicIpInfo []PublicIpInfoID `json:"publicip_info" required:"true"`
}

type PublicIpInfoID struct {
	PublicIpID string `json:"publicip_id" required:"true"`
}

func (opts RemoveOpts) ToBandwidthRemoveMap() (map[string]any, error) {
	return golangsdk.BuildRequestBody(opts, "bandwidth")
}

func Remove(client *golangsdk.ServiceClient, bandwidthID string, opts RemoveOptsBuilder) (r DeleteResult) {
	b, err := opts.ToBandwidthRemoveMap()
	if err != nil {
		r.Err = err
		return
	}

	_, r.Err = client.Post(removeURL(client, bandwidthID), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200, 204},
	})
	return
}
