package testing

import (
	"net/http"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/lts/v2/favorites"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
	fake "github.com/opentelekomcloud/gophertelekomcloud/testhelper/client"
)

func TestDelete(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handleDelete(t, http.StatusOK, "")

	err := favorites.Delete(fake.ServiceClient(), favorites.DeleteOpts{ResourceID: "favorite-resource-id"})
	th.AssertNoErr(t, err)
}

func TestDeleteRejectsMissingResourceID(t *testing.T) {
	err := favorites.Delete(fake.ServiceClient(), favorites.DeleteOpts{})
	if err == nil {
		t.Fatal("expected missing resource ID to return an error")
	}
}

func TestDeleteReturnsBadRequestError(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handleDelete(
		t,
		http.StatusBadRequest,
		`{"message":{"code":"LTS.0009","details":"update favorite failed"}}`,
	)

	err := favorites.Delete(fake.ServiceClient(), favorites.DeleteOpts{ResourceID: "favorite-resource-id"})
	if err == nil {
		t.Fatal("expected bad request error")
	}
}
