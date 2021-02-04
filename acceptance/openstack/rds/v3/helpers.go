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

func createRDS(t *testing.T, client *golangsdk.ServiceClient, region string) *instances.Instance {
	t.Logf("Attempting to create RDSv3")

	prefix := "rds-"
	rdsName := tools.RandomString(prefix, 8)

	az := clients.EnvOS.GetEnv("AVAILABILITY_ZONE")
	if az == "" {
		az = "eu-de-01"
	}

	vpcID := clients.EnvOS.GetEnv("VPC_ID")
	subnetID := clients.EnvOS.GetEnv("NETWORK_ID")
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

		Volume: &instances.Volume{
			Type: "COMMON",
			Size: 100,
		},
		Datastore: &instances.Datastore{
			Type:    "PostgreSQL",
			Version: "11",
		},
	}

	createResult := instances.Create(client, createRdsOpts)
	rds, err := createResult.Extract()
	th.AssertNoErr(t, err)
	jobResponse, err := createResult.ExtractJobResponse()
	th.AssertNoErr(t, err)
	err = instances.WaitForJobCompleted(client, 1200, jobResponse.JobID)
	th.AssertNoErr(t, err)
	t.Logf("Created RDSv3: %s", rds.Instance.Id)

	return &rds.Instance
}

func deleteRDS(t *testing.T, client *golangsdk.ServiceClient, rdsID string) {
	t.Logf("Attempting to delete RDSv3: %s", rdsID)

	_, err := instances.Delete(client, rdsID).ExtractJobResponse()
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
		EnlargeVolume: &instances.EnlargeVolumeSize{
			Size: 200,
		},
	}

	updateResult, err := instances.EnlargeVolume(client, updateOpts, rdsID).ExtractJobResponse()
	th.AssertNoErr(t, err)

	if err := instances.WaitForJobCompleted(client, 1200, updateResult.JobID); err != nil {
		return err
	}
	t.Logf("RDSv3 successfully updated: %s", rdsID)
	return nil
}

func createRDSConfiguration(t *testing.T, client *golangsdk.ServiceClient) *configurations.ConfigurationCreate {
	t.Logf("Attempting to create RDSv3 template configuration")
	prefix := "rds-config-"
	configName := tools.RandomString(prefix, 3)

	configCreateOpts := configurations.CreateOpts{
		Name:        configName,
		Description: "some config description",
		Values: map[string]string{
			"max_connections": "10",
			"autocommit":      "OFF",
		},
		DataStore: configurations.DataStore{
			Type:    "PostgreSQL",
			Version: "11",
		},
	}

	rdsConfiguration, err := configurations.Create(client, configCreateOpts).Extract()
	th.AssertNoErr(t, err)

	t.Logf("Created RDSv3 configuration: %s", rdsConfiguration.ID)

	return rdsConfiguration
}

func deleteRDSConfiguration(t *testing.T, client *golangsdk.ServiceClient, rdsConfigID string) {
	t.Logf("Attempting to delete RDSv3 configuration: %s", rdsConfigID)

	err := configurations.Delete(client, rdsConfigID).Err
	th.AssertNoErr(t, err)

	t.Logf("RDSv3 configuration deleted: %s", rdsConfigID)
}

func updateRDSConfiguration(t *testing.T, client *golangsdk.ServiceClient, rdsConfigID string) {
	t.Logf("Attempting to update name")

	prefix := "rds-update-config-"
	configName := tools.RandomString(prefix, 3)

	updateOpts := configurations.UpdateOpts{
		Name:        configName,
		Description: "some updated description",
	}

	err := configurations.Update(client, rdsConfigID, updateOpts).Err
	th.AssertNoErr(t, err)

	t.Logf("RDSv3 configuration updated")
}
