package dns

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type CreateDomainNameGroupOpts struct {
	// Firewall ID.
	FwInstanceID string `json:"fw_instance_id" required:"true"`
	// Protected object ID, which is used to distinguish between Internet border protection and VPC border protection
	// after a cloud firewall is created. You can obtain the ID by calling the Get function in management package.
	// In the return value, find the ID in ProtectObjects[n].ObjectID.
	// If the value of type is 0, the protected object ID belongs to the Internet border.
	// If the value of type is 1, the protected object ID belongs to the VPC border.
	ObjectID string `json:"object_id" required:"true"`
	// Domain name group name.
	Name string `json:"name" required:"true"`
	// Domain name group description.
	Description string `json:"description,omitempty"`
	// Domain name information list.
	DomainNames []DomainSetInfoDto `json:"domain_names" required:"true"`
	// Domain name group type:
	// 0 - Application domain name group
	// 1 - Network domain name group
	DomainSetType int `json:"domain_set_type,omitempty"`
}

// This function is used to add domain name group.
func CreateDomainNameGroup(client *golangsdk.ServiceClient, opts CreateDomainNameGroupOpts) (*DomainSetResponseData, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// POST /v1/{project_id}/domain-set
	raw, err := client.Post(client.ServiceURL("domain-set"), b, nil, &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: map[string]string{"Content-Type": "application/json"},
	})
	if err != nil {
		return nil, err
	}

	var res CommonDomainNameGroupDataResponse
	return &res.Data, extract.Into(raw.Body, &res)
}
