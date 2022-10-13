package tracker

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
)

type GetOpts struct {
	TrackerName string `q:"tracker_name"`
}

func Get(client *golangsdk.ServiceClient, opts GetOpts) (*Tracker, error) {
	q, err := golangsdk.BuildQueryString(opts)
	if err != nil {
		return nil, err
	}

	raw, err := client.Get(client.ServiceURL("tracker")+q.String(), nil, nil)
	return extra(err, raw)
}
