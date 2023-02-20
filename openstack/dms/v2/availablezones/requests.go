package availablezones

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// Get available zones
func Get(client *golangsdk.ServiceClient) (*GetResponse, error) {
	raw, err := client.Get(getURL(client), nil, nil)
	if err != nil {
		return nil, err
	}
	var res GetResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}
