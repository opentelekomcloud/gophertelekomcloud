package rules

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/elb/v3/policies"
)

type UpdateOpts struct {
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
	CompareType string `json:"compare_type,omitempty"`
	// Specifies whether reverse matching is supported. The value can be true or false.
	//
	// This parameter is unsupported. Please do not use it.
	Invert *bool `json:"invert,omitempty"`
	// Specifies the key of the match item. For example, if an HTTP header is used for matching, key is the name of the HTTP header parameter.
	//
	// This parameter is unsupported. Please do not use it.
	//
	// Minimum: 0
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
	Value string `json:"value,omitempty"`
	// Specifies the matching conditions of the forwarding rule. This parameter will take effect when enhance_l7policy_enable is set to .true. If conditions is specified, key and value will not take effect, and the value of this parameter will contain all conditions configured for the forwarding rule. The keys in the list must be the same, whereas each value must be unique.
	//
	// This parameter is not available in eu-nl region. Please do not use it.
	Conditions []policies.RuleCondition `json:"conditions,omitempty"`
}

func Update(client *golangsdk.ServiceClient, policyID, id string, opts UpdateOpts) (*ForwardingRule, error) {
	b, err := build.RequestBody(opts, "rule")
	if err != nil {
		return nil, err
	}

	raw, err := client.Put(client.ServiceURL("l7policies", policyID, "rules", id), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200, 201},
	})
	return extra(err, raw)
}
