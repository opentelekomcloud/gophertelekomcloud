package trigger

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

func DeleteTriggers(client *golangsdk.ServiceClient, funcURN string) (err error) {
	_, err = client.Delete(client.ServiceURL("fgs", "triggers", funcURN), nil)
	return
}
