package tms

import (
	"os"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/pointerto"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/tms/v1/provider"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestTMSProviderList(t *testing.T) {
	if os.Getenv("RUN_TMS_TAGS") == "" {
		t.Skip("unstable test")
	}
	client, err := clients.NewTmsV1Client()
	th.AssertNoErr(t, err)

	p, err := provider.List(client, provider.ListOpts{
		Provider: "ecs",
		Limit:    pointerto.Int(200),
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, true, len(p) > 0)
}
