package testing

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/cbr/v3/tasks"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
	fake "github.com/opentelekomcloud/gophertelekomcloud/testhelper/client"
)

func TestGetV3Task(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handleTaskGet(t)

	actual, err := tasks.Get(fake.ServiceClient(), "4827f2da-b008-4507-ab7d-42d0df5ed912").Extract()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, expectedGetResponseData, actual)
}
