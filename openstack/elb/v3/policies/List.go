package policies

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type ListOpts struct {
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
	// Specifies the number of records on each page.
	//
	// Minimum: 0
	//
	// Maximum: 2000
	//
	// Default: 2000
	Limit *int `q:"limit"`
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
	// Specifies the forwarding policy ID.
	//
	// Multiple IDs can be queried in the format of id=xxx&id=xxx.
	Id []string `q:"id"`
	// Specifies the forwarding policy name.
	//
	// Multiple names can be queried in the format of name=xxx&name=xxx.
	Name []string `q:"name"`
	// Provides supplementary information about the forwarding policy.
	//
	// Multiple descriptions can be queried in the format of description=xxx&description=xxx.
	Description []string `q:"description"`
	// Specifies the administrative status of the forwarding policy. The default value is true.
	//
	// This parameter is unsupported. Please do not use it.
	AdminStateUp *bool `q:"admin_state_up"`
	// Specifies the ID of the listener to which the forwarding policy is added.
	//
	// Multiple IDs can be queried in the format of listener_id=xxx&listener_id=xxx.
	ListenerId []string `q:"listener_id"`
	// Specifies the forwarding policy priority.
	//
	// Multiple priorities can be queried in the format of position=xxx&position=xxx.
	//
	// This parameter is unsupported. Please do not use it.
	Position []string `q:"position"`
	// Specifies where requests are forwarded. The value can be one of the following:
	//
	// REDIRECT_TO_POOL: Requests are forwarded to another backend server group.
	//
	// REDIRECT_TO_LISTENER: Requests are redirected to an HTTPS listener.
	//
	// REDIRECT_TO_URL: Requests are redirected to another URL.
	//
	// FIXED_RESPONSE: A fixed response body is returned.
	//
	// Multiple values can be queried in the format of action=xxx&action=xxx.
	Action []string `q:"action"`
	// Specifies the URL to which requests will be forwarded. The URL must be in the format of protocol://host:port/path?query.
	//
	// Multiple URLs can be queried in the format of redirect_url=xxx&redirect_url=xxx.
	//
	// This parameter is unsupported. Please do not use it.
	RedirectUrl []string `q:"redirect_url"`
	// Specifies the ID of the backend server group to which requests will be forwarded.
	//
	// Multiple IDs can be queried in the format of redirect_pool_id=xxx&redirect_pool_id=xxx.
	RedirectPoolId []string `q:"redirect_pool_id"`
	// Specifies the ID of the listener to which requests are redirected.
	//
	// Multiple IDs can be queried in the format of redirect_listener_id=xxx&redirect_listener_id=xxx.
	RedirectListenerId []string `q:"redirect_listener_id"`
	// Specifies the provisioning status of the forwarding policy.
	//
	// ACTIVE: The forwarding policy is provisioned successfully.
	//
	// ERROR: The forwarding policy has the same rule as another forwarding policy added to the same listener.
	//
	// Multiple provisioning statuses can be queried in the format of provisioning_status=xxx&provisioning_status=xxx.
	ProvisioningStatus []string `q:"provisioning_status"`
	// Specifies whether to display details about the forwarding rule in the forwarding policy.
	//
	// true: Details about the forwarding rule are displayed.
	//
	// false: Only the rule ID is displayed.
	DisplayAllRules *bool `q:"display_all_rules"`
	// Specifies the forwarding policy priority. A smaller value indicates a higher priority.
	//
	// Multiple priorities can be queried in the format of position=xxx&position=xxx.
	Priority []string `q:"priority"`
}

func List(client *golangsdk.ServiceClient, opts ListOpts) pagination.Pager {
	q, err := golangsdk.BuildQueryString(opts)
	if err != nil {
		return pagination.Pager{Err: err}
	}

	return pagination.NewPager(client, client.ServiceURL("l7policies")+q.String(), func(r pagination.PageResult) pagination.Page {
		return PolicyPage{PageWithInfo: pagination.NewPageWithInfo(r)}
	})
}

type PolicyPage struct {
	pagination.PageWithInfo
}

func (p PolicyPage) IsEmpty() (bool, error) {
	l, err := ExtractPolicies(p)
	if err != nil {
		return false, err
	}
	return len(l) == 0, nil
}

func ExtractPolicies(p pagination.Page) ([]L7Policy, error) {
	var res []L7Policy
	err := extract.IntoSlicePtr(p.(PolicyPage).BodyReader(), &res, "l7policies")
	return res, err
}
