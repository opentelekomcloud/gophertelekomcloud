package global_connection_bandwidth

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

func Delete(client *golangsdk.ServiceClient, id string) error {
	_, err := client.Delete(client.ServiceURL(client.DomainID, "gcb", "gcbandwidths", id), &golangsdk.RequestOpts{
		OkCodes: []int{200, 202, 204},
	})
	return err
}
