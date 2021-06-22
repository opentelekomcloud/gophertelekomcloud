package v2

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/imageservice/v2/members"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestImageServiceV2MemberLifecycle(t *testing.T) {
	client, err := clients.NewImageServiceV2Client()
	th.AssertNoErr(t, err)

	projectID := clients.EnvOS.GetEnv("PROJECT_ID_2")
	privateImageID := clients.EnvOS.GetEnv("PRIVATE_IMAGE_ID")
	if projectID == "" || privateImageID == "" {
		t.Skipf("OS_PROJECT_ID_2 or OS_PRIVATE_IMAGE_ID env vars are missing but IMS member test requires it")
	}
	createOpts := members.CreateOpts{
		Member: projectID,
	}

	share, err := members.Create(client, privateImageID, createOpts).Extract()
	defer func() {
		th.AssertNoErr(t, members.Delete(client, privateImageID, projectID).ExtractErr())
	}()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, createOpts.Member, share.MemberID)
	th.AssertEquals(t, "pending", share.Status)

}
