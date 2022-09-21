package v1

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/cts/v1/tracker"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestTrackersLifecycle(t *testing.T) {
	client, err := clients.NewCTSV1Client()
	th.AssertNoErr(t, err)

	bucketName := createOBSBucket(t)
	t.Cleanup(func() {
		deleteOBSBucket(t, bucketName)
	})

	smn := createSMNTopic(t)
	t.Cleanup(func() {
		deleteSMNTopic(t, smn.TopicUrn)
	})

	t.Logf("Attempting to create CTSv1 Tracker")
	f := false
	createOpts := tracker.CreateOpts{
		BucketName: bucketName,
		Lts: tracker.CreateLts{
			IsLtsEnabled: &f,
		},
	}
	ctsTracker, err := tracker.Create(client, createOpts)
	t.Cleanup(func() {
		t.Logf("Attempting to delete CTSv1 Tracker: %s", ctsTracker.TrackerName)
		err := tracker.Delete(client)
		th.AssertNoErr(t, err)
		t.Logf("Deleted CTSv1 Tracker: %s", ctsTracker.TrackerName)
	})
	th.AssertNoErr(t, err)
	t.Logf("Created CTSv1 Tracker: %s", ctsTracker.TrackerName)
	th.AssertEquals(t, true, ctsTracker.Lts.IsLtsEnabled)
	th.AssertEquals(t, "enabled", ctsTracker.Status)

	t.Logf("Attempting to update CTSv1 Tracker: %s", ctsTracker.TrackerName)
	updateOpts := tracker.UpdateOpts{
		BucketName: bucketName,
		Status:     "disabled",
	}
	_, err = tracker.Update(client, updateOpts)
	th.AssertNoErr(t, err)
	t.Logf("Updated CTSv1 Tracker: %s", ctsTracker.TrackerName)

	listOpts := tracker.ListOpts{
		TrackerName: ctsTracker.TrackerName,
	}
	trackerList, err := tracker.List(client, listOpts)
	th.AssertNoErr(t, err)
	if len(trackerList) == 0 {
		t.Fatalf("CTS tracker wasn't found: %s", ctsTracker.TrackerName)
	}
	th.AssertEquals(t, updateOpts.Status, trackerList[0].Status)
}
