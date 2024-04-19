package trigger

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

func Delete(client *golangsdk.ServiceClient, funcURN, triggerType, triggerId string) (err error) {
	_, err = client.Delete(client.ServiceURL("fgs", "triggers", funcURN, triggerType, triggerId), nil)
	return
}
