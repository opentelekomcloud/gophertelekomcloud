package routetables

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func Get(client *golangsdk.ServiceClient, id string) (*RouteTable, error) {
	url, err := golangsdk.NewURLBuilder().WithEndpoints(client.ProjectID, "routetables", id).Build()
	if err != nil {
		return nil, err
	}
	raw, err := client.Get(client.ServiceURL(url.String()), nil, nil)
	if err != nil {
		return nil, err
	}
	var res struct {
		RouteTable RouteTable `json:"routetable"`
	}
	if err = extract.Into(raw.Body, &res); err != nil {
		return nil, err
	}
	return &res.RouteTable, nil
}
