package v2

import (
	"os"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/image/v2/members"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestImageServiceV2MemberLifecycle(t *testing.T) {
	client, err := clients.NewIMSV2Client()
	th.AssertNoErr(t, err)

	shareProjectID := clients.EnvOS.GetEnv("PROJECT_ID_2")
	privateImageID := clients.EnvOS.GetEnv("PRIVATE_IMAGE_ID")
	if shareProjectID == "" || privateImageID == "" {
		t.Skipf("OS_PROJECT_ID_2 or OS_PRIVATE_IMAGE_ID env vars are missing but IMS member test requires it")
	}
	createOpts := members.CreateOpts{
		ImageId: privateImageID,
		Member:  shareProjectID,
	}

	share, err := members.Create(client, createOpts)
	t.Cleanup(func() {
		th.AssertNoErr(t, members.Delete(client, members.MemberOpts{
			ImageId:  privateImageID,
			MemberId: shareProjectID,
		}))
	})

	th.AssertNoErr(t, err)
	th.AssertEquals(t, createOpts.Member, share.MemberId)
	th.AssertEquals(t, "pending", share.Status)

	newCloud := clients.EnvOS.GetEnv("CLOUD_2")
	if newCloud != "" {
		err = os.Setenv("OS_CLOUD", newCloud)
		th.AssertNoErr(t, err)
		_, err := clients.EnvOS.Cloud(newCloud)
		th.AssertNoErr(t, err)
		newClient, err := clients.NewIMSV2Client()
		th.AssertNoErr(t, err)
		updateOpts := members.UpdateOpts{
			Status: "accepted",
		}
		_, err = members.Update(newClient, privateImageID, shareProjectID, updateOpts)
		th.AssertNoErr(t, err)

		newShare, err := members.Get(client, privateImageID, shareProjectID).Extract()
		th.AssertNoErr(t, err)
		th.AssertEquals(t, updateOpts.Status, newShare.Status)
	}
}
