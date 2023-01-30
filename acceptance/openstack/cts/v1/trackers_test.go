package v1

import (
	"os"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/openstack/cts"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/pointerto"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/cts/v1/tracker"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestTrackersLifecycle(t *testing.T) {
	if os.Getenv("RUN_CTS_TRACKER") == "" {
		t.Skip("unstable test")
	}
	client, err := clients.NewCTSV1Client()
	th.AssertNoErr(t, err)

	bucketName := cts.CreateOBSBucket(t)
	t.Cleanup(func() {
		cts.DeleteOBSBucket(t, bucketName)
	})

	t.Logf("Attempting to create CTSv1 Tracker")
	ctsTracker, err := tracker.Create(client, tracker.CreateOpts{
		BucketName: bucketName,
		Lts: tracker.CreateLts{
			IsLtsEnabled: pointerto.Bool(true),
		},
	})

	t.Cleanup(func() {
		t.Logf("Attempting to delete CTSv1 Tracker: %s", ctsTracker.TrackerName)
		err := tracker.Delete(client, ctsTracker.TrackerName)
		th.AssertNoErr(t, err)
		t.Logf("Deleted CTSv1 Tracker: %s", ctsTracker.TrackerName)
	})

	th.AssertNoErr(t, err)
	t.Logf("Created CTSv1 Tracker: %s", ctsTracker.TrackerName)
	th.AssertEquals(t, true, ctsTracker.Lts.IsLtsEnabled)
	th.AssertEquals(t, "enabled", ctsTracker.Status)

	t.Logf("Attempting to update CTSv1 Tracker: %s", ctsTracker.TrackerName)
	_, err = tracker.Update(client, tracker.UpdateOpts{
		BucketName: bucketName,
		Status:     "disabled",
	}, ctsTracker.TrackerName)
	th.AssertNoErr(t, err)
	t.Logf("Updated CTSv1 Tracker: %s", ctsTracker.TrackerName)

	trackerList, err := tracker.Get(client, ctsTracker.TrackerName)
	th.AssertNoErr(t, err)

	th.AssertEquals(t, tracker.UpdateOpts{
		BucketName: bucketName,
		Status:     "disabled",
	}.Status, trackerList.Status)
}
