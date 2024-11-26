package clusters

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
)

type UpdateIpOpts struct {
	Action    string `json:"action" required:"true"`
	Spec      IpSpec `json:"spec,omitempty"`
	ElasticIp string `json:"elasticIp"`
}

type IpSpec struct {
	ID string `json:"id" required:"true"`
}

type UpdateIpOptsBuilder interface {
	ToMasterIpUpdateMap() (map[string]interface{}, error)
}

func (opts UpdateIpOpts) ToMasterIpUpdateMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "spec")
}

// Update the access information of a specified cluster.
func UpdateMasterIp(c *golangsdk.ServiceClient, id string, opts UpdateIpOptsBuilder) (r UpdateIpResult) {
	b, err := opts.ToMasterIpUpdateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = c.Put(masterIpURL(c, id), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}
