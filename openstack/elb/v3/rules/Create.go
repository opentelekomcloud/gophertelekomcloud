package rules

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/elb/v3/policies"
)

type CreateOpts struct {
	// Specifies the administrative status of the forwarding rule. The default value is true.
	//
	// This parameter is unsupported. Please do not use it.
	AdminStateUp *bool `json:"admin_state_up,omitempty"`
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
	// Specifies the key of match content. For example, if the request header is used for forwarding, key is the request header.
	//
	// This parameter is unsupported. Please do not use it.
	//
	// Minimum: 0
	//
	// Maximum: 255
	Key string `json:"key,omitempty"`
	// Specifies the value of the match content. For example, if a domain name is used for matching, value is the domain name. This parameter is valid only when conditions is left blank.
	//
	// If type is set to HOST_NAME, the value can contain letters, digits, hyphens (-), periods (.), and asterisks (*) and must start with a letter or digit. If you want to use a wildcard domain name, enter an asterisk (*) as the leftmost label of the domain name.
	//
	// If type is set to PATH and compare_type to STARTS_WITH or EQUAL_TO, the value can contain only letters, digits, and special characters _~';@^-%#&$.*+?,=!:|/()[]{}
	//
	// If type is set to METHOD, SOURCE_IP, HEADER, or QUERY_STRING, this parameter will not take effect, and conditions will be used to specify the key and value.
	//
	// Minimum: 1
	//
	// Maximum: 128
	Value string `json:"value" required:"true"`
	// Specifies the project ID.
	//
	// Minimum: 32
	//
	// Maximum: 32
	ProjectId string `json:"project_id,omitempty"`
	// Specifies the match content. The value can be one of the following:
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
	// SOURCE_IP: The source IP address will be used for matching. Note: If type is set to HOST_NAME, PATH, METHOD, or SOURCE_IP, only one forwarding rule can be created for each type.
	Type string `json:"type" required:"true"`
	// Specifies whether reverse matching is supported. The value can be true or false (default).
	//
	// This parameter is unsupported. Please do not use it.
	Invert *bool `json:"invert,omitempty"`
	// Specifies the matching conditions of the forwarding rule. This parameter is available only when enhance_l7policy_enable is set to true.
	//
	// If conditions is specified, parameters key and value will not take effect, and the conditions value will contain all conditions configured for the forwarding rule. The keys in the list must be the same, whereas each value must be unique.
	//
	// This parameter is not available in eu-nl region. Please do not use it.
	Conditions []policies.RuleCondition `json:"conditions,omitempty"`
}

func Create(client *golangsdk.ServiceClient, policyID string, opts CreateOpts) (*ForwardingRule, error) {
	b, err := build.RequestBody(opts, "rule")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL("l7policies", policyID, "rules"), b, nil, nil)
	return extra(err, raw)
}
