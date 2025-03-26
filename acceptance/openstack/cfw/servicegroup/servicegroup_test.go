package servicegroup

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	common "github.com/opentelekomcloud/gophertelekomcloud/acceptance/openstack/cfw"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	managementv1 "github.com/opentelekomcloud/gophertelekomcloud/openstack/cfw/v1/management"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/cfw/v1/servicegroup"
	managementv2 "github.com/opentelekomcloud/gophertelekomcloud/openstack/cfw/v2/management"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestServiceGroupLifecycle(t *testing.T) {
	t.Skip("Too long. Non reproducible in CI")
	clientv1, err := clients.NewCFWV1Client()
	th.AssertNoErr(t, err)
	clientv2, err := clients.NewCFWV2Client()
	th.AssertNoErr(t, err)
	clientv3, err := clients.NewCFWV3Client()
	th.AssertNoErr(t, err)

	instanceName := tools.RandomString("test-acc-firewall-", 3)
	createOpts := managementv2.CreateOpts{
		Name: instanceName,
		Flavor: managementv2.CreateFlavor{
			Version: "standard",
		},
		ChargeInfo: managementv2.ChargeInfo{
			ChargeMode: "postPaid",
		},
	}
	createResp, err := managementv2.Create(clientv2, createOpts)
	th.AssertNoErr(t, err)
	th.AssertNoErr(t, common.WaitForJobCompleted(clientv3, 600, 5, createResp.JobID))
	instanceId := createResp.JobID
	t.Cleanup(func() {
		_, err = managementv2.Delete(clientv2, instanceId)
		th.AssertNoErr(t, err)
	})

	firewall, err := managementv1.Get(clientv1, instanceId, 0)
	th.AssertNoErr(t, err)

	groupName := tools.RandomString("test-acc-group-", 3)
	group, err := servicegroup.CreateServiceGroup(clientv1, servicegroup.CreateOpts{
		ObjectID: firewall.ProtectObjects[0].ObjectID,
		Name:     groupName,
	})
	th.AssertNoErr(t, err)

	_, err = servicegroup.AddGroupMember(clientv1, servicegroup.AddGroupMemberOpts{
		SetId: group.Id,
		ServiceItems: []servicegroup.ServiceItem{
			{
				Protocol:    6,
				SourcePort:  "1",
				DestPort:    "1",
				Description: "test611",
			},
		},
	})
	th.AssertNoErr(t, err)

	groupMember, err := servicegroup.GetGroupMember(clientv1, group.Id, "1", "1", 6)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "test611", groupMember.Description)

	th.AssertNoErr(t, servicegroup.DeleteGroupMember(clientv1, groupMember.ItemID))

	_, err = servicegroup.ListServiceGroups(clientv1, firewall.ProtectObjects[0].ObjectID)
	th.AssertNoErr(t, err)

	_, err = servicegroup.UpdateServiceGroup(clientv1, group.Id, servicegroup.UpdateOpts{
		Description: "Test Group",
	})
	th.AssertNoErr(t, err)

	groupInfo, err := servicegroup.GetServiceGroup(clientv1, group.Id)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, groupName, groupInfo.Name)
	th.AssertEquals(t, "Test Group", groupInfo.Description)

	th.AssertNoErr(t, servicegroup.DeleteServiceGroup(clientv1, group.Id))
}
