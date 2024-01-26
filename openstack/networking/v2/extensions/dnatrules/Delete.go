package dnatrules

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
)

func Delete(client *golangsdk.ServiceClient, id string) (err error) {
	// DELETE /v2.0/dnat_rules/{dnat_rule_id}
	_, err = client.Delete(client.ServiceURL("dnat_rules", id), &golangsdk.RequestOpts{
		OkCodes: []int{204},
	})
	return
}
