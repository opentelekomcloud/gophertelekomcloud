package v3

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/openstack"
	v3 "github.com/opentelekomcloud/gophertelekomcloud/openstack/gaussdb/v3"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/gaussdb/v3/backup"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/gaussdb/v3/instance"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestGaussDBLifecycle(t *testing.T) {
	client, err := clients.NewGaussDBClient()
	th.AssertNoErr(t, err)
	createOpts := openstack.GetCloudServerCreateOpts(t)

	flavors, err := v3.ShowGaussMySqlFlavors(client, v3.ShowGaussMySqlFlavorsOpts{})
	th.AssertNoErr(t, err)
	flav := flavors[0]

	ins, err := instance.CreateGaussMySqlInstance(client, instance.MysqlInstanceOpts{
		Region: "eu-de",
		Name:   "gaussdb-test",
		Datastore: instance.MysqlDatastore{
			Type:    flav.Type,
			Version: flav.VersionName,
		},
		Mode:      "Cluster",
		FlavorRef: flav.SpecCode,
		VpcId:     createOpts.VpcId,
		SubnetId:  createOpts.Nics[0].SubnetId,
		Password:  "gaussdb-test",
	})
	t.Cleanup(func() {
		jobId, err := instance.DeleteGaussMySqlInstance(client, ins.Instance.Id)
		th.AssertNoErr(t, err)
		_, err = v3.WaitForGaussJob(client, jobId, 600)
		th.AssertNoErr(t, err)
	})
	th.AssertNoErr(t, err)
	_, err = v3.WaitForGaussJob(client, ins.JobId, 600)
	th.AssertNoErr(t, err)

	list, err := instance.ListGaussMySqlInstances(client, instance.ListGaussMySqlInstancesOpts{
		Id: ins.Instance.Id,
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, list.Instances[0].Id, ins.Instance.Id)

	dbBackup, err := backup.CreateGaussMySqlBackup(client, backup.MysqlCreateBackupOpts{
		InstanceId: ins.Instance.Id,
		Name:       "gaussdb-test-backup",
	})
	th.AssertNoErr(t, err)
	_, err = v3.WaitForGaussJob(client, dbBackup.JobId, 600)
	th.AssertNoErr(t, err)

	backupList, err := backup.ShowGaussMySqlBackupList(client, backup.ShowGaussMySqlBackupListOpts{
		InstanceId: ins.Instance.Id,
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, backupList.Backups[0].Id, dbBackup.Backup.Id)
}
