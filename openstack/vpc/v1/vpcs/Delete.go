package vpcs

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

func Delete(client *golangsdk.ServiceClient, id string) error {
	_, err := client.Delete(client.ServiceURL(client.ProjectID, "vpcs", id), nil)
	return err
}
