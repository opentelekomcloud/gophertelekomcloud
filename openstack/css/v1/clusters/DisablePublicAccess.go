package clusters

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// DisablePublicAccess function is used to disable public network access
func DisablePublicAccess(client *golangsdk.ServiceClient, opts ManagePublicAccessOpts) (string, error) {
	rawOpts := PublicAccessOpts{
		Eip: Eip{
			Bandwidth: Bandwidth{
				Size: opts.Size,
			},
		},
	}

	b, err := build.RequestBody(rawOpts, "")
	if err != nil {
		return "", err
	}

	raw, err := client.Put(client.ServiceURL("clusters", opts.ClusterId, "public", "close"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return "", err
	}

	var res struct {
		Action string `json:"action"`
	}
	err = extract.IntoStructPtr(raw.Body, &res, "")
	return res.Action, err
}
