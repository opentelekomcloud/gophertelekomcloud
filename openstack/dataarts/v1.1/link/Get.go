package link

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// Get is used to query a link.
// Send request GET /v1.1/{project_id}/clusters/{cluster_id}/cdm/link/{link_name}
func Get(client *golangsdk.ServiceClient, clusterId, linkName string) (*GetQueryResp, error) {
	raw, err := client.Get(client.ServiceURL(clustersEndpoint, clusterId, cdmEndpoint, linkEndpoint, linkName), nil, nil)
	if err != nil {
		return nil, err
	}

	var res GetQueryResp
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type GetQueryResp struct {
	// Links is a list of Link.
	Links []Link `json:"links"`
	// FromToUnMapping is a Source and destination data sources not supported by table/file migration.
	FromToUnMapping string `json:"fromTo-unMapping"`
	// PageSize is a source and destination data sources supported by entire DB migration.
	BatchFromToMapping string `json:"batchFromTo-mapping,"`
}
