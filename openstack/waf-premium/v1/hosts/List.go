package hosts

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func List(client *golangsdk.ServiceClient, opts ListOpts) ([]Host, error) {
	// GET /v1/{project_id}/premium-waf/host
	query, err := golangsdk.BuildQueryString(opts)
	if err != nil {
		return nil, err
	}

	url := client.ServiceURL("host") + query.String()
	raw, err := client.Get(url, nil, nil)
	if err != nil {
		return nil, err
	}

	var res []Host
	err = extract.IntoSlicePtr(raw.Body, &res, "items")
	return res, err
}

type ListOpts struct {
	// Number of records on each page.
	// The maximum value is 100. Default value: 10
	PageSize string `q:"pageSize,omitempty"`
	// Current page number
	Page string `q:"page,omitempty"`
	// Domain name
	Hostname string `q:"hostname,omitempty"`
	// Policy Name
	PolicyName string `q:"policyname,omitempty"`
	// WAF status of the protected domain name. The value can be:
	// -1: Bypassed. Requests are directly sent to the backend servers without passing through WAF.
	// 0: Suspended. WAF only forwards requests for the domain name but does not detect attacks.
	// 1: Enabled. WAF detects attacks based on the configured policy.
	ProtectStatus int `q:"protect_status,omitempty"`
}
