package configurations

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
)

// ApplyOptsBuilder allows extensions to update instance templates to the
// Apply request.
type ApplyOptsBuilder interface {
	ToInstanceTemplateApplyMap() (map[string]interface{}, error)
}

// ApplyOpts contains all the instances needed to apply another template.
type ApplyOpts struct {
	// Specifies the DB instance ID list object.
	InstanceIDs []string `json:"instance_ids" required:"true"`
}

// ToInstanceTemplateApplyMap builds an apply request body from ApplyOpts.
func (opts ApplyOpts) ToInstanceTemplateApplyMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "")
}

// Apply is used to apply a parameter template to one or more DB instances.
func Apply(client *golangsdk.ServiceClient, id string, opts ApplyOptsBuilder) (r ApplyResult) {
	b, err := opts.ToInstanceTemplateApplyMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Put(client.ServiceURL("configurations", id, "apply"), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200, 201, 202},
	})
	return
}
