package kibana

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

func UpdateAccess(client *golangsdk.ServiceClient, clusterId, whitelist string) error {
	config := Config{
		WhiteList: whitelist,
	}
	b, err := build.RequestBody(config, "")
	if err != nil {
		return err
	}

	_, err = client.Post(client.ServiceURL("clusters", clusterId, "publickibana", "whitelist", "update"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
		MoreHeaders: map[string]string{
			"Content-Type": "application/json",
		},
	})

	return err
}

type Config struct {
	WhiteList string `json:"whiteList"`
}
