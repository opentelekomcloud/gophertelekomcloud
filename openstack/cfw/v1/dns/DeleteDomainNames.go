package dns

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type DeleteDomainNameListQueryParams struct {
	// Firewall ID.
	FwInstanceID string `json:"fw_instance_id" required:"true"`
}

type DeleteDomainNameListOpts struct {
	// Protected object ID, which is used to distinguish between Internet border protection and VPC border protection
	// after a cloud firewall is created. You can obtain the ID by calling the Get function in management package.
	// In the return value, find the ID in ProtectObjects[n].ObjectID.
	// If the value of type is 0, the protected object ID belongs to the Internet border.
	// If the value of type is 1, the protected object ID belongs to the VPC border.
	ObjectID string `json:"object_id" required:"true"`
	// Domain name ID list. It can be obtained by using the function ListDomainNames OR GetDomainNameInfo.
	DomainAddressIDs []string `json:"domain_address_ids" required:"true"`
}

// This function is used to delete list of domain names from a domain name group.
// groupId: Domain name group ID. It is the same as ID retuned while creating a Domain Name Group.
// firewallId: Firewall Instance ID.
func DeleteDomainNames(client *golangsdk.ServiceClient, groupId, firewallId string, opts DeleteDomainNameListOpts) error {
	// DELETE /v1/{project_id}/domain-set/domains/{set_id}
	url, err := golangsdk.NewURLBuilder().WithEndpoints("domain-set", "domains", groupId).WithQueryParams(&DeleteDomainNameListQueryParams{
		FwInstanceID: firewallId,
	}).Build()
	if err != nil {
		return err
	}

	b, err := build.RequestBody(opts, "")
	if err != nil {
		return err
	}

	_, err = client.DeleteWithBody(client.ServiceURL(url.String()), b, &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: map[string]string{"Content-Type": "application/json"},
	})
	return err
}
