package region

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func GetRegions(client *golangsdk.ServiceClient, domainId string) ([]Region, error) {
	// GET /v1/resource-manager/domains/{domain_id}/regions
	raw, err := client.Get(client.ServiceURL("resource-manager", "domains", domainId, "regions"), nil, nil)
	if err != nil {
		return nil, err
	}

	var res []Region

	err = extract.IntoSlicePtr(raw.Body, &res, "value")
	return res, err
}

type Region struct {
	RegionId    string `json:"region_id"`
	DisplayName string `json:"display_name"`
}
