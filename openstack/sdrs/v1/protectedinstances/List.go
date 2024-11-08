package protectedinstances

import (
	"bytes"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type ListOpts struct {
	ServerGroupID        string   `q:"server_group_id"`
	ServerGroupIDs       []string `q:"server_group_ids"`
	ProtectedInstanceIDs []string `q:"protected_instance_ids"`
	Limit                int      `q:"limit"`
	Offset               int      `q:"offset"`
	Status               string   `q:"status"`
	Name                 string   `q:"name"`
	QueryType            string   `q:"query_type"`
	AvailabilityZone     string   `q:"availability_zone"`
}

func List(client *golangsdk.ServiceClient, opts ListOpts) ([]Instance, error) {
	// GET /v1/{project_id}/protected-instances
	url, err := golangsdk.NewURLBuilder().WithEndpoints("protected-instances").WithQueryParams(&opts).Build()
	if err != nil {
		return nil, err
	}

	pages, err := pagination.Pager{
		Client:     client,
		InitialURL: client.ServiceURL(url.String()),
		CreatePage: func(r pagination.NewPageResult) pagination.NewPage {
			return InstancePage{NewSinglePageBase: pagination.NewSinglePageBase{NewPageResult: r}}
		},
	}.NewAllPages()

	if err != nil {
		return nil, err
	}
	return ExtractProtectedInstances(pages)
}

func ExtractProtectedInstances(r pagination.NewPage) ([]Instance, error) {
	var listResponse ListResponse
	err := extract.Into(bytes.NewReader((r.(InstancePage)).Body), &listResponse)
	return listResponse.ProtectedInstances, err
}

// InstancePage is a struct which can do the page function
type InstancePage struct {
	pagination.NewSinglePageBase
}

type ListResponse struct {
	// Specifies the information about protected instances.
	ProtectedInstances []Instance `json:"protected_instances"`
	// Specifies the number of protected instances.
	Count int `json:"count"`
}
