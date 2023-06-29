package rules

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type ListOpts struct {
	// Specifies the number of records on each page.
	//
	// Minimum: 0
	//
	// Maximum: 2000
	//
	// Default: 2000
	Limit *int `q:"limit"`
	// Specifies the ID of the last record on the previous page.
	//
	// Note:
	//
	// This parameter must be used together with limit.
	//
	// If this parameter is not specified, the first page will be queried.
	//
	// This parameter cannot be left blank or set to an invalid ID.
	Marker string `q:"marker"`
	// Specifies whether to use reverse query. Values:
	//
	// true: Query the previous page.
	//
	// false (default): Query the next page.
	//
	// Note:
	//
	// This parameter must be used together with limit.
	//
	// If page_reverse is set to true and you want to query the previous page, set the value of marker to the value of previous_marker.
	PageReverse *bool `q:"page_reverse"`
	// Specifies the forwarding rule ID.
	//
	// Multiple IDs can be queried in the format of id=xxx&id=xxx.
	Id []string `q:"id"`
	// Specifies how requests are matched with the domain names or URL. Values:
	//
	// EQUAL_TO: exact match.
	//
	// REGEX: regular expression match
	//
	// STARTS_WITH: prefix match
	//
	// Multiple values can be queried in the format of compare_type=xxx&compare_type=xxx.
	CompareType []string `q:"compare_type"`
	// Specifies the provisioning status of the forwarding rule. The value can only be ACTIVE, indicating that the forwarding rule is provisioned successfully.
	//
	// Multiple provisioning statuses can be queried in the format of provisioning_status=xxx&provisioning_status=xxx.
	ProvisioningStatus []string `q:"provisioning_status"`
	// Specifies whether reverse matching is supported.
	//
	// The value is fixed at false. This parameter can be updated but remains invalid.
	Invert *bool `q:"invert"`
	// Specifies the administrative status of the forwarding rule. The default value is true.
	//
	// This parameter is unsupported. Please do not use it.
	AdminStateUp *bool `q:"admin_state_up"`
	// Specifies the value of the match content.
	//
	// Multiple values can be queried in the format of value=xxx&value=xxx.
	Value []string `q:"value"`
	// Specifies the key of the match content that is used to identify the forwarding rule.
	//
	// Multiple keys can be queried in the format of key=xxx&key=xxx.
	//
	// This parameter is unsupported. Please do not use it.
	Key []string `q:"key"`
	// Specifies the match type. The value can be HOST_NAME or PATH.
	//
	// The type of forwarding rules for the same forwarding policy cannot be the same.
	//
	// Multiple types can be queried in the format of type=xxx&type=xxx.
	Type []string `q:"type"`
	// Specifies the enterprise project ID.
	//
	// If this parameter is not passed, resources in the default enterprise project are queried, and authentication is performed based on the default enterprise project.
	//
	// If this parameter is passed, its value can be the ID of an existing enterprise project (resources in the specific enterprise project are required) or all_granted_eps (resources in all enterprise projects are queried).
	//
	// Multiple IDs can be queried in the format of enterprise_project_id=xxx&enterprise_project_id=xxx.
	//
	// This parameter is unsupported. Please do not use it.
	EnterpriseProjectId []string `q:"enterprise_project_id"`
}

func List(client *golangsdk.ServiceClient, policyID string, opts ListOpts) pagination.Pager {
	query, err := golangsdk.BuildQueryString(opts)
	if err != nil {
		return pagination.Pager{Err: err}
	}

	return pagination.NewPager(client, client.ServiceURL("l7policies", policyID, "rules")+query.String(), func(r pagination.PageResult) pagination.Page {
		return RulePage{PageWithInfo: pagination.NewPageWithInfo(r)}
	})
}

type RulePage struct {
	pagination.PageWithInfo
}

func (p RulePage) IsEmpty() (bool, error) {
	rules, err := ExtractRules(p)
	return len(rules) == 0, err
}

func ExtractRules(p pagination.Page) ([]ForwardingRule, error) {
	var res []ForwardingRule
	err := extract.IntoSlicePtr(p.(RulePage).BodyReader(), &res, "rules")
	return res, err
}
