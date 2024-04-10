package v2

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/fgs/v2/quotas"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestFunctionGraphQuotas(t *testing.T) {
	client, err := clients.NewFuncGraphClient()
	th.AssertNoErr(t, err)

	listQuotasResp, err := quotas.ListQuotas(client)
	th.AssertNoErr(t, err)
	tools.PrintResource(t, listQuotasResp)
}
