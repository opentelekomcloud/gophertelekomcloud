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
	agencyID := clients.EnvOS.GetEnv("AGENCY_ID")
	if agencyID == "" {
		t.Skipf("OS_AGENCY_ID is required for this test")
	}

	client, err := clients.NewCssV1Client()
	th.AssertNoErr(t, err)

	clusterID := createCluster(t, client)
	defer deleteCluster(t, client, clusterID)
	bucketName := createBucket(t)
	defer deleteBucket(t, bucketName)

	basicOpts := snapshots.UpdateConfigurationOpts{
		Bucket:   bucketName,
		Agency:   agencyID,
		BasePath: "css_repository/css-test",
	}
	err = snapshots.UpdateConfiguration(client, clusterID, basicOpts)
	th.AssertNoErr(t, err)

	policyOpts := snapshots.PolicyCreateOpts{
		Prefix:     "snap",
		Period:     "00:00 GMT+03:00",
		KeepDay:    1,
		Enable:     "true",
		DeleteAuto: "true",
	}
	th.AssertNoErr(t, snapshots.PolicyCreate(client, policyOpts, clusterID))

	policy, err := snapshots.PolicyGet(client, clusterID)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, basicOpts.Bucket, policy.Bucket)
	th.AssertEquals(t, basicOpts.SnapshotCmkID, policy.SnapshotCmkID)
	th.AssertEquals(t, policyOpts.Prefix, policy.Prefix)
	tools.PrintResource(t, policy)

	th.AssertNoErr(t, snapshots.Disable(client, clusterID))
}

func createBucket(t *testing.T) string {
	bucketName := "snapshot-sdk-test-bucket"
	createOpts := &obs.CreateBucketInput{
		Bucket: bucketName,
	}
	obsClient, err := clients.NewOBSClient()
	th.AssertNoErr(t, err)

	_, err = obsClient.CreateBucket(createOpts)
	th.AssertNoErr(t, err)
	return bucketName
}

func deleteBucket(t *testing.T, bucket string) {
	obsClient, err := clients.NewOBSClient()
	th.AssertNoErr(t, err)

	objects, err := obsClient.ListObjects(&obs.ListObjectsInput{Bucket: bucket})
	th.AssertNoErr(t, err)

	objectsToDelete := make([]obs.ObjectToDelete, len(objects.Contents))
	for i, obj := range objects.Contents {
		objectsToDelete[i] = obs.ObjectToDelete{
			Key: obj.Key,
		}
	}
	_, err = obsClient.DeleteObjects(&obs.DeleteObjectsInput{
		Bucket:  bucket,
		Quiet:   true,
		Objects: objectsToDelete,
	})
	th.AssertNoErr(t, err)
}

func createCluster(t *testing.T, client *golangsdk.ServiceClient) string {
	vpcID := clients.EnvOS.GetEnv("VPC_ID")
	subnetID := clients.EnvOS.GetEnv("NETWORK_ID")

	if vpcID == "" || subnetID == "" {
		t.Skip("Both `OS_VPC_ID`, `OS_NETWORK_ID` need to be defined")
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
			AvailabilityZone: "eu-de-01",
		},
		InstanceNum: 1,
		DiskEncryption: &clusters.DiskEncryption{
			Encrypted: "0",
		},
		Datastore: &clusters.Datastore{
			Version: "Opensearch_1.3.6",
			Type:    "elasticsearch",
		},
	}
	created, err := clusters.Create(client, opts)
	th.AssertNoErr(t, err)

	th.AssertNoErr(t, clusters.WaitForClusterOperationSucces(client, created.ID, timeout))
	return created.ID
}

func deleteCluster(t *testing.T, client *golangsdk.ServiceClient, id string) {
	err := clusters.Delete(client, id)
	th.AssertNoErr(t, err)
}
