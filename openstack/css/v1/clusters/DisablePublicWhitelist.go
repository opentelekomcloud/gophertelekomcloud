package clusters

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

// DisablePublicWhitelist function is used to disable whitelist for a single ip
func DisablePublicWhitelist(client *golangsdk.ServiceClient, clusterID string) error {
	_, err := client.Put(client.ServiceURL("clusters", clusterID, "public", "whitelist", "close"), nil, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return err
}
