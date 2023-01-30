package v2

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/openstack/cts"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/cts/v1/tracker"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/cts/v2/traces"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestTraces(t *testing.T) {
	cv1, err := clients.NewCTSV1Client()
	th.AssertNoErr(t, err)

	bucketName := cts.CreateOBSBucket(t)
	t.Cleanup(func() {
		cts.DeleteOBSBucket(t, bucketName)
	})

	t.Logf("Attempting to create CTSv1 Tracker")
	ctsTracker, err := tracker.Create(cv1, tracker.CreateOpts{
		BucketName: bucketName,
	})

	t.Cleanup(func() {
		t.Logf("Attempting to delete CTSv1 Tracker: %s", ctsTracker.TrackerName)
		err := tracker.Delete(cv1, ctsTracker.TrackerName)
		th.AssertNoErr(t, err)
		t.Logf("Deleted CTSv1 Tracker: %s", ctsTracker.TrackerName)
	})

	th.AssertNoErr(t, err)
	t.Logf("Created CTSv1 Tracker: %s", ctsTracker.TrackerName)

	cv2, err := clients.NewCTSV2Client()
	th.AssertNoErr(t, err)

	list, err := traces.List(cv2, ctsTracker.TrackerName, traces.ListTracesOpts{})
	th.AssertNoErr(t, err)

	t.Logf("Number of Tracker Traces: %d", len(list.Traces))
	th.AssertEquals(t, true, len(list.Traces) > 0)
}
