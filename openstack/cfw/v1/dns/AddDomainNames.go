package dns

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type AddDomainNameListOpts struct {
	// Firewall ID.
	FwInstanceID string `json:"fw_instance_id" required:"true"`
	// Protected object ID, which is used to distinguish between Internet border protection and VPC border protection
	// after a cloud firewall is created. You can obtain the ID by calling the Get function in management package.
	// In the return value, find the ID in ProtectObjects[n].ObjectID.
	// If the value of type is 0, the protected object ID belongs to the Internet border.
	// If the value of type is 1, the protected object ID belongs to the VPC border.
	ObjectID string `json:"object_id" required:"true"`
	// Domain name information list.
	DomainNames []DomainSetInfoDto `json:"domain_names" required:"true"`
}

// This function is used to add a list of domain names to a domain name group.
// groupId: Domain name group ID. It is the same as ID retuned while creating a Domain Name Group.
func AddDomainNames(client *golangsdk.ServiceClient, groupId string, opts AddDomainNameListOpts) (*DomainSetResponseData, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// POST /v1/{project_id}/domain-set/domains/{set_id}
	raw, err := client.Post(client.ServiceURL("domain-set", "domains", groupId), b, nil, &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: map[string]string{"Content-Type": "application/json"},
	})
	if err != nil {
		return nil, err
	}

	var res CommonDomainNameGroupDataResponse
	return &res.Data, extract.Into(raw.Body, &res)
}
