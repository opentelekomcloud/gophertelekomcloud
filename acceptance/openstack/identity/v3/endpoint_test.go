package v3

import (
	"testing"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/identity/v3/endpoints"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/identity/v3/services"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestEndpointsList(t *testing.T) {
	client, err := clients.NewIdentityV3Client()
	th.AssertNoErr(t, err)

	allPages, err := endpoints.List(client, nil).AllPages()
	th.AssertNoErr(t, err)

	allEndpoints, err := endpoints.ExtractEndpoints(allPages)
	th.AssertNoErr(t, err)

	for _, endpoint := range allEndpoints {
		tools.PrintResource(t, endpoint)
	}
}

func TestEndpointsNavigateCatalog(t *testing.T) {
	client, err := clients.NewIdentityV3Client()
	th.AssertNoErr(t, err)

	// Discover the service we're interested in.
	serviceListOpts := services.ListOpts{
		ServiceType: "compute",
	}

	allPages, err := services.List(client, serviceListOpts).AllPages()
	th.AssertNoErr(t, err)

	allServices, err := services.ExtractServices(allPages)
	th.AssertNoErr(t, err)

	if len(allServices) != 1 {
		t.Fatalf("Expected one service, got %d", len(allServices))
	}

	computeService := allServices[0]
	tools.PrintResource(t, computeService)

	// Enumerate the endpoints available for this service.
	endpointListOpts := endpoints.ListOpts{
		Availability: golangsdk.AvailabilityPublic,
		ServiceID:    computeService.ID,
	}

	allPages, err = endpoints.List(client, endpointListOpts).AllPages()
	th.AssertNoErr(t, err)

	allEndpoints, err := endpoints.ExtractEndpoints(allPages)
	th.AssertNoErr(t, err)

	th.AssertEquals(t, allEndpoints[0].Region, "eu-de")
	th.AssertEquals(t, allEndpoints[1].Region, "eu-nl")
	th.CheckEquals(t, 2, len(allEndpoints))
	tools.PrintResource(t, allEndpoints)

}
