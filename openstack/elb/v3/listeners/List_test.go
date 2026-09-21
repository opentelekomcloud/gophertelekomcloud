package listeners_test

import (
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/pointerto"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/elb/v3/listeners"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestList(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	th.Mux.HandleFunc("/listeners", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodGet)
		q := r.URL.Query()
		th.AssertEquals(t, "true", q.Get("page_reverse"))
		th.AssertEquals(t, "false", q.Get("enable_member_retry"))
		switch q.Get("marker") {
		case "start":
			_, _ = fmt.Fprintf(w, `{"listeners":[%s],"page_info":{"previous_marker":"listener-1","next_marker":"listener-2"}}`,
				strings.ReplaceAll(listenerJSON, "listener-id", "listener-1"))
		case "listener-1":
			_, _ = fmt.Fprintf(w, `{"listeners":[%s],"page_info":{}}`,
				strings.ReplaceAll(listenerJSON, "listener-id", "listener-2"))
		default:
			t.Fatalf("unexpected marker %q", q.Get("marker"))
		}
	})
	actual, err := listeners.List(serviceClient(), listeners.ListOpts{
		Limit: 1, Marker: "start", PageReverse: true, ProtocolPort: []int{443},
		Protocol: []listeners.Protocol{listeners.ProtocolHTTPS}, LoadBalancerID: []string{"loadbalancer-id"},
		EnableMemberRetry: pointerto.Bool(false), ProtectionStatus: []string{"consoleProtection"},
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 2, len(actual))
	th.AssertEquals(t, "quic-id", actual[0].QuicConfig.QuicListenerID)
	th.AssertEquals(t, 3000, actual[0].TracingConfig.TracingSample)
}
