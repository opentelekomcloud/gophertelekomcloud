package quick_search

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ListOpts struct {
	// Raw logs.
	SearchType string `q:"search_type"`
}

func ListCriterias(client *golangsdk.ServiceClient, groupId, streamId string, opts ListOpts) ([]SearchCriteria, error) {
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints("groups", groupId, "topics", streamId, "search-criterias").
		WithQueryParams(&opts).Build()
	if err != nil {
		return nil, err
	}
	// GET /v1.0/{project_id}/groups/{group_id}/topics/{topic_id}/search-criterias
	raw, err := client.Get(client.ServiceURL(url.String()), nil, &golangsdk.RequestOpts{
		MoreHeaders: map[string]string{
			"content-type": "application/json",
		},
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res []SearchCriteria
	err = extract.IntoSlicePtr(raw.Body, &res, "search_criterias")
	return res, err
}

type SearchCriteria struct {
	// Quick search of a field.
	Criteria string `json:"criteria"`
	// Quick search of a name.
	Name string `json:"name"`
	// Quick search ID.
	ID string `json:"id"`
	// Quick search type.
	SearchType string `json:"search_type"`
}
