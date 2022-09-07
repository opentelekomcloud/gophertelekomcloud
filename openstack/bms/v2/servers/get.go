package servers

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// Get requests details on a single server, by ID.
func Get(client *golangsdk.ServiceClient, id string) (*Server, error) {
	raw, err := client.Get(client.ServiceURL("servers", id), nil, &golangsdk.RequestOpts{
		MoreHeaders: map[string]string{"X-OpenStack-Nova-API-Version": "2.26"},
	})
	if err != nil {
		return nil, err
	}

	var res Server
	err = extract.IntoStructPtr(raw.Body, &res, "server")
	return &res, err
}
