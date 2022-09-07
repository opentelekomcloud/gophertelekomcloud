package hypervisors

import (
	"strconv"

	"github.com/opentelekomcloud/gophertelekomcloud"
)

// GetUptime makes a request against the API to get uptime for specific hypervisor.
func GetUptime(client *golangsdk.ServiceClient, hypervisorID int) (r UptimeResult) {
	v := strconv.Itoa(hypervisorID)
	raw, err := client.Get(client.ServiceURL("os-hypervisors", v, "uptime"), nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}
