package bucket

import (
	"strings"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/obs"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestObsBucketLifecycleConfigurationBasic(t *testing.T) {
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

func TestObsBucketLifecycleConfigurationFilter(t *testing.T) {
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

	_, err = client.SetBucketLifecycleConfiguration(
		&obs.SetBucketLifecycleConfigurationInput{
			Bucket: bucketName,
			BucketLifecycleConfiguration: obs.BucketLifecycleConfiguration{
				LifecycleRules: []obs.LifecycleRule{
					{
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
						Filter: obs.LifecycleFilter{
							Prefix: "prefix",
							Tags: []obs.Tag{
								{
									Key:   "tag1",
									Value: "value1",
								},
								{
									Key:   "tag2",
									Value: "value2",
								},
							},
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
	th.AssertEquals(t, config.BucketLifecycleConfiguration.LifecycleRules[0].Filter.Tags[0].Key, "tag1")
	th.AssertEquals(t, config.BucketLifecycleConfiguration.LifecycleRules[0].Filter.Tags[0].Value, "value1")
	th.AssertEquals(t, config.BucketLifecycleConfiguration.LifecycleRules[0].Filter.Tags[1].Key, "tag2")
	th.AssertEquals(t, config.BucketLifecycleConfiguration.LifecycleRules[0].Filter.Tags[1].Value, "value2")

	t.Cleanup(func() {
		_, err := client.DeleteBucketLifecycleConfiguration(bucketName)
		th.AssertNoErr(t, err)
	})
}
