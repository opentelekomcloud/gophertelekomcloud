package notifications_test

import (
	"net/http"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/autoscaling/v1/notifications"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
	"github.com/opentelekomcloud/gophertelekomcloud/testhelper/client"
)

func TestDelete(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	th.Mux.HandleFunc("/scaling_notification/group-id/"+topicURN, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodDelete)
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.AssertEquals(t, "", r.URL.RawQuery)
		th.AssertEquals(t, int64(0), r.ContentLength)
		w.WriteHeader(http.StatusNoContent)
	})
	th.AssertNoErr(t, notifications.Delete(client.ServiceClient(), "group-id", topicURN))
}

func TestDeleteErrors(t *testing.T) {
	for _, status := range []int{http.StatusNotFound, http.StatusConflict, http.StatusForbidden} {
		t.Run(http.StatusText(status), func(t *testing.T) {
			th.SetupHTTP()
			defer th.TeardownHTTP()
			th.Mux.HandleFunc("/scaling_notification/group-id/"+topicURN, func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(status)
			})
			if err := notifications.Delete(client.ServiceClient(), "group-id", topicURN); err == nil {
				t.Fatal("expected deletion error")
			}
		})
	}
}

func TestDeleteInvalidIdentifiers(t *testing.T) {
	for _, ids := range [][2]string{
		{"../scaling_group/other-group", topicURN},
		{"group-id", "../other-topic"},
	} {
		if err := notifications.Delete(client.ServiceClient(), ids[0], ids[1]); err == nil {
			t.Fatal("expected invalid identifier error")
		}
	}
}
