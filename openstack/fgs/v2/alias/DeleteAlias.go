package alias

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

func Delete(client *golangsdk.ServiceClient, funcURN, aliasName string) (err error) {
	_, err = client.Delete(client.ServiceURL("fgs", "functions", funcURN, "aliases", aliasName), nil)
	return
}
