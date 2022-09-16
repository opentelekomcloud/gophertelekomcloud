package tracker

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ListOpts struct {
	TrackerName string `q:"tracker_name"`
}

func List(client *golangsdk.ServiceClient, opts ListOpts) ([]Tracker, error) {
	q, err := golangsdk.BuildQueryString(opts)
	if err != nil {
		return nil, err
	}

	raw, err := client.Get(client.ServiceURL("tracker")+q.String(), nil, nil)
	if err != nil {
		return nil, err
	}

	var res []Tracker
	err = extract.Into(raw.Body, &res)
	return res, err
}
