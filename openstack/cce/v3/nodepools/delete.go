package nodepools

import "github.com/opentelekomcloud/gophertelekomcloud"

// Delete will permanently delete a particular node pool based on its unique ID and cluster ID.
func Delete(client *golangsdk.ServiceClient, clusterid, nodepoolid string) (err error) {
	_, err = client.Delete(client.ServiceURL("clusters", clusterid, "nodepools", nodepoolid), &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: RequestOpts.MoreHeaders, JSONBody: nil,
	})
	return
}
