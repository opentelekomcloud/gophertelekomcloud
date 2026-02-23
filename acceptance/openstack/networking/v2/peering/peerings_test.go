package peering

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/networking/v2/peerings"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestPeeringList(t *testing.T) {
	client, err := clients.NewNetworkV2Client()
	if err != nil {
		t.Fatalf("Unable to create a vpc client: %v", err)
	}

	listOpts := peerings.ListOpts{}
	peering, err := peerings.List(client, listOpts)
	if err != nil {
		t.Fatalf("Unable to list peerings: %v", err)
	}
	for _, peering := range peering {
		tools.PrintResource(t, peering)
	}
}

func TestAcceptPeering(t *testing.T) {

	clientV2, peerClientV2, clientV1, peerClientV1, peeringConn := InitiatePeeringConnCommonTasks(t)

	// Delete a vpc peering connection
	defer DeletePeeringConnNResources(t, clientV2, clientV1, peerClientV1, peeringConn)

	peeringConn1, err := peerings.Accept(peerClientV2, peeringConn.ID).ExtractResult()
	if err != nil {
		t.Fatalf("Unable to accept peering request: %v", err)
	}
	tools.PrintResource(t, peeringConn1)

}

func TestRejectPeering(t *testing.T) {

	clientV2, peerClientV2, clientV1, peerClientV1, peeringConn := InitiatePeeringConnCommonTasks(t)

	// Delete a vpc peering connection
	defer DeletePeeringConnNResources(t, clientV2, clientV1, peerClientV1, peeringConn)

	peerConn1, err := peerings.Reject(peerClientV2, peeringConn.ID).ExtractResult()
	if err != nil {
		t.Fatalf("Unable to Reject peering request: %v", err)
	}
	tools.PrintResource(t, peerConn1)

}

func TestPeeringCRUD(t *testing.T) {

	clientV2, peerClientV2, clientV1, peerClientV1, peeringConn := InitiatePeeringConnCommonTasks(t)
	// Delete a vpc peering connection
	defer DeletePeeringConnNResources(t, clientV2, clientV1, peerClientV1, peeringConn)

	th.AssertEquals(t, "Test Peering", peeringConn.Description)

	updateOpts := peerings.UpdateOpts{
		Name:        "test2",
		Description: "Test Updated",
	}

	_, err := peerings.Update(clientV2, peeringConn.ID, updateOpts).Extract()
	th.AssertNoErr(t, err)

	peeringConnGet, err := peerings.Get(peerClientV2, peeringConn.ID).Extract()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, updateOpts.Description, peeringConnGet.Description)
}
