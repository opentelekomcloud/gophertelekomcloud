package clusters

import "github.com/opentelekomcloud/gophertelekomcloud"

// GetCertWithExpiration retrieves a particular cluster certificate based on its unique ID.
func GetCertWithExpiration(client *golangsdk.ServiceClient, id string, opts ExpirationOpts) (r GetCertResult) {
	b, err := golangsdk.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL("clusters", id, "clustercert"), b, nil, &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: RequestOpts,
	})
	return
}
