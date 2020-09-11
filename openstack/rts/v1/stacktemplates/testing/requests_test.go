package testing

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/rts/v1/stacktemplates"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
	fake "github.com/opentelekomcloud/gophertelekomcloud/testhelper/client"
)

func TestGetTemplate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleGetSuccessfully(t, GetOutput)

	actual, err := stacktemplates.Get(fake.ServiceClient(), "postman_stack", "16ef0584-4458-41eb-87c8-0dc8d5f66c87").Extract()
	th.AssertNoErr(t, err)

	expected := GetExpected
	th.AssertDeepEquals(t, expected, string(actual))
}
