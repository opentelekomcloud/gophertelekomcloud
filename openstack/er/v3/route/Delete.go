package route

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

func Delete(client *golangsdk.ServiceClient, routeTableId, routeId string) (err error) {
	_, err = client.Delete(client.ServiceURL("enterprise-router", "route-tables", routeTableId, "static-routes", routeId), &golangsdk.RequestOpts{
		OkCodes: []int{202},
	})
	return
}
