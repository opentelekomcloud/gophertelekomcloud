package services

import (
	"fmt"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type PortMapping struct {
	// Specifies the port for accessing the VPC endpoint.
	ClientPort int `json:"client_port"`
	// Specifies the port for accessing the VPC endpoint service.
	ServerPort int `json:"server_port"`
	// Specifies the protocol used in port mappings. The value can be TCP or UDP. The default value is TCP.
	Protocol string `json:"protocol"`
}

type CreateOpts struct {
	// Specifies the ID for identifying the backend resource of the VPC endpoint service.
	// The ID is in the form of the UUID.
	PortID string `json:"port_id" required:"true"`

	// Specifies the ID of the cluster associated with the target VPCEP resource.
	PoolID string `json:"pool_id,omitempty"`

	// Specifies the ID of the virtual NIC to which the virtual IP address is bound.
	VIPPortID string `json:"vip_port_id,omitempty"`

	// Specifies the name of the VPC endpoint service.
	// The value contains a maximum of 16 characters, including letters, digits, underscores (_), and hyphens (-).
	//
	//  If you do not specify this parameter, the VPC endpoint service name is in the format: `regionName.serviceId`.
	//  If you specify this parameter, the VPC endpoint service name is in the format: `regionName.serviceName.serviceId`.
	ServiceName string `json:"service_name,omitempty"`

	// Specifies the ID of the VPC (router) to which the backend resource of the VPC endpoint service belongs.
	RouterID string `json:"vpc_id" required:"true"`

	// Specifies whether connection approval is required.
	// The default value is `true`.
	ApprovalEnabled *bool `json:"approval_enabled,omitempty"`

	// Specifies the type of the VPC endpoint service.
	// Only your private services can be configured into interface VPC endpoint services.
	ServiceType ServiceType `json:"service_type,omitempty"`

	// Specifies the backend resource type.
	//  - `VM`: Resource is an ECS. Backend resources of this type serve as servers.
	//  - `VIP`: Resource is a virtual IP address that functions as a physical server hosting virtual resources.
	//  - `LB`: Resource is an enhanced load balancer.
	ServerType ServerType `json:"server_type" required:"true"`

	// Lists the port mappings opened to the VPC endpoint service.
	Ports []PortMapping `json:"ports" required:"true"`

	// Specifies whether the client IP address and port number or `marker_id` information is transmitted to the server.
	//
	// The values are as follows:
	//    close: indicates that the TOA and Proxy Protocol methods are neither used.
	//    toa_open: indicates that the TOA method is used.
	//    proxy_open: indicates that the Proxy Protocol method is used.
	//    open: indicates that the TOA and Proxy Protocol methods are both used.
	// The default value is close.
	TCPProxy string `json:"tcp_proxy,omitempty"`

	// Lists the resource tags.
	Tags []tags.ResourceTag `json:"tags,omitempty"`
}

type CreateOptsBuilder interface {
	ToServiceCreateMap() (map[string]interface{}, error)
}

func (opts CreateOpts) ToServiceCreateMap() (map[string]interface{}, error) {
	return build.RequestBodyMap(opts, "")
}

func Create(client *golangsdk.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToServiceCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(baseURL(client), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200, 201},
	})
	return
}

func Get(client *golangsdk.ServiceClient, id string) (r GetResult) {
	_, r.Err = client.Get(resourceURL(client, id), &r.Body, nil)
	return
}

type ListOptsBuilder interface {
	ToServiceListQuery() (string, error)
}

type ListOpts struct {
	// Specifies the name of the VPC endpoint service. The value is not case-sensitive and supports fuzzy match.
	Name string `q:"endpoint_service_name"`
	// Specifies the unique ID of the VPC endpoint service.
	ID string `q:"id"`
	// Specifies the status of the VPC endpoint service.
	//
	//    creating: indicates the VPC endpoint service is being created.
	//    available: indicates the VPC endpoint service is connectable.
	//    failed: indicates the creation of the VPC endpoint service failed.
	//    deleting: indicates the VPC endpoint service is being deleted.
	Status Status `q:"status"`

	SortKey string `q:"sort_key"`
	SortDir string `q:"sort_dir"`
}

func (opts ListOpts) ToServiceListQuery() (string, error) {
	q, err := build.QueryString(opts)
	if err != nil {
		return "", err
	}
	return q.String(), nil
}

func List(client *golangsdk.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := baseURL(client)
	if opts != nil {
		q, err := opts.ToServiceListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += q
	}
	return pagination.Pager{
		Client:     client,
		InitialURL: url,
		CreatePage: func(r pagination.PageResult) pagination.Page {
			return ServicePage{pagination.OffsetPageBase{PageResult: r}}
		},
	}
}

func ListPublic(client *golangsdk.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := publicURL(client)
	if opts != nil {
		q, err := opts.ToServiceListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += q
	}
	return pagination.Pager{
		Client:     client,
		InitialURL: url,
		CreatePage: func(r pagination.PageResult) pagination.Page {
			return ServicePage{pagination.OffsetPageBase{PageResult: r}}
		},
	}
}

type UpdateOptsBuilder interface {
	ToServiceUpdateMap() (map[string]interface{}, error)
}

type UpdateOpts struct {
	ApprovalEnabled *bool         `json:"approval_enabled,omitempty"`
	ServiceName     string        `json:"service_name,omitempty"`
	Ports           []PortMapping `json:"ports,omitempty"`
	PortID          string        `json:"port_id,omitempty"`
	VIPPortID       string        `json:"vip_port_id,omitempty"`
}

func (opts UpdateOpts) ToServiceUpdateMap() (map[string]interface{}, error) {
	return build.RequestBodyMap(opts, "")
}

func Update(client *golangsdk.ServiceClient, id string, opts UpdateOptsBuilder) (r UpdateResult) {
	b, err := opts.ToServiceUpdateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Put(resourceURL(client, id), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200, 201},
	})
	return
}

func Delete(client *golangsdk.ServiceClient, id string) (r DeleteResult) {
	_, r.Err = client.Delete(resourceURL(client, id), nil)
	return
}

func WaitForServiceStatus(client *golangsdk.ServiceClient, id string, status Status, timeout int) error {
	return golangsdk.WaitFor(timeout, func() (bool, error) {
		srv, err := Get(client, id).Extract()
		if err != nil {
			if _, ok := err.(golangsdk.ErrDefault404); ok && status == StatusDeleted {
				return true, nil
			}
			return false, fmt.Errorf("error waiting for service to have status %s: %w", status, err)
		}
		if srv.Status == status {
			return true, nil
		}
		return false, nil
	})
}
