package endpoints

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/vpcep/v1/services"
)

type Status string

const (
	StatusPendingAcceptance Status = "pendingAcceptance"
	StatusCreating          Status = "creating"
	StatusAccepted          Status = "accepted"
	StatusFailed            Status = "failed"
)

type CreateOpts struct {
	// The value must be the ID of the subnet created in the VPC specified by vpc_id and in the format of the UUID.
	// This parameter is mandatory only if you create a VPC endpoint for connecting to an interface VPC endpoint service.
	NetworkID string `json:"subnet_id,omitempty"`
	// Specifies the ID of the VPC endpoint service.
	ServiceID string `json:"endpoint_service_id" required:"true"`
	// Specifies the ID of the VPC where the VPC endpoint is to be created.
	VpcId string `json:"vpc_id" required:"true"`
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
	// IPv4 addresses or CIDR blocks can be specified to control access when you create a VPC endpoint.
	// This parameter is mandatory only when you create a VPC endpoint for connecting to an interface VPC endpoint service.
	Whitelist []string `json:"whitelist,omitempty"`
	// Specifies whether to enable access control.
	EnableWhitelist *bool `json:"enable_whitelist,omitempty"`
	// Specifies the name of the VPC endpoint specifications.
	SpecificationName string `json:"specification_name,omitempty"`
	// Specifies the policy of the gateway VPC endpoint.
	// This parameter is available only when you create a gateway VPC endpoint.
	// Array length: 0-10
	PolicyStatement []PolicyStatement `json:"policy_statement,omitempty"`
	// Specifies the description of the VPC endpoint.
	// The description can contain a maximum of 128 characters and cannot contain left angle brackets (<) or right angle brackets (>).
	Description string `json:"description,omitempty"`
}

type PolicyStatement struct {
	// Specifies the policy effect, which can be Allow or Deny.
	Effect string `json:"effect" required:"true"`
	// Specifies OBS access permissions.
	Action []string `json:"action" required:"true"`
	// Specifies the OBS object.
	Resource []string `json:"resource" required:"true"`
}

func Create(client *golangsdk.ServiceClient, opts CreateOpts) (*Endpoint, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}
	raw, err := client.Post(client.ServiceURL("vpc-endpoints"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200, 201},
	})
	if err != nil {
		return nil, err
	}
	var res Endpoint
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type Endpoint struct {
	// Specifies the unique ID of the VPC endpoint.
	ID string `json:"id"`
	// Specifies the type of the VPC endpoint service that is associated with the VPC endpoint.
	ServiceType services.ServiceType `json:"service_type"`
	// Specifies the connection status of the VPC endpoint.
	//    pendingAcceptance: indicates that the VPC endpoint is pending acceptance.
	//    creating: indicates the VPC endpoint is being created.
	//    accepted: indicates the VPC endpoint has been accepted.
	//    failed: indicates the creation of the VPC endpoint failed.
	Status Status `json:"status"`
	// Specifies the domain status.
	//    frozen: indicates that the domain is frozen.
	//    active: indicates that the domain is normal.
	ActiveStatus []string `json:"active_status"`
	// Specifies the name of the VPC endpoint service.
	ServiceName string `json:"endpoint_service_name"`
	// Specifies the packet ID of the VPC endpoint.
	MarkerID int `json:"marker_id"`
	// Specifies the ID of the VPC endpoint service.
	ServiceID string `json:"endpoint_service_id"`
	// Specifies whether to create a private domain name.
	EnableDNS bool `json:"enable_dns"`
	// Specifies the domain name for accessing the associated VPC endpoint service.
	DNSNames []string `json:"dns_names"`
	// Specifies the ID of the subnet (OS network) in the VPC specified by `vpc_id`. The value is in the UUID format.
	NetworkID string `json:"subnet_id"`
	// Specifies the ID of the VPC where the VPC endpoint is to be created.
	VpcID string `json:"vpc_id"`
	// Specifies the creation time of the VPC endpoint.
	CreatedAt string `json:"created_at"`
	// Specifies the update time of the VPC endpoint.
	UpdatedAt string `json:"updated_at"`
	// Specifies the project ID.
	ProjectID string `json:"project_id"`
	// Lists the resource tags.
	Tags []tags.ResourceTag `json:"tags"`
	// Specifies the whitelist for controlling access to the VPC endpoint.
	Whitelist []string `json:"whitelist"`
	// Specifies whether to enable access control.
	EnableWhitelist bool `json:"enable_whitelist"`
	// Lists the IDs of route tables.
	RouteTables []string `json:"routetables"`
	// Specifies the name of the VPC endpoint specifications.
	SpecificationName string `json:"specification_name"`
	// Specifies whether to enable the endpoint.
	// enable: The endpoint will be enabled.
	// disable: The endpoint will be disabled.
	EnableStatus string `json:"enable_status"`
	// Specifies the policy of the gateway VPC endpoint.
	PolicyStatement []PolicyStatement `json:"policy_statement"`
	// Specifies the ID of the cluster associated with the VPC endpoint.
	PoolID string `json:"endpoint_pool_id"`
	// Specifies the description of the VPC endpoint.
	Description string `json:"description"`
	// Specifies the IP address for accessing the associated VPC endpoint service.
	IP string `json:"ip"`
}

func WaitForEndpointStatus(client *golangsdk.ServiceClient, id string, status Status, timeout int) error {
	return golangsdk.WaitFor(timeout, func() (bool, error) {
		ep, err := Get(client, id)
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
