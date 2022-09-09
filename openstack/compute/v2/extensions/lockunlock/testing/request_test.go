package testing

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/compute/v2/extensions/lockunlock"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
	"github.com/opentelekomcloud/gophertelekomcloud/testhelper/client"
)

const serverID = "{serverId}"

func TestLock(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	mockStartServerResponse(t, serverID)

	err := lockunlock.Lock(client.ServiceClient(), serverID)
	th.AssertNoErr(t, err)
}

func TestUnlock(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	mockStopServerResponse(t, serverID)

	err := lockunlock.Unlock(client.ServiceClient(), serverID)
	th.AssertNoErr(t, err)
}
