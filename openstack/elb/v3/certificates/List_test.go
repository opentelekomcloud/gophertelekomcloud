package certificates_test

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/pointerto"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/elb/v3/certificates"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
	"github.com/opentelekomcloud/gophertelekomcloud/testhelper/client"
)

func TestListFilters(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	th.Mux.HandleFunc("/certificates", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodGet)
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.AssertDeepEquals(t, url.Values{
			"id": {"id-1", "id-2"}, "name": {"name one", "name&two"},
			"description": {"first description", "second description"}, "type": {"server", "client"},
			"domain": {"example.com", "*.example.org"}, "limit": {"1"},
			"marker": {"start"}, "page_reverse": {"true"}, "admin_state_up": {"false"},
		}, r.URL.Query())
		_, _ = w.Write([]byte(`{"certificates":[` + certificateJSON + `],"page_info":{}}`))
	})
	actual, err := certificates.List(serviceClient(), certificates.ListOpts{
		ID: []string{"id-1", "id-2"}, Name: []string{"name one", "name&two"},
		Description: []string{"first description", "second description"}, Type: []string{"server", "client"},
		Domain: []string{"example.com", "*.example.org"}, Limit: 1, Marker: "start",
		PageReverse: true, AdminStateUp: pointerto.Bool(false),
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 1, len(actual))
	assertCertificate(t, &actual[0])
}

func TestListPagination(t *testing.T) {
	for _, reverse := range []bool{false, true} {
		t.Run(fmt.Sprint(reverse), func(t *testing.T) {
			th.SetupHTTP()
			defer th.TeardownHTTP()
			requests := 0
			th.Mux.HandleFunc("/certificates", func(w http.ResponseWriter, r *http.Request) {
				requests++
				th.AssertEquals(t, "1", r.URL.Query().Get("limit"))
				th.AssertDeepEquals(t, []string{"first", "second"}, r.URL.Query()["name"])
				marker := "next"
				if reverse {
					marker = "previous"
				}
				switch r.URL.Query().Get("marker") {
				case "start":
					_, _ = fmt.Fprintf(w, `{"certificates":[%s],"page_info":{"next_marker":"next","previous_marker":"previous","current_count":1}}`, certificateJSON)
				case marker:
					_, _ = fmt.Fprintf(w, `{"certificates":[%s],"page_info":{}}`,
						strings.ReplaceAll(certificateJSON, "certificate-id", "second-id"))
				default:
					t.Fatal("unexpected page marker")
				}
			})
			actual, err := certificates.List(serviceClient(), certificates.ListOpts{
				Limit: 1, Marker: "start", PageReverse: reverse, Name: []string{"first", "second"},
			})
			th.AssertNoErr(t, err)
			th.AssertEquals(t, 3, requests)
			th.AssertEquals(t, 2, len(actual))
			th.AssertEquals(t, "certificate-id", actual[0].ID)
			th.AssertEquals(t, "second-id", actual[1].ID)
		})
	}
}

func TestListEmpty(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	th.Mux.HandleFunc("/certificates", func(w http.ResponseWriter, r *http.Request) {
		th.AssertDeepEquals(t, url.Values{"limit": {"2000"}}, r.URL.Query())
		_, _ = w.Write([]byte(`{"certificates":[],"page_info":{}}`))
	})
	actual, err := certificates.List(serviceClient(), certificates.ListOpts{})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 0, len(actual))
}

func TestListDefaultPagination(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	requests := 0
	th.Mux.HandleFunc("/certificates", func(w http.ResponseWriter, r *http.Request) {
		requests++
		if r.URL.Query().Get("limit") != "2000" {
			t.Error("pagination requires the documented default limit")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		switch r.URL.Query().Get("marker") {
		case "":
			_, _ = fmt.Fprintf(w, `{"certificates":[%s],"page_info":{"next_marker":"next"}}`, certificateJSON)
		case "next":
			_, _ = fmt.Fprintf(w, `{"certificates":[%s],"page_info":{}}`,
				strings.ReplaceAll(certificateJSON, "certificate-id", "second-id"))
		default:
			t.Error("unexpected page marker")
			w.WriteHeader(http.StatusBadRequest)
		}
	})
	actual, err := certificates.List(serviceClient(), certificates.ListOpts{})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 3, requests)
	th.AssertEquals(t, 2, len(actual))
	th.AssertEquals(t, "second-id", actual[1].ID)
}

func TestListErrors(t *testing.T) {
	for _, status := range []int{http.StatusForbidden, http.StatusOK} {
		t.Run(http.StatusText(status), func(t *testing.T) {
			th.SetupHTTP()
			defer th.TeardownHTTP()
			th.Mux.HandleFunc("/certificates", func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(status)
				_, _ = w.Write([]byte(`{"certificates":`))
			})
			actual, err := certificates.List(serviceClient(), certificates.ListOpts{})
			if err == nil || actual != nil {
				t.Fatalf("expected error and nil certificates, got %#v, %v", actual, err)
			}
		})
	}
}

func TestListLaterPageError(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	th.Mux.HandleFunc("/certificates", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("marker") == "next" {
			w.WriteHeader(http.StatusForbidden)
			return
		}
		_, _ = fmt.Fprintf(w, `{"certificates":[%s],"page_info":{"next_marker":"next"}}`, certificateJSON)
	})
	actual, err := certificates.List(serviceClient(), certificates.ListOpts{Limit: 1})
	if err == nil || actual != nil {
		t.Fatalf("expected later-page error and nil certificates, got %#v, %v", actual, err)
	}
}

func TestListInvalidPageInfo(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	th.Mux.HandleFunc("/certificates", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintf(w, `{"certificates":[%s],"page_info":{"next_marker":42}}`, certificateJSON)
	})
	actual, err := certificates.List(serviceClient(), certificates.ListOpts{Limit: 1})
	if err == nil || actual != nil {
		t.Fatalf("expected page-info error and nil certificates, got %#v, %v", actual, err)
	}
}
