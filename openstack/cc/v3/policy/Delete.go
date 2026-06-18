package policy

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

func Delete(client *golangsdk.ServiceClient, centralNetworkId, policyId string) error {
	_, err := client.Delete(client.ServiceURL(client.DomainID, "gcn", "central-network", centralNetworkId, "policies", policyId), &golangsdk.RequestOpts{
		OkCodes: []int{200, 202, 204},
	})
	return err
}
