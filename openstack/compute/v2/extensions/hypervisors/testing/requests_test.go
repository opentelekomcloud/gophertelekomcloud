package testing

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/compute/v2/extensions/hypervisors"
	"github.com/opentelekomcloud/gophertelekomcloud/testhelper"
	"github.com/opentelekomcloud/gophertelekomcloud/testhelper/client"
)

func TestListHypervisors(t *testing.T) {
	testhelper.SetupHTTP()
	defer testhelper.TeardownHTTP()
	HandleHypervisorListSuccessfully(t)

	actual, err := hypervisors.List(client.ServiceClient())
	testhelper.AssertNoErr(t, err)

	if len(actual) != 2 {
		t.Fatalf("Expected 2 hypervisors, got %d", len(actual))
	}
	testhelper.CheckDeepEquals(t, HypervisorFake, actual[0])
	testhelper.CheckDeepEquals(t, HypervisorFake, actual[1])
}

func TestListAllHypervisors(t *testing.T) {
	testhelper.SetupHTTP()
	defer testhelper.TeardownHTTP()
	HandleHypervisorListSuccessfully(t)

	actual, err := hypervisors.List(client.ServiceClient())
	testhelper.AssertNoErr(t, err)
	testhelper.CheckDeepEquals(t, HypervisorFake, actual[0])
	testhelper.CheckDeepEquals(t, HypervisorFake, actual[1])
}

func TestHypervisorsStatistics(t *testing.T) {
	testhelper.SetupHTTP()
	defer testhelper.TeardownHTTP()
	HandleHypervisorsStatisticsSuccessfully(t)

	expected := HypervisorsStatisticsExpected

	actual, err := hypervisors.GetStatistics(client.ServiceClient())
	testhelper.AssertNoErr(t, err)
	testhelper.CheckDeepEquals(t, &expected, actual)
}

func TestGetHypervisor(t *testing.T) {
	testhelper.SetupHTTP()
	defer testhelper.TeardownHTTP()
	HandleHypervisorGetSuccessfully(t)

	expected := HypervisorFake
	actual, err := hypervisors.Get(client.ServiceClient(), expected.ID)
	testhelper.AssertNoErr(t, err)
	testhelper.CheckDeepEquals(t, &expected, actual)
}

func TestHypervisorsUptime(t *testing.T) {
	testhelper.SetupHTTP()
	defer testhelper.TeardownHTTP()
	HandleHypervisorUptimeSuccessfully(t)

	expected := HypervisorUptimeExpected

	actual, err := hypervisors.GetUptime(client.ServiceClient(), HypervisorFake.ID)
	testhelper.AssertNoErr(t, err)
	testhelper.CheckDeepEquals(t, &expected, actual)
}
