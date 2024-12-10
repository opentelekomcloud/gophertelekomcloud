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

func TestObsObjectAclLifecycle(t *testing.T) {
	domainId := os.Getenv("OS_DOMAIN_ID")
	if domainId == "" {
		t.Skip("OS_DOMAIN_ID is mandatory for this test")
	}
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

	objectAcl := obs.SetObjectAclInput{
		Bucket: bucketName,
		Key:    objectName,
		AccessControlPolicy: obs.AccessControlPolicy{
			Owner: obs.Owner{
				ID: domainId,
			},
			Grants: []obs.Grant{
				{
					Grantee: obs.Grantee{
						ID:   domainId,
						Type: "CanonicalUser",
					},
					Permission: "READ",
				},
				{
					Grantee: obs.Grantee{
						ID:   domainId,
						Type: "CanonicalUser",
					},
					Permission: "WRITE",
				},
				{
					Grantee: obs.Grantee{
						ID:   "1000010020",
						Type: "CanonicalUser",
					},
					Permission: "READ",
				},
				{
					Grantee: obs.Grantee{
						ID:   "1000010021",
						Type: "CanonicalUser",
					},
					Permission: "WRITE_ACP",
				},
				{
					Grantee: obs.Grantee{
						Type: "Group",
						URI:  "AllUsers",
					},
					Permission: "WRITE_ACP",
				},
				{
					Grantee: obs.Grantee{
						Type: "Group",
						URI:  "AllUsers",
					},
					Permission: "READ_ACP",
				},
			},
		},
	}

	_, err = client.SetObjectAcl(&objectAcl)
	th.AssertNoErr(t, err)

	objAcl, err := client.GetObjectAcl(&obs.GetObjectAclInput{
		Bucket: bucketName,
		Key:    objectName,
	})
	th.AssertNoErr(t, err)

	tools.PrintResource(t, objAcl)

}
