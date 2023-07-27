package policies

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type CreateOpts struct {
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
	// If action is set to REDIRECT_TO_POOL, the listener's protocol must be HTTP or HTTPS.
	//
	// If action is set to REDIRECT_TO_LISTENER, the listener's protocol must be HTTP.
	//
	// Minimum: 1
	//
	// Maximum: 255
	Action string `json:"action" required:"true"`
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
	// Specifies the ID of the listener to which the forwarding policy is added.
	//
	// If action is set to REDIRECT_TO_POOL, the forwarding policy can be added to an HTTP or HTTPS listener.
	//
	// If action is set to REDIRECT_TO_LISTENER, the forwarding policy can be added to an HTTP listener.
	ListenerId string `json:"listener_id" required:"true"`
	// Specifies the forwarding policy name.
	//
	// Minimum: 0
	//
	// Maximum: 255
	Name string `json:"name,omitempty"`
	// Specifies the forwarding policy priority. The value cannot be updated.
	//
	// This parameter is unsupported. Please do not use it.
	//
	// Minimum: 1
	//
	// Maximum: 100
	Position *int `json:"position,omitempty"`
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
	// Specifies the ID of the project where the forwarding policy is used.
	//
	// Minimum: 1
	//
	// Maximum: 32
	ProjectId string `json:"project_id,omitempty"`
	// Specifies the ID of the listener to which requests are redirected. This parameter is mandatory when action is set to REDIRECT_TO_LISTENER.
	//
	// Note:
	//
	// The listener's protocol must be HTTPS.
	//
	// A listener added to another load balancer is not allowed.
	//
	// This parameter cannot be passed in the API for adding or updating a forwarding policy if action is set to REDIRECT_TO_POOL.
	//
	// This parameter is unsupported for shared load balancers.
	RedirectListenerId string `json:"redirect_listener_id,omitempty"`
	// Specifies the ID of the backend server group to which the requests are forwarded. This parameter is valid only when action is set to REDIRECT_TO_POOL. Note:
	//
	// If action is set to REDIRECT_TO_POOL, specify either redirect_pool_id or redirect_pools_config. If both are specified, only redirect_pools_config takes effect.
	//
	// This parameter cannot be specified when action is set to REDIRECT_TO_LISTENER.
	RedirectPoolId string `json:"redirect_pool_id,omitempty"`
	// Specifies the configuration of the backend server group that the requests are forwarded to. This parameter is valid only when action is set to REDIRECT_TO_POOL.
	//
	// Note:
	//
	// If action is set to REDIRECT_TO_POOL, specify either redirect_pool_id or redirect_pools_config. If both are specified, only redirect_pools_config takes effect.
	//
	// This parameter cannot be specified when action is set to REDIRECT_TO_LISTENER.
	RedirectPoolsConfig []RedirectPoolOptions `json:"redirect_pools_config,omitempty"`
	// Specifies the URL to which requests are forwarded.
	//
	// Format: protocol://host:port/path?query
	//
	// Minimum: 1
	//
	// Maximum: 255
	RedirectUrl string `json:"redirect_url,omitempty"`
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
	// If type is set to HOST_NAME, PATH, METHOD, or SOURCE_IP, only one forwarding rule can be created for each type. Note:
	//
	// The entire list will be replaced if you update it.
	//
	// If the action of 17policy is set to Redirect to another listener, 17rule cannot be created.
	Rules []L7PolicyRuleOption `json:"rules,omitempty"`
}

type FixedResponseOptions struct {
	// Specifies the fixed HTTP status code configured in the forwarding rule. The value can be any integer in the range of 200-299, 400-499, or 500-599.
	//
	// Minimum: 1
	//
	// Maximum: 16
	StatusCode string `json:"status_code" required:"true"`
	// Specifies the format of the response body.
	//
	// The value can be text/plain, text/css, text/html, application/javascript, or application/json.
	//
	// Default: text/plain
	//
	// Minimum: 0
	//
	// Maximum: 32
	ContentType string `json:"content_type,omitempty"`
	// Specifies the content of the response message body.
	//
	// Minimum: 0
	//
	// Maximum: 1024
	MessageBody string `json:"message_body,omitempty"`
}

type RedirectPoolOptions struct {
	// Specifies the ID of the backend server group.
	PoolId string `json:"pool_id" required:"true"`
	// Specifies the weight of the backend server group. The value ranges from 0 to 100.
	Weight *int `json:"weight" required:"true"`
}

type RedirectUrlOptions struct {
	// Specifies the protocol for redirection.
	//
	// The value can be HTTP, HTTPS, or ${protocol}. The default value is ${protocol}, indicating that the protocol of the request will be used.
	//
	// Default: ${protocol}
	//
	// Minimum: 1
	//
	// Maximum: 36
	Protocol string `json:"protocol,omitempty"`
	// Specifies the host name that requests are redirected to. The value can contain only letters, digits, hyphens (-), and periods (.) and must start with a letter or digit. The default value is ${host}, indicating that the host of the request will be used.
	//
	// Default: ${host}
	//
	// Minimum: 1
	//
	// Maximum: 128
	Host string `json:"host,omitempty"`
	// Specifies the port that requests are redirected to. The default value is ${port}, indicating that the port of the request will be used.
	//
	// Default: ${port}
	//
	// Minimum: 1
	//
	// Maximum: 16
	Port string `json:"port,omitempty"`
	// Specifies the path that requests are redirected to. The default value is ${path}, indicating that the path of the request will be used.
	//
	// The value can contain only letters, digits, and special characters _~';@^- %#&$.*+?,=!:|/()[]{} and must start with a slash (/).
	//
	// Default: ${path}
	//
	// Minimum: 1
	//
	// Maximum: 128
	Path string `json:"path,omitempty"`
	// Specifies the query string set in the URL for redirection. The default value is ${query}, indicating that the query string of the request will be used.
	//
	// For example, in the URL https://www.xxx.com:8080/elb?type=loadbalancer, ${query} indicates type=loadbalancer. If this parameter is set to ${query}&name=my_name, the URL will be redirected to https://www.xxx.com:8080/elb?type=loadbalancer&name=my_name.
	//
	// The value is case-sensitive and can contain only letters, digits, and special characters !$&'()*+,-./:;=?@^_`
	//
	// Default: ${query}
	//
	// Minimum: 0
	//
	// Maximum: 128
	Query string `json:"query,omitempty"`
	// Specifies the status code returned after the requests are redirected.
	//
	// The value can be 301, 302, 303, 307, or 308.
	//
	// Minimum: 1
	//
	// Maximum: 16
	StatusCode string `json:"status_code" required:"true"`
}

type L7PolicyRuleOption struct {
	// Specifies the administrative status of the forwarding rule. The value can be true or false, and the default value is true.
	//
	// This parameter is unsupported. Please do not use it.
	//
	// Default: true
	AdminStateUp *bool `json:"admin_state_up,omitempty"`
	// Specifies the type of the forwarding rule. The value can be one of the following:
	//
	// HOST_NAME: A domain name will be used for matching.
	//
	// PATH: A URL will be used for matching.
	//
	// METHOD: An HTTP request method will be used for matching.
	//
	// HEADER: The request header will be used for matching.
	//
	// QUERY_STRING: A query string will be used for matching.
	//
	// SOURCE_IP: The source IP address will be used for matching.
	//
	// If type is set to HOST_NAME, PATH, METHOD, and SOURCE_IP, only one forwarding rule can be created for each type. If type is set to HEADER and QUERY_STRING, multiple forwarding rules can be created for each type.
	Type string `json:"type" required:"true"`
	// Specifies how requests are matched with the forwarding rule. Values:
	//
	// EQUAL_TO: exact match.
	//
	// REGEX: regular expression match
	//
	// STARTS_WITH: prefix match
	//
	// Note:
	//
	// If type is set to HOST_NAME, the value can only be EQUAL_TO, and asterisks (*) can be used as wildcard characters.
	//
	// If type is set to PATH, the value can be REGEX, STARTS_WITH, or EQUAL_TO.
	//
	// If type is set to METHOD or SOURCE_IP, the value can only be EQUAL_TO.
	//
	// If type is set to HEADER or QUERY_STRING, the value can only be EQUAL_TO, asterisks (*) and question marks (?) can be used as wildcard characters.
	CompareType string `json:"compare_type" required:"true"`
	// Specifies whether reverse matching is supported. The value can be true or false, and the default value is false.
	//
	// This parameter is unsupported. Please do not use it.
	//
	// Default: false
	Invert *bool `json:"invert,omitempty"`
	// Specifies the key of the match item. For example, if an HTTP header is used for matching, key is the name of the HTTP header parameter.
	//
	// This parameter is unsupported. Please do not use it.
	//
	// Minimum: 1
	//
	// Maximum: 255
	Key string `json:"key,omitempty"`
	// Specifies the value of the match item. For example, if a domain name is used for matching, value is the domain name. This parameter will take effect only when conditions is left blank.
	//
	// If type is set to HOST_NAME, the value can contain letters, digits, hyphens (-), and periods (.) and must start with a letter or digit. If you want to use a wildcard domain name, enter an asterisk (*) as the leftmost label of the domain name.
	//
	// If type is set to PATH and compare_type to STARTS_WITH or EQUAL_TO, the value must start with a slash (/) and can contain only letters, digits, and special characters _~';@^-%#&$.*+?,=!:|/()[]{}
	//
	// If type is set to METHOD, SOURCE_IP, HEADER, or QUERY_STRING, this parameter will not take effect, and conditions will be used to specify the key and value.
	//
	// Minimum: 1
	//
	// Maximum: 128
	Value string `json:"value" required:"true"`
	// Specifies the conditions contained in a forwarding rule. This parameter will take effect when enhance_l7policy_enable is set to true.
	//
	// If conditions is specified, key and value will not take effect, and the value of this parameter will contain all conditions configured for the forwarding rule. The keys in the list must be the same, whereas each value must be unique.
	//
	// This parameter is not available in eu-nl region. Please do not use it.
	Conditions []RuleCondition `json:"conditions,omitempty"`
}

type RuleCondition struct {
	// Specifies the key of match item.
	//
	// If type is set to HOST_NAME, PATH, METHOD, or SOURCE_IP, this parameter is left blank.
	//
	// If type is set to HEADER, key indicates the name of the HTTP header parameter. The value can contain 1 to 40 characters, including letters, digits, hyphens (-), and underscores (_).
	//
	// If type is set to QUERY_STRING, key indicates the name of the query parameter. The value is case sensitive and can contain 1 to 128 characters. Spaces, square brackets ([ ]), curly brackets ({ }), angle brackets (< >), backslashes (), double quotation marks (" "), pound signs (#), ampersands (&), vertical bars (|), percent signs (%), and tildes (~) are not supported.
	//
	// All keys in the conditions list in the same rule must be the same.
	//
	// Minimum: 1
	//
	// Maximum: 128
	Key string `json:"key,omitempty"`
	// Specifies the value of the match item.
	//
	// If type is set to HOST_NAME, key is left blank, and value indicates the domain name, which can contain 1 to 128 characters, including letters, digits, hyphens (-), periods (.), and asterisks (*), and must start with a letter, digit, or asterisk (*). If you want to use a wildcard domain name, enter an asterisk (*) as the leftmost label of the domain name.
	//
	// If type is set to PATH, key is left blank, and value indicates the request path, which can contain 1 to 128 characters. If compare_type is set to STARTS_WITH or EQUAL_TO for the forwarding rule, the value must start with a slash (/) and can contain only letters, digits, and special characters _~';@^-%#&$.*+?,=!:|/()[]{}
	//
	// If type is set to HEADER, key indicates the name of the HTTP header parameter, and value indicates the value of the HTTP header parameter. The value can contain 1 to 128 characters. Asterisks (*) and question marks (?) are allowed, but spaces and double quotation marks are not allowed. An asterisk can match zero or more characters, and a question mark can match 1 character.
	//
	// If type is set to QUERY_STRING, key indicates the name of the query parameter, and value indicates the value of the query parameter. The value is case sensitive and can contain 1 to 128 characters. Spaces, square brackets ([ ]), curly brackets ({ }), angle brackets (< >), backslashes (), double quotation marks (" "), pound signs (#), ampersands (&), vertical bars (|), percent signs (%), and tildes (~) are not supported. Asterisks (*) and question marks (?) are allowed. An asterisk can match zero or more characters, and a question mark can match 1 character.
	//
	// If type is set to METHOD, key is left blank, and value indicates the HTTP method. The value can be GET, PUT, POST, DELETE, PATCH, HEAD, or OPTIONS.
	//
	// If type is set to SOURCE_IP, key is left blank, and value indicates the source IP address of the request. The value is an IPv4 or IPv6 CIDR block, for example, 192.168.0.2/32 or 2049::49/64.
	//
	// All keys in the conditions list in the same rule must be the same.
	//
	// Minimum: 1
	//
	// Maximum: 128
	Value string `json:"value" required:"true"`
}

func Create(client *golangsdk.ServiceClient, opts CreateOpts) (*L7Policy, error) {
	b, err := build.RequestBody(opts, "l7policy")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL("l7policies"), b, nil, nil)
	return extra(err, raw)
}
