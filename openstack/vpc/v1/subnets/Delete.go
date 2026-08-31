package subnets

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

func Delete(client *golangsdk.ServiceClient, vpcID, id string) error {
	_, err := client.Delete(client.ServiceURL(client.ProjectID, "vpcs", vpcID, "subnets", id), nil)
	return err
}
