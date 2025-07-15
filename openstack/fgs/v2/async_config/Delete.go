package async_config

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

func Delete(client *golangsdk.ServiceClient, funcURN string) (err error) {
	_, err = client.Delete(client.ServiceURL("fgs", "functions", funcURN, "async-invoke-config"), nil)
	return
}
