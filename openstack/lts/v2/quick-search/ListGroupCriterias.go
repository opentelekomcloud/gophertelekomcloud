package quick_search

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func ListGroupCriterias(client *golangsdk.ServiceClient, groupId string) ([]SearchGroupCriteria, error) {
	// GET /v1.0/{project_id}/lts/groups/{group_id}/search-criterias
	raw, err := client.Get(client.ServiceURL("lts", "groups", groupId, "search-criterias"), nil, &golangsdk.RequestOpts{
		MoreHeaders: map[string]string{
			"content-type": "application/json",
		},
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res []SearchGroupCriteria
	err = extract.IntoSlicePtr(raw.Body, &res, "search_criterias")
	return res, err
}

type SearchGroupCriteria struct {
	// Quick search of a field.
	Criterias []SearchCriteria `json:"criterias"`
	// Log stream ID.
	StreamId string `json:"log_stream_id"`
	// Log stream name.
	StreamName string `json:"log_stream_name"`
	// Quick search type.
	SearchType string `json:"search_type"`
}
