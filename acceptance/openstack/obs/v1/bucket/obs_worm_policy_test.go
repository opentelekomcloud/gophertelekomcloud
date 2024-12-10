package bucket

import (
	"strings"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/obs"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestOBSObjectLock(t *testing.T) {
	client, err := clients.NewOBSClient()
	th.AssertNoErr(t, err)

	bucketName := strings.ToLower(tools.RandomString("obs-sdk-test-", 5))

	_, err = client.CreateBucket(&obs.CreateBucketInput{
		Bucket:            bucketName,
		ObjectLockEnabled: true,
	})
	t.Cleanup(func() {
		_, err = client.DeleteBucket(bucketName)
		th.AssertNoErr(t, err)
	})
	th.AssertNoErr(t, err)

	wormOpts := obs.SetWORMPolicyInput{
		Bucket: bucketName,
		BucketWormPolicy: obs.BucketWormPolicy{
			ObjectLockEnabled: "Enabled",
			Mode:              "COMPLIANCE",
			Days:              "10",
		},
	}
	_, err = client.SetWORMPolicy(&wormOpts)
	th.AssertNoErr(t, err)

	getPolicy, err := client.GetWORMPolicy(bucketName)
	th.AssertNoErr(t, err)

	th.AssertEquals(t, wormOpts.Days, getPolicy.Days)

	// disable object lock
	wormOpts.BucketWormPolicy = obs.BucketWormPolicy{}
	_, err = client.SetWORMPolicy(&wormOpts)
	th.AssertNoErr(t, err)
}
