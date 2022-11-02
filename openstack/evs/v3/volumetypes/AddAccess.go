package volumetypes

import "github.com/opentelekomcloud/gophertelekomcloud"

// AddAccessOptsBuilder allows extensions to add additional parameters to the
// AddAccess requests.
type AddAccessOptsBuilder interface {
	ToVolumeTypeAddAccessMap() (map[string]interface{}, error)
}

// AddAccessOpts represents options for adding access to a volume type.
type AddAccessOpts struct {
	// Project is the project/tenant ID to grant access.
	Project string `json:"project"`
}

// ToVolumeTypeAddAccessMap constructs a request body from AddAccessOpts.
func (opts AddAccessOpts) ToVolumeTypeAddAccessMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "addProjectAccess")
}

// AddAccess grants a tenant/project access to a volume type.
func AddAccess(client *golangsdk.ServiceClient, id string, opts AddAccessOptsBuilder) (r AddAccessResult) {
	b, err := opts.ToVolumeTypeAddAccessMap()
	if err != nil {
		r.Err = err
		return
	}
	resp, err := client.Post(client.ServiceURL("types", id, "action"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{202},
	})
	_, r.Header, r.Err = golangsdk.ParseResponse(resp, err)
	return
}
