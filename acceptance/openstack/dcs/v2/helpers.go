package v2

import (
	"fmt"
	"testing"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/openstack"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dcs/v2/instance"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dcs/v2/others"
	dcsTags "github.com/opentelekomcloud/gophertelekomcloud/openstack/dcs/v2/tags"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dcs/v2/whitelists"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func createDCSInstance(t *testing.T, client *golangsdk.ServiceClient) *instance.DcsInstance {
	t.Logf("Attempting to create DCSv2 instance")

	vpcID := clients.EnvOS.GetEnv("VPC_ID")
	networkID := clients.EnvOS.GetEnv("NETWORK_ID")
	if vpcID == "" || networkID == "" {
		t.Skip("OS_VPC_ID or OS_NETWORK_ID is missing but test requires using existing network")
	}

	availabilityZone, err := others.ListAvailableZones(client)
	th.AssertNoErr(t, err)
	var az string
	for _, v := range availabilityZone.AvailableZones {
		if v.ResourceAvailability != "true" {
			continue
		}
		az = v.ID
	}
	if az == "" {
		t.Skip("Availability Zone ID wasn't found")
	}

	productList, err := others.ListFlavors(client, others.ListFlavorOpts{
		SpecCode: "redis.ha.xu1.tiny.r2.128",
	})
	th.AssertNoErr(t, err)

	var specCode string
	for _, v := range productList {
		if v.SpecCode == "redis.ha.xu1.tiny.r2.128" {
			specCode = v.SpecCode
		}
	}
	if specCode == "" {
		t.Skip("Product ID wasn't found")
	}

	plan := instance.InstanceBackupPolicyOpts{
		SaveDays: 1,
		PeriodicalBackupPlan: &instance.BackupPlan{
			BeginAt:    "00:00-01:00",
			PeriodType: "weekly",
			BackupAt:   []int{1, 2, 3, 4, 5, 6, 7},
		},
	}
	instanceIDs, err := instance.Create(client, instance.CreateOpts{
		Name:            tools.RandomString("dcs-instance-", 3),
		Description:     "some test DCSv2 instance",
		Engine:          "Redis",
		EngineVersion:   "6.0",
		Capacity:        0.125,
		Password:        "Qwerty123!",
		VpcId:           vpcID,
		SubnetId:        networkID,
		AzCodes:         []string{az},
		SpecCode:        specCode,
		SecurityGroupId: openstack.DefaultSecurityGroup(t),
		BackupPolicy:    &plan,
		Tags: []tags.ResourceTag{
			{
				Key:   "muh",
				Value: "kuh",
			},
			{
				Key:   "muh2",
				Value: "kuh2",
			},
		},
	})
	th.AssertNoErr(t, err)
	instanceID := instanceIDs[0].InstanceID
	t.Cleanup(func() {
		deleteDCSInstance(t, client, instanceID)
	})

	err = waitForInstanceAvailable(client, 100, instanceID)
	th.AssertNoErr(t, err)

	t.Logf("DCSv2 instance successfully created: %s", instanceID)

	ins, err := instance.Get(client, instanceID)
	th.AssertNoErr(t, err)

	th.AssertEquals(t, plan.SaveDays, ins.BackupPolicy.Policy.SaveDays)
	th.AssertEquals(t, plan.PeriodicalBackupPlan.BeginAt, ins.BackupPolicy.Policy.PeriodicalBackupPlan.BeginAt)
	th.AssertEquals(t, plan.PeriodicalBackupPlan.PeriodType, ins.BackupPolicy.Policy.PeriodicalBackupPlan.PeriodType)

	return ins
}

func deleteDCSInstance(t *testing.T, client *golangsdk.ServiceClient, instanceID string) {
	t.Logf("Attempting to delete DCSv2 instance: %s", instanceID)

	err := instance.Delete(client, instanceID)
	th.AssertNoErr(t, err)

	err = waitForInstanceDeleted(client, 600, instanceID)
	th.AssertNoErr(t, err)

	th.AssertNoErr(t, err)
	t.Logf("Deleted DCSv2 instance: %s", instanceID)
}

func waitForInstanceAvailable(client *golangsdk.ServiceClient, secs int, instanceID string) error {
	return golangsdk.WaitFor(secs, func() (bool, error) {
		dcsInstances, err := instance.Get(client, instanceID)
		if err != nil {
			return false, err
		}
		if dcsInstances.Status == "RUNNING" {
			return true, nil
		}
		return false, nil
	})
}

func waitForInstanceDeleted(client *golangsdk.ServiceClient, secs int, instanceID string) error {
	return golangsdk.WaitFor(secs, func() (bool, error) {
		_, err := instance.Get(client, instanceID)
		if err != nil {
			if _, ok := err.(golangsdk.ErrDefault404); ok {
				return true, nil
			}
			return false, err
		}

		return false, nil
	})
}

func updateDcsTags(client *golangsdk.ServiceClient, id string, old, new []tags.ResourceTag) error {
	// remove old tags
	if len(old) > 0 {
		err := dcsTags.Delete(client, id, old)
		if err != nil {
			return err
		}
	}
	// add new tags
	if len(new) > 0 {
		err := dcsTags.Create(client, id, new)
		if err != nil {
			return err
		}
	}
	return nil
}

// WaitForAWhitelistToBeRetrieved - wait until whitelist is retrieved
func WaitForAWhitelistToBeRetrieved(client *golangsdk.ServiceClient, id string, timeoutSeconds int) error {
	return golangsdk.WaitFor(timeoutSeconds, func() (bool, error) {
		wl, err := whitelists.Get(client, id)
		if err != nil {
			return false, fmt.Errorf("error retriving whitelist: %w", err)
		}
		if wl.InstanceID != "" {
			return true, nil
		}
		return false, nil
	})
}
