package tag

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

type DeleteClusterTagOpts struct {
	// Resource ID
	ClusterId string
	// Tag key
	Key string
}

func DeleteCluster(client *golangsdk.ServiceClient, opts DeleteClusterTagOpts) error {
	// DELETE /v1.0/{project_id}/clusters/{resource_id}/tags/{key}
	_, err := client.Delete(client.ServiceURL("clusters", opts.ClusterId, "tags", opts.Key), nil)
	return err
}
