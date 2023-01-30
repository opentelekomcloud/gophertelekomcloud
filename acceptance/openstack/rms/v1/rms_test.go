package v1

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/rms/v1/resources"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestRMS(t *testing.T) {
	client, err := clients.NewRmsV1Client()
	th.AssertNoErr(t, err)

	all, err := resources.ListAllResources(client, resources.ListAllResourcesOpts{
		Limit: 1,
	})
	th.AssertNoErr(t, err)
	tools.PrintResource(t, all)
}
