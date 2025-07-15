package advanced

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

func DeleteQuery(client *golangsdk.ServiceClient, domainId, queryId string) (err error) {
	_, err = client.Delete(client.ServiceURL("resource-manager", "domains", domainId, "stored-queries", queryId), &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}
