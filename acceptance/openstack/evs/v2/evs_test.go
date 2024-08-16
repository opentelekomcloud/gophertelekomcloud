package v2

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/blockstorage/v2/volumes"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/evs/v2/cloudvolumes"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
	"testing"
)

func TestEVSv2List(t *testing.T) {
	client, err := clients.NewBlockStorageV2Client()
	th.AssertNoErr(t, err)

	createOpts := volumes.CreateOpts{
		Size: 40,
		Name: tools.RandomString("tf-evs-disk-", 4),
	}

	resp, err := volumes.Create(client, createOpts).Extract()
	th.AssertNoErr(t, err)

	err = waitForEvsAvailable(client, 100, resp.ID)
	th.AssertNoErr(t, err)

	t.Cleanup(func() {
		err = volumes.Delete(client, resp.ID, volumes.DeleteOpts{}).ExtractErr()
		th.AssertNoErr(t, err)
	})

	list, err := cloudvolumes.List(client, cloudvolumes.ListOpts{
		ID: resp.ID,
	})
	th.AssertNoErr(t, err)
	tools.PrintResource(t, list)
}

func waitForEvsAvailable(client *golangsdk.ServiceClient, secs int, volId string) error {
	return golangsdk.WaitFor(secs, func() (bool, error) {
		vol, err := volumes.Get(client, volId).Extract()
		if err != nil {
			return false, err
		}

		if vol.Status == "available" {
			return true, nil
		}
		return false, nil
	})
}
