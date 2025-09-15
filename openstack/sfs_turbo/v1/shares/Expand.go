package shares

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type ExpandOpts struct {
	// Specifies the extend object.
	Extend ExtendOpts `json:"extend" required:"true"`
}

type ExtendOpts struct {
	// Specifies the post-expansion capacity (GB) of the shared file system.
	NewSize int `json:"new_size" required:"true"`
}

// Expand will expand a SFS Turbo based on the values in ExpandOpts.
func Expand(client *golangsdk.ServiceClient, shareId string, opts ExpandOpts) error {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return err
	}

	// POST /v1/{project_id}/sfs-turbo/shares/{share_id}/action
	_, err = client.Post(client.ServiceURL("sfs-turbo", "shares", shareId, "action"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{202},
	})
	return err
}
