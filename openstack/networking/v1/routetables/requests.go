package routetables

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
)

// Update allows route tables to be updated
// func Update(c *golangsdk.ServiceClient, id string, opts UpdateOptsBuilder) (r UpdateResult) {
// 	b, err := opts.ToRouteTableUpdateMap()
// 	if err != nil {
// 		r.Err = err
// 		return
// 	}
// 	_, r.Err = c.Put(resourceURL(c, id), b, &r.Body, &golangsdk.RequestOpts{
// 		OkCodes: []int{200},
// 	})
// 	return
// }

// ActionOptsBuilder allows extensions to add additional parameters to the
// Action request: associate or disassociate subnets with a route table
type ActionOptsBuilder interface {
	ToRouteTableActionMap() (map[string]interface{}, error)
}

// ActionSubnetsOpts contains the subnets list that associate or disassociate with a route tabl
type ActionSubnetsOpts struct {
	Associate    []string `json:"associate,omitempty"`
	Disassociate []string `json:"disassociate,omitempty"`
}

// ActionOpts contains the values used when associating or disassociating subnets with a route table
type ActionOpts struct {
	Subnets ActionSubnetsOpts `json:"subnets" required:"true"`
}

// ToRouteTableActionMap builds an update body based on UpdateOpts.
func (opts ActionOpts) ToRouteTableActionMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "routetable")
}

// Action will associate or disassociate subnets with a particular route table based on its unique ID
// func Action(c *golangsdk.ServiceClient, id string, opts ActionOptsBuilder) (r ActionResult) {
// 	b, err := opts.ToRouteTableActionMap()
// 	if err != nil {
// 		r.Err = err
// 		return
// 	}
//
// 	_, r.Err = c.Post(actionURL(c, id), b, &r.Body, nil)
// 	return
// }
