package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/eps/v1/providers"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
	"github.com/opentelekomcloud/gophertelekomcloud/testhelper/client"
)

func TestListProviders(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/enterprise-projects/providers", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodGet)
		w.Header().Add("Content-Type", "application/json")
		_, _ = fmt.Fprint(w, `{"providers":[{"provider":"ecs","provider_i18n_display_name":"Elastic Cloud Server"},{"provider":"evs","provider_i18n_display_name":"Elastic Volume Service"}]}`)
	})

	actual, err := providers.List(client.ServiceClient(), providers.ListOpts{})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 2, len(actual))
	th.AssertEquals(t, "ecs", actual[0].Provider)
	th.AssertEquals(t, "Elastic Cloud Server", actual[0].ProviderI18nDisplay)
	th.AssertEquals(t, "evs", actual[1].Provider)
}

func TestListProvidersWithOpts(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/enterprise-projects/providers", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodGet)
		th.AssertEquals(t, "en-us", r.URL.Query().Get("locale"))
		th.AssertEquals(t, "10", r.URL.Query().Get("limit"))
		w.Header().Add("Content-Type", "application/json")
		_, _ = fmt.Fprint(w, `{"providers":[{"provider":"ecs","provider_i18n_display_name":"Elastic Cloud Server"}]}`)
	})

	actual, err := providers.List(client.ServiceClient(), providers.ListOpts{
		Locale: "en-us",
		Limit:  10,
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 1, len(actual))
}
