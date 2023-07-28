package services

import (
	"bytes"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type commonResult struct {
	golangsdk.Result
}

type CreateResult struct {
	commonResult
}

type GetResult struct {
	commonResult
}

type ErrorParameters struct {
	ErrorCode    string
	ErrorMessage string
}

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

type Service struct {
	ID              string             `json:"id"`
	PortID          string             `json:"port_id"`
	VIPPortID       string             `json:"vip_port_id"`
	ServiceName     string             `json:"service_name"`
	ServiceType     string             `json:"service_type"`
	ServerType      ServerType         `json:"server_type"`
	RouterID        string             `json:"vpc_id"`
	PoolID          string             `json:"pool_id"`
	ApprovalEnabled bool               `json:"approval_enabled"`
	Status          Status             `json:"status"`
	CreatedAt       string             `json:"created_at"`
	UpdatedAt       string             `json:"updated_at"`
	ProjectID       string             `json:"project_id"`
	CIDRType        string             `json:"cidr_type"` // CIDRType returned only in Create
	Ports           []PortMapping      `json:"ports"`
	TCPProxy        string             `json:"tcp_proxy"`
	Tags            []tags.ResourceTag `json:"tags"`

	// ConnectionCount is set in `Get` and `List` only
	ConnectionCount int `json:"connection_count"`
	// Error is set in `Get` and `List` only
	Error []ErrorParameters `json:"error"`
}

func (r commonResult) Extract() (*Service, error) {
	srv := &Service{}
	err := r.ExtractInto(srv)
	if err != nil {
		return nil, err
	}
	return srv, nil
}

type ServicePage struct {
	pagination.OffsetPageBase
}

func ExtractServices(p pagination.Page) ([]Service, error) {
	var srv []Service

	err := extract.IntoSlicePtr(bytes.NewReader((p.(ServicePage)).Body), &srv, "endpoint_services")
	return srv, err
}

type PublicService struct {
	ID          string      `json:"id"`
	Owner       string      `json:"owner"`
	ServiceName string      `json:"service_name"`
	ServiceType ServiceType `json:"service_type"`
	CreatedAt   string      `json:"created_at"`
	IsCharge    bool        `json:"is_charge"`
}

func ExtractPublicServices(p pagination.Page) ([]PublicService, error) {
	var srv []PublicService

	err := extract.IntoSlicePtr(bytes.NewReader((p.(ServicePage)).Body), &srv, "endpoint_services")
	return srv, err
}

func (p ServicePage) IsEmpty() (bool, error) {
	srv, err := ExtractServices(p)
	if err != nil {
		return false, err
	}
	return len(srv) == 0, err
}

type UpdateResult struct {
	commonResult
}

type DeleteResult struct {
	Err error
}
