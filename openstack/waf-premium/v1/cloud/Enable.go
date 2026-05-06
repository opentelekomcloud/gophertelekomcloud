package cloud

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type EnableOpts struct {
	// Website the account belongs to. The value is dt for Cloud website.
	ConsoleArea string `json:"console_area" required:"true"`
	// Enterprise project ID.
	EnterpriseProjectID string `json:"-" q:"enterprise_project_id,omitempty"`
}

// Enable to enable the pay-per-use billing mode for cloud WAF.
func Enable(client *golangsdk.ServiceClient, opts EnableOpts) (*EnablePostPaidResponse, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	url, err := golangsdk.NewURLBuilder().
		WithEndpoints("waf", "postpaid").
		WithQueryParams(&opts).
		Build()
	if err != nil {
		return nil, err
	}

	// POST /v1/{project_id}/waf/postpaid
	raw, err := client.Post(client.ServiceURL(url.String()), b,
		nil, &golangsdk.RequestOpts{
			OkCodes: []int{200},
			MoreHeaders: map[string]string{
				"Content-Type": "application/json;charset=utf8",
				"region":       client.RegionID,
			},
		})
	if err != nil {
		return nil, err
	}

	var res EnablePostPaidResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type EnablePostPaidResponse struct {
	// The edition for the cloud WAF instance.By default,
	// 22 is returned when the CreateCloudWafPostPaidResource API is called.
	// 22: The pay-per-use edition.
	Type int `json:"type"`
	// The resource list.
	Resources []ResourceResponse `json:"resources"`
	// New user or not.
	IsNewUser bool `json:"isNewUser"`
}

type ResourceResponse struct {
	// Resource ID.
	ID string `json:"resourceId"`
	// Cloud service type.
	CloudServiceType string `json:"cloudServiceType"`
	// Cloud service resource type.
	// hws.resource.type.waf: yearly/monthly cloud-mode WAF
	// hws.resource.type.waf.domain: domain name expansion packages in yearly/monthly cloud-mode WAF
	// hws.resource.type.waf.domain: bandwidth expansion packages in yearly/monthly cloud-mode WAF
	// hws.resource.type.waf.domain: rule expansion packages in yearly/monthly cloud-mode WAF
	// hws.resource.type.waf.instance: dedicated WAF instances
	// hws.resource.type.waf.payperuserequest: requests to pay-per-use WAF instances
	// hws.resource.type.waf.payperusedomain: domain names protected with pay-per-use WAF instances
	// hws.resource.type.waf.payperuserule: rules created in pay-per-use WAF instances
	ResourceType string `json:"resourceType"`
	// Cloud resource specifications.
	ResourceSpecCode string `json:"resourceSpecCode"`
	// Resource status. The value can be:
	// 0: Unfrozen/Normal.
	// 1: Frozen.
	// 2: Deleted.
	Status int `json:"status"`
	// Resource expiration time.
	ExpireTime string `json:"expireTime"`
	// Resource quantity of your resourceType.
	ResourceSize int `json:"resourceSize"`
}
