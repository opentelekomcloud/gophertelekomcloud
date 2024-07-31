package route

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func Get(client *golangsdk.ServiceClient, routeTableId, routeId string) (*Route, error) {
	raw, err := client.Get(client.ServiceURL("enterprise-router", "route-tables", routeTableId, "static-routes", routeId), nil, nil)
	if err != nil {
		return nil, err
	}

	var res Route
	err = extract.IntoStructPtr(raw.Body, &res, "route")
	return &res, err
}
