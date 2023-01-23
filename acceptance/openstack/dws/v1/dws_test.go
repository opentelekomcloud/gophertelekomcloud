package v1

import (
	"os"
	"testing"
	"time"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/openstack"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/pointerto"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dws/v1/cluster"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dws/v1/snapshot"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestDWS(t *testing.T) {
	if os.Getenv("RUN_DWS_LIFECYCLE") == "" {
		t.Skip("too slow to run in zuul")
	}

	client, err := clients.NewDWSV1Client()
	th.AssertNoErr(t, err)

	vpcID := clients.EnvOS.GetEnv("VPC_ID")
	if vpcID == "" {
		t.Skip("OS_VPC_ID is missing but DWS test requires using existing network")
	}

	subnetID := clients.EnvOS.GetEnv("NETWORK_ID")
	if subnetID == "" {
		t.Skip("OS_SUBNET_ID env var is missing but DWS test requires using existing network")
	}

	// clusterId := "57e5b43d-d5a1-47e7-aef1-1f46a9abb3ab"

	t.Log("Creating cluster")
	name := tools.RandomString("dws-test-", 3)
	clusterId, err := cluster.CreateCluster(client, cluster.CreateClusterOpts{
		NodeType:        "dws.m3.xlarge",
		NumberOfNode:    3,
		SubnetId:        subnetID,
		SecurityGroupId: openstack.DefaultSecurityGroup(t),
		VpcId:           vpcID,
		Name:            name,
		UserName:        "dbadmin",
		UserPwd:         "#dbadmin123",
	})
	th.AssertNoErr(t, err)
	t.Cleanup(func() {
		err = golangsdk.WaitFor(1000, func() (bool, error) {
			err = cluster.DeleteCluster(client, cluster.DeleteClusterOpts{
				ClusterId:              clusterId,
				KeepLastManualSnapshot: pointerto.Int(0),
			})
			if err != nil {
				t.Error(err)
				return false, nil
			}
			return true, nil
		})
		th.AssertNoErr(t, err)
	})

	err = cluster.WaitForCreate(client, clusterId, 1000)
	th.AssertNoErr(t, err)

	t.Log("ResetPassword")
	err = cluster.ResetPassword(client, cluster.ResetPasswordOpts{
		ClusterId:   clusterId,
		NewPassword: "#SomePassword123",
	})
	th.AssertNoErr(t, err)

	err = cluster.WaitForRestart(client, clusterId, 1000)
	th.AssertNoErr(t, err)

	t.Log("ResizeCluster")
	err = cluster.ResizeCluster(client, cluster.ResizeClusterOpts{
		ClusterId: clusterId,
		Count:     1,
	})

	err = cluster.WaitForResize(client, clusterId, 1000)
	th.AssertNoErr(t, err)

	t.Log("RestartCluster")
	err = cluster.RestartCluster(client, cluster.RestartClusterOpts{
		ClusterId: clusterId,
		Restart:   struct{}{},
	})
	th.AssertNoErr(t, err)

	err = cluster.WaitForRestart(client, clusterId, 1000)
	th.AssertNoErr(t, err)

	list, err := cluster.ListClusters(client)
	th.AssertNoErr(t, err)
	tools.PrintResource(t, list)

	t.Log("CreateSnapshot")
	snapId, err := snapshot.CreateSnapshot(client, snapshot.Snapshot{
		Name:      name,
		ClusterId: clusterId,
	})
	th.AssertNoErr(t, err)
	t.Cleanup(func() {
		err = snapshot.DeleteSnapshot(client, snapId)
		th.AssertNoErr(t, err)
	})

	err = snapshot.WaitForSnapshot(client, clusterId, snapId, 1000)
	th.AssertNoErr(t, err)

	t.Log("RestoreCluster")
	newName := tools.RandomString("dws-test-", 3)
	resCId, err := snapshot.RestoreCluster(client, snapshot.RestoreClusterOpts{
		SnapshotId: snapId,
		Name:       newName,
	})
	t.Cleanup(func() {
		err = golangsdk.WaitFor(5000, func() (bool, error) {
			err = cluster.DeleteCluster(client, cluster.DeleteClusterOpts{
				ClusterId:              resCId,
				KeepLastManualSnapshot: pointerto.Int(0),
			})
			if err != nil {
				t.Error(err)
				time.Sleep(10 * time.Second)
				return false, nil
			}
			return true, nil
		})
		th.AssertNoErr(t, err)
	})

	err = snapshot.WaitForRestore(client, resCId, 2000)
	th.AssertNoErr(t, err)

	snaps, err := snapshot.ListSnapshot(client)
	th.AssertNoErr(t, err)
	tools.PrintResource(t, snaps)
}
