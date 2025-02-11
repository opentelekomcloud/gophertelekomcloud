package kibana

import (
	"strings"

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

	url := client.ServiceURL("clusters", clusterId, "publickibana", "whitelist", "update")
	convertedURL := strings.Replace(url, "v1.0", "v1.0/extend", 1)

	_, err = client.Post(convertedURL, b, nil, &golangsdk.RequestOpts{
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
