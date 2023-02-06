package ims

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/openstack"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestIMS(t *testing.T) {
	clientV1, err := clients.NewIMSV1Client()
	th.AssertNoErr(t, err)

	clientV2, err := clients.NewIMSV2Client()
	th.AssertNoErr(t, err)

	opts := openstack.GetCloudServerCreateOpts(t)

	th.AssertEquals(t, clientV1, clientV1)
	th.AssertEquals(t, clientV2, clientV2)
	th.AssertEquals(t, opts, opts)
}
