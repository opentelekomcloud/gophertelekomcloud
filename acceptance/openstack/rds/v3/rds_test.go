package v3

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/networking/v2/extensions/security/groups"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/rds/v3/instances"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
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

	var updateOpts instances.EnlargeVolumeRdsOpts
	updateOpts.EnlargeVolume.Size = 200
	_, err = instances.EnlargeVolume(client, updateOpts, rds.Id).Extract()
	if err != nil {
		t.Fatalf("Unable to resize volume: %s", err)
	}

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
		Region:           clients.OS_REGION_NAME,
		AvailabilityZone: clients.OS_AVAILABILITY_ZONE,
		VpcId:            clients.OS_VPC_ID,
		SubnetId:         clients.OS_SUBNET_ID,
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
	if err = instances.WaitForJobCompleted(client, int(600), jobResponse.JobID); err != nil {
		return nil, err
	}
	t.Logf("Created RDSv3: %s", rdsName)

	return &rds.Instance, nil
}

func deleteRDS(t *testing.T, client *golangsdk.ServiceClient, rdsId string) {
	t.Logf("Attempting to delete RDSv3: %s", rdsId)

	jobResponse, err := instances.Delete(client, rdsId).ExtractJobResponse()
	if err != nil {
		t.Fatalf("Unable to extract RDS delete response: %s", err)
	}
	err = instances.WaitForJobCompleted(client, 600, jobResponse.JobID)
	if err != nil {
		t.Fatalf("Error deleting RDSv3: %v", err)
	}

	t.Logf("Deleted RDSv3: %s", rdsId)
}

func createSecGroup(sgName string) (*groups.SecGroup, error) {
	nwClient, err := clients.NewNetworkV2Client()
	if err != nil {
		return nil, err
	}
	sgOpts := groups.CreateOpts{
		Name:     sgName,
		TenantID: clients.OS_TENANT_ID,
	}
	sg, err := groups.Create(nwClient, sgOpts).Extract()
	if err != nil {
		return nil, err
	}
	return sg, nil
}
