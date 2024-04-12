package function

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
)

func UpdateStatus(client *golangsdk.ServiceClient, funcUrn, state string) (err error) {
	_, err = client.Put(client.ServiceURL("fgs", "functions", funcUrn, "collect", state), nil, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}
