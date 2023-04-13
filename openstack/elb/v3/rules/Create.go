package rules

import "github.com/opentelekomcloud/gophertelekomcloud"

type CreateOpts struct {
	// Specifies the conditions contained in a forwarding rule.
	// This parameter will take effect when enhance_l7policy_enable is set to true.
	// If conditions is specified, key and value will not take effect,
	// and the value of this parameter will contain all conditions configured for the forwarding rule.
	// The keys in the list must be the same, whereas each value must be unique.
	Conditions []Condition `json:"conditions,omitempty"`
	// Specifies the match content. The value can be one of the following:
	//
	//    HOST_NAME: A domain name will be used for matching.
	//    PATH: A URL will be used for matching.
	// If type is set to HOST_NAME, PATH, METHOD, or SOURCE_IP, only one forwarding rule can be created for each type.
	Type RuleType `json:"type" required:"true"`

	// Specifies how requests are matched and forwarded.
	//
	//    If type is set to HOST_NAME, this parameter can only be set to EQUAL_TO. Asterisks (*) can be used as wildcard characters.
	//    If type is set to PATH, this parameter can be set to REGEX, STARTS_WITH, or EQUAL_TO.
	CompareType CompareType `json:"compare_type" required:"true"`

	// Specifies the value of the match item. For example, if a domain name is used for matching, value is the domain name.
	//
	//    If type is set to HOST_NAME, the value can contain letters, digits, hyphens (-), and periods (.) and must
	//    start with a letter or digit. If you want to use a wildcard domain name, enter an asterisk (*) as the leftmost
	//    label of the domain name.
	//
	//    If type is set to PATH and compare_type to STARTS_WITH or EQUAL_TO, the value must start with a slash (/) and
	//    can contain only letters, digits, and special characters _~';@^-%#&$.*+?,=!:|/()[]{}
	Value string `json:"value" required:"true"`

	// Specifies the project ID.
	ProjectID string `json:"project_id,omitempty"`
}

type CreateOptsBuilder interface {
	ToRuleCreateMap() (map[string]interface{}, error)
}

func (opts CreateOpts) ToRuleCreateMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "rule")
}

func Create(client *golangsdk.ServiceClient, policyID string, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToRuleCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(client.ServiceURL("l7policies", policyID, "rules"), b, &r.Body, nil)
	return
}
