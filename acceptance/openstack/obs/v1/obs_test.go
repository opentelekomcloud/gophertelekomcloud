package v1

import (
	"fmt"
	"os"
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

func TestObsReplicationLifecycle(t *testing.T) {
	client, err := clients.NewOBSClient()
	th.AssertNoErr(t, err)

	if os.Getenv("OS_AGENCY") == "" || os.Getenv("OS_DESTINATION_BUCKET") == "" {
		t.Skip("Agency or bucket is not provided for the test.")
	}

	bucketName := strings.ToLower(tools.RandomString("obs-sdk-test", 5))
	bucketNameDest := os.Getenv("OS_DESTINATION_BUCKET")
	agencyName := os.Getenv("OS_AGENCY")

	_, err = client.CreateBucket(&obs.CreateBucketInput{
		Bucket: bucketName,
	})
	t.Cleanup(func() {
		_, err = client.DeleteBucket(bucketName)
		th.AssertNoErr(t, err)
	})
	th.AssertNoErr(t, err)

	_, err = client.SetBucketReplication(
		&obs.SetBucketReplicationInput{
			Bucket: bucketName,
			BucketReplicationConfiguration: obs.BucketReplicationConfiguration{
				Agency: agencyName,
				ReplicationRules: []obs.ReplicationRule{
					{
						Prefix:            "",
						Status:            "Enabled",
						DestinationBucket: bucketNameDest,
						DeleteData:        "Enabled",
					},
				},
			},
		})
	th.AssertNoErr(t, err)

	replication, err := client.GetBucketReplication(bucketName)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, replication.StatusCode, 200)
	th.AssertEquals(t, replication.Agency, agencyName)
	th.AssertEquals(t, replication.ReplicationRules[0].DeleteData, obs.EnabledType("Enabled"))
	th.AssertEquals(t, replication.ReplicationRules[0].Status, obs.RuleStatusType("Enabled"))

	_, err = client.DeleteBucketReplication(bucketName)
	th.AssertNoErr(t, err)
}

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

func TestOBSCustomDomain(t *testing.T) {
	client, err := clients.NewOBSClient()
	th.AssertNoErr(t, err)

	bucketName := strings.ToLower(tools.RandomString("obs-sdk-test-", 5))

	_, err = client.CreateBucket(&obs.CreateBucketInput{
		Bucket: bucketName,
	})
	t.Cleanup(func() {
		_, err = client.DeleteBucket(bucketName)
		th.AssertNoErr(t, err)
	})
	th.AssertNoErr(t, err)

	domainName := "www.test.com"

	input := &obs.SetBucketCustomDomainInput{
		Bucket:       bucketName,
		CustomDomain: domainName,
	}
	_, err = client.SetBucketCustomDomain(input)
	th.AssertNoErr(t, err)

	output, err := client.GetBucketCustomDomain(bucketName)
	th.AssertNoErr(t, err)

	tools.PrintResource(t, output)

	inputDelete := &obs.DeleteBucketCustomDomainInput{
		Bucket:       bucketName,
		CustomDomain: domainName,
	}

	_, err = client.DeleteBucketCustomDomain(inputDelete)
	th.AssertNoErr(t, err)
}

func TestOBSInventories(t *testing.T) {
	client, err := clients.NewOBSClient()
	th.AssertNoErr(t, err)

	var (
		configId   = "test-id"
		bucketName = strings.ToLower(tools.RandomString("obs-sdk-test-", 5))
	)

	_, err = client.CreateBucket(&obs.CreateBucketInput{
		Bucket: bucketName,
	})
	t.Cleanup(func() {
		_, err = client.DeleteBucket(bucketName)
		th.AssertNoErr(t, err)
	})
	th.AssertNoErr(t, err)

	inventoryOpts := obs.SetBucketInventoryInput{
		Bucket:            bucketName,
		InventoryConfigId: configId,
		BucketInventoryConfiguration: obs.BucketInventoryConfiguration{
			Id:        configId,
			IsEnabled: true,
			Schedule: obs.InventorySchedule{
				Frequency: "Daily",
			},
			Destination: obs.InventoryDestination{
				Format: "CSV",
				Bucket: bucketName,
				Prefix: "test",
			},
			Filter: obs.InventoryFilter{
				Prefix: "test",
			},
			IncludedObjectVersions: "All",
			OptionalFields: []obs.InventoryOptionalFields{
				{
					Field: "Size",
				},
			},
		},
	}

	_, err = client.SetBucketInventory(&inventoryOpts)
	th.AssertNoErr(t, err)

	getResp, err := client.GetBucketInventory(obs.GetBucketInventoryInput{
		BucketName:        bucketName,
		InventoryConfigId: configId,
	})
	th.AssertNoErr(t, err)

	tools.PrintResource(t, getResp)

	_, err = client.DeleteBucketInventory(&obs.DeleteBucketInventoryInput{
		Bucket:            bucketName,
		InventoryConfigId: configId,
	})
	th.AssertNoErr(t, err)
}
