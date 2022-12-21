package v3

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	networking "github.com/opentelekomcloud/gophertelekomcloud/acceptance/openstack/networking/v1"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/rds/v3/instances"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestRdsList(t *testing.T) {
	client, err := clients.NewRdsV3()
	th.AssertNoErr(t, err)

	rdsInstances, err := instances.List(client, instances.ListOpts{})
	th.AssertNoErr(t, err)

	for _, rds := range rdsInstances.Instances {
		tools.PrintResource(t, rds)
	}
}

func TestRdsLifecycle(t *testing.T) {
	client, err := clients.NewRdsV3()
	th.AssertNoErr(t, err)

	cc, err := clients.CloudAndClient()
	th.AssertNoErr(t, err)

	// Create RDSv3 instance
	rds := createRDS(t, client, cc.RegionName)
	t.Cleanup(func() { deleteRDS(t, client, rds.Id) })
	th.AssertEquals(t, rds.Volume.Size, 100)

	tagList := []tags.ResourceTag{
		{
			Key:   "muh",
			Value: "value-create",
		},
		{
			Key:   "kuh",
			Value: "value-create",
		},
	}
	err = tags.Create(client, "instances", rds.Id, tagList).ExtractErr()
	th.AssertNoErr(t, err)

	err = updateRDS(t, client, rds.Id)
	th.AssertNoErr(t, err)
	tools.PrintResource(t, rds)

	newRds, err := instances.List(client, instances.ListOpts{
		Id: rds.Id,
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, len(newRds.Instances), 1)
	th.AssertEquals(t, newRds.Instances[0].Volume.Size, 200)
	th.AssertEquals(t, len(newRds.Instances[0].Tags), 2)

	if err := instances.WaitForStateAvailable(client, 600, rds.Id); err != nil {
		t.Fatalf("Status available wasn't present")
	}
	restart, err := instances.Restart(client, instances.RestartOpts{InstanceId: rds.Id, Restart: struct{}{}})
	th.AssertNoErr(t, err)
	err = instances.WaitForJobCompleted(client, 1200, *restart)
	th.AssertNoErr(t, err)

	t.Logf("Attempting to create RDSv3 Read Replica")

	prefix := "rds-rr-"
	rdsReplicaName := tools.RandomString(prefix, 8)
	kmsID := clients.EnvOS.GetEnv("KMS_ID")
	az := clients.EnvOS.GetEnv("AVAILABILITY_ZONE")
	if az == "" {
		az = "eu-de-01"
	}

	createOpts := instances.CreateReplicaOpts{
		Name:             rdsReplicaName,
		ReplicaOfId:      rds.Id,
		DiskEncryptionId: kmsID,
		FlavorRef:        "rds.pg.c2.medium.rr",
		Volume: &instances.Volume{
			Type: "COMMON",
			Size: 100,
		},
		AvailabilityZone: az,
	}

	rdsReadReplica, err := instances.CreateReplica(client, createOpts)
	th.AssertNoErr(t, err)
	err = instances.WaitForJobCompleted(client, 1200, rdsReadReplica.JobId)
	th.AssertNoErr(t, err)
	t.Logf("Created RDSv3 Read Replica: %s", rdsReadReplica.Instance.Id)

	t.Cleanup(func() {
		t.Logf("Attempting to delete RDSv3 Read Replica: %s", rdsReadReplica.Instance.Id)
		_, err := instances.Delete(client, rdsReadReplica.Instance.Id)
		th.AssertNoErr(t, err)
		t.Logf("RDSv3 Read Replica instance deleted: %s", rdsReadReplica.Instance.Id)
	})

	elasticIP := networking.CreateEip(t, client, 100)
	t.Cleanup(func() {
		networking.DeleteEip(t, client, elasticIP.ID)
	})

	err = instances.AttachEip(client, instances.AttachEipOpts{
		InstanceId: rds.Id,
		PublicIp:   elasticIP.PublicAddress,
		PublicIpId: elasticIP.ID,
		IsBind:     true,
	})
	th.AssertNoErr(t, err)
}
