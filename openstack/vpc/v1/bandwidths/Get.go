package bandwidths

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// Get retrieves details of a bandwidth by its ID.
func Get(client *golangsdk.ServiceClient, id string) (*BandWidth, error) {
	raw, err := client.Get(client.ServiceURL(client.ProjectID, "bandwidths", id), nil, nil)
	if err != nil {
		return nil, err
	}
	var res struct {
		Bandwidth BandWidth `json:"bandwidth"`
	}
	err = extract.Into(raw.Body, &res)
	return &res.Bandwidth, err
}
