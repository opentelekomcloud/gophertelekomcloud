package services

import (
	"fmt"

	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
)

type Status string
type ServerType string
type ServiceType string

const (
	StatusCreating  Status = "creating"
	StatusAvailable Status = "available"
	StatusFailed    Status = "failed"
	StatusDeleting  Status = "deleting"
	StatusDeleted   Status = "" // this is a special status for missing LB

	ServerTypeVM  ServerType = "VM"
	ServerTypeVIP ServerType = "VIP"
	ServerTypeLB  ServerType = "LB"

	ServiceTypeInterface ServiceType = "interface"
	ServiceTypeGateway   ServiceType = "gateway"
)

type CreateOpts struct {
	// Specifies the ID for identifying the backend resource of the VPC endpoint service.
	// The ID is in the form of the UUID.
	PortID string `json:"port_id" required:"true"`
	// Specifies the ID of the cluster associated with the target VPCEP resource.
	PoolID string `json:"pool_id,omitempty"`
	// Specifies the name of the VPC endpoint service.
	// The value contains a maximum of 16 characters, including letters, digits, underscores (_), and hyphens (-).
	//  If you do not specify this parameter, the VPC endpoint service name is in the format: `regionName.serviceId`.
	//  If you specify this parameter, the VPC endpoint service name is in the format: `regionName.serviceName.serviceId`.
	ServiceName string `json:"service_name,omitempty"`
	// Specifies the ID of the VPC (router) to which the backend resource of the VPC endpoint service belongs.
	VpcId string `json:"vpc_id" required:"true"`
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
	// The values are as follows:
	//    close: indicates that the TOA and Proxy Protocol methods are neither used.
	//    toa_open: indicates that the TOA method is used.
	//    proxy_open: indicates that the Proxy Protocol method is used.
	//    open: indicates that the TOA and Proxy Protocol methods are both used.
	// The default value is close.
	TCPProxy string `json:"tcp_proxy,omitempty"`
	// Lists the resource tags.
	Tags []tags.ResourceTag `json:"tags,omitempty"`
	// Supplementary information about the VPC endpoint service.
	Description string `json:"description,omitempty"`
	// Specifies the ID of the virtual NIC to which the virtual IP address is bound.
	VIPPortID string `json:"vip_port_id,omitempty"`
}

type PortMapping struct {
	// Specifies the port for accessing the VPC endpoint.
	ClientPort int `json:"client_port"`
	// Specifies the port for accessing the VPC endpoint service.
	ServerPort int `json:"server_port"`
	// Specifies the protocol used in port mappings. The value can be TCP or UDP. The default value is TCP.
	Protocol string `json:"protocol,omitempty"`
}

func Create(client *golangsdk.ServiceClient, opts CreateOpts) (*Service, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}
	raw, err := client.Post(client.ServiceURL("vpc-endpoint-services"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200, 201},
	})
	if err != nil {
		return nil, err
	}
	var res Service
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type Service struct {
	// Specifies the unique ID of the VPC endpoint service.
	ID string `json:"id"`
	// Specifies the ID for identifying the backend resource of the VPC endpoint service.
	PortID string `json:"port_id"`
	// Specifies the name of the VPC endpoint service.
	ServiceName string `json:"service_name"`
	// Specifies the type of the VPC endpoint service.
	ServiceType string `json:"service_type"`
	// Specifies the resource type.
	ServerType ServerType `json:"server_type"`
	// Specifies the ID of the VPC to which the backend resource of the VPC endpoint service belongs.
	VpcID string `json:"vpc_id"`
	// Specifies the ID of the cluster associated with the target VPCEP resource.
	PoolID string `json:"pool_id"`
	// Specifies whether connection approval is required.
	ApprovalEnabled bool `json:"approval_enabled"`
	// Specifies the status of the VPC endpoint service.
	Status Status `json:"status"`
	// Specifies the creation time of the VPC endpoint service.
	// The UTC time format is used: YYYY-MM-DDTHH:MM:SSZ.
	CreatedAt string `json:"created_at"`
	// Specifies the update time of the VPC endpoint service.
	// The UTC time format is used: YYYY-MM-DDTHH:MM:SSZ.
	UpdatedAt string `json:"updated_at"`
	// Specifies the project ID.
	ProjectID string `json:"project_id"`
	// Lists the port mappings opened to the VPC endpoint service.
	Ports []PortMapping `json:"ports"`
	// Specifies whether the client IP address and port number or marker_id information is transmitted to the server.
	TCPProxy string `json:"tcp_proxy"`
	// Lists the resource tags.
	Tags []tags.ResourceTag `json:"tags"`
	// Supplementary information about the VPC endpoint service.
	Description string `json:"description,omitempty"`

	CIDRType string `json:"cidr_type"` // CIDRType returned only in Create
	// ConnectionCount is set in `Get` and `List` only
	ConnectionCount int `json:"connection_count"`
	// Error is set in `Get` and `List` only
	Error     []ErrorParameters `json:"error"`
	VIPPortID string            `json:"vip_port_id"`
}

type ErrorParameters struct {
	// Specifies the error code.
	ErrorCode string
	// Specifies the error message.
	ErrorMessage string
}

func WaitForServiceStatus(client *golangsdk.ServiceClient, id string, status Status, timeout int) error {
	return golangsdk.WaitFor(timeout, func() (bool, error) {
		srv, err := Get(client, id)
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
