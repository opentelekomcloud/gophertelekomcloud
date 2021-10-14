package vpcep

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestEndpointVersions(t *testing.T) {
	client, err := clients.NewVPCEndpointV1Client()
	th.AssertNoErr(t, err)

	var resp interface{}
	_, err = client.Get(client.Endpoint, &resp, nil)
	th.AssertNoErr(t, err)
	t.Log(resp)
}
