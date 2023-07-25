package dnatrules

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
)

// DnatRule is a struct that represents a dnat rule
type DnatRule struct {
	ID                  string `json:"id"`
	TenantID            string `json:"tenant_id"`
	NatGatewayID        string `json:"nat_gateway_id"`
	PortID              string `json:"port_id"`
	PrivateIp           string `json:"private_ip"`
	InternalServicePort int    `json:"internal_service_port"`
	FloatingIpID        string `json:"floating_ip_id"`
	FloatingIpAddress   string `json:"floating_ip_address"`
	ExternalServicePort int    `json:"external_service_port"`
	Protocol            string `json:"protocol"`
	Status              string `json:"status"`
	AdminStateUp        bool   `json:"admin_state_up"`
	CreatedAt           string `json:"created_at"`
}

// GetResult is a return struct of get method
type GetResult struct {
	golangsdk.Result
}

func (r GetResult) Extract() (*DnatRule, error) {
	s := new(DnatRule)
	err := r.ExtractIntoStructPtr(s, "dnat_rule")
	if err != nil {
		return nil, err
	}
	return s, nil
}

// CreateResult is a return struct of create method
type CreateResult struct {
	golangsdk.Result
}

func (r CreateResult) Extract() (*DnatRule, error) {
	s := new(DnatRule)
	err := r.ExtractIntoStructPtr(s, "dnat_rule")
	if err != nil {
		return nil, err
	}
	return s, nil
}

// DeleteResult is a return struct of delete method
type DeleteResult struct {
	golangsdk.ErrResult
}
