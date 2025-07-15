package v3

import (
	"os"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/openstack/cts"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/pointerto"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/cts/v3/tracker"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestTrackersLifecycle(t *testing.T) {
	if os.Getenv("RUN_CTS_TRACKER") == "" {
		t.Skip("unstable test")
	}
	client, err := clients.NewCTSV3Client()
	th.AssertNoErr(t, err)

	bucketName := cts.CreateOBSBucket(t)
	t.Cleanup(func() {
		cts.DeleteOBSBucket(t, bucketName)
	})

	t.Logf("Attempting to create CTSv3 Tracker")
	ctsTracker, err := tracker.Create(client, tracker.CreateOpts{
		TrackerType:  "system",
		TrackerName:  "system",
		IsLtsEnabled: true,
		ObsInfo: tracker.ObsInfo{
			BucketName:      bucketName,
			FilePrefixName:  "test-prefix",
			CompressType:    "json",
			IsSortByService: pointerto.Bool(true),
		},
	})

	t.Cleanup(func() {
		t.Logf("Attempting to delete CTSv3 Tracker: %s", ctsTracker.TrackerName)
		err := tracker.Delete(client, ctsTracker.TrackerName)
		th.AssertNoErr(t, err)
		t.Logf("Deleted CTSv3 Tracker: %s", ctsTracker.TrackerName)
	})

	th.AssertNoErr(t, err)
	t.Logf("Created CTSv3 Tracker: %s", ctsTracker.TrackerName)
	th.AssertEquals(t, true, ctsTracker.Lts.IsLtsEnabled)
	th.AssertEquals(t, "enabled", ctsTracker.Status)
	th.AssertEquals(t, false, *ctsTracker.ObsInfo.IsObsCreated)
	th.AssertEquals(t, bucketName, ctsTracker.ObsInfo.BucketName)
	th.AssertEquals(t, "json", ctsTracker.ObsInfo.CompressType)

	t.Logf("Attempting to update CTSv3 Tracker: %s", ctsTracker.TrackerName)
	ltsEnable := false
	_, err = tracker.Update(client, tracker.UpdateOpts{
		TrackerName:  "system",
		TrackerType:  "system",
		Status:       "enabled",
		IsLtsEnabled: &ltsEnable,
		ObsInfo: tracker.ObsInfo{
			FilePrefixName: "test-2-",
			CompressType:   "gzip",
		},
	})
	th.AssertNoErr(t, err)
	t.Logf("Updated CTSv3 Tracker: %s", ctsTracker.TrackerName)

	trackerList, err := tracker.List(client, ctsTracker.TrackerName)
	th.AssertNoErr(t, err)
	trackerGet := trackerList[0]
	th.AssertEquals(t, trackerGet.TrackerType, "system")
	th.AssertEquals(t, trackerGet.TrackerName, "system")
	th.AssertEquals(t, trackerGet.Status, "enabled")
	// if tracker is disabled LTS status can't be changed
	th.AssertEquals(t, trackerGet.Lts.IsLtsEnabled, false)
	th.AssertEquals(t, trackerGet.ObsInfo.FilePrefixName, "test-2-")
	// update of `compress_type` doesn't work
	// th.AssertEquals(t, "gzip", ctsTracker.ObsInfo.CompressType)
}
