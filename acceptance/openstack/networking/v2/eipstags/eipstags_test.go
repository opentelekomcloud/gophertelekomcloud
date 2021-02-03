package eipstags

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/networking/v2/eipstags"
)

func TestEipsTagsLifecycle(t *testing.T) {
	clientV1, err := clients.NewNetworkV1Client()
	if err != nil {
		t.Fatalf("Unable to create NetworkingV1 client: %v", err)
	}
	eip := CreateEip(t, clientV1)
	defer DeleteEip(t, clientV1, eip.ID)

	clientV2, err := clients.NewNetworkV2Client()
	if err != nil {
		t.Fatalf("Unable to create NetworkingV2 client: %v", err)
	}

	tagKey := "muh"
	CreateTag(t, clientV2, eip.ID, tagKey)
	defer DeleteTag(t, clientV2, eip.ID, tagKey)

	tagList, err := eipstags.List(clientV2, eip.ID).Extract()
	if err != nil {
		t.Fatalf("Unable to get EIP tags: %v", err)
	}
	tools.PrintResource(t, tagList[0])
}
