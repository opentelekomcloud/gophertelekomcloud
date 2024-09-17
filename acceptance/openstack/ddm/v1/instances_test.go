package v1

import (
	"testing"

	// golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/ddm/v1/instances"

	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestQueryInstances(t *testing.T) {
	client, err := clients.NewDDMV1Client()
	th.AssertNoErr(t, err)

	queryOpts := instances.QueryInstancesOpts{}
	_, err = instances.QueryInstances(client, queryOpts)
	th.AssertNoErr(t, err)
}
