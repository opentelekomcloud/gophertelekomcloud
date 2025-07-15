package quick_search

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type DeleteOpts struct {
	// Enterprise project ID.
	EpsId string `json:"eps_id,omitempty"`
	// Quick search ID.
	ID string `json:"id" required:"true"`
}

func Delete(client *golangsdk.ServiceClient, groupId, streamId string, opts DeleteOpts) (err error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return err
	}

	// DELETE /v1.0/{project_id}/groups/{group_id}/topics/{topic_id}/search-criterias
	_, err = client.DeleteWithBody(client.ServiceURL("groups", groupId, "topics", streamId, "search-criterias"), b, &golangsdk.RequestOpts{
		MoreHeaders: map[string]string{
			"content-type": "application/json",
		},
		OkCodes: []int{204},
	})
	return
}
