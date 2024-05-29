package events

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func Get(client *golangsdk.ServiceClient, funcURN, eventId string) (*Event, error) {
	raw, err := client.Get(client.ServiceURL("fgs", "functions", funcURN, "events", eventId), nil, nil)
	if err != nil {
		return nil, err
	}

	var res Event
	err = extract.Into(raw.Body, &res)
	return &res, err
}
