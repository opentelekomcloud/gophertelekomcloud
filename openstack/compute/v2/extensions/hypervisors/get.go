package hypervisors

import (
	"strconv"

	"github.com/opentelekomcloud/gophertelekomcloud"
)

// Get makes a request against the API to get details for specific hypervisor.
func Get(client *golangsdk.ServiceClient, hypervisorID int) (r HypervisorResult) {
	v := strconv.Itoa(hypervisorID)
	raw, err := client.Get(hypervisorsGetURL(client, v), nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}
