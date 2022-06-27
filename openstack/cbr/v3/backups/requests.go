package backups

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

func Get(client *golangsdk.ServiceClient, id string) (r GetResult) {
	_, r.Err = client.Get(singleURL(client, id), &r.Body, nil)
	return
}
