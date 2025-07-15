package tracker

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
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
	if err != nil {
		return nil, err
	}
	var res []Tracker

	err = extract.IntoSlicePtr(raw.Body, &res, "trackers")
	return res, err
}
