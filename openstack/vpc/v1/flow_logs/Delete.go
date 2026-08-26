package flow_logs

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

func Delete(client *golangsdk.ServiceClient, flowLogID string) error {
	_, err := client.Delete(
		client.ServiceURL(client.ProjectID, "fl", "flow_logs", flowLogID),
		&golangsdk.RequestOpts{OkCodes: []int{204}},
	)
	return err
}
