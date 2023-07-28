package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/structs"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/elb/v3/policies"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
	"github.com/opentelekomcloud/gophertelekomcloud/testhelper/client"
)

func expectedResult() *policies.Policy {
	return &policies.Policy{
		ID:                 "cf4360fd-8631-41ff-a6f5-b72c35da74be",
		Action:             "REDIRECT_TO_LISTENER",
		ListenerID:         "e2220d2a-3faf-44f3-8cd6-0c42952bd0ab",
		Position:           100,
		ProjectID:          "99a3fff0d03c428eac3678da6a7d0f24",
		Status:             "ACTIVE",
		RedirectListenerID: "48a97732-449e-4aab-b561-828d29e45050",
		Rules:              []structs.ResourceRef{},
	}
}

func TestCreateRequest(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/l7policies", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.TestJSONRequest(t, r, createRequestBody)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		_, _ = fmt.Fprint(w, createResponseBody)
	})

	opts := policies.CreateOpts{
		Action:             "REDIRECT_TO_LISTENER",
		ListenerID:         "e2220d2a-3faf-44f3-8cd6-0c42952bd0ab",
		RedirectListenerID: "48a97732-449e-4aab-b561-828d29e45050",
	}
	created, err := policies.Create(client.ServiceClient(), opts).Extract()
	th.AssertNoErr(t, err)

	th.AssertDeepEquals(t, expectedResult(), created)
}

func TestGetRequest(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	expected := expectedResult()
	th.Mux.HandleFunc(fmt.Sprintf("/l7policies/%s", expected.ID), func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprint(w, createResponseBody)
	})

	policy, err := policies.Get(client.ServiceClient(), expected.ID).Extract()
	th.AssertNoErr(t, err)

	th.AssertDeepEquals(t, expected, policy)
}

func TestListRequest(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/l7policies", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprint(w, listResponseBody)
	})

	pages, err := policies.List(client.ServiceClient(), nil).AllPages()
	th.AssertNoErr(t, err)

	policySlice, err := policies.ExtractPolicies(pages)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 2, len(policySlice))
}

func TestDeleteRequest(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	id := expectedResult().ID
	th.Mux.HandleFunc(fmt.Sprintf("/l7policies/%s", id), func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusNoContent)
	})

	th.AssertNoErr(t, policies.Delete(client.ServiceClient(), id).Err)
}
