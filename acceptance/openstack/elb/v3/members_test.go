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

func iInt(v int) *int {
	return &v
}

func TestMemberLifecycle(t *testing.T) {
	client, err := clients.NewElbV3Client()
	th.AssertNoErr(t, err)

	loadbalancerID := createLoadBalancer(t, client)
	defer deleteLoadbalancer(t, client, loadbalancerID)

	poolID := createPool(t, client, loadbalancerID)
	defer deletePool(t, client, poolID)

	t.Logf("Attempting to create ELBv3 Member")
	memberName := tools.RandomString("create-member-", 3)

	createOpts := members.CreateOpts{
		Address:      openstack.ValidIP(t, clients.EnvOS.GetEnv("NETWORK_ID")),
		ProtocolPort: 89,
		Name:         memberName,
		Weight:       iInt(1),
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
	th.AssertEquals(t, *createOpts.Weight, member.Weight)
	t.Logf("Created ELBv3 Member: %s", member.ID)

	t.Logf("Attempting to update ELBv3 Member: %s", member.ID)
	memberName = ""
	updateOpts := members.UpdateOpts{
		Name:   &memberName,
		Weight: iInt(0),
	}
	_, err = members.Update(client, poolID, member.ID, updateOpts).Extract()
	th.AssertNoErr(t, err)
	t.Logf("Updated ELBv3 Member: %s", member.ID)

	newMember, err := members.Get(client, poolID, member.ID).Extract()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, *updateOpts.Name, newMember.Name)
	th.AssertEquals(t, *updateOpts.Weight, newMember.Weight)

	updateOptsPool := pools.UpdateOpts{
		DeletionProtectionEnable: pointerto.Bool(false),
	}
	_, err = pools.Update(client, poolID, updateOptsPool).Extract()
	th.AssertNoErr(t, err)
	t.Logf("Updated ELBv3 Pool: %s", poolID)
}
