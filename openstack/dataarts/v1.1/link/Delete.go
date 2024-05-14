package link

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
)

// Delete is used to delete a link.
// Send request DELETE /v1.1/{project_id}/clusters/{cluster_id}/cdm/link/{link_name}
func Delete(client *golangsdk.ServiceClient, clusterId, linkName string) error {
	_, err := client.Delete(client.ServiceURL(clustersEndpoint, clusterId, cdmEndpoint, linkEndpoint, linkName), &golangsdk.RequestOpts{OkCodes: []int{200}})
	if err != nil {
		return err
	}

	return nil
}
