package testing

import (
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/cbr/v3/policies"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
	fake "github.com/opentelekomcloud/gophertelekomcloud/testhelper/client"
)

func TestCreateV3PolicyMarshall(t *testing.T) {
	res, err := build.RequestBodyMap(createOpts, "policy")
	th.AssertNoErr(t, err)
	th.AssertJSONEquals(t, expectedRequest, res)
}

func TestCreateV3Policy(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handlePolicyCreation(t)

	actual, err := policies.Create(fake.ServiceClient(), createOpts)
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, expectedCreateResponseData, actual)
}

func TestDeleteV3Policy(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handlePolicyDeletion(t)

	err := policies.Delete(fake.ServiceClient(), "d32019d3-bc6e-4319-9c1d-6722fc136a22")
	th.AssertNoErr(t, err)
}

func TestUpdateV3Policy(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handlePolicyUpdate(t)

	updateId := "cbb3ce6f-3332-4e7c-b98e-77290d8471ff"
	actual, err := policies.Update(fake.ServiceClient(), updateId, updateOpts)
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, expectedCreateResponseData, actual)
}
