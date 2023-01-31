package ims

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestIMS(t *testing.T) {
	clientV1, err := clients.NewIMSV1Client()
	th.AssertNoErr(t, err)

	clientV2, err := clients.NewIMSV2Client()
	th.AssertNoErr(t, err)
}
