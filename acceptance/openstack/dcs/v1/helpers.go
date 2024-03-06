package v1

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dcs/v1/others"

	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/openstack"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dcs/v1/lifecycle"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func createDCSInstance(t *testing.T, client *golangsdk.ServiceClient) *lifecycle.Instance {
	t.Logf("Attempting to create DCSv1 instance")

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

	productList, err := others.GetProducts(client)
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

	plan := lifecycle.InstanceBackupPolicy{
		SaveDays: 1,
		PeriodicalBackupPlan: lifecycle.PeriodicalBackupPlan{
			BeginAt:    "00:00-01:00",
			PeriodType: "weekly",
			BackupAt:   []int{1, 2, 3, 4, 5, 6, 7},
		},
	}
	instanceID, err := lifecycle.Create(client, lifecycle.CreateOps{
		Name:                 tools.RandomString("dcs-instance-", 3),
		Description:          "some test DCSv1 instance",
		Engine:               "Redis",
		EngineVersion:        "5.0",
		Capacity:             0.125,
		Password:             "Qwerty123!",
		VPCId:                vpcID,
		SubnetID:             networkID,
		AvailableZones:       []string{az},
		SpecCode:             specCode,
		SecurityGroupID:      openstack.DefaultSecurityGroup(t),
		InstanceBackupPolicy: &plan,
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
	t.Cleanup(func() {
		deleteDCSInstance(t, client, instanceID)
	})

	err = waitForInstanceAvailable(client, 100, instanceID)
	th.AssertNoErr(t, err)

	t.Logf("DCSv1 instance successfully created: %s", instanceID)

	ins, err := lifecycle.Get(client, instanceID)
	th.AssertNoErr(t, err)

	th.AssertEquals(t, plan.SaveDays, ins.InstanceBackupPolicy.Policy.SaveDays)
	th.AssertEquals(t, plan.PeriodicalBackupPlan.BeginAt, ins.InstanceBackupPolicy.Policy.PeriodicalBackupPlan.BeginAt)
	th.AssertEquals(t, plan.PeriodicalBackupPlan.PeriodType, ins.InstanceBackupPolicy.Policy.PeriodicalBackupPlan.PeriodType)

	return ins
}

func deleteDCSInstance(t *testing.T, client *golangsdk.ServiceClient, instanceID string) {
	t.Logf("Attempting to delete DCSv1 instance: %s", instanceID)

	err := lifecycle.Delete(client, instanceID)
	th.AssertNoErr(t, err)

	err = waitForInstanceDeleted(client, 600, instanceID)
	th.AssertNoErr(t, err)

	th.AssertNoErr(t, err)
	t.Logf("Deleted DCSv1 instance: %s", instanceID)
}

func waitForInstanceAvailable(client *golangsdk.ServiceClient, secs int, instanceID string) error {
	return golangsdk.WaitFor(secs, func() (bool, error) {
		dcsInstances, err := lifecycle.Get(client, instanceID)
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
		_, err := lifecycle.Get(client, instanceID)
		if err != nil {
			if _, ok := err.(golangsdk.ErrDefault404); ok {
				return true, nil
			}
			return false, err
		}

		return false, nil
	})
}
