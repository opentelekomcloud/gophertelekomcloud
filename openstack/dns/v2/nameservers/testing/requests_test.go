package testing

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dns/v2/nameservers"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
	"github.com/opentelekomcloud/gophertelekomcloud/testhelper/client"
)

func TestList(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListSuccessfully(t)

	zoneID := "2c9eb155587194ec01587224c9f90149"
	actual, err := nameservers.List(client.ServiceClient(), zoneID).Extract()

	th.AssertNoErr(t, err)
	th.CheckEquals(t, 2, len(actual))
}
