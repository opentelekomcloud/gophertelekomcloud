package proxy

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ChangeFlavorOpts struct {
	InstanceID string `json:"-"`
	ProxyID    string `json:"-"`
	FlavorRef  string `json:"flavor_ref" required:"true"`
}

func ChangeFlavor(client *golangsdk.ServiceClient, opts ChangeFlavorOpts) (string, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return "", err
	}

	raw, err := client.Put(client.ServiceURL("instances", opts.InstanceID, "proxy", opts.ProxyID, "flavor"), b, nil, nil)
	if err != nil {
		return "", err
	}

	var res JobId
	err = extract.Into(raw.Body, &res)
	return res.JobId, err
}
