package v3

import (
	"os"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/openstack"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/pointerto"
	v3 "github.com/opentelekomcloud/gophertelekomcloud/openstack/gaussdb/v3"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/gaussdb/v3/backup"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/gaussdb/v3/instance"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestGaussDBLifecycle(t *testing.T) {
	if os.Getenv("RUN_GAUSS") == "" {
		t.Skip("long test")
	}
	client, err := clients.NewGaussDBClient()
	th.AssertNoErr(t, err)
	createOpts := openstack.GetCloudServerCreateOpts(t)

	// API is not registered in the backend
	// flavors, err := v3.ShowGaussMySqlFlavors(client, v3.ShowGaussMySqlFlavorsOpts{})
	// th.AssertNoErr(t, err)
	// flav := flavors[0]

	name := tools.RandomString("gaussdb-test-", 5)

	ins, err := instance.CreateInstance(client, instance.CreateInstanceOpts{
		Region:                 "eu-de",
		Name:                   name,
		AvailabilityZoneMode:   "multi",
		MasterAvailabilityZone: "eu-de-01",
		Datastore: instance.Datastore{
			Type:    "gaussdb-mysql",
			Version: "8.0",
		},
		Mode:            "Cluster",
		FlavorRef:       "gaussdb.mysql.xlarge.x86.8",
		VpcId:           createOpts.VpcId,
		SubnetId:        createOpts.Nics[0].SubnetId,
		Password:        "gaussdb1!-test",
		SlaveCount:      pointerto.Int(1),
		SecurityGroupId: openstack.DefaultSecurityGroup(t),
	})
	th.AssertNoErr(t, err)
	id := ins.Instance.Id

	t.Cleanup(func() {
		_, err = instance.DeleteInstance(client, ins.Instance.Id)
		th.AssertNoErr(t, err)
	})
	th.AssertNoErr(t, err)
	_, err = v3.WaitForGaussJob(client, ins.JobId, 600)
	th.AssertNoErr(t, err)

	getInstance, err := instance.GetInstance(client, ins.Instance.Id)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, getInstance.Name, name)

	list, err := instance.ListInstances(client, instance.ListInstancesOpts{
		Id: ins.Instance.Id,
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, list.Instances[0].Id, ins.Instance.Id)
	tools.PrintResource(t, list)

	nameOpts := instance.UpdateNameOpts{
		InstanceId: id,
		Name:       name + "_updated",
	}

	jobId, err := instance.UpdateName(client, nameOpts)
	th.AssertNoErr(t, err)

	_, err = v3.WaitForGaussJob(client, *jobId, 600)
	th.AssertNoErr(t, err)

	passwdOpts := instance.ResetPwdOpts{
		InstanceId: id,
		Password:   "gaussdb1!-test-2",
	}
	err = instance.ResetPassword(client, passwdOpts)

	updateOpts := instance.UpdateSpecOpts{
		InstanceId: id,
		ResizeFlavor: instance.ResizeFlavor{
			SpecCode: "gaussdb.mysql.2xlarge.x86.8",
		},
	}

	updateResp, err := instance.UpdateInstance(client, updateOpts)
	th.AssertNoErr(t, err)
	tools.PrintResource(t, updateResp)

	_, err = v3.WaitForGaussJob(client, updateResp.JobId, 1200)
	th.AssertNoErr(t, err)

	getInstance, err = instance.GetInstance(client, ins.Instance.Id)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, getInstance.Name, nameOpts.Name)

}

func TestGaussDBReplicationLifecycle(t *testing.T) {
	if os.Getenv("RUN_GAUSS") == "" {
		t.Skip("long test")
	}
	client, err := clients.NewGaussDBClient()
	th.AssertNoErr(t, err)
	createOpts := openstack.GetCloudServerCreateOpts(t)

	name := tools.RandomString("gaussdb-test-", 5)

	ins, err := instance.CreateInstance(client, instance.CreateInstanceOpts{
		Region:                 "eu-de",
		Name:                   name,
		AvailabilityZoneMode:   "multi",
		MasterAvailabilityZone: "eu-de-01",
		Datastore: instance.Datastore{
			Type:    "gaussdb-mysql",
			Version: "8.0",
		},
		Mode:            "Cluster",
		FlavorRef:       "gaussdb.mysql.xlarge.x86.8",
		VpcId:           createOpts.VpcId,
		SubnetId:        createOpts.Nics[0].SubnetId,
		Password:        "gaussdb1!-test",
		SlaveCount:      pointerto.Int(1),
		SecurityGroupId: openstack.DefaultSecurityGroup(t),
	})
	id := ins.Instance.Id

	t.Cleanup(func() {
		_, err = instance.DeleteInstance(client, ins.Instance.Id)
		th.AssertNoErr(t, err)
	})
	th.AssertNoErr(t, err)
	_, err = v3.WaitForGaussJob(client, ins.JobId, 600)
	th.AssertNoErr(t, err)

	replicaOpts := instance.CreateNodeOpts{
		InstanceId: id,
		Priorities: []int{1},
	}
	node, err := instance.CreateReplica(client, replicaOpts)
	th.AssertNoErr(t, err)

	_, err = v3.WaitForGaussJob(client, node.JobId, 600)
	th.AssertNoErr(t, err)

	getInstance, err := instance.GetInstance(client, ins.Instance.Id)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, getInstance.Name, name)

	nodeElem := *getInstance.Nodes

	job, err := instance.DeleteReplica(client, id, nodeElem[0].Id)
	th.AssertNoErr(t, err)

	_, err = v3.WaitForGaussJob(client, *job, 600)
	th.AssertNoErr(t, err)
}

func TestGaussDBBackupLifecycle(t *testing.T) {
	if os.Getenv("RUN_GAUSS") == "" {
		t.Skip("long test")
	}
	client, err := clients.NewGaussDBClient()
	th.AssertNoErr(t, err)
	createOpts := openstack.GetCloudServerCreateOpts(t)

	ins, err := instance.CreateInstance(client, instance.CreateInstanceOpts{
		Region:                 "eu-de",
		Name:                   "gaussdb-test",
		AvailabilityZoneMode:   "multi",
		MasterAvailabilityZone: "eu-de-01",
		Datastore: instance.Datastore{
			Type:    "gaussdb-mysql",
			Version: "8.0",
		},
		Mode:            "Cluster",
		FlavorRef:       "gaussdb.mysql.xlarge.x86.8",
		VpcId:           createOpts.VpcId,
		SubnetId:        createOpts.Nics[0].SubnetId,
		Password:        "gaussdb1!-test",
		SlaveCount:      pointerto.Int(1),
		SecurityGroupId: openstack.DefaultSecurityGroup(t),
	})
	t.Cleanup(func() {
		jobId, err := instance.DeleteInstance(client, ins.Instance.Id)
		th.AssertNoErr(t, err)
		_, err = v3.WaitForGaussJob(client, *jobId, 600)
		th.AssertNoErr(t, err)
	})
	th.AssertNoErr(t, err)
	_, err = v3.WaitForGaussJob(client, ins.JobId, 600)
	th.AssertNoErr(t, err)

	dbBackup, err := backup.CreateBackup(client, backup.CreateBackupOpts{
		InstanceId: ins.Instance.Id,
		Name:       "gaussdb-test-backup",
	})
	th.AssertNoErr(t, err)
	_, err = v3.WaitForGaussJob(client, dbBackup.JobId, 600)
	th.AssertNoErr(t, err)

	backupList, err := backup.ListBackups(client, backup.BackupListOpts{
		InstanceId: ins.Instance.Id,
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, backupList.Backups[0].Id, dbBackup.Backup.Id)
}
