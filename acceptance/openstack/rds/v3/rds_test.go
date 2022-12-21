package v3

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/openstack"
	networking "github.com/opentelekomcloud/gophertelekomcloud/acceptance/openstack/networking/v1"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/rds/v3/backups"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/rds/v3/instances"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/rds/v3/logs"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/rds/v3/security"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestRdsList(t *testing.T) {
	client, err := clients.NewRdsV3()
	th.AssertNoErr(t, err)

	rdsInstances, err := instances.List(client, instances.ListOpts{})
	th.AssertNoErr(t, err)
	tools.PrintResource(t, rdsInstances)

	collations, err := instances.ListCollations(client)
	th.AssertNoErr(t, err)
	tools.PrintResource(t, collations)
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
	// rds := struct{ Id string }{Id: "f4d563c17b9f4d70815b93a88f13600ain03"}

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

	rdsName := tools.RandomString("rds-test-", 8)
	err = instances.UpdateInstanceName(client, instances.UpdateInstanceNameOpts{
		InstanceId: rds.Id,
		Name:       rdsName,
	})
	th.AssertNoErr(t, err)

	_, err = security.SetSecurityGroup(client, security.SetSecurityGroupOpts{
		InstanceId:      rds.Id,
		SecurityGroupId: openstack.DefaultSecurityGroup(t),
	})
	th.AssertNoErr(t, err)

	err = security.SwitchSsl(client, security.SwitchSslOpts{
		InstanceId: rds.Id,
		SslOption:  true,
	})
	th.AssertEquals(t, true, err != nil)

	port, err := security.UpdatePort(client, security.UpdatePortOpts{
		InstanceId: rds.Id,
		Port:       3306,
	})
	th.AssertNoErr(t, err)
	err = instances.WaitForJobCompleted(client, 600, *port)
	th.AssertNoErr(t, err)

	restart, err := instances.Restart(client, instances.RestartOpts{InstanceId: rds.Id, Restart: struct{}{}})
	th.AssertNoErr(t, err)
	err = instances.WaitForJobCompleted(client, 600, *restart)
	th.AssertNoErr(t, err)

	t.Logf("Attempting to create RDSv3 Read Replica")

	rdsReplicaName := tools.RandomString("rds-rr-", 8)
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

	replica, err := instances.CreateReplica(client, createOpts)
	th.AssertNoErr(t, err)
	err = instances.WaitForJobCompleted(client, 1200, replica.JobId)
	th.AssertNoErr(t, err)

	t.Logf("Created RDSv3 Read Replica: %s", replica.Instance.Id)

	t.Cleanup(func() {
		t.Logf("Attempting to delete RDSv3 Read Replica: %s", replica.Instance.Id)
		_, err := instances.Delete(client, replica.Instance.Id)
		th.AssertNoErr(t, err)
		t.Logf("RDSv3 Read Replica instance deleted: %s", replica.Instance.Id)
	})

	ip, err := security.UpdateDataIp(client, security.UpdateDataIpOpts{
		InstanceId: rds.Id,
		NewIp:      "192.168.30.254",
	})
	th.AssertNoErr(t, err)
	err = instances.WaitForJobCompleted(client, 600, *ip)
	th.AssertNoErr(t, err)

	netClient, err := clients.NewNetworkV1Client()
	th.AssertNoErr(t, err)

	elasticIP := networking.CreateEip(t, netClient, 100)
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
	t.Cleanup(func() {
		err = instances.AttachEip(client, instances.AttachEipOpts{
			InstanceId: rds.Id,
			IsBind:     false,
		})
		th.AssertNoErr(t, err)
	})

	stop, err := instances.StopInstance(client, rds.Id)
	th.AssertNoErr(t, err)
	err = instances.WaitForJobCompleted(client, 600, *stop)
	th.AssertNoErr(t, err)

	start, err := instances.StartupInstance(client, rds.Id)
	th.AssertNoErr(t, err)
	err = instances.WaitForJobCompleted(client, 600, *start)
	th.AssertNoErr(t, err)

	err = instances.ChangeOpsWindow(client, instances.ChangeOpsWindowOpts{
		InstanceId: rds.Id,
		StartTime:  "22:00",
		EndTime:    "02:00",
	})
	th.AssertNoErr(t, err)

	follower, err := instances.MigrateFollower(client, instances.MigrateFollowerOpts{
		InstanceId: rds.Id,
		NodeId:     replica.Instance.Id,
		AzCode:     az,
	})
	th.AssertNoErr(t, err)
	err = instances.WaitForJobCompleted(client, 600, *follower)
	th.AssertNoErr(t, err)

	resize, err := instances.Resize(client, instances.ResizeOpts{
		InstanceId: rds.Id,
		SpecCode:   "rds.pg.c2.large",
	})
	th.AssertNoErr(t, err)
	err = instances.WaitForJobCompleted(client, 600, *resize)
	th.AssertNoErr(t, err)

	ha, err := instances.SingleToHa(client, instances.SingleToHaOpts{
		InstanceId:    rds.Id,
		AzCodeNewNode: az,
	})
	th.AssertNoErr(t, err)
	err = instances.WaitForJobCompleted(client, 600, *ha)
	th.AssertNoErr(t, err)

	mode, err := instances.ChangeFailoverMode(client, instances.ChangeFailoverModeOpts{
		InstanceId: rds.Id,
		Mode:       "sync",
	})
	th.AssertNoErr(t, err)
	err = instances.WaitForJobCompleted(client, 600, mode.WorkflowId)
	th.AssertNoErr(t, err)

	err = instances.ChangeFailoverStrategy(client, instances.ChangeFailoverStrategyOpts{
		InstanceId:     rds.Id,
		RepairStrategy: "reliability",
	})
	th.AssertNoErr(t, err)

	failover, err := instances.StartFailover(client, rds.Id)
	th.AssertNoErr(t, err)
	err = instances.WaitForJobCompleted(client, 600, failover.WorkflowId)
	th.AssertNoErr(t, err)

	backup, err := backups.Create(client, backups.CreateOpts{
		InstanceID: rds.Id,
		Name:       tools.RandomString("rds-backup-test-", 5),
	})
	th.AssertNoErr(t, err)
	t.Log("Backup creation started")
	// backup := struct{ ID string }{ID: "6d85318246d644af95c0666e2a6ae1debr03"}

	t.Cleanup(func() {
		th.AssertNoErr(t, backups.Delete(client, backup.ID))
		t.Log("Backup deleted")
	})

	err = backups.WaitForBackup(client, rds.Id, backup.ID, backups.StatusCompleted)
	th.AssertNoErr(t, err)
	t.Log("Backup creation complete")

	backupList, err := backups.List(client, backups.ListOpts{InstanceID: rds.Id, BackupID: backup.ID})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 1, len(backupList))
	tools.PrintResource(t, backupList[0])

	times, err := backups.ListRestoreTimes(client, backups.ListRestoreTimesOpts{
		InstanceId: rds.Id,
	})
	th.AssertNoErr(t, err)
	tools.PrintResource(t, times)

	if err := instances.WaitForStateAvailable(client, 600, rds.Id); err != nil {
		t.Fatalf("Status available wasn't present")
	}

	pitr, err := backups.RestorePITR(client, backups.RestorePITROpts{
		Source: backups.Source{
			BackupID:   backupList[0].ID,
			InstanceID: backupList[0].InstanceID,
			Type:       "backup",
		},
		Target: backups.Target{
			InstanceID: rds.Id,
		},
	})
	th.AssertNoErr(t, err)
	err = instances.WaitForJobCompleted(client, 600, pitr)
	th.AssertNoErr(t, err)

	toNew, err := backups.RestoreToNew(client, backups.RestoreToNewOpts{
		Name:      rdsName,
		Password:  "acc-test-password1!",
		FlavorRef: "rds.pg.c2.large",
		Volume: &instances.Volume{
			Type: "COMMON",
			Size: 200,
		},
		AvailabilityZone: az,
		VpcId:            clients.EnvOS.GetEnv("VPC_ID"),
		SubnetId:         clients.EnvOS.GetEnv("NETWORK_ID"),
		SecurityGroupId:  openstack.DefaultSecurityGroup(t),
		RestorePoint: backups.RestorePoint{
			InstanceID: rds.Id,
			Type:       "backup",
			BackupID:   backupList[0].ID,
		},
	})
	th.AssertNoErr(t, err)
	err = instances.WaitForJobCompleted(client, 600, toNew.JobId)
	th.AssertNoErr(t, err)
	t.Cleanup(func() {
		deleteRDS(t, client, toNew.Instance.Id)
	})

	policy, err := backups.ShowBackupPolicy(client, rds.Id)
	th.AssertNoErr(t, err)
	tools.PrintResource(t, policy)

	err = backups.Update(client, backups.UpdateOpts{
		InstanceId: rds.Id,
		KeepDays:   7,
	})
	th.AssertNoErr(t, err)

	log, err := logs.ListErrorLog(client, logs.DbErrorlogOpts{
		InstanceId: rds.Id,
		Limit:      "1",
	})
	th.AssertNoErr(t, err)
	tools.PrintResource(t, log)

	slowLog, err := logs.ListSlowLog(client, logs.DbSlowLogOpts{
		InstanceId: rds.Id,
		Limit:      "1",
	})
	th.AssertNoErr(t, err)
	tools.PrintResource(t, slowLog)
}
