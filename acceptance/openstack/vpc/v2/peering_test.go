package v2

import (
	"testing"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/vpc/v1/vpcs"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/vpc/v2/peerings"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestPeeringDeleteNotFound(t *testing.T) {
	client, err := clients.NewVPCV2Client()
	if err != nil {
		t.Fatal(err)
	}
	if err = peerings.Delete(client, "00000000-0000-0000-0000-000000000000"); err == nil {
		t.Fatal("expected deleting a synthetic peering ID to fail")
	}
}

func TestPeeringAcceptNotFound(t *testing.T) {
	client, err := clients.NewVPCV2Client()
	if err != nil {
		t.Fatal(err)
	}
	if _, err = peerings.Accept(client, "00000000-0000-0000-0000-000000000000"); err == nil {
		t.Fatal("expected accepting a synthetic peering ID to fail")
	}
}

func TestPeeringRejectNotFound(t *testing.T) {
	client, err := clients.NewVPCV2Client()
	if err != nil {
		t.Fatal(err)
	}
	if _, err = peerings.Reject(client, "00000000-0000-0000-0000-000000000000"); err == nil {
		t.Fatal("expected rejecting a synthetic peering ID to fail")
	}
}

func TestPeeringLifecycle(t *testing.T) {
	v1Client, err := clients.NewVPCV1Client()
	th.AssertNoErr(t, err)
	v2Client, err := clients.NewVPCV2Client()
	th.AssertNoErr(t, err)

	requestVpc := createPeeringVPC(t, v1Client, "192.168.100.0/24")
	acceptVpc := createPeeringVPC(t, v1Client, "192.168.101.0/24")

	peering, err := peerings.Create(v2Client, peerings.CreateOpts{
		Name:           tools.RandomString("peering-acc-", 3),
		Description:    "VPC peering acceptance test",
		RequestVpcInfo: peerings.VpcInfo{VpcID: requestVpc.ID},
		AcceptVpcInfo:  peerings.VpcInfo{VpcID: acceptVpc.ID},
	})
	th.AssertNoErr(t, err)
	t.Cleanup(func() {
		err := peerings.Delete(v2Client, peering.ID)
		if _, ok := err.(golangsdk.ErrDefault404); !ok {
			th.AssertNoErr(t, err)
		}
	})
	th.AssertEquals(t, requestVpc.ID, peering.RequestVpcInfo.VpcID)
	th.AssertEquals(t, acceptVpc.ID, peering.AcceptVpcInfo.VpcID)

	found, err := peerings.Get(v2Client, peering.ID)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, peering.ID, found.ID)

	all, err := peerings.List(v2Client, peerings.ListOpts{ID: peering.ID})
	th.AssertNoErr(t, err)
	if len(all) != 1 {
		t.Fatalf("expected one peering, got %d", len(all))
	}
	th.AssertEquals(t, peering.ID, all[0].ID)

	updatedName := tools.RandomString("peering-updated-acc-", 3)
	updated, err := peerings.Update(v2Client, peering.ID, peerings.UpdateOpts{
		Name: &updatedName,
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, updatedName, updated.Name)
}

func createPeeringVPC(t *testing.T, client *golangsdk.ServiceClient, cidr string) *vpcs.Vpc {
	t.Helper()
	vpc, err := vpcs.Create(client, vpcs.CreateOpts{
		Name: tools.RandomString("peering-vpc-acc-", 3),
		CIDR: cidr,
	})
	th.AssertNoErr(t, err)
	t.Cleanup(func() {
		deleteSecurityGroupVPC(t, client, vpc.ID)
	})
	return vpc
}
