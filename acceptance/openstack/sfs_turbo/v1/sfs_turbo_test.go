package v1

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/openstack"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/sfs_turbo/v1/shares"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestSFSTurboList(t *testing.T) {
	client, err := clients.NewSharedFileSystemTurboV1Client()
	th.AssertNoErr(t, err)

	listOpts := shares.ListOpts{}
	allSharesPages, err := shares.List(client, listOpts).AllPages()
	th.AssertNoErr(t, err)

	sfsTurboShares, err := shares.ExtractTurbos(allSharesPages)
	th.AssertNoErr(t, err)

	for _, share := range sfsTurboShares {
		tools.PrintResource(t, share)
	}
}

func TestSFSTurboLifecycle(t *testing.T) {
	client, err := clients.NewSharedFileSystemTurboV1Client()
	th.AssertNoErr(t, err)

	secGroupID := openstack.CreateSecurityGroup(t)
	defer openstack.DeleteSecurityGroup(t, secGroupID)

	share := createShare(t, client)
	defer deleteShare(t, client, share.ID)

	tools.PrintResource(t, share)

	share = expandShare(t, client, share.ID)
	th.AssertEquals(t, "1000.00", share.Size)
	share = changeShareSG(t, client, share.ID, secGroupID)

	newShare, err := shares.Get(client, share.ID).Extract()
	th.AssertNoErr(t, err)
	tools.PrintResource(t, newShare)
}
