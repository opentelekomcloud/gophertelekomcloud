package v1

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestRMS(t *testing.T) {
	client, err := clients.NewRmsV1Client()
	th.AssertNoErr(t, err)

	client.Get("resources", nil, nil)
}
