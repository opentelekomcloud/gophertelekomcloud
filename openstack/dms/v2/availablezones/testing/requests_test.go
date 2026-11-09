package testing

import (
	"fmt"
	"net/http"
	"testing"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dms/v2/availablezones"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
	"github.com/opentelekomcloud/gophertelekomcloud/testhelper/client"
)

const projectID = "0123456789abcdef0123456789abcdef"

func dmsClient() *golangsdk.ServiceClient {
	sc := client.ServiceClient()
	sc.Endpoint = th.Endpoint() + "v2/" + projectID + "/"
	return sc
}

func TestGet(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v2/available-zones", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprint(w, `
{
  "region_id" : "eu-de",
  "available_zones" : [ {
    "id" : "d539378ec1314c85b76fefa3f7071458",
    "code" : "eu-de-01",
    "name" : "AZ 2.",
    "port" : "8003",
    "resource_availability" : "true"
  }, {
    "id" : "9f1c5806706d4c1fb0eb72f0a9b18c77",
    "code" : "eu-de-02",
    "name" : "AZ 3.",
    "port" : "443",
    "resource_availability" : "false"
  } ]
}`)
	})

	res, err := availablezones.Get(dmsClient())
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, &availablezones.GetResponse{
		RegionID: "eu-de",
		AvailableZones: []availablezones.AvailableZone{
			{
				ID:                   "d539378ec1314c85b76fefa3f7071458",
				Code:                 "eu-de-01",
				Name:                 "AZ 2.",
				Port:                 "8003",
				ResourceAvailability: "true",
			},
			{
				ID:                   "9f1c5806706d4c1fb0eb72f0a9b18c77",
				Code:                 "eu-de-02",
				Name:                 "AZ 3.",
				Port:                 "443",
				ResourceAvailability: "false",
			},
		},
	}, res)
}
