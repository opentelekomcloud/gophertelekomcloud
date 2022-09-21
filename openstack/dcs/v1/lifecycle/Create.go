package lifecycle

import "github.com/opentelekomcloud/gophertelekomcloud"

// Create an instance with given parameters.
func Create(client *golangsdk.ServiceClient, ops CreateOpsBuilder) (r CreateResult) {
	b, err := ops.ToInstanceCreateMap()
	if err != nil {
		r.Err = err
		return
	}

	_, r.Err = client.Post(client.ServiceURL("instances"), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}
