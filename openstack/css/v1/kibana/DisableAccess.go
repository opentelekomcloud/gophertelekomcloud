package kibana

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
)

func DisableAccess(client *golangsdk.ServiceClient, clusterId string) error {
	_, err := client.Put(client.ServiceURL("clusters", clusterId, "publickibana", "whitelist", "close"), nil, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
		MoreHeaders: map[string]string{
			"Content-Type": "application/json",
		},
	})

	return err
}
