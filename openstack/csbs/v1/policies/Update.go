package policies

import "github.com/opentelekomcloud/gophertelekomcloud"

// Update allows backup policies to be updated.
// call the Extract method on the UpdateResult.
func Update(c *golangsdk.ServiceClient, policyId string, opts UpdateOptsBuilder) (r UpdateResult) {
	b, err := opts.ToPoliciesUpdateMap()
	if err != nil {
		r.Err = err
		return
	}
	raw, err := c.Put(c.ServiceURL("policies", policyId), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}
