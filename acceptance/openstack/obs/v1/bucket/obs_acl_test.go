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

func TestObsBucketLifecycleAcl(t *testing.T) {
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

	params := &obs.SetBucketAclInput{
		Bucket: bucketName,
	}

	grants := []obs.Grant{
		{
			Permission: "FULL_CONTROL",
			Grantee: obs.Grantee{
				ID:   domainId,
				Type: "CanonicalUser",
			},
		},
		{
			Permission: "READ",
			Grantee: obs.Grantee{
				ID:   "1000010021",
				Type: "CanonicalUser",
			},
		},
		{
			Permission: "READ_ACP",
			Grantee: obs.Grantee{
				ID:   "1000010021",
				Type: "CanonicalUser",
			},
		},
		{
			Permission: "READ_ACP",
			Grantee: obs.Grantee{
				ID:   "1000010020",
				Type: "CanonicalUser",
			},
		},
		{
			Permission: "READ",
			Grantee: obs.Grantee{
				Type: "Group",
				URI:  "LogDelivery",
			},
		},
		{
			Permission: "WRITE",
			Grantee: obs.Grantee{
				Type: "Group",
				URI:  "LogDelivery",
			},
		},
	}

	params.Owner.ID = domainId
	params.Grants = grants

	_, err = client.SetBucketAcl(params)
	th.AssertNoErr(t, err)

	output, err := client.GetBucketAcl(bucketName)
	th.AssertNoErr(t, err)
	tools.PrintResource(t, output)
}
