package v2

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/ces/v2/resourcegroups"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestResourceGroupsCRUD(t *testing.T) {
	client, err := clients.NewCesV2Client()
	th.AssertNoErr(t, err)

	t.Log("Attempting to create resource group")
	createOpts := resourcegroups.CreateOpts{
		GroupName: "test-rg-acc",
		Type:      "Manual",
	}

	groupId, err := resourcegroups.Create(client, createOpts)
	th.AssertNoErr(t, err)
	t.Logf("Created resource group: %s", groupId)

	t.Cleanup(func() {
		t.Log("Attempting to delete resource group")
		_, err := resourcegroups.Delete(client, resourcegroups.DeleteOpts{
			GroupIds: []string{groupId},
		})
		th.AssertNoErr(t, err)
	})

	t.Log("Attempting to get resource group details")
	group, err := resourcegroups.Get(client, groupId)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, group.GroupName, "test-rg-acc")
	th.AssertEquals(t, group.Type, "Manual")

	t.Log("Attempting to list resource groups")
	listResp, err := resourcegroups.List(client, resourcegroups.ListOpts{
		GroupName: "test-rg-acc",
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, listResp.Count >= 1, true)

	t.Log("Attempting to update resource group")
	err = resourcegroups.Update(client, groupId, resourcegroups.UpdateOpts{
		GroupName: "test-rg-acc-updated",
	})
	th.AssertNoErr(t, err)

	t.Log("Attempting to verify resource group was updated")
	group, err = resourcegroups.Get(client, groupId)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, group.GroupName, "test-rg-acc-updated")
}

func TestResourceGroupsWithTags(t *testing.T) {
	client, err := clients.NewCesV2Client()
	th.AssertNoErr(t, err)

	t.Log("Attempting to create resource group with tags")
	createOpts := resourcegroups.CreateOpts{
		GroupName: "test-rg-tag-acc",
		Type:      "TAG",
		Tags: []resourcegroups.ResourceGroupTag{
			{
				Key:   "Environment",
				Value: "Test",
			},
			{
				Key:   "Project",
				Value: "Acceptance",
			},
		},
	}

	groupId, err := resourcegroups.Create(client, createOpts)
	th.AssertNoErr(t, err)
	t.Logf("Created resource group with tags: %s", groupId)

	t.Cleanup(func() {
		t.Log("Attempting to delete resource group")
		_, err := resourcegroups.Delete(client, resourcegroups.DeleteOpts{
			GroupIds: []string{groupId},
		})
		th.AssertNoErr(t, err)
	})

	t.Log("Attempting to get resource group details")
	group, err := resourcegroups.Get(client, groupId)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, group.GroupName, "test-rg-tag-acc")
	th.AssertEquals(t, group.Type, "TAG")
	th.AssertEquals(t, len(group.Tags), 2)

	t.Log("Attempting to update resource group tags")
	err = resourcegroups.Update(client, groupId, resourcegroups.UpdateOpts{
		GroupName: "test-rg-tag-acc-updated",
		Tags: []resourcegroups.ResourceGroupTag{
			{
				Key:   "Environment",
				Value: "Production",
			},
		},
	})
	th.AssertNoErr(t, err)

	t.Log("Attempting to verify resource group was updated")
	group, err = resourcegroups.Get(client, groupId)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, group.GroupName, "test-rg-tag-acc-updated")
}

func TestResourceGroupsList(t *testing.T) {
	client, err := clients.NewCesV2Client()
	th.AssertNoErr(t, err)

	t.Log("Attempting to list all resource groups")
	listResp, err := resourcegroups.List(client, resourcegroups.ListOpts{
		Limit: 10,
	})
	th.AssertNoErr(t, err)

	t.Logf("Found %d resource groups", listResp.Count)
	for _, rg := range listResp.ResourceGroups {
		t.Logf("Resource Group ID: %s, Name: %s, Type: %s",
			rg.GroupId, rg.GroupName, rg.Type)
	}
}
