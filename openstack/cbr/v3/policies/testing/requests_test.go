package testing

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/cbr/v3/policies"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
	fake "github.com/opentelekomcloud/gophertelekomcloud/testhelper/client"
)

func TestCreateV3PolicyMarshall(t *testing.T) {
	res, err := createOpts.ToPolicyCreateMap()
	th.AssertNoErr(t, err)
	th.AssertJSONEquals(t, expectedRequest, res)
}

func TestCreateV3Policy(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handlePolicyCreation(t)

	actual, err := policies.Create(fake.ServiceClient(), createOpts).Extract()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, expectedCreateResponseData, actual)

}
