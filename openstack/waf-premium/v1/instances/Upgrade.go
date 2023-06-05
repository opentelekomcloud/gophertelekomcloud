package instances

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type UpgradeOpts struct {
	// Operation name.
	// upgrade: Upgrade the	software version of the dedicated WAF engine
	Action string `json:"action"`
}

// Upgrade will upgrade instance software.
func Upgrade(client *golangsdk.ServiceClient, id string) (*Instance, error) {
	b, err := build.RequestBody(UpgradeOpts{
		Action: "upgrade",
	}, "")
	if err != nil {
		return nil, err
	}

	// POST /v1/{project_id}/premium-waf/instance
	raw, err := client.Post(client.ServiceURL("instance", id, "action"), b,
		nil, &golangsdk.RequestOpts{
			OkCodes: []int{200},
		})
	if err != nil {
		return nil, err
	}

	var res Instance
	return &res, extract.Into(raw.Body, &res)
}
