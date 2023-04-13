package rules

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type RuleType string
type CompareType string

const (
	HostName RuleType = "HOST_NAME"
	Path     RuleType = "PATH"

	EqualTo    CompareType = "EQUAL_TO"
	Regex      CompareType = "REGEX"
	StartsWith CompareType = "STARTS_WITH"
)

type Condition struct {
	// Specifies the key of match item.
	Key string `json:"key"`

	// Specifies the value of the match item.
	Value string `json:"value" required:"true"`
}

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

func Get(client *golangsdk.ServiceClient, policyID, id string) (r GetResult) {
	_, r.Err = client.Get(client.ServiceURL("l7policies", policyID, "rules", id), &r.Body, nil)
	return
}

type UpdateOptsBuilder interface {
	ToUpdateRuleMap() (map[string]interface{}, error)
}

type UpdateOpts struct {
	CompareType CompareType `json:"compare_type,omitempty"`
	Value       string      `json:"value,omitempty"`
	Conditions  []Condition `json:"conditions,omitempty"`
}

func (opts UpdateOpts) ToUpdateRuleMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "rule")
}

func Update(client *golangsdk.ServiceClient, policyID, id string, opts UpdateOptsBuilder) (r UpdateResult) {
	b, err := opts.ToUpdateRuleMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Put(client.ServiceURL("l7policies", policyID, "rules", id), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200, 201},
	})
	return
}

type ListOptsBuilder interface {
	ToRuleListQuery() (string, error)
}

type ListOpts struct {
	ID          []string      `q:"id"`
	CompareType []CompareType `q:"compare_type"`
	Value       []string      `q:"value"`
	Type        []RuleType    `q:"type"`

	Limit       int    `q:"limit"`
	Marker      string `q:"marker"`
	PageReverse bool   `q:"page_reverse"`
}

func List(client *golangsdk.ServiceClient, policyID string, opts ListOptsBuilder) pagination.Pager {
	url := client.ServiceURL("l7policies", policyID, "rules")

	if opts != nil {
		q, err := opts.ToRuleListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += q
	}

	return pagination.NewPager(client, url, func(r pagination.PageResult) pagination.Page {
		return RulePage{PageWithInfo: pagination.NewPageWithInfo(r)}
	})
}

func Delete(client *golangsdk.ServiceClient, policyID, id string) (r DeleteResult) {
	_, r.Err = client.Delete(client.ServiceURL("l7policies", policyID, "rules", id), nil)
	return
}
