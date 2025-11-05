package v3

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/gemini/v3/quota"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestGeminiGetQuotas(t *testing.T) {
	client, err := clients.NewGeminiDBClient()
	th.AssertNoErr(t, err)

	t.Logf("Attempting to get Gemini db quotas")
	listResp, err := quota.GetQuotas(client)
	th.AssertNoErr(t, err)
	tools.PrintResource(t, listResp)
}
