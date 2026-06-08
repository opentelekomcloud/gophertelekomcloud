package bucket

import (
	"strings"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/obs"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

// A dotted bucket (auto path-style) must still serve SSE KMS, which needs x-obs-*
// headers — i.e. the OBS signature is kept, not downgraded to V2.
func TestObsPathStyleSSEKMS(t *testing.T) {
	// No WithPathStyle: the dotted bucket name alone must trigger path-style.
	client, err := clients.NewOBSClient()
	th.AssertNoErr(t, err)

	bucketName := strings.ToLower(tools.RandomString("obs-sdk", 5)) + ".path.style"

	_, err = client.CreateBucket(&obs.CreateBucketInput{
		Bucket: bucketName,
	})
	th.AssertNoErr(t, err)
	t.Cleanup(func() {
		_, err = client.DeleteBucket(bucketName)
		th.AssertNoErr(t, err)
	})

	// Fails if path-style had silently downgraded the signature to V2.
	_, err = client.SetBucketEncryption(&obs.SetBucketEncryptionInput{
		Bucket: bucketName,
		BucketEncryptionConfiguration: obs.BucketEncryptionConfiguration{
			SSEAlgorithm: "kms",
		},
	})
	th.AssertNoErr(t, err)

	encryption, err := client.GetBucketEncryption(bucketName)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, encryption.SSEAlgorithm, "kms")
}
