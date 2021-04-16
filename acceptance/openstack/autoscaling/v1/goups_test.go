package v1

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/openstack"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/autoscaling/v1/groups"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestGroupsList(t *testing.T) {
	client, err := clients.NewAutoscalingV1Client()
	th.AssertNoErr(t, err)

	listOpts := groups.ListOpts{}

	allPages, err := groups.List(client, listOpts).AllPages()
	th.AssertNoErr(t, err)

	asGroups, err := groups.ExtractGroups(allPages)
	th.AssertNoErr(t, err)

	for _, group := range asGroups {
		tools.PrintResource(t, group)
	}
}

func TestGroupLifecycle(t *testing.T) {
	client, err := clients.NewAutoscalingV1Client()
	th.AssertNoErr(t, err)

	asGroupCreateName := tools.RandomString("as-group-create-", 3)
	networkID := clients.EnvOS.GetEnv("NETWORK_ID")
	vpcID := clients.EnvOS.GetEnv("VPC_ID")
	if networkID == "" {
		t.Skip("OS_NETWORK_ID or OS_VPC_ID env vars are missing but AS Group test requires")
	}

	secGroupID := openstack.CreateSecurityGroup(t)
	defer openstack.DeleteSecurityGroup(t, secGroupID)

	defaultSGID := openstack.DefaultSecurityGroup(t)
	deletePublicIP := true

	createOpts := groups.CreateOpts{
		Name: asGroupCreateName,
		Networks: []groups.NetworkOpts{
			{
				ID: networkID,
			},
		},
		SecurityGroup: []groups.SecurityGroupOpts{
			{
				ID: defaultSGID,
			},
		},
		VpcID:          vpcID,
		DeletePublicIP: &deletePublicIP,
	}
	t.Logf("Attempting to create AutoScaling Group")
	groupID, err := groups.Create(client, createOpts).Extract()
	th.AssertNoErr(t, err)
	t.Logf("Created AutoScaling Group: %s", groupID)
	defer func() {
		t.Logf("Attempting to delete AutoScaling Group")
		err := groups.Delete(client, groupID).ExtractErr()
		th.AssertNoErr(t, err)
		t.Logf("Deleted AutoScaling Group: %s", groupID)
	}()

	group, err := groups.Get(client, groupID).Extract()
	th.AssertNoErr(t, err)
	tools.PrintResource(t, group)
	th.AssertEquals(t, asGroupCreateName, group.Name)
	th.AssertEquals(t, 1, len(group.SecurityGroups))
	th.AssertEquals(t, true, group.DeletePublicIP)

	t.Logf("Attempting to update AutoScaling Group")
	asGroupUpdateName := tools.RandomString("as-group-update-", 3)
	deletePublicIP = false

	updateOpts := groups.UpdateOpts{
		Name: asGroupUpdateName,
		SecurityGroup: []groups.SecurityGroupOpts{
			{
				ID: secGroupID,
			},
		},
		DeletePublicIP: &deletePublicIP,
	}

	groupID, err = groups.Update(client, group.ID, updateOpts).Extract()
	th.AssertNoErr(t, err)
	t.Logf("Updated AutoScaling Group")

	group, err = groups.Get(client, groupID).Extract()
	th.AssertNoErr(t, err)
	tools.PrintResource(t, group)
	th.AssertEquals(t, asGroupUpdateName, group.Name)
	th.AssertEquals(t, secGroupID, group.SecurityGroups[0].ID)
	th.AssertEquals(t, false, group.DeletePublicIP)
}
