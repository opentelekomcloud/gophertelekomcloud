package v3

import (
	"testing"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/openstack"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/rds/v3/configurations"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/rds/v3/instances"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func CreateRDS(t *testing.T, client *golangsdk.ServiceClient, region string) *instances.Instance {
	t.Logf("Attempting to create RDSv3")

	rdsName := tools.RandomString("rds-test-", 8)

	az := clients.EnvOS.GetEnv("AVAILABILITY_ZONE")
	if az == "" {
		az = "eu-de-01"
	}

	vpcID := clients.EnvOS.GetEnv("VPC_ID")
	subnetID := clients.EnvOS.GetEnv("NETWORK_ID")
	kmsID := clients.EnvOS.GetEnv("KMS_ID")
	if vpcID == "" || subnetID == "" {
		t.Skip("One of OS_VPC_ID or OS_NETWORK_ID env vars is missing but RDS test requires using existing network")
	}

	createRdsOpts := instances.CreateRdsOpts{
		Name:             rdsName,
		Port:             "8635",
		Password:         "acc-test-password1!",
		FlavorRef:        "rds.pg.c2.medium",
		Region:           region,
		AvailabilityZone: az,
		VpcId:            vpcID,
		SubnetId:         subnetID,
		SecurityGroupId:  openstack.DefaultSecurityGroup(t),
		DiskEncryptionId: kmsID,

		Volume: &instances.Volume{
			Type: "COMMON",
			Size: 100,
		},
		Datastore: &instances.Datastore{
			Type:    "PostgreSQL",
			Version: "11",
		},
		UnchangeableParam: &instances.Param{
			LowerCaseTableNames: "0",
		},
	}

	rds, err := instances.Create(client, createRdsOpts)
	th.AssertNoErr(t, err)
	err = instances.WaitForJobCompleted(client, 1200, rds.JobId)
	th.AssertNoErr(t, err)
	t.Logf("Created RDSv3: %s", rds.Instance.Id)

	return &rds.Instance
}

func DeleteRDS(t *testing.T, client *golangsdk.ServiceClient, rdsID string) {
	t.Logf("Attempting to delete RDSv3: %s", rdsID)

	err := golangsdk.WaitFor(1000, func() (bool, error) {
		_, err := instances.Delete(client, rdsID)
		if err != nil {
			return false, nil
		}
		return true, nil
	})
	th.AssertNoErr(t, err)

	t.Logf("RDSv3 instance deleted: %s", rdsID)
}

func updateRDS(t *testing.T, client *golangsdk.ServiceClient, rdsID string) error {
	t.Logf("Attempting to increase volume size")

	t.Logf("Update volume size could be done only in status `available`")
	if err := instances.WaitForStateAvailable(client, 600, rdsID); err != nil {
		t.Fatalf("Status available wasn't present")
	}

	updateOpts := instances.EnlargeVolumeRdsOpts{
		InstanceId: rdsID,
		Size:       200,
	}

	updateResult, err := instances.EnlargeVolume(client, updateOpts)
	th.AssertNoErr(t, err)

	if err := instances.WaitForJobCompleted(client, 1200, *updateResult); err != nil {
		return err
	}
	t.Logf("RDSv3 successfully updated: %s", rdsID)
	return nil
}

func createRDSConfiguration(t *testing.T, client *golangsdk.ServiceClient) *configurations.Configuration {
	t.Logf("Attempting to create RDSv3 template configuration")
	prefix := "rds-config-"
	configName := tools.RandomString(prefix, 3)

	configCreateOpts := configurations.CreateOpts{
		Name:        configName,
		Description: "some config description",
		Values: map[string]string{
			"autocommit": "OFF",
		},
		DataStore: configurations.DataStore{
			Type:    "PostgreSQL",
			Version: "11",
		},
	}

	rdsConfiguration, err := configurations.Create(client, configCreateOpts)
	th.AssertNoErr(t, err)

	t.Logf("Created RDSv3 configuration: %s", rdsConfiguration.ID)

	return rdsConfiguration
}

func deleteRDSConfiguration(t *testing.T, client *golangsdk.ServiceClient, rdsConfigID string) {
	t.Logf("Attempting to delete RDSv3 configuration: %s", rdsConfigID)

	err := configurations.Delete(client, rdsConfigID)
	th.AssertNoErr(t, err)

	t.Logf("RDSv3 configuration deleted: %s", rdsConfigID)
}

func updateRDSConfiguration(t *testing.T, client *golangsdk.ServiceClient, rdsConfigID string) {
	t.Logf("Attempting to update name")

	prefix := "rds-update-config-"
	configName := tools.RandomString(prefix, 3)

	updateOpts := configurations.UpdateOpts{
		ConfigId:    rdsConfigID,
		Name:        configName,
		Description: "some updated description",
		Values: map[string]string{
			"autocommit": "ON",
		},
	}

	err := configurations.Update(client, updateOpts)
	th.AssertNoErr(t, err)

	t.Logf("RDSv3 configuration updated")
}
