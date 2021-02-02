package eipstags

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/networking/v2/eipstags"
)

func TestEipsTagsLifecycle(t *testing.T) {
	client, err := clients.NewNetworkV2Client()
	if err != nil {
		t.Fatalf("Unable to create NetworkingV2 client: %v", err)
	}

	eipstags.Action(client, "123", eipstags.BatchActionOpts{})

}
