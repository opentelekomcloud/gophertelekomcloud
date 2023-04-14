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
