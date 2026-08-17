package testing

import (
	"net/http"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/lts/v2/favorites"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
	fake "github.com/opentelekomcloud/gophertelekomcloud/testhelper/client"
)

func TestCreate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handleCreate(t, http.StatusCreated, createResponse)

	actual, err := favorites.Create(fake.ServiceClient(), createOpts)
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, expectedFavorite, actual)
}

func TestCreateExtractsZeroValues(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handleCreate(t, http.StatusCreated, `{}`)

	actual, err := favorites.Create(fake.ServiceClient(), createOpts)
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, &favorites.Favorite{}, actual)
}

func TestCreateRejectsMissingRequiredInput(t *testing.T) {
	_, err := favorites.Create(fake.ServiceClient(), favorites.CreateOpts{})
	if err == nil {
		t.Fatal("expected missing required input to return an error")
	}
}

func TestCreateReturnsBadRequestError(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handleCreate(t, http.StatusBadRequest, `{"message":{"code":"LTS.0603","details":"group or stream not exist"}}`)

	_, err := favorites.Create(fake.ServiceClient(), createOpts)
	if err == nil {
		t.Fatal("expected bad request error")
	}
}
