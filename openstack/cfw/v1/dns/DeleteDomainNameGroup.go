package dns

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
)

type DeleteDomainNameGroupQueryParams struct {
	// Firewall ID.
	FwInstanceID string `q:"fw_instance_id" required:"true"`
}

// This function is used to delete a Domain Name Group.
// groupId: Domain name group ID. It is the same as ID retuned while creating a Domain Name Group.
// firewallId: Firewall Instance ID.
func DeleteDomainNameGroup(client *golangsdk.ServiceClient, groupId, firewallId string) error {
	// DELETE /v1/{project_id}/domain-set/{set_id}
	url, err := golangsdk.NewURLBuilder().WithEndpoints("domain-set", groupId).WithQueryParams(&DeleteDomainNameGroupQueryParams{
		FwInstanceID: firewallId,
	}).Build()
	if err != nil {
		return err
	}

	_, err = client.Delete(client.ServiceURL(url.String()), &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: map[string]string{"Content-Type": "application/json"},
	})
	return err
}
