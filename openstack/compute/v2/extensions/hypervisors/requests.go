package hypervisors

import (
	"strconv"

	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

// List makes a request against the API to list hypervisors.
func List(client *golangsdk.ServiceClient) pagination.Pager {
	return pagination.NewPager(client, hypervisorsListDetailURL(client), func(r pagination.PageResult) pagination.Page {
		return HypervisorPage{pagination.SinglePageBase(r)}
	})
}

// Statistics makes a request against the API to get hypervisors statistics.
func GetStatistics(client *golangsdk.ServiceClient) (r StatisticsResult) {
	raw, err := client.Get(hypervisorsStatisticsURL(client), nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

// Get makes a request against the API to get details for specific hypervisor.
func Get(client *golangsdk.ServiceClient, hypervisorID int) (r HypervisorResult) {
	v := strconv.Itoa(hypervisorID)
	raw, err := client.Get(hypervisorsGetURL(client, v), nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

// GetUptime makes a request against the API to get uptime for specific hypervisor.
func GetUptime(client *golangsdk.ServiceClient, hypervisorID int) (r UptimeResult) {
	v := strconv.Itoa(hypervisorID)
	raw, err := client.Get(hypervisorsUptimeURL(client, v), nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}
