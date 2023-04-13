package policies

import (
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/elb/v3/rules"
)

type Action string

const (
	ActionRedirectToPool     Action = "REDIRECT_TO_POOL"
	ActionRedirectToListener Action = "REDIRECT_TO_LISTENER"
	ActionFixedResponse      Action = "FIXED_RESPONSE"
	ActionUrlRedirect        Action = "REDIRECT_TO_URL"
)

type Rule struct {
	// Specifies the match content. The value can be one of the following:
	//
	//    HOST_NAME: A domain name will be used for matching.
	//    PATH: A URL will be used for matching.
	// If type is set to HOST_NAME, PATH, METHOD, or SOURCE_IP, only one forwarding rule can be created for each type.
	Type rules.RuleType `json:"type" required:"true"`

	// Specifies how requests are matched with the domain name or URL.
	//
	// If type is set to HOST_NAME, this parameter can only be set to EQUAL_TO (exact match).
	//
	// If type is set to PATH, this parameter can be set to REGEX (regular expression match), STARTS_WITH (prefix match), or EQUAL_TO (exact match).
	CompareType rules.CompareType `json:"compare_type" required:"true"`

	// Specifies the value of the match item. For example, if a domain name is used for matching, value is the domain name.
	//
	//    If type is set to HOST_NAME, the value can contain letters, digits, hyphens (-), and periods (.) and must start with a letter or digit.
	//    If you want to use a wildcard domain name, enter an asterisk (*) as the leftmost label of the domain name.
	//
	//    If type is set to PATH and compare_type to STARTS_WITH or EQUAL_TO, the value must start with a slash (/)
	//    and can contain only letters, digits, and special characters _~';@^-%#&$.*+?,=!:|/()[]{}
	Value string `json:"value" required:"true"`
}

type RedirectPoolOptions struct {
	// Specifies the ID of the backend server group.
	PoolId string `json:"pool_id" required:"true"`

	// Specifies the weight of the backend server group. The value ranges from 0 to 100.
	Weight string `json:"weight" required:"true"`
}

type FixedResponseOptions struct {
	// Specifies the fixed HTTP status code configured in the forwarding rule.
	// The value can be any integer in the range of 200-299, 400-499, or 500-599.
	StatusCode string `json:"status_code" required:"true"`

	// Specifies the format of the response body.
	//
	// The value can be text/plain, text/css, text/html, application/javascript, or application/json.
	ContentType string `json:"content_type"`

	// Specifies the content of the response message body.
	MessageBody string `json:"message_body"`
}

type RedirectUrlOptions struct {
	// Specifies the protocol for redirection.
	// The value can be HTTP, HTTPS, or ${protocol}. The default value is ${protocol}, indicating that the protocol of the request will be used.
	Protocol string `json:"protocol"`

	// Specifies the host name that requests are redirected to.
	// The value can contain only letters, digits, hyphens (-), and periods (.) and must start with a letter or digit.
	// The default value is ${host}, indicating that the host of the request will be used.
	Host string `json:"host"`

	// Specifies the port that requests are redirected to.
	// The default value is ${port}, indicating that the port of the request will be used.
	Port string `json:"port"`

	// Specifies the path that requests are redirected to. The default value is ${path}, indicating that the path of the request will be used.
	//
	// The value can contain only letters, digits, and special characters _~';@^- %#&$.*+?,=!:|/()[]{} and must start with a slash (/).
	Path string `json:"path"`

	// Specifies the query string set in the URL for redirection. The default value is ${query}, indicating that the query string of the request will be used.
	//
	// For example, in the URL https://www.xxx.com:8080/elb?type=loadbalancer, ${query} indicates type=loadbalancer.
	// If this parameter is set to ${query}&name=my_name, the URL will be redirected to https://www.xxx.com:8080/elb?type=loadbalancer&name=my_name.
	//
	// The value is case-sensitive and can contain only letters, digits, and special characters !$&'()*+,-./:;=?@^_`
	Query string `json:"query"`

	// Specifies the status code returned after the requests are redirected.
	//
	// The value can be 301, 302, 303, 307, or 308.
	StatusCode string `json:"status_code" required:"true"`
}
