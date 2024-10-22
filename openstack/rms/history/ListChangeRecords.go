package history

import (
	"bytes"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type ListAllOpts struct {
	DomainId   string `json:"-"`
	ResourceId string `json:"-"`
	// Specifies the maximum number of resources to return.
	Limit *int `q:"limit"`
	// Specifies the pagination parameter.
	Marker string `q:"marker"`
	// Specifies the start time of the query. If this parameter is not set, the earliest time is used by default.
	EarlierTime *int64 `q:"earlier_time"`
	// Specifies the end time of the query. If this parameter is not set, the current time is used by default.
	LaterTime *int64 `q:"later_time"`
	// Specifies the time sequence of the data to be returned. The default value is Reverse.
	ChronologicalOrder *int64 `q:"chronological_order"`
}

func ListAllRecords(client *golangsdk.ServiceClient, opts ListAllOpts) ([]HistoryItem, error) {
	// GET /v1/resource-manager/domains/{domain_id}/resources/{resource_id}/history
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints("resource-manager", "domains", opts.DomainId, "resources", opts.ResourceId).
		WithQueryParams(&opts).Build()
	if err != nil {
		return nil, err
	}
	pages, err := pagination.Pager{
		Client:     client,
		InitialURL: client.ServiceURL(url.String()),
		CreatePage: func(r pagination.NewPageResult) pagination.NewPage {
			return ResPage{NewSinglePageBase: pagination.NewSinglePageBase{NewPageResult: r}}
		},
	}.NewAllPages()
	if err != nil {
		return nil, err
	}
	return ExtractResources(pages)
}

type ResPage struct {
	pagination.NewSinglePageBase
}

func ExtractResources(r pagination.NewPage) ([]HistoryItem, error) {
	var s struct {
		Items []HistoryItem `json:"items"`
	}
	err := extract.Into(bytes.NewReader((r.(ResPage)).Body), &s)
	return s.Items, err
}

type HistoryItem struct {
	DomainId     string         `json:"domain_id"`
	ResourceId   string         `json:"resource_id"`
	ResourceType string         `json:"resource_type"`
	CaptureTime  string         `json:"capture_time"`
	Status       string         `json:"status"`
	Relations    []Relations    `json:"relations"`
	Resource     ResourceEntity `json:"resource"`
}

type Relations struct {
	RelationType     string `json:"relation_type"`
	FromResourceType string `json:"from_resource_type"`
	ToResourceType   string `json:"to_resource_type"`
	FromResourceId   string `json:"from_resource_id"`
	ToResourceId     string `json:"to_resource_id"`
}

type ResourceEntity struct {
	ID                string                 `json:"id"`
	Name              string                 `json:"name"`
	Provider          string                 `json:"provider"`
	Type              string                 `json:"type"`
	RegionID          string                 `json:"region_id"`
	ProjectID         string                 `json:"project_id"`
	ProjectName       string                 `json:"project_name"`
	EpID              string                 `json:"ep_id"`
	EpName            string                 `json:"ep_name"`
	Checksum          string                 `json:"checksum"`
	Created           string                 `json:"created"`
	Updated           string                 `json:"updated"`
	ProvisioningState string                 `json:"provisioning_state"`
	State             string                 `json:"state"`
	Tags              map[string]string      `json:"tags"`
	Properties        map[string]interface{} `json:"properties"`
	OSType            string                 `json:"osType,omitempty"`
	KeyName           string                 `json:"keyName,omitempty"`
	SchedulerHints    interface{}            `json:"schedulerHints,omitempty"`
}
