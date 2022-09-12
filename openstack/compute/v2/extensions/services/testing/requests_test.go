package testing

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/compute/v2/extensions/services"
	"github.com/opentelekomcloud/gophertelekomcloud/testhelper"
	"github.com/opentelekomcloud/gophertelekomcloud/testhelper/client"
)

func TestListServices(t *testing.T) {
	testhelper.SetupHTTP()
	defer testhelper.TeardownHTTP()
	HandleListSuccessfully(t)

	actual, err := services.List(client.ServiceClient())
	testhelper.AssertNoErr(t, err)

	if len(actual) != 4 {
		t.Fatalf("Expected 4 services, got %d", len(actual))
	}
	testhelper.CheckDeepEquals(t, FirstFakeService, actual[0])
	testhelper.CheckDeepEquals(t, SecondFakeService, actual[1])
	testhelper.CheckDeepEquals(t, ThirdFakeService, actual[2])
	testhelper.CheckDeepEquals(t, FourthFakeService, actual[3])
}
