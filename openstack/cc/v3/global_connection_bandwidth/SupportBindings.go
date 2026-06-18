package global_connection_bandwidth

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ListSupportBindingsOpts struct {
	// Number of records returned on each page. Value range: 1 to 2000.
	Limit int `q:"limit"`
	// ID of the last record on the previous page.
	Marker string `q:"marker"`
	// Filter by enterprise project IDs.
	EnterpriseProjectId []string `q:"enterprise_project_id"`
	// Local access point.
	LocalArea string `q:"local_area"`
	// Remote access point.
	RemoteArea string `q:"remote_area"`
	// Bound instance type. Value options: CC, GEIP, GCN, GSN.
	BindingService string `q:"binding_service" required:"true"`
}

// ListSupportBindings queries the list of global connection bandwidths that meet the binding conditions.
func ListSupportBindings(client *golangsdk.ServiceClient, opts ListSupportBindingsOpts) (*ListGlobalConnectionBandwidthResp, error) {
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints(client.DomainID, "gcb", "gcbandwidths", "support-bindings").
		WithQueryParams(&opts).Build()
	if err != nil {
		return nil, err
	}

	raw, err := client.Get(client.ServiceURL(url.String()), nil, nil)
	if err != nil {
		return nil, err
	}

	var res ListGlobalConnectionBandwidthResp
	err = extract.Into(raw.Body, &res)
	return &res, err
}
