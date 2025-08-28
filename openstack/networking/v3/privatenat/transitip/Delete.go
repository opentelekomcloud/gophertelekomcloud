package transitip

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

// This function is used to release a transit IP address.
func Delete(client *golangsdk.ServiceClient, transitIpId string) error {
	// DELETE /v3/{project_id}/private-nat/transit-ips/{transit_ip_id}
	_, err := client.Delete(client.ServiceURL("private-nat", "transit-ips", transitIpId), &golangsdk.RequestOpts{
		OkCodes: []int{200, 204},
	})
	return err
}
