package flow_logs

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

func Delete(client *golangsdk.ServiceClient, routerID, flowLogID string) (err error) {
	_, err = client.Delete(client.ServiceURL("enterprise-router", routerID, "flow-logs", flowLogID), &golangsdk.RequestOpts{
		OkCodes: []int{202},
	})
	return
}
