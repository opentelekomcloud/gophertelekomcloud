package er

import (
	"os"
	"strings"
	"testing"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	fl "github.com/opentelekomcloud/gophertelekomcloud/openstack/er/v3/flow-logs"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/er/v3/vpc"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/lts/v2/groups"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/lts/v2/streams"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/lts/v2/transfers"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/obs"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestFlowLogsLifeCycle(t *testing.T) {
	routerID := os.Getenv("ER_ROUTER_ID")
	if routerID == "" {
		t.Skip("ER_ROUTER_ID is required for this test")
	}

	vpcId = os.Getenv("OS_VPC_ID")
	networkId = os.Getenv("OS_NETWORK_ID")
	vpcName = tools.RandomString("acctest_vpc_attachments-", 4)
	description = "test vpc attachment"
	client, err := clients.NewERClient()
	th.AssertNoErr(t, err)

	logGroupId, logStreamId, err := createLtsObjects(t)
	th.AssertNoErr(t, err)

	createVpcOpts := vpc.CreateOpts{
		Name:                vpcName,
		RouterID:            routerID,
		VpcId:               vpcId,
		SubnetId:            networkId,
		Description:         description,
		AutoCreateVpcRoutes: true,
	}

	t.Logf("Attempting to create vpc attachment")
	createVpcResp, err := vpc.Create(client, createVpcOpts)
	th.AssertNoErr(t, err)

	err = waitForVpcAttachmentsAvailable(client, 100, routerID, createVpcResp.ID)
	th.AssertNoErr(t, err)

	t.Cleanup(func() {
		t.Logf("Attempting to delete vpc attachment")
		err = vpc.Delete(client, routerID, createVpcResp.ID)
		th.AssertNoErr(t, err)
		err = waitForVpcAttachmentsDeleted(client, 500, routerID, createVpcResp.ID)
		th.AssertNoErr(t, err)
	})

	flowLog, err := fl.Create(client, fl.CreateOpts{
		RouterID: routerID,
		FlowLog: &fl.FlowLog{
			Name:         tools.RandomString("flow-log-attach-", 4),
			Description:  "low log test",
			ResourceType: "attachment",
			ResourceId:   createVpcResp.ID,
			LogGroupId:   logGroupId,
			LogStreamId:  logStreamId,
			LogStoreType: "LTS",
		},
	})
	th.AssertNoErr(t, err)

	err = waitForFlowLogAvailable(client, 100, routerID, flowLog.ID)
	th.AssertNoErr(t, err)

	t.Cleanup(func() {
		t.Logf("Attempting to delete flow log")
		err = fl.Delete(client, routerID, flowLog.ID)
		th.AssertNoErr(t, err)
		err = waitForFlowLogDelete(client, 500, routerID, flowLog.ID)
		th.AssertNoErr(t, err)
	})

	t.Logf("Attempting to list flow logs")
	listFl, err := fl.List(client, routerID, fl.ListOpts{})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 1, len(listFl))

	t.Logf("Attempting to update flow log")
	flUpdated, err := fl.Update(client, flowLog.ID, fl.UpdateOpts{
		RouterID:    routerID,
		Description: "updated",
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "updated", flUpdated.Description)

	t.Logf("Attempting to enable flow log")
	enabled, err := fl.Enable(client, routerID, flowLog.ID)
	th.AssertNoErr(t, err)
	err = waitForFlowLogAvailable(client, 100, routerID, enabled.ID)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, true, enabled.Enabled)

	t.Logf("Attempting to disable flow log")
	disabled, err := fl.Disable(client, routerID, flowLog.ID)
	th.AssertNoErr(t, err)
	err = waitForFlowLogAvailable(client, 100, routerID, disabled.ID)

	t.Logf("Attempting to get flow log")
	getDisabled, err := fl.Get(client, routerID, flowLog.ID)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, false, getDisabled.Enabled)
}

func createLtsObjects(t *testing.T) (string, string, error) {
	client, err := clients.NewLtsV2Client()
	th.AssertNoErr(t, err)

	name := tools.RandomString("test-group-", 3)
	createOpts := groups.CreateOpts{
		LogGroupName: name,
		TTLInDays:    7,
	}
	t.Logf("Attempting to create LTS log group")
	logId, err := groups.CreateLogGroup(client, createOpts)
	th.AssertNoErr(t, err)

	t.Cleanup(func() {
		t.Logf("Attempting to delete LTS log group")
		err = groups.DeleteLogGroup(client, logId)
		th.AssertNoErr(t, err)
	})

	sname := tools.RandomString("test-stream-", 3)
	t.Logf("Attempting to create LTS log stream")
	streamId, err := streams.CreateLogStream(client, streams.CreateOpts{
		GroupId:       logId,
		LogStreamName: sname,
	})
	th.AssertNoErr(t, err)

	t.Cleanup(func() {
		t.Logf("Attempting to delete LTS log stream")
		err = streams.DeleteLogStream(client, streams.DeleteOpts{
			GroupId:  logId,
			StreamId: streamId,
		})
		th.AssertNoErr(t, err)
	})

	obsClient, err := clients.NewOBSClient()
	th.AssertNoErr(t, err)

	bucketName := strings.ToLower(tools.RandomString("obs-flow-log-test", 5))

	_, err = obsClient.CreateBucket(&obs.CreateBucketInput{
		Bucket: bucketName,
	})
	t.Cleanup(func() {
		_, err = obsClient.DeleteBucket(bucketName)
		th.AssertNoErr(t, err)
	})
	th.AssertNoErr(t, err)

	switchOn := false
	createTransferOpts := transfers.CreateLogDumpObsOpts{
		LogGroupId: logId,
		LogStreamIds: []string{
			streamId,
		},
		ObsBucketName: bucketName,
		Type:          "cycle",
		StorageFormat: "RAW",
		SwitchOn:      &switchOn,
		PrefixName:    "test",
		Period:        3,
		PeriodUnit:    "hour",
	}
	logDumpId, err := transfers.CreateLogDumpObs(client, createTransferOpts)
	th.AssertNoErr(t, err)
	t.Logf("Obs log dump created, id: %s", logDumpId)

	t.Cleanup(func() {
		err = transfers.DeleteTransfer(client, logDumpId)
		th.AssertNoErr(t, err)
	})
	return logId, streamId, nil
}

func waitForFlowLogDelete(client *golangsdk.ServiceClient, secs int, erId, flowId string) error {
	return golangsdk.WaitFor(secs, func() (bool, error) {
		_, err := fl.Get(client, erId, flowId)
		if err != nil {
			if _, ok := err.(golangsdk.ErrDefault404); ok {
				return true, nil
			}
			return false, err
		}

		return false, nil
	})
}

func waitForFlowLogAvailable(client *golangsdk.ServiceClient, secs int, erId, flowId string) error {
	return golangsdk.WaitFor(secs, func() (bool, error) {
		flowLog, err := fl.Get(client, erId, flowId)
		if err != nil {
			return false, err
		}
		if flowLog.Status == "available" {
			return true, nil
		}
		return false, nil
	})
}
