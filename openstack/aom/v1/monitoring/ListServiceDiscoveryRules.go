package monitoring

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ListServiceDiscoveryRulesOpts struct {
	// Service discovery rule ID, which corresponds to a service discovery rule.
	// If this parameter is left blank, all service discovery rules in the project are returned.
	Id string `q:"id,omitempty"`
}

// ListServiceDiscoveryRules is used to query existing service discovery rules in the system.
// If Id is left blank, all service discovery rules in the project are returned.
func ListServiceDiscoveryRules(client *golangsdk.ServiceClient, opts ListServiceDiscoveryRulesOpts) ([]AppRule, error) {
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints("inv", "servicediscoveryrules").
		WithQueryParams(&opts).Build()
	if err != nil {
		return nil, err
	}

	// GET /v1/{project_id}/inv/servicediscoveryrules?id=<rule_id>
	raw, err := client.Get(client.ServiceURL(url.String()), nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res struct {
		AppRules []AppRule `json:"appRules"`
	}
	err = extract.Into(raw.Body, &res)
	return res.AppRules, err
}
