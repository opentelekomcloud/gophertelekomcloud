package v1

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/vpc/v1/quotas"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestQuotaList(t *testing.T) {
	client, err := clients.NewVPCV1Client()
	th.AssertNoErr(t, err)

	resources, err := quotas.List(client, quotas.ListOpts{Type: "publicIp"})
	th.AssertNoErr(t, err)
	if len(resources) == 0 {
		t.Fatal("expected at least one publicIp quota resource")
	}
	th.AssertEquals(t, "publicIp", resources[0].Type)
}
