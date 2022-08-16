package configurations

import "github.com/opentelekomcloud/gophertelekomcloud"

func Get(client *golangsdk.ServiceClient, id string) (r GetResult) {
	_, r.Err = client.Get(client.ServiceURL("scaling_configuration", id), &r.Body, nil)
	return
}
