package route_table

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

func Delete(client *golangsdk.ServiceClient, erId, routeTableId string) (err error) {
	_, err = client.Delete(client.ServiceURL("enterprise-router", erId, "route-tables", routeTableId), &golangsdk.RequestOpts{
		OkCodes: []int{202},
	})
	return
}
