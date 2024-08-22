package snatrules

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
)

func Delete(client *golangsdk.ServiceClient, id string) (err error) {
	// DELETE /v2.0/snat_rules/{snat_rule_id}
	_, err = client.Delete(client.ServiceURL("snat_rules", id), &golangsdk.RequestOpts{
		OkCodes: []int{204},
	})
	return
}
