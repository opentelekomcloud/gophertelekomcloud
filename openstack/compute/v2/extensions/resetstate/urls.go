package resetstate

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
)

func actionURL(client *golangsdk.ServiceClient, id string) string {
	return client.ServiceURL("servers", id, "action")
}
