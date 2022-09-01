package clusters

import "github.com/opentelekomcloud/gophertelekomcloud"

// GetCert retrieves a particular cluster certificate based on its unique ID.
func GetCert(client *golangsdk.ServiceClient, id string) (r GetCertResult) {
	raw, err := client.Get(client.ServiceURL("clusters", id, "clustercert"), nil, &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: RequestOpts,
	})
	return
}
