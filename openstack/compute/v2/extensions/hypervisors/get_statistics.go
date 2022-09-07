package hypervisors

import "github.com/opentelekomcloud/gophertelekomcloud"

// Statistics makes a request against the API to get hypervisors statistics.
func GetStatistics(client *golangsdk.ServiceClient) (r StatisticsResult) {
	raw, err := client.Get(hypervisorsStatisticsURL(client), nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}
