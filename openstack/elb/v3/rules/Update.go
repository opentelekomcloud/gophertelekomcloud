package rules

import "github.com/opentelekomcloud/gophertelekomcloud"

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
