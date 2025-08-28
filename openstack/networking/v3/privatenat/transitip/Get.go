package transitip

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// This function is used to query details about a specified transit IP address.
func Get(client *golangsdk.ServiceClient, transitIpId string) (*TransitIPCommonResponse, error) {
	// GET /v3/{project_id}/private-nat/transit-ips/{transit_ip_id}
	raw, err := client.Get(client.ServiceURL("private-nat", "transit-ips", transitIpId), nil, nil)
	if err != nil {
		return nil, err
	}

	var res TransitIPCommonResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}
