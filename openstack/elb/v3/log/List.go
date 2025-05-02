package log

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack"
)

type ListOpts struct {
	// Specifies the number of records on each page.
	Limit int `q:"limit"`
	// Specifies the ID of the last record on the previous page.
	Marker string `q:"marker"`
	// Specifies whether to use reverse query. Values:
	// true: Query the previous page.
	// false (default): Query the next page.
	PageReverse bool `q:"page_reverse"`
	// Specifies the enterprise project ID.
	EnterpriseProjectId []string `q:"enterprise_project_id"`
	// Specifies the ID of the log tank.
	ID []string `q:"id"`
	// Specifies the ID of a load balancer.
	LoadbalancerId []string `q:"loadbalancer_id"`
	// Specifies the log group ID.
	LogGroupId []string `q:"log_group_id"`
	// Specifies the log stream ID.
	LogStreamId []string `q:"log_topic_id"`
}

func List(client *golangsdk.ServiceClient, opts ListOpts) ([]Logtank, error) {
	// GET /v3/{project_id}/elb/logtanks
	url, err := golangsdk.NewURLBuilder().WithEndpoints("logtanks").WithQueryParams(&opts).Build()
	if err != nil {
		return nil, err
	}

	raw, err := client.Get(client.ServiceURL(url.String()), nil, openstack.StdRequestOpts())
	if err != nil {
		return nil, err
	}

	var res []Logtank
	err = extract.IntoSlicePtr(raw.Body, &res, "logtanks")
	return res, err
}
