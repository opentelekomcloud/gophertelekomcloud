package endpoints

import (
	"bytes"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/vpcep/v1/services"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type Status string

const (
	StatusPendingAcceptance Status = "pendingAcceptance"
	StatusCreating          Status = "creating"
	StatusAccepted          Status = "accepted"
	StatusFailed            Status = "failed"
)

type Endpoint struct {
	// Specifies the unique ID of the VPC endpoint.
	ID string `json:"id"`

	// Specifies the type of the VPC endpoint service that is associated with the VPC endpoint.
	ServiceType services.ServiceType `json:"service_type"`

	// Specifies the packet ID of the VPC endpoint.
	MarkerID int `json:"marker_id"`

	// Specifies the connection status of the VPC endpoint.
	//
	//    pendingAcceptance: indicates that the VPC endpoint is pending acceptance.
	//    creating: indicates the VPC endpoint is being created.
	//    accepted: indicates the VPC endpoint has been accepted.
	//    failed: indicates the creation of the VPC endpoint failed.
	Status Status `json:"status"`

	// Specifies the domain status.
	//    frozen: indicates that the domain is frozen.
	//    active: indicates that the domain is normal.
	ActiveStatus []string `json:"active_status"`

	// Specifies the ID of the VPC where the VPC endpoint is to be created.
	RouterID string `json:"vpc_id"`

	// Specifies the ID of the subnet (OS network) in the VPC specified by `vpc_id`. The value is in the UUID format.
	NetworkID string `json:"subnet_id"`

	// Specifies whether to create a private domain name.
	EnableDNS bool `json:"enable_dns"`

	// Specifies the domain name for accessing the associated VPC endpoint service.
	DNSNames []string `json:"dns_names"`

	// Specifies the IP address for accessing the associated VPC endpoint service.
	IP string `json:"ip"`

	// Specifies the name of the VPC endpoint service.
	ServiceName string `json:"endpoint_service_name"`

	// Specifies the ID of the VPC endpoint service.
	ServiceID string `json:"endpoint_service_id"`

	// Specifies the project ID.
	ProjectID string `json:"project_id"`

	// Specifies the whitelist for controlling access to the VPC endpoint.
	Whitelist []string `json:"whitelist"`

	// Specifies whether to enable access control.
	EnableWhitelist bool `json:"enable_whitelist"`

	// Lists the IDs of route tables.
	RouteTables []string `json:"routetables"`

	// Specifies the creation time of the VPC endpoint.
	CreatedAt string `json:"created_at"`

	// Specifies the update time of the VPC endpoint.
	UpdatedAt string `json:"updated_at"`

	// Lists the resource tags.
	Tags []tags.ResourceTag `json:"tags"`
}

type commonResult struct {
	golangsdk.Result
}

func (r commonResult) Extract() (*Endpoint, error) {
	ep := new(Endpoint)
	err := r.ExtractIntoStructPtr(ep, "")
	if err != nil {
		return nil, err
	}
	return ep, nil
}

type CreateResult struct {
	commonResult
}

type GetResult struct {
	commonResult
}

type UpdateResult struct {
	commonResult
}

type DeleteResult struct {
	golangsdk.ErrResult
}

type EndpointPage struct {
	pagination.OffsetPageBase
}

func (p EndpointPage) IsEmpty() (bool, error) {
	eps, err := ExtractEndpoints(p)
	if err != nil {
		return false, err
	}
	return len(eps) == 0, nil
}

func ExtractEndpoints(p pagination.Page) ([]Endpoint, error) {
	var eps []Endpoint

	err := extract.IntoSlicePtr(bytes.NewReader(p.(EndpointPage).Body), &eps, "endpoints")
	return eps, err
}
