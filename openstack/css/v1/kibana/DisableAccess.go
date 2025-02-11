package kibana

import (
	"strings"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
)

func DisableAccess(client *golangsdk.ServiceClient, clusterId string) error {
	url := client.ServiceURL("clusters", clusterId, "publickibana", "whitelist", "close")
	convertedURL := strings.Replace(url, "v1.0", "v1.0/extend", 1)

	_, err := client.Put(convertedURL, nil, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
		MoreHeaders: map[string]string{
			"Content-Type": "application/json",
		},
	})

	return err
}
