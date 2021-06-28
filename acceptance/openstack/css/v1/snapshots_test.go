package v1

import (
	"testing"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/openstack"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/css/v1/clusters"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/css/v1/snapshots"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/obs"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestSnapshotWorkflow(t *testing.T) {
	client, err := clients.NewCssV1Client()
	th.AssertNoErr(t, err)

	clusterID := createCluster(t, client)
	defer deleteCluster(t, client, clusterID)

	th.AssertNoErr(t, snapshots.Enable(client, clusterID).ExtractErr())
	defer func() {
		th.AssertNoErr(t, snapshots.Disable(client, clusterID).ExtractErr())
	}()

	policyOpts := snapshots.PolicyCreateOpts{
		Prefix:     "snap-",
		Period:     "00:00 GMT+03:00",
		KeepDay:    1,
		Enable:     "true",
		DeleteAuto: "true",
	}
	th.AssertNoErr(t, snapshots.PolicyCreate(client, policyOpts, clusterID).ExtractErr())

	bucketName := "snapshot-sdk-test-bucket"
	createOpts := &obs.CreateBucketInput{
		Bucket: bucketName,
	}
	obsClient, err := clients.NewOBSClient()
	th.AssertNoErr(t, err)

	_, err = obsClient.CreateBucket(createOpts)
	th.AssertNoErr(t, err)

	defer func() {
		_, err = obsClient.DeleteBucket(bucketName)
		th.AssertNoErr(t, err)
	}()

	basicOpts := snapshots.UpdateConfigurationOpts{
		Bucket: bucketName,
	}
	err = snapshots.UpdateConfiguration(client, clusterID, basicOpts).ExtractErr()
	th.AssertNoErr(t, err)

	policy, err := snapshots.PolicyGet(client, clusterID).Extract()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, basicOpts.Bucket, policy.Bucket)
	th.AssertEquals(t, policyOpts.Prefix, policy.Prefix)
	tools.PrintResource(t, policy)
}

func createCluster(t *testing.T, client *golangsdk.ServiceClient) string {
	vpcID := clients.EnvOS.GetEnv("VPC_ID")
	subnetID := clients.EnvOS.GetEnv("NETWORK_ID")

	if vpcID == "" || subnetID == "" {
		t.Skip("Both `VPC_ID` and `NETWORK_ID` need to be defined")
	}

	sgID := openstack.DefaultSecurityGroup(t)

	opts := clusters.CreateOpts{
		Name: tools.RandomString("snap-cluster-", 4),
		Instance: &clusters.InstanceSpec{
			Flavor: "css.medium.8",

			Volume: &clusters.Volume{
				Type: "COMMON",
				Size: 40,
			},
			Nics: &clusters.Nics{
				VpcID:           vpcID,
				SubnetID:        subnetID,
				SecurityGroupID: sgID,
			},
			AvailabilityZone: "eu-de-02",
		},
		InstanceNum: 1,
		DiskEncryption: &clusters.DiskEncryption{
			Encrypted: "0",
		},
	}
	created, err := clusters.Create(client, opts).Extract()
	th.AssertNoErr(t, err)

	th.AssertNoErr(t, clusters.WaitForClusterOperationSucces(client, created.ID, timeout))
	return created.ID
}

func deleteCluster(t *testing.T, client *golangsdk.ServiceClient, id string) {
	err := clusters.Delete(client, id).ExtractErr()
	th.AssertNoErr(t, err)
}
