package dns

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type UpdateOpts struct {
	// domain name group name.
	Name string `json:"name" required:"true"`
	// Description.
	Description string `json:"description,omitempty"`
}

type UpdateQueryParams struct {
	// Firewall ID.
	FwInstanceID string `q:"fw_instance_id" required:"true"`
}

// This function is used to update domain name group information.
// groupId: Domain name group ID. It is the same as ID retuned while creating an address group.
// firewallId: Firewall Instance ID.
func UpdateDomainNameGroup(client *golangsdk.ServiceClient, groupId, firewallId string, opts UpdateOpts) (*DomainSetResponseData, error) {

	// PUT /v1/{project_id}/domain-set/{set_id}
	url, err := golangsdk.NewURLBuilder().WithEndpoints("domain-set", groupId).WithQueryParams(&UpdateQueryParams{
		FwInstanceID: firewallId,
	}).Build()
	if err != nil {
		return nil, err
	}

	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Put(client.ServiceURL(url.String()), b, nil, &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: map[string]string{"Content-Type": "application/json"},
	})
	if err != nil {
		return nil, err
	}

	var res CommonDomainNameGroupDataResponse
	return &res.Data, extract.Into(raw.Body, &res)
}
