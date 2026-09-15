package v1

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/vpc/v1/routetables"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestRouteTableLifecycle(t *testing.T) {
	client, err := clients.NewVPCV1Client()
	th.AssertNoErr(t, err)

	vpc := createSubnetVPC(t, client)
	t.Cleanup(func() { deleteTestVPC(t, client, vpc.ID) })

	table, err := routetables.Create(client, routetables.CreateOpts{
		Name: tools.RandomString("route-table-acc-", 3), VpcID: vpc.ID, Description: "route table acceptance test",
	})
	th.AssertNoErr(t, err)
	tableID := table.ID
	t.Cleanup(func() {
		if tableID != "" {
			th.AssertNoErr(t, routetables.Delete(client, tableID))
		}
	})

	found, err := routetables.Get(client, tableID)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, tableID, found.ID)

	all, err := routetables.List(client, routetables.ListOpts{ID: tableID, VpcID: vpc.ID})
	th.AssertNoErr(t, err)
	if len(all) != 1 {
		t.Fatalf("expected one route table, got %d", len(all))
	}

	updatedName := tools.RandomString("route-table-updated-acc-", 3)
	updated, err := routetables.Update(client, tableID, routetables.UpdateOpts{Name: updatedName})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, updatedName, updated.Name)

	subnet := createTestSubnet(t, client, vpc.ID)
	t.Cleanup(func() { deleteTestSubnet(t, client, vpc.ID, subnet.ID) })

	associated, err := routetables.Associate(client, tableID, routetables.AssociateOpts{Subnets: []string{subnet.NetworkID}})
	th.AssertNoErr(t, err)
	subnetAssociated := true
	t.Cleanup(func() {
		if subnetAssociated && tableID != "" {
			_, err := routetables.Disassociate(client, tableID, routetables.DisassociateOpts{Subnets: []string{subnet.NetworkID}})
			th.AssertNoErr(t, err)
		}
	})
	if !containsRouteTableSubnet(associated, subnet.NetworkID) {
		t.Fatalf("route table does not contain associated subnet %s", subnet.NetworkID)
	}

	_, err = routetables.Disassociate(client, tableID, routetables.DisassociateOpts{Subnets: []string{subnet.NetworkID}})
	th.AssertNoErr(t, err)
	subnetAssociated = false

	th.AssertNoErr(t, routetables.Delete(client, tableID))
	tableID = ""
}

func containsRouteTableSubnet(table *routetables.RouteTable, id string) bool {
	for _, subnet := range table.Subnets {
		if subnet.ID == id {
			return true
		}
	}
	return false
}
