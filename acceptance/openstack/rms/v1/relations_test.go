package v1

import (
	"os"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/rms/relations"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestRelationsList(t *testing.T) {
	t.Skip("You are not authorized with rms:resources:list impossible to run within CI")
	var resourceId string
	if resourceId = os.Getenv("RESOURCE_ID"); resourceId == "" {
		t.Skip("RESOURCE_ID is required for this test")
	}
	client, err := clients.NewRMSClient()
	th.AssertNoErr(t, err)

	listOpts := relations.ListAllOpts{
		DomainId:   client.DomainID,
		ResourceId: resourceId,
		Direction:  "out",
	}

	listRelations, err := relations.ListRelations(
		client,
		listOpts,
	)
	th.AssertNoErr(t, err)
	tools.AssertLengthGreaterThan(t, listRelations, 1)
}
