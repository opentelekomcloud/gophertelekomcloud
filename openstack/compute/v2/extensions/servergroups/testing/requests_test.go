package testing

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/compute/v2/extensions/servergroups"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
	"github.com/opentelekomcloud/gophertelekomcloud/testhelper/client"
)

func TestList(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListSuccessfully(t)

	actual, err := servergroups.List(client.ServiceClient())
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, ExpectedServerGroupSlice, actual)
}

func TestCreate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleCreateSuccessfully(t)

	actual, err := servergroups.Create(client.ServiceClient(), servergroups.CreateOpts{
		Name:     "test",
		Policies: []string{"anti-affinity"},
	})
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &CreatedServerGroup, actual)
}

func TestGet(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleGetSuccessfully(t)

	actual, err := servergroups.Get(client.ServiceClient(), "4d8c3732-a248-40ed-bebc-539a6ffd25c0")
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &FirstServerGroup, actual)
}

func TestDelete(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleDeleteSuccessfully(t)

	err := servergroups.Delete(client.ServiceClient(), "616fb98f-46ca-475e-917e-2563e5a8cd19")
	th.AssertNoErr(t, err)
}
