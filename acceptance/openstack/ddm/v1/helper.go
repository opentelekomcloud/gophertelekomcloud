package v1

import (
	"testing"
	"time"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/openstack"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"

	ddminstancesv1 "github.com/opentelekomcloud/gophertelekomcloud/openstack/ddm/v1/instances"
	ddminstancesv2 "github.com/opentelekomcloud/gophertelekomcloud/openstack/ddm/v2/instances"
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
	secGroupId := openstack.DefaultSecurityGroup(t)
	if vpcID == "" || subnetID == "" {
		t.Skip("One of OS_VPC_ID or OS_NETWORK_ID env vars is missing but RDS test requires using existing network")
	}

	createRdsOpts := rdsinstances.CreateRdsOpts{
		Name:             rdsName,
		Port:             "3306",
		Password:         "acc-test-password1!",
		FlavorRef:        "rds.mysql.c2.medium",
		Region:           region,
		AvailabilityZone: az,
		VpcId:            vpcID,
		SubnetId:         subnetID,
		SecurityGroupId:  secGroupId,
		DiskEncryptionId: kmsID,

		Volume: &rdsinstances.Volume{
			Type: "COMMON",
			Size: 100,
		},
		Datastore: &rdsinstances.Datastore{
			Type:    "MySQL",
			Version: "8.0",
		},
		UnchangeableParam: &rdsinstances.Param{
			LowerCaseTableNames: "1",
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

func CreateDDMInstance(t *testing.T, client *golangsdk.ServiceClient) *ddminstancesv1.Instance {
	vpcID := clients.EnvOS.GetEnv("VPC_ID")
	subnetID := clients.EnvOS.GetEnv("NETWORK_ID")
	if subnetID == "" || vpcID == "" {
		t.Skip("OS_NETWORK_ID or OS_VPC_ID env vars are missing but are required for DDM instances")
	}
	secGroupId := openstack.DefaultSecurityGroup(t)

	engineId := "367b68a3-b48b-3d8a-b3a1-4c463a75a4b4"
	clientV2, err := clients.NewDDMV2Client()
	th.AssertNoErr(t, err)
	engines, err := ddminstancesv2.QueryEngineInfo(clientV2, ddminstancesv2.QueryEngineOpts{})
	th.AssertNoErr(t, err)
	if len(engines) != 0 {
		engineId = engines[0].ID
	}

	flavorId := "941b5a6d-3485-329e-902c-ffd49d352f16"
	classes, err := ddminstancesv2.QueryNodeClasses(clientV2, ddminstancesv2.QueryNodeClassesOpts{
		EngineId: engineId,
	})
	th.AssertNoErr(t, err)
	if len(classes.ComputeFlavorGroups) != 0 {
		flavorId = classes.ComputeFlavorGroups[0].ComputeFlavors[0].ID
	}

	instanceName := tools.RandomString("ddm-instance-", 3)
	instanceDetails := ddminstancesv1.CreateInstanceDetail{
		Name:            instanceName,
		FlavorId:        flavorId,
		NodeNum:         2,
		EngineId:        engineId,
		AvailableZones:  []string{"eu-de-01", "eu-de-02", "eu-de-03"},
		VpcId:           vpcID,
		SubnetId:        subnetID,
		SecurityGroupId: secGroupId,
	}
	createOpts := ddminstancesv1.CreateOpts{
		Instance: instanceDetails,
	}

	t.Logf("Creating DDM Instance: %s", instanceName)
	ddmInstance, err := ddminstancesv1.Create(client, createOpts)
	th.AssertNoErr(t, err)
	err = golangsdk.WaitFor(1200, func() (bool, error) {
		instanceDetails, err := ddminstancesv1.QueryInstanceDetails(client, ddmInstance.Id)
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
	th.AssertNoErr(t, err)
	t.Logf("Created DDM Instance: %s\nDDM instance ID: %s", instanceName, ddmInstance.Id)
	return ddmInstance
}

func DeleteDDMInstance(t *testing.T, client *golangsdk.ServiceClient, instanceId string) {
	t.Logf("Attempting to delete DDM Instance with ID: %s", instanceId)

	_, err := ddminstancesv1.Delete(client, instanceId, true)
	th.AssertNoErr(t, err)
	err = waitForInstanceDeleted(client, 600, instanceId)
	th.AssertNoErr(t, err)
	t.Logf("Deleted DDM Instance with ID: %s", instanceId)
}

func WaitForInstanceInRunningState(client *golangsdk.ServiceClient, instanceID string) error {
	return golangsdk.WaitFor(1200, func() (bool, error) {
		instanceDetails, err := ddminstancesv1.QueryInstanceDetails(client, instanceID)
		if err != nil {
			return false, err
		}
		time.Sleep(5 * time.Second)
		if instanceDetails.Status == "RUNNING" {
			return true, nil
		}
		return false, nil
	})
}

func waitForInstanceDeleted(client *golangsdk.ServiceClient, secs int, instanceID string) error {
	return golangsdk.WaitFor(secs, func() (bool, error) {
		_, err := ddminstancesv1.QueryInstanceDetails(client, instanceID)
		if err != nil {
			if _, ok := err.(golangsdk.ErrDefault404); ok {
				return true, nil
			}
			return false, err
		}

		return false, nil
	})
}
