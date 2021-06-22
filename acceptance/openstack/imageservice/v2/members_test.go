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

	projectID := clients.EnvOS.GetEnv("PROJECT_ID")
	privateImageID := clients.EnvOS.GetEnv("PRIVATE_IMAGE_ID")
	if projectID == "" || privateImageID == "" {
		t.Skipf("OS_PROJECT_ID or OS_PRIVATE_IMAGE_ID env vars are missing but IMS member test requires it")
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

	updateOpts := members.UpdateOpts{
		Status: "accepted",
	}
	_, err = members.Update(client, privateImageID, projectID, updateOpts).Extract()
	th.AssertNoErr(t, err)

	newShare, err := members.Get(client, privateImageID, projectID).Extract()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, updateOpts.Status, newShare.Status)
}
