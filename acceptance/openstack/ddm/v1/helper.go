package v1

import (
	"testing"
	"time"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/openstack"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"

	// common "github.com/opentelekomcloud/gophertelekomcloud/openstack/ddm/v1"
	ddminstances "github.com/opentelekomcloud/gophertelekomcloud/openstack/ddm/v1/instances"
	rdsinstances "github.com/opentelekomcloud/gophertelekomcloud/openstack/rds/v3/instances"

	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func CreateRDS(t *testing.T, client *golangsdk.ServiceClient, region string) *rdsinstances.Instance {
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

	createRdsOpts := rdsinstances.CreateRdsOpts{
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

		Volume: &rdsinstances.Volume{
			Type: "COMMON",
			Size: 100,
		},
		Datastore: &rdsinstances.Datastore{
			Type:    "PostgreSQL",
			Version: "11",
		},
		UnchangeableParam: &rdsinstances.Param{
			LowerCaseTableNames: "0",
		},
	}

	rds, err := rdsinstances.Create(client, createRdsOpts)
	th.AssertNoErr(t, err)
	err = rdsinstances.WaitForJobCompleted(client, 1200, rds.JobId)
	th.AssertNoErr(t, err)
	t.Logf("Created RDSv3: %s", rds.Instance.Id)

	return &rds.Instance
}

func DeleteRDS(t *testing.T, client *golangsdk.ServiceClient, rdsID string) {
	t.Logf("Attempting to delete RDSv3: %s", rdsID)

	err := golangsdk.WaitFor(1000, func() (bool, error) {
		_, err := rdsinstances.Delete(client, rdsID)
		if err != nil {
			return false, nil
		}
		return true, nil
	})
	th.AssertNoErr(t, err)

	t.Logf("RDSv3 instance deleted: %s", rdsID)
}

func CreateDDMInstance(t *testing.T, client *golangsdk.ServiceClient) *ddminstances.Instance {
	vpcID := clients.EnvOS.GetEnv("VPC_ID")
	subnetID := clients.EnvOS.GetEnv("NETWORK_ID")
	secGroupId := clients.EnvOS.GetEnv("SECURITY_GROUP")
	if subnetID == "" || vpcID == "" || secGroupId == "" {
		t.Skip("OS_NETWORK_ID or OS_VPC_ID or OS_SECURITY_GROUP env vars are missing but are required for DDM instances")
	}

	instanceName := tools.RandomString("ddm-instance-", 3)
	instanceDetails := ddminstances.CreateInstanceDetail{
		Name:            instanceName,
		FlavorId:        "941b5a6d-3485-329e-902c-ffd49d352f16",
		NodeNum:         2,
		EngineId:        "367b68a3-b48b-3d8a-b3a1-4c463a75a4b4",
		AvailableZones:  []string{"eu-de-01", "eu-de-02", "eu-de-03"},
		VpcId:           vpcID,
		SubnetId:        subnetID,
		SecurityGroupId: secGroupId,
	}
	createOpts := ddminstances.CreateOpts{
		Instance: instanceDetails,
	}

	t.Logf("Creating DDM Instance: %s", instanceName)
	ddmInstance, err := ddminstances.Create(client, createOpts)
	th.AssertNoErr(t, err)
	golangsdk.WaitFor(1200, func() (bool, error) {
		instanceDetails, err := ddminstances.QueryInstanceDetails(client, ddmInstance.Id)
		if err != nil {
			return false, err
		}
		time.Sleep(5 * time.Second)
		if instanceDetails.Status == "RUNNING" {
			th.AssertEquals(t, instanceDetails.Name, instanceName)
			return true, nil
		}
		return false, nil
	})
	t.Logf("Created DDM Instance: %s\nDDM instance ID: %s", instanceName, ddmInstance.Id)
	return ddmInstance
}

func DeleteDDMInstance(t *testing.T, client *golangsdk.ServiceClient, instanceId string) {
	t.Logf("Attempting to delete DDM Instance with ID: %s", instanceId)

	err := golangsdk.WaitFor(1000, func() (bool, error) {
		_, err := ddminstances.Delete(client, instanceId, true)
		if err != nil {
			return false, nil
		}
		return true, nil
	})
	th.AssertNoErr(t, err)
	t.Logf("Deleted DDM Instance with ID: %s", instanceId)
}
