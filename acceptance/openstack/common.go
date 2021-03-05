// Package openstack contains common functions that can be used
// across all OpenStack components for acceptance testing.
package openstack

import (
	"fmt"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/blockstorage/v2/volumes"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/extensions"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/compute/v2/extensions/secgroups"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

// PrintExtension prints an extension and all of its attributes.
func PrintExtension(t *testing.T, extension *extensions.Extension) {
	t.Logf("Name: %s", extension.Name)
	t.Logf("Namespace: %s", extension.Namespace)
	t.Logf("Alias: %s", extension.Alias)
	t.Logf("Description: %s", extension.Description)
	t.Logf("Updated: %s", extension.Updated)
	t.Logf("Links: %v", extension.Links)
}

func DefaultSecurityGroup(t *testing.T) string {
	client, err := clients.NewComputeV2Client()
	th.AssertNoErr(t, err)

	securityGroupPages, err := secgroups.List(client).AllPages()
	th.AssertNoErr(t, err)
	securityGroups, err := secgroups.ExtractSecurityGroups(securityGroupPages)
	th.AssertNoErr(t, err)
	var sgId string
	for _, val := range securityGroups {
		if val.Name == "default" {
			sgId = val.ID
			break
		}
	}
	if sgId == "" {
		t.Fatalf("Unable to find default secgroup")
	}
	return sgId
}

func CreateSecurityGroup(t *testing.T) string {
	client, err := clients.NewComputeV2Client()
	th.AssertNoErr(t, err)

	createSGOpts := secgroups.CreateOpts{
		Name:        tools.RandomString("acc-sg-", 3),
		Description: "security group for acceptance testing",
	}
	secGroup, err := secgroups.Create(client, createSGOpts).Extract()
	th.AssertNoErr(t, err)

	t.Logf("Security group %s was created", secGroup.ID)
	return secGroup.ID
}

func DeleteSecurityGroup(t *testing.T, secGroupID string) {
	client, err := clients.NewComputeV2Client()
	th.AssertNoErr(t, err)

	err = secgroups.DeleteWithRetry(client, secGroupID, 600)
	th.AssertNoErr(t, err)

	t.Logf("Security group %s was deleted", secGroupID)
}

func CreateVolume(t *testing.T) *volumes.Volume {
	client, err := clients.NewBlockStorageV3Client()
	th.AssertNoErr(t, err)
	vol, err := volumes.Create(client, volumes.CreateOpts{
		Name:       tools.RandomString("test-vol-", 6),
		Size:       10,
		VolumeType: "SSD",
	}).Extract()
	th.AssertNoErr(t, err)

	err = golangsdk.WaitFor(300, func() (bool, error) {
		volume, err := volumes.Get(client, vol.ID).Extract()
		if err != nil {
			return false, err
		}
		if volume.Status == "available" {
			return true, nil
		}
		if volume.Status == "error" {
			return false, fmt.Errorf("error creating a volume")
		}
		return false, nil
	})
	th.AssertNoErr(t, err)

	return vol
}

func DeleteVolume(t *testing.T, id string) {
	client, err := clients.NewBlockStorageV3Client()
	th.AssertNoErr(t, err)
	th.AssertNoErr(t, volumes.Delete(client, id, volumes.DeleteOpts{Cascade: true}).ExtractErr())
}
