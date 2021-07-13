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
	defer deleteOBSBucket(t, bucketName)

	smn := createSMNTopic(t)
	defer deleteSMNTopic(t, smn.TopicUrn)

	supportSMN := true
	sendAllKey := true
	createOpts := tracker.CreateOptsWithSMN{
		BucketName: bucketName,
		SimpleMessageNotification: tracker.SimpleMessageNotification{
			IsSupportSMN:          &supportSMN,
			TopicID:               smn.TopicUrn,
			IsSendAllKeyOperation: &sendAllKey,
		},
	}
	ctsTracker, err := tracker.Create(client, createOpts).Extract()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, supportSMN, ctsTracker.SimpleMessageNotification.IsSupportSMN)
	th.AssertEquals(t, supportSMN, ctsTracker.SimpleMessageNotification.IsSendAllKeyOperation)
	th.AssertEquals(t, "enabled", ctsTracker.Status)

	updateOpts := tracker.UpdateOpts{
		Status: "disabled",
	}
	_, err = tracker.Update(client, updateOpts).Extract()
	th.AssertNoErr(t, err)

	listOpts := tracker.ListOpts{
		TrackerName: ctsTracker.TrackerName,
		BucketName:  bucketName,
		Status:      updateOpts.Status,
	}
	trackerList, err := tracker.List(client, listOpts)
	if len(trackerList) == 0 {
		t.Fatal("CTS tracker wasn't found")
	}
	th.AssertEquals(t, updateOpts.Status, trackerList[0].Status)
}
