package v1

import (
	"testing"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/openstack"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/sfs_turbo/v1/shares"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func createShare(t *testing.T, client *golangsdk.ServiceClient) *shares.Turbo {
	t.Logf("Attempting to create SFSTurboV1")

	vpcID := clients.EnvOS.GetEnv("VPC_ID")
	subnetID := clients.EnvOS.GetEnv("NETWORK_ID")
	az := clients.EnvOS.GetEnv("AVAILABILITY_ZONE")
	if vpcID == "" || subnetID == "" || az == "" {
		t.Skip(`One of OS_VPC_ID or OS_NETWORK_ID of OS_AVAILABILITY_ZONE
env vars is missing but SFS Turbo test requires`)
	}

	createOpts := shares.CreateOpts{
		Name:             tools.RandomString("acc-share-", 3),
		ShareType:        "STANDARD",
		Size:             500,
		AvailabilityZone: az,
		VpcID:            vpcID,
		SubnetID:         subnetID,
		SecurityGroupID:  openstack.DefaultSecurityGroup(t),
		Description:      "some interesting description",
	}

	share, err := shares.Create(client, createOpts).Extract()
	th.AssertNoErr(t, err)

	t.Logf("Waiting for SFS Turbo %s to be active", share.ID)
	err = waitForShareToActive(client, share.ID, 600)
	th.AssertNoErr(t, err)

	newShare, err := shares.Get(client, share.ID).Extract()
	th.AssertNoErr(t, err)

	t.Logf("Created SFS turbo: %s", newShare.ID)

	return newShare
}

func deleteShare(t *testing.T, client *golangsdk.ServiceClient, shareID string) {
	t.Logf("Attempting to delete SFS Turbo: %s", shareID)

	err := shares.Delete(client, shareID).Err
	th.AssertNoErr(t, err)

	err = waitForShareToDelete(client, shareID, 600)
	th.AssertNoErr(t, err)

	t.Logf("Deleted SFS Turbo: %s", shareID)
}

func expandShare(t *testing.T, client *golangsdk.ServiceClient, shareID string) *shares.Turbo {
	t.Logf("Attempting to expand SFS Turbo: %s", shareID)

	expandOpts := shares.ExpandOpts{
		Extend: shares.ExtendOpts{
			NewSize: 1000,
		},
	}

	err := shares.Expand(client, shareID, expandOpts).Err
	th.AssertNoErr(t, err)

	err = waitForShareSubStatusSuccess(client, shareID, 600)
	th.AssertNoErr(t, err)

	newShare, err := shares.Get(client, shareID).Extract()
	th.AssertNoErr(t, err)

	t.Logf("Expanded SFS turbo: %s", shareID)

	return newShare
}

func changeShareSG(t *testing.T, client *golangsdk.ServiceClient, shareID string, secGroupID string) *shares.Turbo {
	t.Logf("Attempting to change SG SFS Turbo: %s", shareID)

	changeSGOpts := shares.ChangeSGOpts{
		ChangeSecurityGroup: shares.SecurityGroupOpts{
			SecurityGroupID: secGroupID,
		},
	}

	err := shares.ChangeSG(client, shareID, changeSGOpts).Err
	th.AssertNoErr(t, err)

	err = waitForShareSubStatusSuccess(client, shareID, 600)
	th.AssertNoErr(t, err)

	newShare, err := shares.Get(client, shareID).Extract()
	th.AssertNoErr(t, err)

	t.Logf("Change SG SFS turbo: %s", shareID)

	return newShare
}

func waitForShareToActive(client *golangsdk.ServiceClient, shareID string, secs int) error {
	return golangsdk.WaitFor(secs, func() (bool, error) {
		share, err := shares.Get(client, shareID).Extract()
		if err != nil {
			return false, err
		}
		if share.Status == "200" {
			return true, nil
		}

		return false, nil
	})
}

func waitForShareToDelete(client *golangsdk.ServiceClient, shareID string, secs int) error {
	return golangsdk.WaitFor(secs, func() (bool, error) {
		_, err := shares.Get(client, shareID).Extract()
		if err != nil {
			if _, ok := err.(golangsdk.ErrDefault404); ok {
				return true, nil
			}
			return false, err
		}
		return false, nil
	})
}

func waitForShareSubStatusSuccess(client *golangsdk.ServiceClient, shareID string, secs int) error {
	return golangsdk.WaitFor(secs, func() (bool, error) {
		share, err := shares.Get(client, shareID).Extract()
		if err != nil {
			return false, err
		}

		if share.SubStatus == "221" || share.SubStatus == "232" {
			return true, nil
		}

		return false, nil
	})
}
