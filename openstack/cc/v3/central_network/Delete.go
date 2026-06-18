package central_network

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

func Delete(client *golangsdk.ServiceClient, centralNetworkId string) error {
	_, err := client.Delete(client.ServiceURL(client.DomainID, "gcn", "central-networks", centralNetworkId), &golangsdk.RequestOpts{
		OkCodes: []int{200, 202},
	})
	return err
}
