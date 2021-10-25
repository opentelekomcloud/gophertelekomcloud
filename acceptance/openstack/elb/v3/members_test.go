package v3

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/openstack"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/elb/v3/members"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestMemberLifecycle(t *testing.T) {
	client, err := clients.NewElbV3Client(t)
	th.AssertNoErr(t, err)

	loadbalancerID := createLoadBalancer(t, client)
	defer deleteLoadbalancer(t, client, loadbalancerID)

	computeClient, err := clients.NewComputeV1Client()
	th.AssertNoErr(t, err)

	ecs := openstack.CreateCloudServer(t, computeClient, openstack.GetCloudServerCreateOpts(t))
	defer openstack.DeleteCloudServer(t, computeClient, ecs.ID)

	poolID := createPool(t, client, loadbalancerID)
	defer deletePool(t, client, loadbalancerID)

	t.Logf("Attempting to create ELBv3 Member")
	memberName := tools.RandomString("create-member-", 3)
	createOpts := members.CreateOpts{
		Address:      ecs.AccessIPv4,
		ProtocolPort: 89,
		Name:         memberName,
	}

	member, err := members.Create(client, poolID, createOpts).Extract()
	th.AssertNoErr(t, err)
	defer func() {
		t.Logf("Attempting to delete ELBv3 Member: %s", member.ID)
		err := members.Delete(client, poolID, member.ID).ExtractErr()
		th.AssertNoErr(t, err)
		t.Logf("Deleted ELBv3 Member: %s", member.ID)
	}()
	th.AssertEquals(t, createOpts.Name, member.Name)
	th.AssertEquals(t, createOpts.ProtocolPort, member.ProtocolPort)
	th.AssertEquals(t, createOpts.Address, member.Address)
	t.Logf("Created ELBv3 Member: %s", member.ID)

	t.Logf("Attempting to update ELBv3 Member: %s", member.ID)
	memberName = tools.RandomString("update-member-", 3)
	updateOpts := members.UpdateOpts{
		Name: memberName,
	}
	_, err = members.Update(client, poolID, member.ID, updateOpts).Extract()
	th.AssertNoErr(t, err)
	t.Logf("Updated ELBv3 Member: %s", member.ID)

	newMember, err := members.Get(client, poolID, member.ID).Extract()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, updateOpts.Name, newMember.Name)
}
