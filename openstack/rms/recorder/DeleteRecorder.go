package recorder

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
)

func DeleteRecorder(client *golangsdk.ServiceClient, domainId string) (err error) {
	// DELETE /v1/resource-manager/domains/{domain_id}/tracker-config
	_, err = client.Delete(client.ServiceURL("resource-manager", "domains", domainId, "tracker-config"), &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}
