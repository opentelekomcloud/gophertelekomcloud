package logs

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
)

// DisableAutomaticBackups function will disable the automatic log backup policy for a CSS cluster.
func DisableAutomaticBackups(client *golangsdk.ServiceClient, clusterID string) error {
	_, err := client.Put(client.ServiceURL("clusters", clusterID, "logs", "policy", "close"), nil, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return err
}
