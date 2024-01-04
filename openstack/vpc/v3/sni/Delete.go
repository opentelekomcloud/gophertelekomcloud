package sni

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
)

// Delete is used to delete a supplementary network interface.
func Delete(client *golangsdk.ServiceClient, id string) (err error) {
	// DELETE /v3/{project_id}/vpc/sub-network-interfaces/{sub_network_interface_id}
	_, err = client.Delete(client.ServiceURL("sub-network-interfaces", id), nil)
	return
}
