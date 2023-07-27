package flavors

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// Get returns additional information about a Flavor, given its ID.
func Get(client *golangsdk.ServiceClient, flavorID string) (*Flavor, error) {
	// GET /v3/{project_id}/elb/flavors/{flavor_id}
	raw, err := client.Get(client.ServiceURL("flavors", flavorID), nil, nil)
	if err != nil {
		return nil, err
	}

	var res Flavor
	err = extract.IntoStructPtr(raw.Body, &res, "flavor")
	return &res, err
}

type Flavor struct {
	// Specifies the ID of the flavor.
	ID string `json:"id"`
	// Specifies the info of the flavor.
	Info FlavorInfo `json:"info"`
	// Specifies the name of the flavor.
	Name string `json:"name"`
	// Specifies whether shared.
	Shared bool `json:"shared"`
	// Specifies the type of the flavor.
	Type string `json:"type"`
	// Specifies whether sold out.
	SoldOut bool `json:"flavor_sold_out"`
}

type FlavorInfo struct {
	// Specifies the connection
	Connection int `json:"connection"`
	// Specifies the cps.
	Cps int `json:"cps"`
	// Specifies the qps
	Qps int `json:"qps"`
	// Specifies the https_cps
	HttpsCps int `json:"https_cps"`
	// Specifies the lcu
	Lcu int `json:"lcu"`
	// Specifies the bandwidth
	Bandwidth int `json:"bandwidth"`
}
