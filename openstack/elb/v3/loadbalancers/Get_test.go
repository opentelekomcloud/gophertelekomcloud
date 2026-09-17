package loadbalancers_test

import (
	"net/http"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/elb/v3/loadbalancers"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestGet(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	th.Mux.HandleFunc("/loadbalancers/loadbalancer-id", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodGet)
		_, _ = w.Write([]byte(`{"loadbalancer":` + loadBalancerJSON + `}`))
	})
	actual, err := loadbalancers.Get(serviceClient(), "loadbalancer-id")
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "loadbalancer-id", actual.ID)
	th.AssertEquals(t, "endpoint-service-id", actual.ProxyProtocolExtensions[0].Extension.EpServiceID)
}

func TestGetInvalidID(t *testing.T) {
	actual, err := loadbalancers.Get(serviceClient(), "../listeners/listener-id")
	if err == nil || actual != nil {
		t.Fatalf("expected invalid ID error, got %#v, %v", actual, err)
	}
}
