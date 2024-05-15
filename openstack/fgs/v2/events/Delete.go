package events

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

func Delete(client *golangsdk.ServiceClient, funcURN, eventId string) (err error) {
	_, err = client.Delete(client.ServiceURL("fgs", "functions", funcURN, "events", eventId), nil)
	return
}
