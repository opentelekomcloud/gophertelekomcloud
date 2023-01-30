package relations

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/rms/v1/resources"
)

type ShowResourceRelationsOpts struct {
	// Specifies the resource ID.
	// Maximum length: 256
	ResourceId string
	// Specifies the resource relationship direction.
	// Possible values are as follows:
	// - in
	// - out
	Direction string `q:"direction"`
	// Specifies the maximum number of records to return.
	// Minimum value: 1
	// Maximum value: 1000
	Limit int `q:"limit,omitempty"`
	// Specifies the pagination parameter.
	// You can use the marker value returned to the previous request as the number of the first page of records to return in this request.
	// Minimum length: 4
	// Maximum length: 400
	Marker string `q:"marker,omitempty"`
}

func ShowResourceRelations(client *golangsdk.ServiceClient, opts ShowResourceRelationsOpts) (*ShowResourceRelationsResponse, error) {
	q, err := golangsdk.BuildQueryString(opts)
	if err != nil {
		return nil, err
	}

	// GET /v1/resource-manager/domains/{domain_id}/resources/{resource_id}/relations
	raw, err := client.Get(client.ServiceURL("resources", opts.ResourceId, "relations")+q.String(), nil, nil)
	if err != nil {
		return nil, err
	}

	var res ShowResourceRelationsResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type ShowResourceRelationsResponse struct {
	// Specifies the list of the resource relationships.
	Relations []ResourceRelation `json:"relations,omitempty"`
	// Specifies the pagination object.
	PageInfo resources.PageInfo `json:"page_info,omitempty"`
}

type ResourceRelation struct {
	// Specifies the relationship type.
	RelationType string `json:"relation_type,omitempty"`
	// Specifies the type of the source resource.
	FromResourceType string `json:"from_resource_type,omitempty"`
	// Specifies the type of the associated resource.
	ToResourceType string `json:"to_resource_type,omitempty"`
	// Specifies the source resource ID.
	FromResourceId string `json:"from_resource_id,omitempty"`
	// Specifies the ID of the associated resource.
	ToResourceId string `json:"to_resource_id,omitempty"`
}
