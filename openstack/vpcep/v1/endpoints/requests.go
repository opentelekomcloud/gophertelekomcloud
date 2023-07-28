package endpoints

import (
	"net/http"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type CreateOpts struct {
	// The value must be the ID of the subnet created in the VPC specified by vpc_id and in the format of the UUID.
	// This parameter is mandatory only if you create a VPC endpoint for connecting to an interface VPC endpoint service.
	NetworkID string `json:"subnet_id,omitempty"`

	// Specifies the ID of the VPC endpoint service.
	ServiceID string `json:"endpoint_service_id" required:"true"`

	// Specifies the ID of the VPC where the VPC endpoint is to be created.
	RouterID string `json:"vpc_id" required:"true"`

	// Specifies whether to create a private domain name.
	EnableDNS bool `json:"enable_dns"`

	// Lists the resource tags.
	Tags []tags.ResourceTag `json:"tags,omitempty"`

	// Lists the IDs of route tables.
	// This parameter is mandatory only if you create a VPC endpoint for connecting to a `gateway` VPC endpoint service.
	RouteTables []string `json:"routetables,omitempty"`

	// Specifies the IP address for accessing the associated VPC endpoint service.
	// This parameter is mandatory only if you create a VPC endpoint for connecting to an `interface` VPC endpoint service.
	PortIP string `json:"port_ip,omitempty"`

	// Specifies the whitelist for controlling access to the VPC endpoint.
	//
	// IPv4 addresses or CIDR blocks can be specified to control access when you create a VPC endpoint.
	//
	// This parameter is mandatory only when you create a VPC endpoint for connecting to an interface VPC endpoint service.
	Whitelist []string `json:"whitelist,omitempty"`

	// Specifies whether to enable access control.
	EnableWhitelist *bool `json:"enable_whitelist,omitempty"`
}

type CreateOptsBuilder interface {
	ToEndpointCreateMap() (map[string]interface{}, error)
}

func (opts CreateOpts) ToEndpointCreateMap() (map[string]interface{}, error) {
	return build.RequestBodyMap(opts, "")
}

func Create(client *golangsdk.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToEndpointCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(baseURL(client), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{http.StatusOK, http.StatusCreated},
	})
	return
}

func Get(client *golangsdk.ServiceClient, id string) (r GetResult) {
	_, r.Err = client.Get(resourceURL(client, id), &r.Body, nil)
	return
}

func Delete(client *golangsdk.ServiceClient, id string) (r DeleteResult) {
	_, r.Err = client.Delete(resourceURL(client, id), &golangsdk.RequestOpts{
		OkCodes: []int{http.StatusNoContent},
	})
	return
}

type ListOptsBuilder interface {
	ToEndpointListQuery() (string, error)
}

type ListOpts struct {
	ServiceName string `q:"endpoint_service_name"`
	RouterID    string `q:"vpc_id"`
	ID          string `q:"id"`
	SortKey     string `q:"sort_key"`
	SortDir     string `q:"sort_dir"`
}

func (opts ListOpts) ToEndpointListQuery() (string, error) {
	q, err := build.QueryString(opts)
	if err != nil {
		return "", err
	}
	return q.String(), nil
}

func List(client *golangsdk.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := baseURL(client)
	if opts != nil {
		q, err := opts.ToEndpointListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += q
	}
	return pagination.Pager{
		Client:     client,
		InitialURL: url,
		CreatePage: func(r pagination.PageResult) pagination.Page {
			return EndpointPage{OffsetPageBase: pagination.OffsetPageBase{PageResult: r}}
		},
	}
}

func WaitForEndpointStatus(client *golangsdk.ServiceClient, id string, status Status, timeout int) error {
	return golangsdk.WaitFor(timeout, func() (bool, error) {
		ep, err := Get(client, id).Extract()
		if err != nil {
			if _, ok := err.(golangsdk.ErrDefault404); ok && status == "" {
				return true, nil
			}
			return false, err
		}
		if ep.Status == status {
			return true, nil
		}
		return false, nil
	})
}
