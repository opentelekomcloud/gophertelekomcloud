package cc

import (
	"os"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	acceptanceer "github.com/opentelekomcloud/gophertelekomcloud/acceptance/openstack/er"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/cc/v3/central_network"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/cc/v3/connection"
	gcb "github.com/opentelekomcloud/gophertelekomcloud/openstack/cc/v3/global_connection_bandwidth"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/cc/v3/policy"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/pointerto"
	az "github.com/opentelekomcloud/gophertelekomcloud/openstack/er/v3/availability-zones"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/er/v3/instance"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestCentralNetworkPolicyLifeCycle(t *testing.T) {
	if os.Getenv("RUN_CC_LIFECYCLE") == "" {
		t.Skip("too slow to run in zuul")
	}
	client, err := clients.NewCCClient()
	th.AssertNoErr(t, err)

	erClient, err := clients.NewERClient()
	th.AssertNoErr(t, err)

	azs, err := az.List(erClient, az.ListOpts{})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, true, len(azs) > 0)
	azCode := azs[0].Code

	erInstance, err := instance.Create(erClient, instance.CreateOpts{
		Name:                     tools.RandomString("acctest-cc-er-", 4),
		Description:              "created by gophertelekomcloud acceptance test",
		Asn:                      64512,
		EnableDefaultAssociation: pointerto.Bool(true),
		EnableDefaultPropagation: pointerto.Bool(true),
		AvailabilityZoneIDs:      []string{azCode},
	})
	th.AssertNoErr(t, err)
	erID := erInstance.Instance.ID
	t.Cleanup(func() {
		th.AssertNoErr(t, instance.Delete(erClient, erID))
		th.AssertNoErr(t, acceptanceer.WaitForInstanceDeleted(erClient, 500, erID))
	})
	th.AssertNoErr(t, acceptanceer.WaitForInstanceAvailable(erClient, 300, erID))

	got, err := instance.Get(erClient, erID)
	th.AssertNoErr(t, err)
	erTableID := got.Instance.DefaultAssociationRouteTableID
	projectID := got.Instance.ProjectID
	regionID := erClient.RegionID

	name := tools.RandomString("acctest_cc_cn-", 4)
	cn, err := central_network.Create(client, central_network.CreateOpts{
		Name:        name,
		Description: "created by gophertelekomcloud acceptance test",
	})
	th.AssertNoErr(t, err)
	t.Cleanup(func() {
		th.AssertNoErr(t, WaitForCentralNetworkAvailable(client, 300, cn.ID))
		th.AssertNoErr(t, central_network.Delete(client, cn.ID))
	})

	policyOpts := policy.CreateOpts{
		CentralNetworkId: cn.ID,
		DefaultPlane:     "default",
		Planes: []policy.PlaneDocument{
			{
				Name: "default",
				AssociateErTables: []policy.AssociateErTable{
					{
						ProjectId:               projectID,
						RegionId:                regionID,
						EnterpriseRouterId:      erID,
						EnterpriseRouterTableId: erTableID,
					},
				},
			},
		},
		ErInstances: []policy.AssociateErInstance{
			{
				EnterpriseRouterId: erID,
				ProjectId:          projectID,
				RegionId:           regionID,
			},
		},
	}

	created, err := policy.Create(client, policyOpts)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, cn.ID, created.CentralNetworkId)

	list, err := policy.List(client, policy.ListOpts{CentralNetworkId: cn.ID})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, true, len(list.CentralNetworkPolicies) > 0)

	_, err = policy.ListChangeSet(client, cn.ID, created.ID)
	th.AssertNoErr(t, err)

	th.AssertNoErr(t, policy.Delete(client, cn.ID, created.ID))

	applyPolicy, err := policy.Create(client, policyOpts)
	th.AssertNoErr(t, err)

	applied, err := policy.Apply(client, cn.ID, applyPolicy.ID)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, applyPolicy.ID, applied.CentralNetworkPolicy.ID)

	th.AssertNoErr(t, WaitForPolicyAvailable(client, 300, cn.ID, applyPolicy.ID))
	th.AssertNoErr(t, WaitForCentralNetworkAvailable(client, 300, cn.ID))

	conns, err := connection.List(client, connection.ListOpts{CentralNetworkId: cn.ID})
	th.AssertNoErr(t, err)

	band, err := gcb.Create(client, gcb.CreateOpts{
		Name:        tools.RandomString("acctest_cc_gcb-", 4),
		Description: "created by gophertelekomcloud acceptance test",
		Bordercross: pointerto.Bool(false),
		Type:        "Region",
		ChargeMode:  "bwd",
		Size:        5,
	})
	th.AssertNoErr(t, err)
	t.Cleanup(func() {
		th.AssertNoErr(t, gcb.Delete(client, band.ID))
	})

	conn := conns.CentralNetworkConnections[0]
	updated, err := connection.Update(client, connection.UpdateOpts{
		CentralNetworkId:            cn.ID,
		ConnectionId:                conn.ID,
		BandwidthType:               "BandwidthPackage",
		BandwidthSize:               5,
		GlobalConnectionBandwidthId: band.ID,
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, band.ID, updated.GlobalConnectionBandwidthId)
}
