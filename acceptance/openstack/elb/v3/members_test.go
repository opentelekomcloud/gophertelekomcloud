package v3

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/openstack"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/pointerto"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/elb/v3/members"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/elb/v3/pools"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestMemberLifecycle(t *testing.T) {
	client, err := clients.NewElbV3Client()
	th.AssertNoErr(t, err)

	loadbalancerID := createLoadBalancer(t, client)
	t.Cleanup(func() {
		deleteLoadbalancer(t, client, loadbalancerID)
	})

	poolID := createPool(t, client, loadbalancerID)
	t.Cleanup(func() { deletePool(t, client, poolID) })

	t.Logf("Attempting to create ELBv3 Member")
	memberName := tools.RandomString("create-member-", 3)

	createOpts := members.CreateOpts{
		Address:      openstack.ValidIP(t, clients.EnvOS.GetEnv("NETWORK_ID")),
		ProtocolPort: pointerto.Int(89),
		Name:         memberName,
		Weight:       pointerto.Int(1),
	}

	member, err := members.Create(client, poolID, createOpts)
	th.AssertNoErr(t, err)
	t.Cleanup(func() {
		t.Logf("Attempting to delete ELBv3 Member: %s", member.Id)
		err := members.Delete(client, poolID, member.Id)
		th.AssertNoErr(t, err)
		t.Logf("Deleted ELBv3 Member: %s", member.Id)
	})
	th.AssertEquals(t, createOpts.Name, member.Name)
	th.AssertEquals(t, createOpts.ProtocolPort, member.ProtocolPort)
	th.AssertEquals(t, createOpts.Address, member.Address)
	th.AssertEquals(t, *createOpts.Weight, member.Weight)
	t.Logf("Created ELBv3 Member: %s", member.Id)

	t.Logf("Attempting to update ELBv3 Member: %s", member.Id)
	memberName = ""
	updateOpts := members.UpdateOpts{
		Name:   memberName,
		Weight: pointerto.Int(0),
	}
	_, err = members.Update(client, poolID, member.Id, updateOpts)
	th.AssertNoErr(t, err)
	t.Logf("Updated ELBv3 Member: %s", member.Id)

	newMember, err := members.Get(client, poolID, member.Id)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, updateOpts.Name, newMember.Name)
	th.AssertEquals(t, updateOpts.Weight, newMember.Weight)

	updateOptsPool := pools.UpdateOpts{
		MemberDeletionProtectionEnable: pointerto.Bool(false),
	}
	_, err = pools.Update(client, poolID, updateOptsPool)
	th.AssertNoErr(t, err)
	t.Logf("Updated ELBv3 Pool: %s", poolID)
}
