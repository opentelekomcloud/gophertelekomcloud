package upgrade

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type GetAvailableVersionOpts struct {
	InstanceId string `json:"-"`
}

type AvailableVersions struct {
	AvailableVersions []string `json:"available_versions"`
}

func GetAvailableVersion(client *golangsdk.ServiceClient, opts GetAvailableVersionOpts) (*AvailableVersions, error) {
	raw, err := client.Get(client.ServiceURL("instances", opts.InstanceId, "major-version", "available-version"), nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res AvailableVersions
	err = extract.Into(raw.Body, &res)
	return &res, err
}
