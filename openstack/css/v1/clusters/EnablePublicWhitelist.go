package clusters

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

// EnablePublicWhitelist function is used to enable public whitelist
func EnablePublicWhitelist(client *golangsdk.ServiceClient, clusterId, whitelist string) error {
	opts := struct {
		WhiteList string `json:"whiteList"`
	}{whitelist}

	b, err := build.RequestBody(opts, "")
	if err != nil {
		return err
	}

	_, err = client.Post(client.ServiceURL("clusters", clusterId, "public", "whitelist", "update"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return err
}
