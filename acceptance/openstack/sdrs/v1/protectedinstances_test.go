package v1

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/sdrs/v1/protectedinstances"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestSDRSInstanceList(t *testing.T) {
	client, err := clients.NewSDRSV1()
	th.AssertNoErr(t, err)

	listOpts := protectedinstances.ListOpts{}
	allPages, err := protectedinstances.List(client, listOpts).AllPages()
	th.AssertNoErr(t, err)

	sdrsInstances, err := protectedinstances.ExtractInstances(allPages)
	th.AssertNoErr(t, err)

	for _, instance := range sdrsInstances {
		tools.PrintResource(t, instance)
	}
}

func TestSDRSInstanceLifecycle(t *testing.T) {
	_, err := clients.NewSDRSV1()
	th.AssertNoErr(t, err)
}
