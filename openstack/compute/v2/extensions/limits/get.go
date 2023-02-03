package limits

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// GetOpts enables retrieving limits by a specific tenant.
type GetOpts struct {
	// The tenant ID to retrieve limits for.
	TenantID string `q:"tenant_id"`
}

// Get returns the limits about the currently scoped tenant.
func Get(client *golangsdk.ServiceClient, opts GetOpts) (*Limits, error) {
	query, err := golangsdk.BuildQueryString(opts)
	if err != nil {
		return nil, err
	}

	raw, err := client.Get(client.ServiceURL("limits")+query.String(), nil, nil)
	if err != nil {
		return nil, err
	}

	var res Limits
	err = extract.IntoStructPtr(raw.Body, &res, "limits")
	return &res, err
}
