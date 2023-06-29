package policies

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type UpdateOpts struct {
	// Specifies the administrative status of the forwarding policy. The default value is true.
	//
	// This parameter is unsupported. Please do not use it.
	AdminStateUp *bool `json:"admin_state_up,omitempty"`
	// Provides supplementary information about the forwarding policy.
	//
	// Minimum: 0
	//
	// Maximum: 255
	Description string `json:"description,omitempty"`
	// Specifies the forwarding policy name.
	//
	// Minimum: 0
	//
	// Maximum: 255
	Name string `json:"name,omitempty"`
	// Specifies the ID of the listener to which requests are redirected.
	//
	// Note:
	//
	// This parameter cannot be left blank or set to null when action is set to REDIRECT_TO_LISTENER.
	//
	// The listener's protocol must be HTTPS.
	//
	// A listener added to another load balancer is not allowed.
	//
	// This parameter cannot be passed in the API for adding or updating a forwarding policy if action is set to REDIRECT_TO_POOL.
	RedirectListenerId string `json:"redirect_listener_id,omitempty"`
	// Specifies the ID of the backend server group that requests will be forwarded to.
	//
	// This parameter is valid and mandatory only when action is set to REDIRECT_TO_POOL. The specified backend server group cannot be the default backend server group associated with the listener, or any backend server group associated with the forwarding policies of other listeners.
	//
	// This parameter cannot be specified when action is set to REDIRECT_TO_LISTENER.
	RedirectPoolId string `json:"redirect_pool_id,omitempty"`
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
	RedirectUrlConfig RedirectUrlOptions `json:"redirect_url_config,omitempty"`
	// Specifies the configuration of the page that will be returned. This parameter will take effect when enhance_l7policy_enable is set to true. If this parameter is passed and enhance_l7policy_enable is set to false, an error will be returned.
	//
	// This parameter is mandatory when action is set to FIXED_RESPONSE. It cannot be specified if the value of action is not FIXED_RESPONSE.
	//
	// For shared load balancers, this parameter is unsupported. If it is passed, an error will be returned.
	//
	// This parameter is not available in eu-nl region. Please do not use it.
	FixedResponseConfig FixedResponseOptions `json:"fixed_response_config,omitempty"`
	// Lists the forwarding rules in the forwarding policy.
	//
	// The list can contain a maximum of 10 forwarding rules (if conditions is specified, a condition is considered as a rule).
	//
	// If type is set to HOST_NAME, PATH, METHOD, or SOURCE_IP, only one forwarding rule can be created for each type.
	//
	// For details, see the description of l7rule.
	Rules []L7PolicyRuleOption `json:"rules,omitempty"`
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
	//
	// Minimum: 0
	//
	// Maximum: 10000
	Priority *int `json:"priority,omitempty"`
	// Specifies the configuration of the backend server group that the requests are forwarded to. This parameter is valid only when action is set to REDIRECT_TO_POOL.
	//
	// Note:
	//
	// If action is set to REDIRECT_TO_POOL, specify either redirect_pool_id or redirect_pools_config. If both are specified, only redirect_pools_config takes effect.
	//
	// This parameter cannot be specified when action is set to REDIRECT_TO_LISTENER. All configuration will be overwritten.
	RedirectPoolsConfig []RedirectPoolOptions `json:"redirect_pools_config,omitempty"`
}

func Update(client *golangsdk.ServiceClient, id string, opts UpdateOpts) (*L7Policy, error) {
	b, err := build.RequestBody(opts, "l7policy")
	if err != nil {
		return nil, err
	}

	raw, err := client.Put(client.ServiceURL("l7policies", id), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200, 201},
	})
	return extra(err, raw)
}
