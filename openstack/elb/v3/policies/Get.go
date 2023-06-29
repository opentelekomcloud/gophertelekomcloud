package policies

import (
	"net/http"

	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func Get(client *golangsdk.ServiceClient, id string) (*L7Policy, error) {
	raw, err := client.Get(client.ServiceURL("l7policies", id), nil, nil)
	return extra(err, raw)
}

func extra(err error, raw *http.Response) (*L7Policy, error) {
	if err != nil {
		return nil, err
	}

	var res L7Policy
	err = extract.IntoStructPtr(raw.Body, &res, "l7policy")
	return &res, err
}

type L7Policy struct {
	// Specifies where requests will be forwarded. The value can be one of the following:
	//
	// REDIRECT_TO_POOL: Requests will be forwarded to another backend server group.
	//
	// REDIRECT_TO_LISTENER: Requests will be redirected to an HTTPS listener.
	//
	// REDIRECT_TO_URL: Requests will be redirected to another URL.
	//
	// FIXED_RESPONSE: A fixed response body will be returned.
	//
	// REDIRECT_TO_LISTENER has the highest priority. If requests are to be redirected to an HTTPS listener, other forwarding policies of the listener will become invalid.
	//
	// Note:
	//
	// If action is set to REDIRECT_TO_POOL, the listener's protocol must be HTTP, HTTPS, or TERMINATED_HTTPS.
	//
	// If action is set to REDIRECT_TO_LISTENER, the listener's protocol must be HTTP.
	Action string `json:"action"`
	// Specifies the administrative status of the forwarding policy. The default value is true.
	//
	// This parameter is unsupported. Please do not use it.
	AdminStateUp *bool `json:"admin_state_up"`
	// Provides supplementary information about the forwarding policy.
	Description string `json:"description"`
	// Specifies the forwarding policy ID.
	Id string `json:"id"`
	// Specifies the ID of the listener to which the forwarding policy is added.
	ListenerId string `json:"listener_id"`
	// Specifies the forwarding policy name.
	//
	// Minimum: 1
	//
	// Maximum: 255
	Name string `json:"name"`
	// Specifies the forwarding policy priority. This parameter cannot be updated.
	//
	// This parameter is unsupported. Please do not use it.
	//
	// Minimum: 1
	//
	// Maximum: 100
	Position *int `json:"position"`
	// Specifies the forwarding policy priority. A smaller value indicates a higher priority. The value must be unique for forwarding policies of the same listener. This parameter will take effect only when enhance_l7policy_enable is set to true. If this parameter is passed and enhance_l7policy_enable is set to false, an error will be returned. This parameter is unsupported for shared load balancers.
	//
	// If action is set to REDIRECT_TO_LISTENER, the value can only be 0, indicating REDIRECT_TO_LISTENER has the highest priority.
	//
	// If enhance_l7policy_enable is not enabled, forwarding policies are automatically prioritized based on the original policy sorting logic. The priorities of domain names are independent from each other. For the same domain name, the priorities are sorted in the order of exact match (EQUAL_TO), prefix match (STARTS_WITH), and regular expression match (REGEX). If the matching types are the same, the longer the URL is, the higher the priority is. If a forwarding policy contains only a domain name without a path specified, the path is /, and prefix match is used by default.
	//
	// If enhance_l7policy_enable is set to true and this parameter is not passed, the priority will be a sum of 1 and the highest priority of existing forwarding policy in the same listener by default. If the highest priority of existing forwarding policies is the maximum (10,000), the forwarding policy will fail to be created because the final priority for creating the forwarding policy is the sum of 1 and 10,000, which exceeds the maximum. In this case, specify a value or adjust the priorities of existing forwarding policies. If no forwarding policies exist, the highest priority of existing forwarding policies will be set to 1 by default.
	//
	// This parameter is invalid for shared load balancers.
	//
	// This parameter is not available in eu-nl region. Please do not use it.
	Priority *int `json:"priority"`
	// Specifies the project ID of the forwarding policy.
	ProjectId string `json:"project_id"`
	// Specifies the provisioning status of the forwarding policy.
	//
	// The value can be ACTIVE or ERROR.
	//
	// ACTIVE (default): The forwarding policy is provisioned successfully.
	ProvisioningStatus string `json:"provisioning_status"`
	// Specifies the ID of the backend server group that requests will be forwarded to.
	//
	// This parameter is valid and mandatory only when action is set to REDIRECT_TO_POOL.
	//
	// If both redirect_pools_config and redirect_pool_id are specified, redirect_pools_config will take effect.
	RedirectPoolId string `json:"redirect_pool_id"`
	// Specifies the configuration of the backend server group that the requests are forwarded to. This parameter is valid only when action is set to REDIRECT_TO_POOL.
	RedirectPoolsConfig []RedirectPoolOptions `json:"redirect_pools_config"`
	// Specifies the ID of the listener to which requests are redirected. This parameter is mandatory when action is set to REDIRECT_TO_LISTENER.
	//
	// Note:
	//
	// The listener's protocol must be HTTPS or TERMINATED_HTTPS.
	//
	// A listener added to another load balancer is not allowed.
	//
	// This parameter cannot be passed in the API for adding or updating a forwarding policy if action is set to REDIRECT_TO_POOL.
	RedirectListenerId string `json:"redirect_listener_id"`
	// Specifies the URL to which requests are forwarded.
	//
	// Format: protocol://host:port/path?query
	//
	// This parameter is unsupported. Please do not use it.
	RedirectUrl string `json:"redirect_url"`
	// Lists the forwarding rules in the forwarding policy.
	Rules []RuleCondition `json:"rules"`
	// Specifies the URL to which requests are forwarded.
	//
	// For dedicated load balancers, this parameter will take effect only when advanced forwarding is enabled (enhance_l7policy_enable is set to true). If it is passed when enhance_l7policy_enable is set to false, an error will be returned.
	//
	// This parameter is mandatory when action is set to REDIRECT_TO_URL. It cannot be specified if the value of action is not REDIRECT_TO_URL.
	//
	// Format: protocol://host:port/path?query
	//
	// At least one of the four parameters (protocol, host, port, and path) must be passed, or their values cannot be set to ${xxx} at the same time. (${xxx} indicates that the value in the request will be inherited. For example, ${host} indicates the host in the URL to be redirected.)
	//
	// The values of protocol and port cannot be the same as those of the associated listener, and either host or path must be passed or their values cannot be ${xxx} at the same time.
	//
	// For shared load balancers, this parameter is unsupported. If it is passed, an error will be returned.
	//
	// This parameter is not available in eu-nl region. Please do not use it.
	RedirectUrlConfig RedirectUrlOptions `json:"redirect_url_config"`
	// Specifies the configuration of the page that will be returned. This parameter will take effect when enhance_l7policy_enable is set to true. If this parameter is passed and enhance_l7policy_enable is set to false, an error will be returned.
	//
	// This parameter is mandatory when action is set to FIXED_RESPONSE. It cannot be specified if the value of action is not FIXED_RESPONSE.
	//
	// For shared load balancers, this parameter is unsupported. If it is passed, an error will be returned.
	//
	// This parameter is not available in eu-nl region. Please do not use it.
	FixedResponseConfig FixedResponseOptions `json:"fixed_response_config"`
	// Specifies the time when the forwarding policy was added. The format is yyyy-MM-dd'T'HH:mm:ss'Z' (UTC time).
	//
	// This is a new field in this version, and it will not be returned for resources associated with existing dedicated load balancers and for resources associated with existing and new shared load balancers.
	CreatedAt string `json:"created_at"`
	// Specifies the time when the forwarding policy was updated. The format is yyyy-MM-dd'T'HH:mm:ss'Z' (UTC time).
	//
	// This is a new field in this version, and it will not be returned for resources associated with existing dedicated load balancers and for resources associated with existing and new shared load balancers.
	UpdatedAt string `json:"updated_at"`
}
