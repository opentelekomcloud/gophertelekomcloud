package tracker

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
)

type tracker struct {
	Tracker string `q:"tracker_name"`
}

func Get(client *golangsdk.ServiceClient, trackerName string) (*Tracker, error) {

	url, err := golangsdk.NewURLBuilder().WithEndpoints("tracker").WithQueryParams(&tracker{Tracker: trackerName}).Build()
	if err != nil {
		return nil, err
	}
	// GET /v1.0/{project_id}/tracker?tracker_name={tracker_name}
	raw, err := client.Get(client.ServiceURL(url.String()), nil, nil)
	return extra(err, raw)
}
