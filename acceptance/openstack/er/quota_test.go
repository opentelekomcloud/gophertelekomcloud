package er

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/er/v3/quota"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestQuotasList(t *testing.T) {
	client, err := clients.NewERClient()
	th.AssertNoErr(t, err)

	listOpts := quota.ListOpts{}
	quotas, err := quota.List(client, listOpts)
	th.AssertNoErr(t, err)

	for _, q := range quotas {
		tools.PrintResource(t, q)
	}
}
