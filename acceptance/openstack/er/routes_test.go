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
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/er/v3/route"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/er/v3/route_table"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/er/v3/vpc"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestRoutesLifeCycle(t *testing.T) {
	if os.Getenv("RUN_ER_LIFECYCLE") == "" {
		t.Skip("too slow to run in zuul")
	}
	client, err := clients.NewERClient()
	th.AssertNoErr(t, err)

	routerName := tools.RandomString("acctest_er_router-", 4)

	createOpts := instance.CreateOpts{
		Name:                        routerName,
		Description:                 "terraform created enterprise router",
		Asn:                         64512,
		EnableDefaultAssociation:    pointerto.Bool(true),
		EnableDefaultPropagation:    pointerto.Bool(true),
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
	}

	t.Logf("Attempting to create route table")
	createRtResp, err := route_table.Create(client, createRouteTableOpts)
	th.AssertNoErr(t, err)

	err = waitForRouteTableAvailable(client, 300, createResp.Instance.ID, createRtResp.ID)
	th.AssertNoErr(t, err)

	rtId := createRtResp.ID

	t.Cleanup(func() {
		t.Logf("Attempting to delete route table")
		err = route_table.Delete(client, createResp.Instance.ID, createRtResp.ID)
		th.AssertNoErr(t, err)
		err = waitForRouteTableDeleted(client, 500, createResp.Instance.ID, createRtResp.ID)
	})

	t.Logf("Attempting to create route")

	routeResp, err := route.Create(client, route.CreateOpts{
		RouteTableId: rtId,
		Destination:  "192.168.0.0/16",
		IsBlackhole:  pointerto.Bool(true),
	})

	th.AssertNoErr(t, err)
	th.AssertEquals(t, routeResp.Destination, "192.168.0.0/16")
	th.AssertEquals(t, routeResp.IsBlackhole, true)
	th.AssertEquals(t, routeResp.State, "pending")

	err = waitForRouteAvailable(client, 100, rtId, routeResp.ID)
	th.AssertNoErr(t, err)

	t.Cleanup(func() {
		t.Logf("Attempting to delete route")
		err = route.Delete(client, rtId, routeResp.ID)
		th.AssertNoErr(t, err)
	})

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

	t.Logf("Attempting to create vpc attachemnt")
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

	updateRoute, err := route.Update(client, route.UpdateOpts{
		RouteTableId: rtId,
		RouteId:      routeResp.ID,
		IsBlackhole:  pointerto.Bool(false),
		AttachmentId: createVpcResp.ID,
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, updateRoute.Destination, "192.168.0.0/16")
	th.AssertEquals(t, updateRoute.IsBlackhole, false)
	th.AssertEquals(t, updateRoute.State, "modifying")
	th.AssertEquals(t, updateRoute.Attachments[0].ResourceType, "vpc")

	listResp, err := route.List(client, route.ListOpts{
		RouteTableId: rtId,
		Destination: []string{
			"192.168.0.0/16",
		},
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, len(listResp.Routes), 1)
	th.AssertEquals(t, listResp.Routes[0].Destination, "192.168.0.0/16")

	listStaticResp, err := route.ListStatic(client, route.ListStaticOpts{
		RouteTableId: rtId,
	})
	th.AssertNoErr(t, err)
	tools.PrintResource(t, listStaticResp)
}

func waitForRouteAvailable(client *golangsdk.ServiceClient, secs int, rtId, routeId string) error {
	return golangsdk.WaitFor(secs, func() (bool, error) {
		rtInstance, err := route.Get(client, rtId, routeId)
		if err != nil {
			return false, err
		}
		if rtInstance.State == "available" {
			return true, nil
		}
		return false, nil
	})
}
