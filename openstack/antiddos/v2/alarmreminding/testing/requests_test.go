package testing

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/antiddos/v2/alarmreminding"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
	"github.com/opentelekomcloud/gophertelekomcloud/testhelper/client"
)

func TestQueryTraffic(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleWarnAlertSuccessfully(t)

	actual, err := alarmreminding.WarnAlert(client.ServiceClient()).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &WarnAlertResponse, actual)
}
