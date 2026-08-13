package clusters

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
)

func Delete(client *golangsdk.ServiceClient, id string) error {
	_, err := client.Delete(client.ServiceURL("clusters", id), nil)
	return err
}
