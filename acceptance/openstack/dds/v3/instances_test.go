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
	ddsInstance, err := createDdsInstance(t, client)
	if err != nil {
		t.Fatalf("Unable to create DDSv3 instance: %s", err)
	}
	defer deleteDdsInstance(t, client, ddsInstance.Id)

	tools.PrintResource(t, ddsInstance)
	listOpts := instances.ListInstanceOpts{Id: ddsInstance.Id}
	allPages, err := instances.List(client, listOpts).AllPages()
	if err != nil {
		t.Fatalf("Unable to query DDSv3 instance: %s", err)
	}
	newDdsInstance, err := instances.ExtractInstances(allPages)
	if err != nil {
		t.Fatalf("Error extracting DDSv3 instances: %s", err)
	}
	if newDdsInstance.TotalCount == 0 {
		t.Fatalf("DDSv3 instance wasn't found: %s", err)
	}
	tools.PrintResource(t, newDdsInstance.Instances[0])
}

func createDdsInstance(t *testing.T, client *golangsdk.ServiceClient) (*instances.Instance, error) {
	t.Logf("Attempting to create DDSv3 replica set instance")
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
	ddsInstance, err := instances.Create(client, createOpts).Extract()
	if err != nil {
		return nil, err
	}
	err = waitForInstanceAvailable(client, 600, ddsInstance.Id)
	if err != nil {
		return nil, err
	}
	t.Logf("DDSv3 replica set instance successfully created: %s", ddsInstance.Id)

	return ddsInstance, nil
}

func deleteDdsInstance(t *testing.T, client *golangsdk.ServiceClient, instanceId string) {
	t.Logf("Attempting to delete DDSv3 instance: %s", instanceId)

	_, err := instances.Delete(client, instanceId).Extract()
	if err != nil {
		t.Fatalf("Unable to delete DDSv3 instance: %s", err)
	}
	err = waitForInstanceDelete(client, 600, instanceId)
	if err != nil {
		t.Fatalf("Error waiting delete DDSv3 instance: %s", err)
	}
	t.Logf("DDSv3 instance deleted successfully: %s", instanceId)
}

func waitForInstanceAvailable(client *golangsdk.ServiceClient, secs int, instanceId string) error {
	return golangsdk.WaitFor(secs, func() (bool, error) {
		listOpts := instances.ListInstanceOpts{
			Id: instanceId,
		}
		allPages, err := instances.List(client, listOpts).AllPages()
		if err != nil {
			return false, err
		}
		ddsInstances, err := instances.ExtractInstances(allPages)
		if err != nil {
			return false, err
		}
		if ddsInstances.TotalCount == 1 {
			return true, nil
		}
		return false, nil
	})
}

func waitForInstanceDelete(client *golangsdk.ServiceClient, secs int, instanceId string) error {
	return golangsdk.WaitFor(secs, func() (bool, error) {
		listOpts := instances.ListInstanceOpts{
			Id: instanceId,
		}
		allPages, err := instances.List(client, listOpts).AllPages()
		if err != nil {
			return false, err
		}
		ddsInstances, err := instances.ExtractInstances(allPages)
		if err != nil {
			if _, ok := err.(golangsdk.ErrDefault404); ok {
				return true, nil
			}
			return false, err
		}
		if ddsInstances.TotalCount == 0 {
			return true, nil
		}
		return false, nil
	})
}
