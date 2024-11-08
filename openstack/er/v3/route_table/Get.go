package route_table

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func Get(client *golangsdk.ServiceClient, erID, routeTableId string) (*RouteTable, error) {
	raw, err := client.Get(client.ServiceURL("enterprise-router", erID, "route-tables", routeTableId), nil, nil)
	if err != nil {
		return nil, err
	}

	var res RouteTable
	err = extract.IntoStructPtr(raw.Body, &res, "route_table")
	return &res, err
}
