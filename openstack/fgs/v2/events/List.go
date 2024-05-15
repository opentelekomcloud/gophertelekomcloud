package events

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func List(client *golangsdk.ServiceClient, funcUrn string) (*ListEventsResponse, error) {
	url, err := golangsdk.NewURLBuilder().WithEndpoints("fgs", "functions", funcUrn, "events").Build()
	if err != nil {
		return nil, err
	}

	raw, err := client.Get(client.ServiceURL(url.String()), nil, nil)
	if err != nil {
		return nil, err
	}

	var res ListEventsResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type ListEventsResponse struct {
	Events     []Event `json:"events"`
	NextMarker int     `json:"next_marker"`
	Count      int     `json:"count"`
}

type Event struct {
	Id           string `json:"id"`
	Content      string `json:"content"`
	LastModified int    `json:"last_modified"`
	Name         string `json:"name"`
}
