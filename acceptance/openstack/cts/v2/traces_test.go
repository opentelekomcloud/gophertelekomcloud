package v2

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/openstack/cts"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/cts/v2/traces"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestTraces(t *testing.T) {
	bucketName := cts.CreateOBSBucket(t)
	t.Cleanup(func() {
		cts.DeleteOBSBucket(t, bucketName)
	})

	cv2, err := clients.NewCTSV2Client()
	th.AssertNoErr(t, err)

	list, err := traces.List(cv2, "system", traces.ListTracesOpts{})
	th.AssertNoErr(t, err)

	t.Logf("Number of Tracker Traces: %d", len(list.Traces))
	th.AssertEquals(t, true, len(list.Traces) > 0)
	tools.PrintResource(t, list)
}
