package alarm

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

// DeleteActionRule is used to delete alarm action rules.
// ruleNames contains the names of the rules to be deleted.
func DeleteActionRule(client *golangsdk.ServiceClient, ruleNames []string) error {
	b, err := build.RequestBody(ruleNames, "")
	if err != nil {
		return err
	}

	// POST /v2/{project_id}/alert/action-rules/delete
	_, err = client.Post(client.ServiceURL("alert", "action-rules", "delete"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{204},
	})
	return err
}
