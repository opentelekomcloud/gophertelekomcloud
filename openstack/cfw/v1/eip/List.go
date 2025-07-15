package eip

import (
	"fmt"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ListOpts struct {
	// Protected object ID, which is used to distinguish between Internet border protection and VPC border protection
	// after a cloud firewall is created. You can obtain the ID by calling the Get function in management package.
	// In the return value, find the ID in ProtectObjects[n].ObjectID.
	// If the value of type is 0, the protected object ID belongs to the Internet border.
	// If the value of type is 1, the protected object ID belongs to the VPC border.
	ObjectID string `q:"object_id" required:"true"`
	// Keyword for querying the protected EIP list. You can set an EIP ID or an EIP.
	KeyWord string `q:"key_word,omitempty"`
	// Protection status: null (all), 0 (enabled), or 1 (disabled).
	Status string `q:"status,omitempty"`
	// Enterprise project ID of a project based on organizations. If not enabled, the value is 0.
	EnterpriseProjectID string `q:"enterprise_project_id,omitempty"`
	// Device keyword, which is the name or ID of the asset bound to an EIP.
	DeviceKey string `q:"device_key,omitempty"`
	// Internet protocol type of an address: 0 (IPv4), 1 (IPv6).
	AddressType *int `q:"address_type,omitempty"`
	// Firewall ID, which can be obtained by referring to "Obtaining a Firewall ID".
	FwInstanceID string `q:"fw_instance_id,omitempty"`
	// Firewall keyword, which can be queried based on the firewall ID or name.
	FwKeyWord string `q:"fw_key_word,omitempty"`
	// Enterprise project ID of the EIP. If not enabled, the value is 0.
	EpsID string `q:"eps_id,omitempty"`
	// Tags that can be obtained by querying in the EIP console.
	Tags string `q:"tags,omitempty"`
}

// This function function is used to query the EIP list and display their protection status among other details.
func List(client *golangsdk.ServiceClient, opts ListOpts) ([]EipResource, error) {
	// GET /v1/{project_id}/eips/protect
	url, err := golangsdk.NewURLBuilder().WithEndpoints("eips", "protect").WithQueryParams(opts).Build()
	if err != nil {
		return nil, err
	}
	urlFixed := fmt.Sprintf("%s&limit=1024&offset=0&sync=1", url.String())

	raw, err := client.Get(client.ServiceURL(urlFixed), nil, &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: map[string]string{"Content-Type": "application/json"},
	})
	if err != nil {
		return nil, err
	}

	var res ListEIPResponse
	err = extract.Into(raw.Body, &res)
	return res.Data.Records, err
}

type ListEIPResponse struct {
	Data ListEIPResponseData `json:"data"`
}

type ListEIPResponseData struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
	// EIP resource record.
	Records []EipResource `json:"records"`
	Total   int           `json:"total"`
}

type EipResource struct {
	// Type of the associated instance: NATGW, ELB, or PORT.
	AssociateInstanceType string `json:"associate_instance_type"`
	// ID of the device (such as ECS and NAT) bound to the EIP.
	DeviceID string `json:"device_id"`
	// Name of the device (such as ECS and NAT) bound to the EIP
	DeviceName string `json:"device_name"`
	// Owner of the device (such as ECS and NAT) bound to the EIP.
	DeviceOwner string `json:"device_owner"`
	// ID of the user that an EIP belongs to.
	DomainID string `json:"domain_id"`
	// Enterprise project ID of the account that the EIP belongs to.
	EnterpriseProjectID string `json:"enterprise_project_id"`
	// User that a firewall belongs to.
	FwDomainID string `json:"fw_domain_id"`
	// Enterprise project ID of the firewall bound to the EIP.
	FwEnterpriseProjectID string `json:"fw_enterprise_project_id"`
	// Firewall instance ID, which is automatically generated after a CFW instance is created.
	FwInstanceID string `json:"fw_instance_id"`
	// Firewall name.
	FwInstanceName string `json:"fw_instance_name"`
	// EIP ID.
	ID string `json:"id"`
	// Protected object ID.
	ObjectID string `json:"object_id"`
	// EIP
	PublicIP string `json:"public_ip"`
	// EIP (IPV6)
	PublicIPV6 string `json:"public_ipv6"`
	// EIP protection status: 0 (protected), 1 (unprotected).
	Status int `json:"status"`
	// Tag list.
	Tags string `json:"tags"`
}
