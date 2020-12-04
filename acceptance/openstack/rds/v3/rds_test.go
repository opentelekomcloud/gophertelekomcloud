package v3

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/networking/v2/extensions/security/groups"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/rds/v3/instances"
	"testing"
)

func TestRdsList(t *testing.T) {
	client, err := clients.NewRdsV3()
	if err != nil {
		t.Fatalf("Unable to create a RDSv3 client: %s", err)
	}

	listOpts := instances.ListRdsInstanceOpts{}
	allRdsPages, err := instances.List(client, listOpts).AllPages()
	if err != nil {
		t.Fatalf("Unable to fetch RDSv3 pages: %s", err)
	}
	rdsInstances, err := instances.ExtractRdsInstances(allRdsPages)
	if err != nil {
		t.Fatalf("Unable to extract RDSv3 pages: %s", err)
	}
	for _, rds := range rdsInstances.Instances {
		tools.PrintResource(t, rds)
	}
}

func TestRdsCRUD(t *testing.T) {
	client, err := clients.NewRdsV3()
	if err != nil {
		t.Fatalf("Unable to create a RDSv3 client: %s", err)
	}

	// Create RDSv3 instance
	rds, err := createRDS(t, client)
	if err != nil {
		t.Fatalf("Unable to create create: %s", err)
	}
	defer deleteRDS(t, client, rds.Id)

	tools.PrintResource(t, rds)

	err = updateRDS(t, client, rds.Id)
	if err != nil {
		t.Fatalf("Unable to update RDSv3 instance: %s", err)
	}
	tools.PrintResource(t, rds)

	listOpts := instances.ListRdsInstanceOpts{
		Id: rds.Id,
	}
	allPages, err := instances.List(client, listOpts).AllPages()
	if err != nil {
		t.Fatalf("Unable to get all RDS pages: %s", err)
	}
	newRds, err := instances.ExtractRdsInstances(allPages)
	if err != nil {
		t.Fatalf("Unable to extract RDS pages: %s", err)
	}
	tools.PrintResource(t, newRds)
}

func createRDS(t *testing.T, client *golangsdk.ServiceClient) (*instances.Instance, error) {
	rdsName := tools.RandomString("test-acc-", 8)
	sgName := tools.RandomString("test-acc", 8)
	sg, err := createSecGroup(sgName)
	if err != nil {
		return nil, err
	}

	createRdsOpts := instances.CreateRdsOpts{
		Name:             rdsName,
		Port:             "8635",
		Password:         "acc-test-password1!",
		FlavorRef:        "rds.pg.c2.medium",
		Region:           clients.EnvOS.GetEnv("REGION_NAME"),
		AvailabilityZone: clients.EnvOS.GetEnv("AVAILABILITY_ZONE"),
		VpcId:            clients.EnvOS.GetEnv("VPC_ID"),
		SubnetId:         clients.EnvOS.GetEnv("NETWORK_ID"),
		SecurityGroupId:  sg.ID,

		Volume: &instances.Volume{
			Type: "COMMON",
			Size: 100,
		},
		Datastore: &instances.Datastore{
			Type:    "PostgreSQL",
			Version: "10",
		},
	}
	createResult := instances.Create(client, createRdsOpts)
	rds, err := createResult.Extract()
	if err != nil {
		return nil, err
	}
	jobResponse, err := createResult.ExtractJobResponse()
	if err != nil {
		return nil, err
	}
	if err = instances.WaitForJobCompleted(client, 1200, jobResponse.JobID); err != nil {
		return nil, err
	}
	t.Logf("Created RDSv3: %s", rds.Instance.Id)

	return &rds.Instance, nil
}

func deleteRDS(t *testing.T, client *golangsdk.ServiceClient, rdsId string) {
	t.Logf("Attempting to delete RDSv3: %s", rdsId)
	listOpts := instances.ListRdsInstanceOpts{
		Id: rdsId,
	}
	allPages, err := instances.List(client, listOpts).AllPages()
	if err != nil {
		t.Fatalf("Unable to get all RDS pages: %s", err)
	}
	rds, err := instances.ExtractRdsInstances(allPages)
	if err != nil {
		t.Fatalf("Unable to extract RDS pages: %s", err)
	}

	_, err = instances.Delete(client, rdsId).ExtractJobResponse()
	if err != nil {
		t.Fatalf("Unable to delete RDSv3: %s", err)
	}
	t.Logf("RDSv3 instance deleted: %s", rdsId)

	deleteSecGroup(t, rds.Instances[0].SecurityGroupId)
}

func updateRDS(t *testing.T, client *golangsdk.ServiceClient, rdsId string) error {
	t.Logf("Attempting to increase volume size")

	t.Logf("Update volume size could be done only in status `available`")
	if err := instances.WaitForStateAvailable(client, 600, rdsId); err != nil {
		t.Fatalf("Status available wasn't present")
	}

	updateOpts := instances.EnlargeVolumeRdsOpts{
		EnlargeVolume: &instances.EnlargeVolumeSize{
			Size: 200,
		},
	}

	updateResult, err := instances.EnlargeVolume(client, updateOpts, rdsId).ExtractJobResponse()
	if err != nil {
		return err
	}
	if err := instances.WaitForJobCompleted(client, 1200, updateResult.JobID); err != nil {
		return err
	}
	t.Logf("RDSv3 successfully updated: %s", rdsId)
	return nil
}

func createSecGroup(sgName string) (*groups.SecGroup, error) {
	nwClient, err := clients.NewNetworkV2Client()
	if err != nil {
		return nil, err
	}
	sgOpts := groups.CreateOpts{
		Name:     sgName,
		TenantID: clients.EnvOS.GetEnv("OS_TENANT_ID"),
	}
	sg, err := groups.Create(nwClient, sgOpts).Extract()
	if err != nil {
		return nil, err
	}
	return sg, nil
}

func deleteSecGroup(t *testing.T, sgID string) {
	t.Logf("Attempting to delete networking_secgroup: %s", sgID)
	nwClient, err := clients.NewNetworkV2Client()
	if err != nil {
		t.Fatalf("Unable to create Networkv2 client: %s", err)
	}
	err = groups.DeleteWithRetry(nwClient, sgID, 600)
	if err != nil {
		t.Fatalf("Unable to delete networking_secgroup: %s, err: %s", sgID, err)
	}
	t.Logf("Deleted networking_secgroup: %s", sgID)
}
