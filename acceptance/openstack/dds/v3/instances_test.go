package v3

import (
	"fmt"
	"testing"
	"time"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/openstack"
	networking "github.com/opentelekomcloud/gophertelekomcloud/acceptance/openstack/networking/v1"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dds/v3/instances"
	ddsjob "github.com/opentelekomcloud/gophertelekomcloud/openstack/dds/v3/job"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestDdsList(t *testing.T) {
	client, err := clients.NewDdsV3Client()
	th.AssertNoErr(t, err)

	listOpts := instances.ListInstanceOpts{}
	_, err = instances.List(client, listOpts)
	th.AssertNoErr(t, err)
}

func TestDdsSingleLifeCycle(t *testing.T) {
	client, err := clients.NewDdsV3Client()
	th.AssertNoErr(t, err)

	t.Logf("Attempting to create DDSv3 single instance")
	ddsInstance := createDdsSingleInstance(t, client)
	defer deleteDdsInstance(t, client, ddsInstance.Id)

	listOpts := instances.ListInstanceOpts{Id: ddsInstance.Id}
	newDdsInstance, err := instances.List(client, listOpts)
	th.AssertNoErr(t, err)
	if newDdsInstance.TotalCount == 0 {
		t.Fatalf("No DDSv3 instance was found: %s", err)
	}

	updateDdsInstance(t, client, newDdsInstance.Instances[0])
}

func TestDdsClusterLifeCycle(t *testing.T) {
	client, err := clients.NewDdsV3Client()
	th.AssertNoErr(t, err)

	t.Logf("Attempting to create DDSv3 cluster instance")
	ddsInstance := createDdsClusterInstance(t, client)
	defer deleteDdsInstance(t, client, ddsInstance.Id)

	listOpts := instances.ListInstanceOpts{Id: ddsInstance.Id}
	newDdsInstance, err := instances.List(client, listOpts)
	th.AssertNoErr(t, err)
	if newDdsInstance.TotalCount == 0 {
		t.Fatalf("No DDSv3 instance was found: %s", err)
	}

	t.Logf("Attempting to add 2 mongo nodes to cluster")
	_, err = instances.AddNode(client, instances.AddNodeOpts{
		InstanceId: newDdsInstance.Instances[0].Id,
		Type:       "mongos",
		SpecCode:   "dds.mongodb.s2.large.4.mongos",
		Num:        2,
	})
	th.AssertNoErr(t, err)
	err = waitForInstanceAvailable(client, 600, ddsInstance.Id)
	th.AssertNoErr(t, err)

	t.Logf("Attempting to add 2 shard nodes to cluster")
	_, err = instances.AddNode(client, instances.AddNodeOpts{
		InstanceId: newDdsInstance.Instances[0].Id,
		Type:       "shard",
		SpecCode:   "dds.mongodb.s2.large.4.shard",
		Num:        2,
		Volume: &instances.VolumeNode{
			Size: 70,
		},
	})
	th.AssertNoErr(t, err)
	err = waitForInstanceAvailable(client, 600, ddsInstance.Id)
	th.AssertNoErr(t, err)

	t.Log("Enable config IP")
	err = instances.EnableConfigIp(client, instances.EnableConfigIpOpts{
		InstanceId: ddsInstance.Id,
		Type:       "config",
		Password:   "5ecurePa55w0rd@@",
	})
	th.AssertNoErr(t, err)
	err = waitForInstanceAvailable(client, 600, ddsInstance.Id)
	th.AssertNoErr(t, err)
}

func TestDdsReplicaLifeCycle(t *testing.T) {
	client, err := clients.NewDdsV3Client()
	th.AssertNoErr(t, err)

	ddsInstance := createDdsReplicaInstance(t, client)
	defer deleteDdsInstance(t, client, ddsInstance.Id)

	listOpts := instances.ListInstanceOpts{Id: ddsInstance.Id}
	newDdsInstance, err := instances.List(client, listOpts)
	th.AssertNoErr(t, err)
	if newDdsInstance.TotalCount == 0 {
		t.Fatalf("No DDSv3 instance was found: %s", err)
	}

	err = instances.ChangePassword(client, instances.ChangePasswordOpt{
		InstanceId: ddsInstance.Id,
		UserPwd:    "5ecurePa55w0rd@@",
	})
	th.AssertNoErr(t, err)
	err = waitForInstanceAvailable(client, 600, ddsInstance.Id)
	th.AssertNoErr(t, err)
}

func updateDdsInstance(t *testing.T, client *golangsdk.ServiceClient, instance instances.InstanceResponse) {
	t.Log("Update name")

	err := instances.UpdateName(client,
		instances.UpdateNameOpt{
			InstanceId:      instance.Id,
			NewInstanceName: tools.RandomString("dds-acc-2-", 8),
		})
	th.AssertNoErr(t, err)

	netClient, err := clients.NewNetworkV1Client()
	th.AssertNoErr(t, err)

	elasticIP := networking.CreateEip(t, netClient, 100)
	t.Cleanup(func() {
		networking.DeleteEip(t, netClient, elasticIP.ID)
	})

	t.Log("AttachEip")

	job, err := instances.BindEIP(client, instances.BindEIPOpts{
		PublicIpId: elasticIP.ID,
		PublicIp:   elasticIP.PublicAddress,
		NodeId:     instance.Groups[0].Nodes[0].Id,
	})
	th.AssertNoErr(t, err)

	err = waitForJobCompleted(client, 600, *job)
	th.AssertNoErr(t, err)

	t.Log("UnbindEip")

	job, err = instances.UnBindEIP(client, instance.Groups[0].Nodes[0].Id)
	th.AssertNoErr(t, err)

	err = waitForJobCompleted(client, 600, *job)
	th.AssertNoErr(t, err)

	err = waitForInstanceAvailable(client, 600, instance.Id)
	th.AssertNoErr(t, err)
	t.Log("Enable the SSL")
	job, err = instances.SwitchSSL(client, instances.SSLOpt{InstanceId: instance.Id,
		SSL: "1"})
	th.AssertNoErr(t, err)
	err = waitForJobCompleted(client, 600, *job)
	th.AssertNoErr(t, err)

	t.Log("Modify instance internal IP")
	job, err = instances.ModifyInternalIp(client, instances.ModifyInternalIpOpts{
		InstanceId: instance.Id,
		NewIp:      "192.168.1.42",
		NodeId:     instance.Groups[0].Nodes[0].Id,
	})
	th.AssertNoErr(t, err)
	err = waitForJobCompleted(client, 600, *job)
	th.AssertNoErr(t, err)

	t.Log("Modify instance port")
	job, err = instances.ModifyPort(client, instances.ModifyPortOpt{
		InstanceId: instance.Id,
		Port:       8636,
	})
	th.AssertNoErr(t, err)
	err = waitForJobCompleted(client, 600, *job)
	th.AssertNoErr(t, err)

	t.Log("Modify instance SG")
	_, err = instances.ModifySG(client, instances.ModifySGOpt{
		InstanceId:      instance.Id,
		SecurityGroupId: openstack.DefaultSecurityGroup(t),
	})
	th.AssertNoErr(t, err)
	err = waitForJobCompleted(client, 600, *job)
	th.AssertNoErr(t, err)

	t.Log("Modify instance specs")
	job, err = instances.ModifySpec(client, instances.ModifySpecOpt{
		InstanceId:     instance.Id,
		TargetId:       instance.Id,
		TargetSpecCode: "dds.mongodb.s2.large.4.single",
	})
	th.AssertNoErr(t, err)
	err = waitForJobCompleted(client, 600, *job)
	th.AssertNoErr(t, err)

	_, err = instances.Restart(client, instances.RestartOpts{
		InstanceId: instance.Id,
		TargetId:   instance.Id,
	})
	th.AssertNoErr(t, err)
	err = waitForInstanceAvailable(client, 600, instance.Id)
	th.AssertNoErr(t, err)

	t.Log("Modify instance volume size")
	job, err = instances.ScaleStorage(client, instances.ScaleStorageOpt{
		InstanceId: instance.Id,
		Size:       "60",
		GroupId:    instance.Groups[0].Id,
	})
	th.AssertNoErr(t, err)
	err = waitForJobCompleted(client, 600, *job)
	th.AssertNoErr(t, err)
}

func createDdsSingleInstance(t *testing.T, client *golangsdk.ServiceClient) *instances.Instance {
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
			Version:       "4.0",
			StorageEngine: "wiredTiger",
		},
		Region:           cloud.RegionName,
		AvailabilityZone: az,
		VpcId:            vpcID,
		SubnetId:         subnetID,
		SecurityGroupId:  openstack.DefaultSecurityGroup(t),
		Password:         "5ecurePa55w0rd@",
		Mode:             "Single",
		Flavor: []instances.Flavor{
			{
				Type:     "single",
				Num:      1,
				Storage:  "ULTRAHIGH",
				Size:     20,
				SpecCode: "dds.mongodb.s2.medium.4.single",
			},
		},
		BackupStrategy: instances.BackupStrategy{
			StartTime: "08:15-09:15",
		},
	}
	ddsInstance, err := instances.Create(client, createOpts)
	th.AssertNoErr(t, err)
	err = waitForInstanceAvailable(client, 600, ddsInstance.Id)
	th.AssertNoErr(t, err)
	t.Logf("DDSv3 replica set instance successfully created")
	return ddsInstance
}

func createDdsReplicaInstance(t *testing.T, client *golangsdk.ServiceClient) *instances.Instance {
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
	ddsInstance, err := instances.Create(client, createOpts)
	th.AssertNoErr(t, err)
	err = waitForInstanceAvailable(client, 600, ddsInstance.Id)
	th.AssertNoErr(t, err)
	t.Logf("DDSv3 replica set instance successfully created")
	return ddsInstance
}

func createDdsClusterInstance(t *testing.T, client *golangsdk.ServiceClient) *instances.Instance {
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
		Mode:             "Sharding",
		Flavor: []instances.Flavor{
			{
				Type:     "mongos",
				Num:      2,
				Storage:  "ULTRAHIGH",
				SpecCode: "dds.mongodb.s2.medium.4.mongos",
			},
			{
				Type:     "shard",
				Num:      2,
				Storage:  "ULTRAHIGH",
				Size:     20,
				SpecCode: "dds.mongodb.s2.medium.4.shard",
			},
			{
				Type:     "config",
				Num:      1,
				Storage:  "ULTRAHIGH",
				Size:     20,
				SpecCode: "dds.mongodb.s2.large.2.config",
			},
		},
		BackupStrategy: instances.BackupStrategy{
			StartTime: "08:15-09:15",
		},
	}
	ddsInstance, err := instances.Create(client, createOpts)
	th.AssertNoErr(t, err)
	err = waitForInstanceAvailable(client, 600, ddsInstance.Id)
	th.AssertNoErr(t, err)
	t.Logf("DDSv3 replica set instance successfully created: %s", ddsInstance.Id)
	return ddsInstance
}

func deleteDdsInstance(t *testing.T, client *golangsdk.ServiceClient, instanceId string) {
	t.Logf("Attempting to delete DDSv3 instance: %s", instanceId)

	_, err := instances.Delete(client, instanceId)
	if err != nil {
		t.Fatalf("Unable to delete DDSv3 instance: %s", err)
	}
	err = waitForInstanceDelete(client, 600, instanceId)
	if err != nil {
		t.Fatalf("Error waiting delete DDSv3 instance: %s", err)
	}
	t.Logf("DDSv3 instance deleted successfully: %s", instanceId)
}

func waitForJobCompleted(client *golangsdk.ServiceClient, secs int, jobID string) error {
	jobClient := *client
	jobClient.ResourceBase = jobClient.Endpoint

	return golangsdk.WaitFor(secs, func() (bool, error) {
		job, err := ddsjob.Get(client, jobID)
		if err != nil {
			return false, err
		}

		if job.Status == "Completed" {
			return true, nil
		}
		if job.Status == "Failed" {
			err = fmt.Errorf("Job failed %s.\n", job.Status)
			return false, err
		}

		time.Sleep(5 * time.Second)
		return false, nil
	})
}

func waitForInstanceAvailable(client *golangsdk.ServiceClient, secs int, instanceId string) error {
	return golangsdk.WaitFor(secs, func() (bool, error) {
		listOpts := instances.ListInstanceOpts{
			Id: instanceId,
		}
		ddsInstances, err := instances.List(client, listOpts)
		if err != nil {
			return false, err
		}
		if ddsInstances.TotalCount == 1 {
			dds := ddsInstances.Instances
			if len(dds) == 1 && len(dds[0].Actions) == 0 && dds[0].Status == "normal" {
				return true, nil
			}
			return false, nil
		}
		return false, nil
	})
}

func waitForInstanceDelete(client *golangsdk.ServiceClient, secs int, instanceId string) error {
	return golangsdk.WaitFor(secs, func() (bool, error) {
		listOpts := instances.ListInstanceOpts{
			Id: instanceId,
		}
		ddsInstances, err := instances.List(client, listOpts)
		if err != nil {
			return false, err
		}
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
