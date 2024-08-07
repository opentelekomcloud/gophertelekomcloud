package er

import (
	"os"
	"testing"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/pointerto"
	tag "github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/er/v3/instance"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/er/v3/propagation"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/er/v3/route_table"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/er/v3/vpc"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestERPropagationsLifeCycle(t *testing.T) {
	if os.Getenv("RUN_ER_LIFECYCLE") == "" {
		t.Skip("too slow to run in zuul")
	}

	client, err := clients.NewERClient()
	th.AssertNoErr(t, err)

	createOpts := instance.CreateOpts{
		Name:                        tools.RandomString("acctest_er_router-", 4),
		Asn:                         64512,
		AutoAcceptSharedAttachments: pointerto.Bool(true),
		AvailabilityZoneIDs: []string{
			"eu-de-01",
			"eu-de-02",
		},
	}

	t.Logf("Attempting to create enterprise router")

	createResp, err := instance.Create(client, createOpts)
	th.AssertNoErr(t, err)

	err = waitForInstanceAvailable(client, 100, createResp.Instance.ID)
	th.AssertNoErr(t, err)

	t.Cleanup(func() {
		t.Logf("Attempting to delete enterprise router")
		err = instance.Delete(client, createResp.Instance.ID)
		th.AssertNoErr(t, err)
		err = waitForInstanceDeleted(client, 500, createResp.Instance.ID)
	})

	createRouteTableOpts := route_table.CreateOpts{
		Name:        rtName,
		RouterID:    createResp.Instance.ID,
		Description: &descriptionRt,
		Tags: []tag.ResourceTag{
			{
				Key:   "muh",
				Value: "muh",
			},
			{
				Key:   "test",
				Value: "test",
			},
		},
	}

	t.Logf("Attempting to create route table")
	createRtResp, err := route_table.Create(client, createRouteTableOpts)
	th.AssertNoErr(t, err)

	err = waitForRouteTableAvailable(client, 300, createResp.Instance.ID, createRtResp.ID)
	th.AssertNoErr(t, err)

	t.Cleanup(func() {
		t.Logf("Attempting to delete route table")
		err = route_table.Delete(client, createResp.Instance.ID, createRtResp.ID)
		th.AssertNoErr(t, err)
		err = waitForRouteTableDeleted(client, 500, createResp.Instance.ID, createRtResp.ID)
	})

	t.Logf("Attempting to create vpc attachemnt")

	createVpcOpts := vpc.CreateOpts{
		Name:                vpcName,
		RouterID:            createResp.Instance.ID,
		VpcId:               vpcId,
		SubnetId:            networkId,
		Description:         description,
		AutoCreateVpcRoutes: true,
		Tags: []tag.ResourceTag{
			{
				Key:   "muh",
				Value: "muh",
			},
			{
				Key:   "test",
				Value: "test",
			},
		},
	}

	createVpcResp, err := vpc.Create(client, createVpcOpts)
	th.AssertNoErr(t, err)

	err = waitForVpcAttachmentsAvailable(client, 100, createResp.Instance.ID, createVpcResp.ID)
	th.AssertNoErr(t, err)

	t.Cleanup(func() {
		t.Logf("Attempting to delete vpc attachemnt")
		err = vpc.Delete(client, createResp.Instance.ID, createVpcResp.ID)
		th.AssertNoErr(t, err)
		err = waitForVpcAttachmentsDeleted(client, 500, createResp.Instance.ID, createVpcResp.ID)
		th.AssertNoErr(t, err)
	})

	t.Logf("Attempting to create enterprise route propagation")

	propagateOpts := propagation.CreateOpts{
		RouterID:     createResp.Instance.ID,
		RouteTableID: createRtResp.ID,
		AttachmentID: createVpcResp.ID,
	}

	propagateResp, err := propagation.Create(client, propagateOpts)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, propagateOpts.RouteTableID, propagateResp.RouteTableID)
	th.AssertEquals(t, propagateResp.State, "pending")

	err = waitForVpcPropagationAvailable(client, 100, createResp.Instance.ID, createRtResp.ID)
	th.AssertNoErr(t, err)

	t.Cleanup(func() {
		t.Logf("Attempting to disable propagation")
		err = propagation.Delete(client, propagation.DeleteOpts{
			RouterID:     createResp.Instance.ID,
			RouteTableID: createRtResp.ID,
			AttachmentID: createVpcResp.ID,
		})
		th.AssertNoErr(t, err)
	})

}

func waitForVpcPropagationAvailable(client *golangsdk.ServiceClient, secs int, erId, rtId string) error {
	return golangsdk.WaitFor(secs, func() (bool, error) {
		listPropagationResp, err := propagation.List(client, propagation.ListOpts{
			RouterId:     erId,
			RouteTableId: rtId,
		})
		if err != nil {
			return false, err
		}
		if listPropagationResp.Propagations[0].State == "available" {
			return true, nil
		}
		return false, nil
	})
}
