package v1

import (
	"os"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/pointerto"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/rms/recorder"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestRecorderLifecycle(t *testing.T) {
	var agency, bucket string
	if agency = os.Getenv("OS_AGENCY_NAME"); agency == "" {
		t.Skip("OS_AGENCY_NAME is required for this test")
	}
	if bucket = os.Getenv("OS_BUCKET_NAME"); bucket == "" {
		t.Skip("OS_BUCKET_NAME is required for this test")
	}
	client, err := clients.NewRMSClient()
	th.AssertNoErr(t, err)

	createOpts := recorder.UpdateOpts{
		DomainId: client.DomainID,
		Channel: recorder.ChannelConfigBody{
			Obs: &recorder.TrackerObsConfigBody{
				BucketName: bucket,
				RegionId:   client.RegionID,
			},
		},
		Selector: recorder.SelectorConfigBody{
			AllSupported:  true,
			ResourceTypes: []string{},
		},
		AgencyName:    agency,
		RetentionDays: pointerto.Int(2557),
	}

	err = recorder.UpdateRecorder(client, createOpts)
	th.AssertNoErr(t, err)
	t.Cleanup(func() {
		err = recorder.DeleteRecorder(client, client.DomainID)
	})
	th.AssertNoErr(t, err)

	resp, err := recorder.GetRecorder(client, client.DomainID)
	th.AssertNoErr(t, err)

	th.AssertEquals(t, createOpts.Channel.Obs.BucketName, resp.Channel.Obs.BucketName)
	th.AssertEquals(t, createOpts.Channel.Obs.RegionId, resp.Channel.Obs.RegionId)
	th.AssertEquals(t, createOpts.Selector.AllSupported, resp.Selector.AllSupported)
	th.AssertEquals(t, createOpts.AgencyName, resp.AgencyName)
	th.AssertEquals(t, *createOpts.RetentionDays, resp.RetentionPeriod)
}
