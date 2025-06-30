package er

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	az "github.com/opentelekomcloud/gophertelekomcloud/openstack/er/v3/availability-zones"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestAZsList(t *testing.T) {
	client, err := clients.NewERClient()
	th.AssertNoErr(t, err)

	listOpts := az.ListOpts{}
	azs, err := az.List(client, listOpts)
	th.AssertNoErr(t, err)

	for _, a := range azs {
		tools.PrintResource(t, a)
	}
}
