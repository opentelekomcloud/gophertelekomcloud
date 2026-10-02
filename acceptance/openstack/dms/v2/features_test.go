package v2

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dms/v2/features"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestFeaturesList(t *testing.T) {
	client, err := clients.NewDmsV2Client()
	th.AssertNoErr(t, err)

	listed, err := features.List(client)
	th.AssertNoErr(t, err)
	tools.PrintResource(t, listed)
	th.AssertEquals(t, true, len(listed.Features) > 0)
	th.AssertEquals(t, len(listed.Features), listed.TotalRecord)
	for _, f := range listed.Features {
		th.AssertEquals(t, true, f.FeatureID != "")
	}
}
