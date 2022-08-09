package v1

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/ces/v1/quotas"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestQuotas(t *testing.T) {
	client, err := clients.NewCesV1Client()
	th.AssertNoErr(t, err)

	quota, err := quotas.ShowQuotas(client).Extract()
	th.AssertNoErr(t, err)

	th.AssertEquals(t, quota.Quotas.Resources[0].Type, "alarm")
}
