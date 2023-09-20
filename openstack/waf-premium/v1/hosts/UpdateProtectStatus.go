package hosts

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type ProtectUpdateOpts struct {
	// WAF status of the protected domain name.
	// 0: The WAF protection is suspended.
	// WAF only forwards requests destined for the domain name and does not detect attacks.
	// 1: The WAF protection is enabled. WAF detects attacks based on the policy you configure.
	ProtectStatus int `json:"protect_status"`
}

func UpdateProtectStatus(client *golangsdk.ServiceClient, id string, opts ProtectUpdateOpts) (err error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return
	}

	// PUT /v1/{project_id}/premium-waf/host/{id}/protect_status
	_, err = client.Put(client.ServiceURL("premium-waf", "host", id, "protect_status"), b, nil, &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: map[string]string{"Content-Type": "application/json;charset=utf8"},
	})
	return
}
