package rules

import (
	"net/http"

	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/elb/v3/policies"
)

func Get(client *golangsdk.ServiceClient, policyID, id string) (*ForwardingRule, error) {
	raw, err := client.Get(client.ServiceURL("l7policies", policyID, "rules", id), nil, nil)
	return extra(err, raw)
}

func extra(err error, raw *http.Response) (*ForwardingRule, error) {
	if err != nil {
		return nil, err
	}

	var rule ForwardingRule
	err = extract.IntoStructPtr(raw.Body, &rule, "rule")
	return &rule, err
}

type ForwardingRule struct {
	// Specifies the administrative status of the forwarding rule. The default value is true.
	//
	// This parameter is unsupported. Please do not use it.
	AdminStateUp *bool `json:"admin_state_up"`
	// Specifies how requests are matched with the domain name or URL.
	//
	// If type is set to HOST_NAME, this parameter can only be set to EQUAL_TO.
	//
	// If type is set to PATH, the value can be REGEX, STARTS_WITH, or EQUAL_TO.
	CompareType string `json:"compare_type"`
	// Specifies the key of the match content. This parameter will not take effect if type is set to HOST_NAME or PATH.
	//
	// Minimum: 1
	//
	// Maximum: 255
	Key string `json:"key"`
	// Specifies the project ID.
	ProjectId string `json:"project_id"`
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
	// Note:
	//
	// If type is set to HOST_NAME, PATH, METHOD, and SOURCE_IP, only one forwarding rule can be created for each type. If type is set to HEADER and QUERY_STRING, multiple forwarding rules can be created for each type.
	Type string `json:"type"`
	// Specifies the value of the match item. For example, if a domain name is used for matching, value is the domain name. This parameter will take effect only when conditions is left blank.
	//
	// If type is set to HOST_NAME, the value can contain letters, digits, hyphens (-), and periods (.) and must start with a letter or digit. If you want to use a wildcard domain name, enter an asterisk (*) as the leftmost label of the domain name.
	//
	// If type is set to PATH and compare_type to STARTS_WITH or EQUAL_TO, the value must start with a slash (/) and can contain only letters, digits, and special characters _~';@^-%#&$.*+?,=!:|/()[]{}
	//
	// If type is set to METHOD, SOURCE_IP, HEADER, or QUERY_STRING, this parameter will not take effect, and condition_pair will be used to specify the key and value.
	//
	// Minimum: 1
	//
	// Maximum: 128
	Value string `json:"value"`
	// Specifies the provisioning status of the forwarding rule.
	//
	// The value can only be ACTIVE (default), PENDING_CREATE, or ERROR.
	//
	// This parameter is unsupported. Please do not use it.
	ProvisioningStatus string `json:"provisioning_status"`
	// Specifies whether reverse matching is supported. The value is fixed at false. This parameter can be updated but will not take effect.
	Invert *bool `json:"invert"`
	// Specifies the forwarding policy ID.
	Id string `json:"id"`
	// Specifies the matching conditions of the forwarding rule. This parameter will take effect when enhance_l7policy_enable is set to .true. If conditions is specified, key and value will not take effect, and the value of this parameter will contain all conditions configured for the forwarding rule. The keys in the list must be the same, whereas each value must be unique.
	//
	// This parameter is not available in eu-nl region. Please do not use it.
	Conditions []policies.RuleCondition `json:"conditions"`
	// Specifies the time when the forwarding rule was added. The format is yyyy-MM-dd'T'HH:mm:ss'Z' (UTC time).
	//
	// This is a new field in this version, and it will not be returned for resources associated with existing dedicated load balancers and for resources associated with existing and new shared load balancers.
	CreatedAt string `json:"created_at"`
	// Specifies the time when the forwarding rule was updated. The format is yyyy-MM-dd'T'HH:mm:ss'Z' (UTC time).
	//
	// This is a new field in this version, and it will not be returned for resources associated with existing dedicated load balancers and for resources associated with existing and new shared load balancers.
	UpdatedAt string `json:"updated_at"`
}
