package v3

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/pointerto"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/elb/v3/ipgroups"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/elb/v3/listeners"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestIpGroupList(t *testing.T) {
	client, err := clients.NewElbV3Client()
	th.AssertNoErr(t, err)

	ipgroupsLists, err := ipgroups.List(client, ipgroups.ListOpts{}).AllPages()
	th.AssertNoErr(t, err)
	ipgroupsList, err := ipgroups.ExtractIpGroups(ipgroupsLists)
	th.AssertNoErr(t, err)

	for _, gr := range ipgroupsList {
		tools.PrintResource(t, gr)
	}
}

func TestIpGroupsLifecycle(t *testing.T) {
	client, err := clients.NewElbV3Client()
	th.AssertNoErr(t, err)

	loadbalancerID := createLoadBalancer(t, client)
	defer deleteLoadbalancer(t, client, loadbalancerID)

	t.Logf("Attempting to create ELBv3 IpGroup")
	ipGroup, err := ipgroups.Create(client, ipgroups.CreateOpts{
		Description: "some interesting description",
		Name:        tools.RandomString("create-ip-group-", 3),
		IpList: []ipgroups.IpGroupOption{ipgroups.IpGroupOption{
			Ip:          "192.168.10.10",
			Description: "first",
		}},
	})
	th.AssertNoErr(t, err)

	t.Cleanup(func() {
		t.Logf("Attempting to delete ELBv3 IpGroup: %s", ipGroup.ID)
		th.AssertNoErr(t, ipgroups.Delete(client, ipGroup.ID))
		t.Logf("Deleted ELBv3 IpGroup: %s", ipGroup.ID)
	})

	t.Logf("Attempting to update ELBv3 IpGroup: %s", ipGroup.ID)
	_, err = ipgroups.Update(client, ipGroup.ID, ipgroups.UpdateOpts{
		Name: tools.RandomString("update-ip-group-", 3),
		IpList: []ipgroups.IpGroupOption{
			{
				Ip:          "192.168.10.12",
				Description: "third",
			},
			{
				Ip:          "192.168.10.13",
				Description: "fourth",
			},
		},
	})
	th.AssertNoErr(t, err)
	t.Logf("Updated ELBv3 ipGroup: %s", ipGroup.ID)

	updatedIpGroup, err := ipgroups.Get(client, ipGroup.ID)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, tools.RandomString("update-ip-group-", 3), updatedIpGroup.Name)

	ipGroupsSlices, err := ipgroups.List(client, ipgroups.ListOpts{}).AllPages()
	th.AssertNoErr(t, err)
	ipGroupsSlice, err := ipgroups.ExtractIpGroups(ipGroupsSlices)
	th.AssertNoErr(t, err)

	th.AssertEquals(t, 1, len(ipGroupsSlice))
	th.AssertDeepEquals(t, *updatedIpGroup, ipGroupsSlice[0])

	t.Logf("Attempting to create ELBv3 Listener with ipGroup association")
	listener, err := listeners.Create(client, listeners.CreateOpts{
		LoadbalancerID:  loadbalancerID,
		Protocol:        "HTTP",
		ProtocolPort:    80,
		EnhanceL7policy: pointerto.Bool(true),
		IpGroup: &listeners.IpGroup{
			IpGroupID: ipGroup.ID,
			Enable:    pointerto.Bool(true),
		},
	})
	th.AssertNoErr(t, err)
	t.Cleanup(func() {
		deleteListener(t, client, listener.ID)
	})

	updatedIpList, err := ipgroups.UpdateIpList(client, ipGroup.ID, ipgroups.UpdateOpts{
		IpList: []ipgroups.IpGroupOption{
			{
				Ip:          "192.168.10.12",
				Description: "third",
			},
			{
				Ip:          "192.168.10.13",
				Description: "fourth",
			},
			{
				Ip:          "192.168.10.14",
				Description: "fifth",
			},
		}})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 3, len(updatedIpList.IpList))
	th.AssertEquals(t, listener.ID, updatedIpList.Listeners[0].ID)

	deletedIps, err := ipgroups.DeleteIpFromList(client,
		ipGroup.ID,
		ipgroups.BatchDeleteOpts{IpList: []ipgroups.IpList{
			{
				Ip: "192.168.10.12",
			},
			{
				Ip: "192.168.10.13",
			},
		}})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 1, len(deletedIps.IpList))
}
