package v1

import (
	"testing"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/openstack"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dds/v3/instances"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestDcsConfigLifeCycle(t *testing.T) {
	client, err := clients.NewDcsV1Client()
	th.AssertNoErr(t, err)

	ddsInstance := createDdsInstance(t, client)
	defer deleteDdsInstance(t, client, ddsInstance.Id)

	tools.PrintResource(t, ddsInstance)
	listOpts := instances.ListInstanceOpts{Id: ddsInstance.Id}
	allPages, err := instances.List(client, listOpts).AllPages()
	th.AssertNoErr(t, err)
	newDdsInstance, err := instances.ExtractInstances(allPages)
	th.AssertNoErr(t, err)
	if newDdsInstance.TotalCount == 0 {
		t.Fatalf("No DDSv3 instance was found: %s", err)
	}
	tools.PrintResource(t, newDdsInstance.Instances[0])
}

func createDdsInstance(t *testing.T, client *golangsdk.ServiceClient) *instances.Instance {
	t.Logf("Attempting to create DDSv3 replica set instance")
	prefix := "dds-acc-"
	ddsName := tools.RandomString(prefix, 8)
	az := clients.EnvOS.GetEnv("AVAILABILITY_ZONE")
	if az == "" {
		az = "eu-de-01"
	}
	cloud, err := clients.EnvOS.Cloud()
	th.AssertNoErr(t, err)

	vpcID := clients.EnvOS.GetEnv("VPC_ID")
	subnetID := clients.EnvOS.GetEnv("NETWORK_ID")
	if vpcID == "" || subnetID == "" {
		t.Skip("One of OS_VPC_ID or OS_NETWORK_ID env vars is missing but RDS test requires using existing network")
	}

	createOpts := instances.CreateOpts{
		Name: ddsName,
		DataStore: instances.DataStore{
			Type:          "DDS-Community",
			Version:       "3.4",
			StorageEngine: "wiredTiger",
		},
		Region:           cloud.RegionName,
		AvailabilityZone: az,
		VpcId:            vpcID,
		SubnetId:         subnetID,
		SecurityGroupId:  openstack.DefaultSecurityGroup(t),
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
	th.AssertNoErr(t, err)
	err = waitForInstanceAvailable(client, 600, ddsInstance.Id)
	th.AssertNoErr(t, err)
	t.Logf("DDSv3 replica set instance successfully created: %s", ddsInstance.Id)
	return ddsInstance
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

// func waitForInstanceAvailable(client *golangsdk.ServiceClient, secs int, instanceId string) error {
// 	return golangsdk.WaitFor(secs, func() (bool, error) {
// 		listOpts := instances.ListInstanceOpts{
// 			Id: instanceId,
// 		}
// 		allPages, err := instances.List(client, listOpts).AllPages()
// 		if err != nil {
// 			return false, err
// 		}
// 		ddsInstances, err := instances.ExtractInstances(allPages)
// 		if err != nil {
// 			return false, err
// 		}
// 		if ddsInstances.TotalCount == 1 {
// 			dds := ddsInstances.Instances
// 			if len(dds) == 1 && len(dds[0].Actions) == 0 {
// 				return true, nil
// 			}
// 			return false, nil
// 		}
// 		return false, nil
// 	})
// }
//
// func waitForInstanceDelete(client *golangsdk.ServiceClient, secs int, instanceId string) error {
// 	return golangsdk.WaitFor(secs, func() (bool, error) {
// 		listOpts := instances.ListInstanceOpts{
// 			Id: instanceId,
// 		}
// 		allPages, err := instances.List(client, listOpts).AllPages()
// 		if err != nil {
// 			return false, err
// 		}
// 		ddsInstances, err := instances.ExtractInstances(allPages)
// 		if err != nil {
// 			if _, ok := err.(golangsdk.ErrDefault404); ok {
// 				return true, nil
// 			}
// 			return false, err
// 		}
// 		if ddsInstances.TotalCount == 0 {
// 			return true, nil
// 		}
// 		return false, nil
// 	})
// }
