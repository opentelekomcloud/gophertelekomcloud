package versions

import (
	"strings"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type Version struct {
	ID     string `json:"id"`
	Links  []Link `json:"links"`
	Status string `json:"status"`
	Update string `json:"updated"`
}

type Link struct {
	Href string `json:"href"`
	Rel  string `json:"rel"`
}

// List queries available API versions at the EPS service root.
func List(client *golangsdk.ServiceClient) ([]Version, error) {
	// Strip version path to query the root: https://eps.xxx/v1.0/ → https://eps.xxx/
	baseURL := client.Endpoint
	if idx := strings.Index(baseURL, "/v"); idx > 0 {
		baseURL = baseURL[:idx] + "/"
	}

	raw, err := client.Get(baseURL, nil, nil)
	if err != nil {
		return nil, err
	}

	var res []Version
	err = extract.IntoSlicePtr(raw.Body, &res, "versions")
	return res, err
}
