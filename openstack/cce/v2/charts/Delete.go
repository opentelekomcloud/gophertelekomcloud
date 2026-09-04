package charts

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
)

// This function is used to delete a Helm chart.
func Delete(client *golangsdk.ServiceClient, chartId string) error {
	// DELETE /v2/charts/{chart_id}
	_, err := client.Delete(client.ServiceURL("charts", chartId), &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return err
}
