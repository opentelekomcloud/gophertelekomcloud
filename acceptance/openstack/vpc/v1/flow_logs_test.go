package v1

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/pointerto"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/lts/v2/groups"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/lts/v2/streams"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/networking/v1/vpcs"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/vpc/v1/flow_logs"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestVPCFlowLogLifecycle(t *testing.T) {
	vpcClient, err := clients.NewNetworkV1Client()
	th.AssertNoErr(t, err)
	ltsClient, err := clients.NewLtsV2Client()
	th.AssertNoErr(t, err)

	vpc, err := vpcs.Create(vpcClient, vpcs.CreateOpts{
		Name: tools.RandomString("vpc-flow-log-", 4),
	}).Extract()
	th.AssertNoErr(t, err)
	t.Cleanup(func() {
		th.AssertNoErr(t, vpcs.Delete(vpcClient, vpc.ID).ExtractErr())
	})

	logGroupID, err := groups.Create(ltsClient, groups.CreateOpts{
		LogGroupName: tools.RandomString("vpc-flow-log-", 4),
		TTLInDays:    7,
	})
	th.AssertNoErr(t, err)
	t.Cleanup(func() {
		th.AssertNoErr(t, groups.Delete(ltsClient, logGroupID))
	})

	logTopicID, err := streams.Create(ltsClient, streams.CreateOpts{
		GroupId:       logGroupID,
		LogStreamName: tools.RandomString("vpc-flow-log-", 4),
	})
	th.AssertNoErr(t, err)
	t.Cleanup(func() {
		th.AssertNoErr(t, streams.Delete(ltsClient, streams.DeleteOpts{
			GroupId:  logGroupID,
			StreamId: logTopicID,
		}))
	})

	created, err := flow_logs.Create(vpcClient, flow_logs.CreateOpts{
		Name:         tools.RandomString("vpc-flow-log-", 4),
		Description:  "VPC flow log acceptance test",
		ResourceType: "vpc",
		ResourceID:   vpc.ID,
		TrafficType:  "all",
		LogGroupID:   logGroupID,
		LogTopicID:   logTopicID,
		IndexEnabled: pointerto.Bool(true),
	})
	th.AssertNoErr(t, err)
	flowLogID := created.ID
	t.Cleanup(func() {
		if flowLogID != "" {
			th.AssertNoErr(t, flow_logs.Delete(vpcClient, flowLogID))
		}
	})

	got, err := flow_logs.Get(vpcClient, created.ID)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, created.ID, got.ID)

	listed, err := flow_logs.List(vpcClient, flow_logs.ListOpts{ID: created.ID})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 1, len(listed))
	th.AssertEquals(t, created.ID, listed[0].ID)

	updated, err := flow_logs.Update(vpcClient, created.ID, flow_logs.UpdateOpts{
		Name:        pointerto.String(tools.RandomString("vpc-flow-log-updated-", 4)),
		Description: pointerto.String("Updated VPC flow log acceptance test"),
		AdminState:  pointerto.Bool(false),
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, false, updated.AdminState)

	th.AssertNoErr(t, flow_logs.Delete(vpcClient, created.ID))
	flowLogID = ""
}
