package v3

import (
	"os"
	"testing"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/openstack"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/rds/v3/instances"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/rds/v3/upgrade"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func CreatePostgreSQLRDS(t *testing.T, client *golangsdk.ServiceClient, region string) *instances.Instance {
	t.Logf("Attempting to create PostgreSQL RDS instance")

	rdsName := tools.RandomString("rds-pg-test-", 8)

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
		Port:             "5432",
		Password:         "Postgres!120521",
		FlavorRef:        "rds.pg.n1.large.2",
		Region:           region,
		AvailabilityZone: az,
		VpcId:            vpcID,
		SubnetId:         subnetID,
		SecurityGroupId:  openstack.DefaultSecurityGroup(t),
		DiskEncryptionId: kmsID,

		Volume: &instances.Volume{
			Type: "CLOUDSSD",
			Size: 100,
		},
		Datastore: &instances.Datastore{
			Type:    "PostgreSQL",
			Version: "13",
		},
	}

	rds, err := instances.Create(client, createRdsOpts)
	th.AssertNoErr(t, err)
	err = instances.WaitForJobCompleted(client, 1200, rds.JobId)
	th.AssertNoErr(t, err)
	t.Logf("Created PostgreSQL RDS: %s", rds.Instance.Id)

	return &rds.Instance
}

func TestGetAvailableVersion(t *testing.T) {
	if os.Getenv("RUN_RDS_UPGRADE") == "" {
		t.Skip("RUN_RDS_UPGRADE not set, skipping RDS upgrade tests")
	}

	client, err := clients.NewRdsV3()
	th.AssertNoErr(t, err)

	cc, err := clients.CloudAndClient()
	th.AssertNoErr(t, err)

	rds := CreatePostgreSQLRDS(t, client, cc.RegionName)
	t.Cleanup(func() { DeleteRDS(t, client, rds.Id) })

	th.AssertNoErr(t, instances.WaitForStateAvailable(client, 600, rds.Id))

	t.Log("Attempting to query available upgrade versions")

	versions, err := upgrade.GetAvailableVersion(client, upgrade.GetAvailableVersionOpts{
		InstanceId: rds.Id,
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, len(versions.AvailableVersions) > 0, true)
}

func TestMajorVersionUpgradeWorkflow(t *testing.T) {
	if os.Getenv("RUN_RDS_UPGRADE") == "" {
		t.Skip("RUN_RDS_UPGRADE not set, skipping RDS upgrade tests")
	}

	client, err := clients.NewRdsV3()
	th.AssertNoErr(t, err)

	cc, err := clients.CloudAndClient()
	th.AssertNoErr(t, err)

	rds := CreatePostgreSQLRDS(t, client, cc.RegionName)
	t.Cleanup(func() { DeleteRDS(t, client, rds.Id) })

	th.AssertNoErr(t, instances.WaitForStateAvailable(client, 600, rds.Id))

	t.Log("Attempting to get available versions")
	versions, err := upgrade.GetAvailableVersion(client, upgrade.GetAvailableVersionOpts{
		InstanceId: rds.Id,
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, len(versions.AvailableVersions) > 0, true)

	if len(versions.AvailableVersions) == 0 {
		t.Skip("No available versions for upgrade")
	}

	targetVersion := versions.AvailableVersions[0]

	t.Log("Attempting to perform pre-check")
	reportId, err := upgrade.PreCheck(client, upgrade.PreCheckOpts{
		InstanceId:    rds.Id,
		TargetVersion: targetVersion,
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, reportId != "", true)

	t.Log("Attempting to check pre-check status")
	status, err := upgrade.GetStatus(client, upgrade.GetStatusOpts{
		InstanceId: rds.Id,
		Action:     "check",
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, status.TargetVersion, targetVersion)

	t.Log("Attempting to query inspection histories")
	limit := 10
	histories, err := upgrade.GetInspectionHistories(client, upgrade.GetInspectionHistoriesOpts{
		InstanceId: rds.Id,
		Limit:      &limit,
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, histories.TotalCount > 0, true)

	t.Log("Attempting to perform major version upgrade")
	th.AssertNoErr(t, instances.WaitForStateAvailable(client, 600, rds.Id))

	jobResp, err := upgrade.UpgradeMajorVersion(client, upgrade.UpgradeMajorVersionOpts{
		InstanceId:               rds.Id,
		TargetVersion:            targetVersion,
		IsChangePrivateIp:        false,
		StatisticsCollectionMode: ""})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, jobResp.JobId != "", true)

	t.Log("Attempting to monitor upgrade status")
	upgradeStatus, err := upgrade.GetStatus(client, upgrade.GetStatusOpts{
		InstanceId: rds.Id,
		Action:     "upgrade"})
	th.AssertNoErr(t, err)
	t.Logf("Upgrade status: %s", upgradeStatus.Status)
}

func TestGetUpgradeHistories(t *testing.T) {
	if os.Getenv("RUN_RDS_UPGRADE") == "" {
		t.Skip("RUN_RDS_UPGRADE not set, skipping RDS upgrade tests")
	}

	client, err := clients.NewRdsV3()
	th.AssertNoErr(t, err)

	cc, err := clients.CloudAndClient()
	th.AssertNoErr(t, err)

	rds := CreatePostgreSQLRDS(t, client, cc.RegionName)
	t.Cleanup(func() { DeleteRDS(t, client, rds.Id) })

	th.AssertNoErr(t, instances.WaitForStateAvailable(client, 600, rds.Id))

	t.Log("Attempting to query upgrade histories")

	limit := 10
	histories, err := upgrade.GetUpgradeHistories(client, upgrade.GetUpgradeHistoriesOpts{
		InstanceId: rds.Id,
		Limit:      &limit,
		Order:      "DESC",
		SortField:  "start_time",
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, histories.TotalCount >= 0, true)
}
