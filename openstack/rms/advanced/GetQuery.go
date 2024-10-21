package advanced

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func GetQuery(client *golangsdk.ServiceClient, domainId, queryId string) (*Query, error) {
	// GET  /v1/resource-manager/domains/{domain_id}/stored-queries/{query_id}
	raw, err := client.Get(client.ServiceURL(
		"resource-manager", "domains", domainId, "stored-queries", queryId), nil, nil)
	if err != nil {
		return nil, err
	}

	var res Query
	err = extract.Into(raw.Body, &res)
	return &res, err
}
