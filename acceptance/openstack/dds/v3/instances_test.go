package v3

import (
	"testing"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dds/v3/instances"
)

func TestDdsList(t *testing.T) {
	client, err := clients.NewDdsV3Client()
	if err != nil {
		t.Fatalf("Unable to create a DDSv3 client: %s", err)
	}

	listOpts := instances.ListInstanceOpts{}
	ddsAllPages, err := instances.List(client, listOpts).AllPages()
	if err != nil {
		t.Fatalf("Unable to get DDSv3 instance pages: %s", err)
	}
	ddsInstances, err := instances.ExtractInstances(ddsAllPages)
	if err != nil {
		t.Fatalf("Unable to extract DDSv3 instances: %s", err)
	}
	for _, val := range ddsInstances.Instances {
		tools.PrintResource(t, val)
	}
}

func TestDdsLifeCycle(t *testing.T) {
	client, err := clients.NewDdsV3Client()
	if err != nil {
		t.Fatalf("Unable to create a DDSv3 client: %s", err)
	}
	_, err = createDdsInstance(t, client)

	return
}

func createDdsInstance(t *testing.T, client *golangsdk.ServiceClient) (*instances.Instance, error) {
	ddsName := tools.RandomString("test-acc-", 8)

	// This value got from tenant default security group
	defaultSgId := "88a47a36-1b69-41b5-bef8-f74a2a85933f"

	createOpts := instances.CreateOpts{
		Name: ddsName,
		DataStore: instances.DataStore{
			Type:          "DDS-Community",
			Version:       "3.4",
			StorageEngine: "wiredTiger",
		},
		Region:           clients.OS_REGION_NAME,
		AvailabilityZone: clients.OS_AVAILABILITY_ZONE,
		VpcId:            clients.OS_VPC_ID,
		SubnetId:         clients.OS_NETWORK_ID,
		SecurityGroupId:  defaultSgId,
		Password:         "5ecurePa55w0rd@",
		Mode:             "ReplicaSet",
		Flavor: []instances.Flavor{
			{
				Type:     "replica",
				Num:      1,
				Storage:  "ULTRAHIGH",
				Size:     20,
				SpecCode: "dds.mongodb.s2.medium.4.repset",
			},
		},
		BackupStrategy: instances.BackupStrategy{
			StartTime: "08:15-09:15",
		},
	}
	createResult := instances.Create(client, createOpts)
	ddsInstance, err := createResult.Extract()
	if err != nil {
		return nil, err
	}
	jobResponse, err := createResult.ExtractJobResponse()
	if err != nil {
		return nil, err
	}

	if err = instances.WaitForJobCompleted(client, 600, jobResponse.JobID); err != nil {
		return nil, err
	}

	return ddsInstance, nil
}
