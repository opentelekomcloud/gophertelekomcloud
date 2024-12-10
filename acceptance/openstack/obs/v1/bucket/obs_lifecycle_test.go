package bucket

import (
	"strings"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/obs"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestObsBucketLifecycleConfiguration(t *testing.T) {
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
