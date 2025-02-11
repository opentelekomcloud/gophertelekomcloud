package clusters

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ManagePublicAccessOpts struct {
	ClusterId string `json:"-"`
	Size      int    `json:"-" required:"true"`
}

// EnablePublicAccess function is used to enable public network access
func EnablePublicAccess(client *golangsdk.ServiceClient, opts ManagePublicAccessOpts) (string, error) {
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

	raw, err := client.Post(client.ServiceURL("clusters", opts.ClusterId, "public", "open"), b, nil, &golangsdk.RequestOpts{
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

type PublicAccessOpts struct {
	Eip Eip `json:"eip"`
}

type Eip struct {
	Bandwidth Bandwidth `json:"bandWidth" required:"true"`
}

type Bandwidth struct {
	Size int `json:"size" required:"true"`
}
