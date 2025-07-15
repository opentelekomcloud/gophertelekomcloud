package vpc_endpoint

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
)

// EnableVpcEndpoint function is used to Disable VCP endpoint of the cluster.
func Disable(client *golangsdk.ServiceClient, clusterID string) error {
	_, err := client.Put(client.ServiceURL("clusters", clusterID, "vpcepservice", "close"), nil, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return err
}
