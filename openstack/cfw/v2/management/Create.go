package management

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type CreateOpts struct {
	// Specifies the firewall name
	Name string `json:"name" required:"true"`
	// Specifies the enterprise project ID, which is the ID of a project planned based on organizations.
	// You can obtain the enterprise project ID by referring to
	// https://docs.otc.t-systems.com/cloud-firewall/api-ref/appendix/obtaining_an_enterprise_project_id.html
	// If the enterprise project function is not enabled, the value is 0.
	EnterpriseProjectId string `json:"enterprise_project_id,omitempty"`
	// List of service resource tags. After tags are added to firewall resources, you can query resources and combine CDRs by key and value.
	Tags []CreateTags `json:"tags,omitempty"`
	// Specifies the firewall specifications.
	Flavor CreateFlavor `json:"flavor" required:"true"`
	// Specifies the billing type, which can be yearly/monthly or pay-per-use (default setting).
	ChargeInfo ChargeInfo `json:"charge_info" required:"true"`
}

type CreateFlavor struct {
	// Firewall edition. Only the professional edition is supported.
	Version string `json:"version" required:"true"`
}

type ChargeInfo struct {
	// Billing mode. The value can only be "postPaid", indicating pay-per-use billing.
	ChargeMode string `json:"charge_mode" required:"true"`
}

// This function is used to create a firewall
func Create(client *golangsdk.ServiceClient, opts CreateOpts) (*CreateResponse, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// POST /v2/{project_id}/firewall
	raw, err := client.Post(client.ServiceURL("firewall"), b, nil, &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: map[string]string{"Content-Type": "application/json"},
	})
	if err != nil {
		return nil, err
	}

	var res CreateResponse
	return &res, extract.Into(raw.Body, &res)
}

// ResponseBody represents the response body parameters.
type CreateResponse struct {
	// job_id: Instance creation task ID. This parameter is returned only when pay-per-use instances are created.
	// Note for developer: This is the same as instance ID.
	JobID string `json:"job_id"`
	// order_id: Order ID. This parameter is returned only when yearly/monthly instances are created.
	OrderID string `json:"order_id"`
	// data: Request body for creating a firewall.
	Data CreateResponseData `json:"data"`
}

type CreateResponseData struct {
	// Specifies the firewall name
	Name string `json:"name"`
	// Specifies the enterprise project ID, which is the ID of a project planned based on organizations.
	EnterpriseProjectId string `json:"enterprise_project_id"`
	// Specifies the number of nodes
	Tags []Tags `json:"tags"`
	// Specifies the firewall specifications.
	Flavor FlavorResp `json:"flavor"`
	// Specifies the billing type.
	ChargeInfo ChargeInfoResp `json:"charge_info"`
}

type FlavorResp struct {
	// Firewall edition.
	Version string `json:"version"`
}

type ChargeInfoResp struct {
	// Billing mode. The value can only be "postPaid", indicating pay-per-use billing.
	ChargeMode  string `json:"charge_mode" required:"true"`
	PeriodType  string `json:"period_type"`
	PeriodNum   string `json:"period_num"`
	IsAutoRenew string `json:"is_auto_renew"`
	IsAutoPay   string `json:"is_auto_pay"`
}
