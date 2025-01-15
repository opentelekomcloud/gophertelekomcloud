package cert

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

func Delete(client *golangsdk.ServiceClient, certID string) (err error) {
	_, err = client.Delete(client.ServiceURL("apigw", "certificates", certID), nil)
	return
}
