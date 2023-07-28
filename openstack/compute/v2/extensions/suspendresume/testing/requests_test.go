package testing

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/compute/v2/extensions/suspendresume"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
	"github.com/opentelekomcloud/gophertelekomcloud/testhelper/client"
)

const serverID = "{serverId}"

func TestSuspend(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	mockSuspendServerResponse(t, serverID)

	err := suspendresume.Suspend(client.ServiceClient(), serverID).Err
	th.AssertNoErr(t, err)
}

func TestResume(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	mockResumeServerResponse(t, serverID)

	err := suspendresume.Resume(client.ServiceClient(), serverID).Err
	th.AssertNoErr(t, err)
}
