package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/bms/v2/tags"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
	fake "github.com/opentelekomcloud/gophertelekomcloud/testhelper/client"
)

func TestCreateTag(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/servers/2bff7a8a-3934-4f79-b1d6-53dc5540f00e/tags", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestJSONRequest(t, r, `
{
    "tags": [
        "__type_baremetal"
    ]
}
			`)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		_, _ = fmt.Fprint(w, `
{
    "tags": [
        "__type_baremetal"
    ]
}	`)
	})

	options := tags.CreateOpts{
		Tag: []string{"__type_baremetal"},
	}
	n, err := tags.Create(fake.ServiceClient(), "2bff7a8a-3934-4f79-b1d6-53dc5540f00e", options)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "__type_baremetal", n[0])
}

func TestDeleteTag(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/servers/2bff7a8a-3934-4f79-b1d6-53dc5540f00e/tags", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		w.WriteHeader(http.StatusNoContent)
	})

	res := tags.Delete(fake.ServiceClient(), "2bff7a8a-3934-4f79-b1d6-53dc5540f00e")
	th.AssertNoErr(t, res)
}

func TestGetTags(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/servers/2bff7a8a-3934-4f79-b1d6-53dc5540f00e/tags", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		_, _ = fmt.Fprint(w, `
{
    "tags": [
        "__type_baremetal"
    ]
}
		`)
	})

	n, err := tags.Get(fake.ServiceClient(), "2bff7a8a-3934-4f79-b1d6-53dc5540f00e")
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "__type_baremetal", n[0])

}
