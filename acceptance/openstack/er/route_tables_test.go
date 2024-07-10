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
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/er/v3/route_table"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

var (
	rtName        = tools.RandomString("acctest_route-table-", 4)
	descriptionRt = "test vpc attachment"
)

func TestRouteTableLifeCycle(t *testing.T) {
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

	defer func(client *golangsdk.ServiceClient, id string) {
		t.Logf("Attempting to delete enterprise router")
		err = instance.Delete(client, id)
		th.AssertNoErr(t, err)
		err = waitForInstanceDeleted(client, 500, createResp.Instance.ID)
	}(client, createResp.Instance.ID)

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
	th.AssertEquals(t, createRouteTableOpts.Name, createRtResp.Name)
	th.AssertEquals(t, *createRouteTableOpts.Description, createRtResp.Description)

	err = waitForRouteTableAvailable(client, 300, createResp.Instance.ID, createRtResp.ID)
	th.AssertNoErr(t, err)

	defer func(client *golangsdk.ServiceClient, erId, rtId string) {
		t.Logf("Attempting to delete route table")
		err = route_table.Delete(client, erId, rtId)
		th.AssertNoErr(t, err)
		err = waitForRouteTableDeleted(client, 500, erId, rtId)
	}(client, createResp.Instance.ID, createRtResp.ID)

	t.Logf("Attempting to update route table")
	updateRtResp, err := route_table.Update(client, route_table.UpdateOpts{
		RouterID:     createResp.Instance.ID,
		RouteTableId: createRtResp.ID,
		Name:         rtName + "-updated",
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, updateRtResp.Name, rtName+"-updated")

	t.Logf("Attempting to list route table")
	listResp, err := route_table.List(client, route_table.ListOpts{
		RouterId: createResp.Instance.ID,
	})
	th.AssertNoErr(t, err)
	tools.PrintResource(t, listResp)
}

func waitForRouteTableAvailable(client *golangsdk.ServiceClient, secs int, erId, routeTableId string) error {
	return golangsdk.WaitFor(secs, func() (bool, error) {
		rtInstance, err := route_table.Get(client, erId, routeTableId)
		if err != nil {
			return false, err
		}
		if rtInstance.State == "available" {
			return true, nil
		}
		return false, nil
	})
}

func waitForRouteTableDeleted(client *golangsdk.ServiceClient, secs int, erId, rtId string) error {
	return golangsdk.WaitFor(secs, func() (bool, error) {
		_, err := route_table.Get(client, erId, rtId)
		if err != nil {
			if _, ok := err.(golangsdk.ErrDefault404); ok {
				return true, nil
			}
			return false, err
		}

		return false, nil
	})
}
