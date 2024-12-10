package bucket

import (
	"os"
	"strings"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/obs"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

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
