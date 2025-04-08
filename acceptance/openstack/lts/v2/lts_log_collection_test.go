package v2

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	lc "github.com/opentelekomcloud/gophertelekomcloud/openstack/lts/v2/log-collection"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestLtsLogCollectionLifecycle(t *testing.T) {
	client, err := clients.NewLtsV2Client()
	th.AssertNoErr(t, err)

	t.Logf("Attempt to Enable to continue collect logs when the free quota runs out")
	err = lc.Enable(client)
	th.AssertNoErr(t, err)

	t.Logf("Attempt to Disable to continue collect logs when the free quota runs out")
	err = lc.Disable(client)
	th.AssertNoErr(t, err)
}
