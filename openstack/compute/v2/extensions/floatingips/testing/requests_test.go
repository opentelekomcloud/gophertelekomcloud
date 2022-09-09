package testing

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/compute/v2/extensions/floatingips"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
	"github.com/opentelekomcloud/gophertelekomcloud/testhelper/client"
)

func TestList(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListSuccessfully(t)

	actual, err := floatingips.List(client.ServiceClient())
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, ExpectedFloatingIPsSlice, actual)
}

func TestCreate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleCreateSuccessfully(t)

	actual, err := floatingips.Create(client.ServiceClient(), floatingips.CreateOpts{
		Pool: "nova",
	})
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &CreatedFloatingIP, actual)
}

func TestCreateWithNumericID(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleCreateWithNumericIDSuccessfully(t)

	actual, err := floatingips.Create(client.ServiceClient(), floatingips.CreateOpts{
		Pool: "nova",
	})
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &CreatedFloatingIP, actual)
}

func TestGet(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleGetSuccessfully(t)

	actual, err := floatingips.Get(client.ServiceClient(), "2")
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &SecondFloatingIP, actual)
}

func TestDelete(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleDeleteSuccessfully(t)

	err := floatingips.Delete(client.ServiceClient(), "1")
	th.AssertNoErr(t, err)
}

func TestAssociate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleAssociateSuccessfully(t)

	associateOpts := floatingips.AssociateOpts{
		FloatingIP: "10.10.10.2",
	}

	err := floatingips.AssociateInstance(client.ServiceClient(), "4d8c3732-a248-40ed-bebc-539a6ffd25c0", associateOpts)
	th.AssertNoErr(t, err)
}

func TestAssociateFixed(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleAssociateFixedSuccessfully(t)

	associateOpts := floatingips.AssociateOpts{
		FloatingIP: "10.10.10.2",
		FixedIP:    "166.78.185.201",
	}

	err := floatingips.AssociateInstance(client.ServiceClient(), "4d8c3732-a248-40ed-bebc-539a6ffd25c0", associateOpts)
	th.AssertNoErr(t, err)
}

func TestDisassociateInstance(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleDisassociateSuccessfully(t)

	disassociateOpts := floatingips.DisassociateOpts{
		FloatingIP: "10.10.10.2",
	}

	err := floatingips.DisassociateInstance(client.ServiceClient(), "4d8c3732-a248-40ed-bebc-539a6ffd25c0", disassociateOpts)
	th.AssertNoErr(t, err)
}
