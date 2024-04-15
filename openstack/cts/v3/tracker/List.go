package tracker

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
)

type tracker struct {
	Tracker string `q:"tracker_name"`
}

func List(client *golangsdk.ServiceClient, trackerName string) ([]Tracker, error) {
	url, err := golangsdk.NewURLBuilder().WithEndpoints("trackers").WithQueryParams(&tracker{Tracker: trackerName}).Build()
	if err != nil {
		return []Tracker{}, err
	}
	// GET /v3/{project_id}/trackers
	raw, err := client.Get(client.ServiceURL(url.String()), nil, nil)
	return extraStruct(err, raw)
}
