package bucket

import (
	"strings"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/obs"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

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
