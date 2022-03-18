package falsealarmmasking_rules

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack"
)

// CreateOptsBuilder allows extensions to add additional parameters to the
// Create request.
type CreateOptsBuilder interface {
	ToAlarmMaskingCreateMap() (map[string]interface{}, error)
}

// CreateOpts contains all the values needed to create a new falsealarmmasking rule.
type CreateOpts struct {
	Path    string `json:"path" required:"true"`
	EventID string `json:"rule" required:"true"`
}

// ToAlarmMaskingCreateMap builds a create request body from CreateOpts.
func (opts CreateOpts) ToAlarmMaskingCreateMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "")
}

// Create will create a new falsealarmmasking rule based on the values in CreateOpts.
func Create(c *golangsdk.ServiceClient, policyID string, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToAlarmMaskingCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	reqOpt := &golangsdk.RequestOpts{OkCodes: []int{200}}
	_, r.Err = c.Post(rootURL(c, policyID), b, &r.Body, reqOpt)
	return
}

// List retrieves falsealarmmasking rules.
func List(c *golangsdk.ServiceClient, policyID string) (r ListResult) {
	_, r.Err = c.Get(rootURL(c, policyID), &r.Body, openstack.StdRequestOpts())
	return
}

// Delete will permanently delete a particular falsealarmmasking rule based on its unique ID.
func Delete(c *golangsdk.ServiceClient, policyID, ruleID string) (r DeleteResult) {
	reqOpt := &golangsdk.RequestOpts{OkCodes: []int{204},
		MoreHeaders: openstack.StdRequestOpts().MoreHeaders}
	_, r.Err = c.Delete(resourceURL(c, policyID, ruleID), reqOpt)
	return
}
