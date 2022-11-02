package volumetypes

import "github.com/opentelekomcloud/gophertelekomcloud"

// RemoveAccessOptsBuilder allows extensions to add additional parameters to the
// RemoveAccess requests.
type RemoveAccessOptsBuilder interface {
	ToVolumeTypeRemoveAccessMap() (map[string]interface{}, error)
}

// RemoveAccessOpts represents options for removing access to a volume type.
type RemoveAccessOpts struct {
	// Project is the project/tenant ID to remove access.
	Project string `json:"project"`
}

// ToVolumeTypeRemoveAccessMap constructs a request body from RemoveAccessOpts.
func (opts RemoveAccessOpts) ToVolumeTypeRemoveAccessMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "removeProjectAccess")
}

// RemoveAccess removes/revokes a tenant/project access to a volume type.
func RemoveAccess(client *golangsdk.ServiceClient, id string, opts RemoveAccessOptsBuilder) (r RemoveAccessResult) {
	b, err := opts.ToVolumeTypeRemoveAccessMap()
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
