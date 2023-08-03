package v1

import (
	"fmt"
	"strings"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/obs"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestObsBucketLifecycle(t *testing.T) {
	client, err := clients.NewOBSClient()
	th.AssertNoErr(t, err)

	bucketName := strings.ToLower(tools.RandomString("obs-sdk-test", 5))

	_, err = client.CreateBucket(&obs.CreateBucketInput{
		Bucket: bucketName,
	})
	t.Cleanup(func() {
		_, err = client.DeleteBucket(bucketName)
		th.AssertNoErr(t, err)
	})
	th.AssertNoErr(t, err)

	_, err = client.SetBucketEncryption(&obs.SetBucketEncryptionInput{
		Bucket: bucketName,
		BucketEncryptionConfiguration: obs.BucketEncryptionConfiguration{
			SSEAlgorithm: "kms",
		},
	})
	th.AssertNoErr(t, err)

	bucketHead, err := client.GetBucketMetadata(&obs.GetBucketMetadataInput{
		Bucket: bucketName,
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, bucketHead.FSStatus, obs.FSStatusDisabled)
	th.AssertEquals(t, bucketHead.Version, "3.0")
}

func TestObsBucketLifecyclePolicy(t *testing.T) {
	client, err := clients.NewOBSClient()
	th.AssertNoErr(t, err)

	bucketName := strings.ToLower(tools.RandomString("obs-sdk-test", 5))

	_, err = client.CreateBucket(&obs.CreateBucketInput{
		Bucket: bucketName,
	})
	t.Cleanup(func() {
		_, err = client.DeleteBucket(bucketName)
		th.AssertNoErr(t, err)
	})
	th.AssertNoErr(t, err)

	objectName := tools.RandomString("test-obs-", 5)

	_, err = client.PutObject(&obs.PutObjectInput{
		PutObjectBasicInput: obs.PutObjectBasicInput{
			ObjectOperationInput: obs.ObjectOperationInput{
				Bucket: bucketName,
				Key:    objectName,
			},
		},
	})
	th.AssertNoErr(t, err)

	t.Cleanup(func() {
		_, err = client.DeleteObject(&obs.DeleteObjectInput{
			Bucket: bucketName,
			Key:    objectName,
		})
		th.AssertNoErr(t, err)
	})

	_, err = client.SetBucketLifecycleConfiguration(
		&obs.SetBucketLifecycleConfigurationInput{
			Bucket: bucketName,
			BucketLifecycleConfiguration: obs.BucketLifecycleConfiguration{
				LifecycleRules: []obs.LifecycleRule{
					{
						Prefix: "path1/",
						Status: "Enabled",
						Transitions: []obs.Transition{
							{
								Days:         30,
								StorageClass: "COLD",
							},
						},
						Expiration: obs.Expiration{
							Days: 60,
						},
					},
				},
			},
		},
	)
	th.AssertNoErr(t, err)

	config, err := client.GetBucketLifecycleConfiguration(bucketName)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, config.BucketLifecycleConfiguration.LifecycleRules[0].Expiration.Days, 60)

	t.Cleanup(func() {
		_, err := client.DeleteBucketLifecycleConfiguration(bucketName)
		th.AssertNoErr(t, err)
	})

}

func TestObsPolicyLifecycle(t *testing.T) {
	client, err := clients.NewOBSClient()
	th.AssertNoErr(t, err)

	bucketName := strings.ToLower(tools.RandomString("obs-sdk-test", 5))

	_, err = client.CreateBucket(&obs.CreateBucketInput{
		Bucket: bucketName,
	})
	t.Cleanup(func() {
		_, err = client.DeleteBucket(bucketName)
		th.AssertNoErr(t, err)
	})
	th.AssertNoErr(t, err)

	objectName := tools.RandomString("test-obs-", 5)

	_, err = client.PutObject(&obs.PutObjectInput{
		PutObjectBasicInput: obs.PutObjectBasicInput{
			ObjectOperationInput: obs.ObjectOperationInput{
				Bucket: bucketName,
				Key:    objectName,
			},
		},
	})
	th.AssertNoErr(t, err)

	t.Cleanup(func() {
		_, err = client.DeleteObject(&obs.DeleteObjectInput{
			Bucket: bucketName,
			Key:    objectName,
		})
		th.AssertNoErr(t, err)
	})

	policy := fmt.Sprintf(
		`{
  "Statement": [
    {
      "Effect": "Allow",
      "Principal": {
        "ID": [
          "*"
        ]
      },
      "Action": [
        "*"
      ],
      "Resource": [
        "%[1]s/*",
        "%[1]s"
      ]
    }
  ]
}`, bucketName)

	policyInput := &obs.SetBucketPolicyInput{
		Bucket: bucketName,
		Policy: policy,
	}
	_, err = client.SetBucketPolicy(policyInput)
	th.AssertNoErr(t, err)

	t.Cleanup(func() {
		_, err = client.DeleteBucketPolicy(bucketName)
		th.AssertNoErr(t, err)
	})

	_, err = client.GetBucketPolicy(bucketName)
	th.AssertNoErr(t, err)
}
