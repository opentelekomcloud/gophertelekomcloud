package v1

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/vpc/v1/privateips"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestPrivateIPLifecycle(t *testing.T) {
	client, err := clients.NewVPCV1Client()
	th.AssertNoErr(t, err)

	vpc := createSubnetVPC(t, client)
	t.Cleanup(func() {
		deleteTestVPC(t, client, vpc.ID)
	})

	subnet := createTestSubnet(t, client, vpc.ID)
	t.Cleanup(func() {
		deleteTestSubnet(t, client, subnet.VpcID, subnet.ID)
	})

	created, err := privateips.Create(client, privateips.CreateOpts{
		PrivateIPs: []privateips.PrivateIPRequest{{SubnetID: subnet.NetworkID}},
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 1, len(created))
	privateIPID := created[0].ID
	t.Cleanup(func() {
		if privateIPID != "" {
			th.AssertNoErr(t, privateips.Delete(client, privateIPID))
		}
	})

	actual, err := privateips.Get(client, privateIPID)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, privateIPID, actual.ID)
	th.AssertEquals(t, subnet.NetworkID, actual.SubnetID)

	list, err := privateips.List(client, subnet.NetworkID, privateips.ListOpts{})
	th.AssertNoErr(t, err)
	found := false
	for _, privateIP := range list {
		if privateIP.ID == privateIPID {
			found = true
			break
		}
	}
	th.AssertEquals(t, true, found)

	th.AssertNoErr(t, privateips.Delete(client, privateIPID))
	privateIPID = ""
}
