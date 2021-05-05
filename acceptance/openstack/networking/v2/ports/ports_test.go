package ports

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/networking/v2/ports"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestPortList(t *testing.T) {
	client, err := clients.NewNetworkV2Client()
	th.AssertNoErr(t, err)

	listOpts := ports.ListOpts{}
	portPages, err := ports.List(client, listOpts).AllPages()
	th.AssertNoErr(t, err)
	portList, err := ports.ExtractPorts(portPages)
	th.AssertNoErr(t, err)

	for _, port := range portList {
		tools.PrintResource(t, port)
	}
}

func TestPortLifecycle(t *testing.T) {
	client, err := clients.NewNetworkV2Client()
	th.AssertNoErr(t, err)

	network := CreateNetwork(t, client)
	defer DeleteNetwork(t, client, network.ID)

	createName := tools.RandomString("create-port-", 3)
	adminStateUp := true
	portSecurity := false
	createOpts := ports.CreateOpts{
		NetworkID:    network.ID,
		Name:         createName,
		AdminStateUp: &adminStateUp,
		PortSecurity: &portSecurity,
	}

	port, err := ports.Create(client, createOpts).Extract()
	th.AssertNoErr(t, err)
	tools.PrintResource(t, port)
	th.AssertEquals(t, true, port.AdminStateUp)
	th.AssertEquals(t, false, port.PortSecurity)
	defer func() {
		err := ports.Delete(client, port.ID).ExtractErr()
		th.AssertNoErr(t, err)
	}()

	updateName := tools.RandomString("update-port-", 3)
	portSecurity = true
	updateOpts := &ports.UpdateOpts{
		Name:         updateName,
		PortSecurity: &portSecurity,
	}

	_, err = ports.Update(client, port.ID, updateOpts).Extract()
	th.AssertNoErr(t, err)

	newPort, err := ports.Get(client, port.ID).Extract()
	th.AssertNoErr(t, err)
	tools.PrintResource(t, newPort)
	th.AssertEquals(t, updateName, newPort.Name)
	th.AssertEquals(t, true, newPort.PortSecurity)
}
