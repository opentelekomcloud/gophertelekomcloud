package hypervisors

import (
	"strconv"

	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// GetUptime makes a request against the API to get uptime for specific hypervisor.
func GetUptime(client *golangsdk.ServiceClient, hypervisorID int) (*Uptime, error) {
	raw, err := client.Get(client.ServiceURL("os-hypervisors", strconv.Itoa(hypervisorID), "uptime"), nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res Uptime
	err = extract.IntoStructPtr(raw.Body, &res, "hypervisor")
	return &res, err
}

// Uptime represents uptime and additional info for a specific hypervisor.
type Uptime struct {
	// The hypervisor host name provided by the Nova virt driver.
	// For the Ironic driver, it is the Ironic node uuid.
	HypervisorHostname string `json:"hypervisor_hostname"`
	// The id of the hypervisor.
	ID int `json:"id"`
	// The state of the hypervisor. One of up or down.
	State string `json:"state"`
	// The status of the hypervisor. One of enabled or disabled.
	Status string `json:"status"`
	// The total uptime of the hypervisor and information about average load.
	Uptime string `json:"uptime"`
}
