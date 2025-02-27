package eip

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ChangeEIPProtectionQueryParams struct {
	// Enterprise project ID, which is the ID of a project planned based on organizations.
	EnterpriseProjectId string `q:"enterprise_project_id,omitempty"`
	// Firewall ID
	FwInstanceId string `q:"fw_instance_id,omitempty"`
}

type ChangeEIPProtectionOpts struct {
	// Protected object ID, which is used to distinguish between Internet border protection and VPC border protection
	// after a cloud firewall is created. You can obtain the ID by calling the Get function in management package.
	// In the return value, find the ID in ProtectObjects[n].ObjectID.
	// If the value of type is 0, the protected object ID belongs to the Internet border.
	// If the value of type is 1, the protected object ID belongs to the VPC border.
	ObjectID string `json:"object_id" required:"true"`
	// Status that an EIP will be changed to: 0 (protected), 1 (unprotected).
	// Note for dev: Don't add required tag to this field since Status: 0 will return missing input error during validation.
	Status int `json:"status"`
	// List of EIPs whose protection status is changed.
	IPInfos []IPInfo `json:"ip_infos" required:"true"`
}

// IPInfo represents an individual EIP whose protection status is changed.
type IPInfo struct {
	// EIP ID
	ID string `json:"id,omitempty"`
	// EIP IPv4 address.
	PublicIP string `json:"public_ip,omitempty"`
	// EIP IPv6 address.
	PublicIPv6 string `json:"public_ipv6,omitempty"`
}

// This function is used to enable or disable EIP protection.
// After a customer purchases an EIP, the customer needs to call ListEips to synchronize EIPs asset before enabling EIP protection for the first time.
// The sync field should be set to 1.
func ChangeEIPProtection(client *golangsdk.ServiceClient, firewallId string, opts ChangeEIPProtectionOpts) (*EIPSwitchStatusVO, error) {
	// POST /v1/{project_id}/eip/protect
	url, err := golangsdk.NewURLBuilder().WithEndpoints("eip", "protect").WithQueryParams(&ChangeEIPProtectionQueryParams{
		FwInstanceId: firewallId,
	}).Build()
	if err != nil {
		return nil, err
	}

	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL(url.String()), b, nil, &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: map[string]string{"Content-Type": "application/json"},
	})
	if err != nil {
		return nil, err
	}

	var res ChangeEIPProtectionResponse
	return &res.Data, extract.Into(raw.Body, &res)
}

type ChangeEIPProtectionResponse struct {
	Data EIPSwitchStatusVO `json:"data"`
}

type EIPSwitchStatusVO struct {
	// Protected object ID, which is used to distinguish between Internet border protection and VPC border protection
	// after a cloud firewall is created.
	ObjectID string `json:"object_id"`
	// List of EIP protection statuses that fail to be modified.
	// The status can be "successful" or "fail".
	FailEIPIDList []string `json:"fail_eip_id_list"`
	// List of failures to modify the EIP protection status.
	FailEIPList []FailedEIPInfo `json:"fail_eip_list"`
	// Firewall ID, which can be obtained by referring to "Obtaining a Firewall ID".
	ID string `json:"id"`
}

type FailedEIPInfo struct {
	// ID of an EIP whose status fails to be changed.
	ID string `json:"id"`
	// Error code of a status change failure.
	ErrorMessage string `json:"error_message"`
}
