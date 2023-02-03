package testing

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/compute/v2/extensions/networks"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
	"github.com/opentelekomcloud/gophertelekomcloud/testhelper/client"
)

func TestList(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListSuccessfully(t)

	actual, err := networks.List(client.ServiceClient())
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, ExpectedNetworkSlice, actual)
}

func TestGet(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleGetSuccessfully(t)

	actual, err := networks.Get(client.ServiceClient(), "20c8acc0-f747-4d71-a389-46d078ebf000")
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &SecondNetwork, actual)
}
