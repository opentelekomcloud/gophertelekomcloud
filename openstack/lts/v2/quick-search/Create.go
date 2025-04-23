package quick_search

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type CreateOpts struct {
	// Quick search field. Enter the statement to be queried.
	Criteria string `json:"criteria" required:"true"`
	// Enterprise project ID.
	EpsId string `json:"eps_id,omitempty"`
	// Quick search name, which contains 1 to 64 characters,
	// including only letters, digits, underscores (_), hyphens (-), and periods (.).
	// Do not start with a period or underscore or end with a period.
	Name string `json:"name" required:"true"`
	// Search type, for example, raw logs.
	SearchType string `json:"search_type" required:"true"`
}

func Create(client *golangsdk.ServiceClient, groupId, streamId string, opts CreateOpts) (string, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return "", err
	}

	// POST /v1.0/{project_id}/groups/{group_id}/topics/{topic_id}/search-criterias
	raw, err := client.Post(client.ServiceURL("groups", groupId, "topics", streamId, "search-criterias"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{201},
	})
	if err != nil {
		return "", err
	}

	var res struct {
		ID string `json:"id"`
	}
	err = extract.Into(raw.Body, &res)
	return res.ID, err
}
