package er

import (
	"os"
	"testing"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/pointerto"
	tag "github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/er/v3/association"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/er/v3/instance"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/er/v3/route_table"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/er/v3/vpc"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestERAssociationsLifeCycle(t *testing.T) {
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

	t.Logf("Attempting to associate route table with vpc attachemnt")

	associateOpts := association.CreateOpts{
		RouterID:     createResp.Instance.ID,
		RouteTableID: createRtResp.ID,
		AttachmentID: createVpcResp.ID,
	}

	associationResp, err := association.Create(client, associateOpts)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, associateOpts.RouteTableID, associationResp.RouteTableID)
	th.AssertEquals(t, associationResp.State, "pending")

	err = waitForVpcAssociationAvailable(client, 100, createResp.Instance.ID, createRtResp.ID)
	th.AssertNoErr(t, err)

	t.Cleanup(func() {
		t.Logf("Attempting to dissasociate vpc attachemnt")
		err = association.Delete(client, association.DeleteOpts{
			RouterID:     createResp.Instance.ID,
			RouteTableID: createRtResp.ID,
			AttachmentID: createVpcResp.ID,
		})
		th.AssertNoErr(t, err)
	})

}

func waitForVpcAssociationAvailable(client *golangsdk.ServiceClient, secs int, erId, rtId string) error {
	return golangsdk.WaitFor(secs, func() (bool, error) {
		listAssotiationResp, err := association.List(client, association.ListOpts{
			RouterId:     erId,
			RouteTableId: rtId,
		})
		if err != nil {
			return false, err
		}
		if listAssotiationResp.Associations[0].State == "available" {
			return true, nil
		}
		return false, nil
	})
}
