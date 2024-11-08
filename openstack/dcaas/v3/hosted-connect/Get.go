package hosted_connect

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// Get retrieves a particular hosted connect based on its unique ID.
func Get(c *golangsdk.ServiceClient, id string) (*HostedConnect, error) {
	// GET /v3/{project_id}/dcaas/hosted-connects/{hosted_connect_id}
	raw, err := c.Get(c.ServiceURL("dcaas", "hosted-connects", id), nil, nil)
	if err != nil {
		return nil, err
	}

	var res HostedConnect
	err = extract.IntoStructPtr(raw.Body, &res, "hosted_connect")
	return &res, err
}
