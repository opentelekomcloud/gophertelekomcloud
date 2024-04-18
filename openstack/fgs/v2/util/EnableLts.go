package util

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
)

func EnableFuncLts(client *golangsdk.ServiceClient) error {
	_, err := client.Post(client.ServiceURL("fgs", "functions", "enable-lts-logs"), nil, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return err
}
