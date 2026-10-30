package alarm

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type DeleteSilenceRuleName struct {
	// Name of the silence rule to be deleted.
	Name string `json:"name" required:"true"`
}

// DeleteSilenceRule is used to delete silence rules.
// rules contains the names of the silence rules to be deleted.
func DeleteSilenceRule(client *golangsdk.ServiceClient, rules []DeleteSilenceRuleName) error {
	b, err := build.RequestBody(rules, "")
	if err != nil {
		return err
	}

	// POST /v2/{project_id}/alert/mute-rules/delete
	_, err = client.Post(client.ServiceURL("alert", "mute-rules", "delete"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{204},
	})
	return err
}
