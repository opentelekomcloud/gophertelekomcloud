package bucket

import (
	"io"
	"strings"
	"testing"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/eps/v1/resources"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/obs"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestObsBucketEnterpriseProjectMigration(t *testing.T) {
	enterpriseProjectID := clients.EnvOS.GetEnv("ENTERPRISE_PROJECT_ID")
	if enterpriseProjectID == "" || enterpriseProjectID == "0" {
		t.Skip("OS_ENTERPRISE_PROJECT_ID must specify a non-default enterprise project")
	}
	projectID := clients.EnvOS.GetEnv("PROJECT_ID")
	region := clients.EnvOS.GetEnv("REGION_NAME")
	if projectID == "" || region == "" {
		t.Skip("OS_PROJECT_ID and OS_REGION_NAME are required for OBS migration")
	}

	client, err := clients.NewOBSClient()
	th.AssertNoErr(t, err)

	epsClient, err := clients.NewEPSV1Client()
	th.AssertNoErr(t, err)

	bucketName := strings.ToLower(tools.RandomString("obs-sdk-eps-test-", 8))
	_, err = client.CreateBucket(&obs.CreateBucketInput{
		Bucket:         bucketName,
		Epid:           "0",
		BucketLocation: obs.BucketLocation{Location: region},
	})
	th.AssertNoErr(t, err)
	t.Cleanup(func() {
		_, err := client.DeleteBucket(bucketName)
		th.AssertNoErr(t, err)
	})

	metadata, err := client.GetBucketMetadata(&obs.GetBucketMetadataInput{Bucket: bucketName})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "0", metadata.Epid)

	const key = "migration-test.txt"
	const content = "object content preserved during enterprise project migration"
	_, err = client.PutObject(&obs.PutObjectInput{
		PutObjectBasicInput: obs.PutObjectBasicInput{
			ObjectOperationInput: obs.ObjectOperationInput{Bucket: bucketName, Key: key},
		},
		Body: strings.NewReader(content),
	})
	th.AssertNoErr(t, err)
	t.Cleanup(func() {
		_, err := client.DeleteObject(&obs.DeleteObjectInput{Bucket: bucketName, Key: key})
		th.AssertNoErr(t, err)
	})

	original := readObsMigrationObject(t, client, bucketName, key, content)
	for _, target := range []string{enterpriseProjectID, "0"} {
		t.Logf("Migrating OBS bucket %s to enterprise project %s", bucketName, target)
		err = resources.Migrate(epsClient, target, resources.MigrateOpts{
			ProjectID:    projectID,
			ResourceID:   bucketName,
			ResourceType: "bucket",
			RegionID:     region,
		})
		th.AssertNoErr(t, err)

		err = golangsdk.WaitFor(300, func() (bool, error) {
			metadata, err := client.GetBucketMetadata(&obs.GetBucketMetadataInput{Bucket: bucketName})
			if err != nil {
				return false, err
			}
			return metadata.Epid == target, nil
		})
		th.AssertNoErr(t, err)

		object := readObsMigrationObject(t, client, bucketName, key, content)
		th.AssertEquals(t, original.ETag, object.ETag)
		if !object.LastModified.Equal(original.LastModified) {
			t.Fatal("object was modified during enterprise project migration")
		}
	}
}

func readObsMigrationObject(t *testing.T, client *obs.ObsClient, bucket, key, content string) obs.GetObjectMetadataOutput {
	t.Helper()
	output, err := client.GetObject(&obs.GetObjectInput{
		GetObjectMetadataInput: obs.GetObjectMetadataInput{Bucket: bucket, Key: key},
	})
	th.AssertNoErr(t, err)
	body, err := io.ReadAll(output.Body)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, content, string(body))
	return output.GetObjectMetadataOutput
}
