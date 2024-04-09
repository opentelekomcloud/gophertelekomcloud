package function

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
)

func Delete(client *golangsdk.ServiceClient, funcURN string) (err error) {
	_, err = client.Delete(client.ServiceURL("fgs", "functions", funcURN), nil)
	return
}
