package servers

import "github.com/opentelekomcloud/gophertelekomcloud"

// ConfirmResize confirms a previous resize operation on a server. See Resize() for more details.
func ConfirmResize(client *golangsdk.ServiceClient, id string) (err error) {
	_, err = client.Post(client.ServiceURL("servers", id, "action"), map[string]interface{}{"confirmResize": nil},
		nil, &golangsdk.RequestOpts{
			OkCodes: []int{201, 202, 204},
		})
	return
}
