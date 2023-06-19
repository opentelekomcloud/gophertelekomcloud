package extensions

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/evs/extensions/schedulerstats"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestSchedulerStatsList(t *testing.T) {
	t.Skip("The API does not exist or has not been published in the environment")

	blockClient, err := clients.NewBlockStorageV3Client()
	th.AssertNoErr(t, err)

	listOpts := schedulerstats.ListOpts{
		Detail: true,
	}

	allStats, err := schedulerstats.List(blockClient, listOpts)
	th.AssertNoErr(t, err)

	for _, stat := range allStats {
		tools.PrintResource(t, stat)
	}
}
