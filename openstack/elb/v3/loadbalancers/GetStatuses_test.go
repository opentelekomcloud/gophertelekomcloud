package loadbalancers_test

import (
	"net/http"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/elb/v3/loadbalancers"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

const statusTreeJSON = `{
	"statuses": {
		"loadbalancer": {
			"name": "loadbalancer-test",
			"provisioning_status": "ACTIVE",
			"listeners": [{
				"name": "listener-test",
				"provisioning_status": "ACTIVE",
				"pools": [{
					"name": "pool-test",
					"provisioning_status": "ACTIVE",
					"healthmonitor": {"type":"TCP","id":"monitor-id","name":"monitor-test","provisioning_status":"ACTIVE"},
					"members": [{"protocol_port":80,"address":"192.0.2.30","id":"member-id","operating_status":"ONLINE","provisioning_status":"ACTIVE"}],
					"id": "pool-id",
					"operating_status": "ONLINE"
				}],
				"l7policies": [{
					"action": "REDIRECT_TO_POOL",
					"id": "policy-id",
					"provisioning_status": "ACTIVE",
					"name": "policy-test",
					"rules": [{"id":"rule-id","type":"PATH","provisioning_status":"ACTIVE"}]
				}],
				"id": "listener-id",
				"operating_status": "ONLINE"
			}],
			"pools": [],
			"id": "loadbalancer-id",
			"operating_status": "ONLINE"
		}
	},
	"request_id": "request-id"
}`

func TestGetStatuses(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	th.Mux.HandleFunc("/loadbalancers/loadbalancer-id/statuses", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodGet)
		_, _ = w.Write([]byte(statusTreeJSON))
	})
	actual, err := loadbalancers.GetStatuses(serviceClient(), "loadbalancer-id")
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "loadbalancer-id", actual.LoadBalancer.ID)
	th.AssertEquals(t, "listener-id", actual.LoadBalancer.Listeners[0].ID)
	th.AssertEquals(t, "pool-id", actual.LoadBalancer.Listeners[0].Pools[0].ID)
	th.AssertEquals(t, "monitor-id", actual.LoadBalancer.Listeners[0].Pools[0].HealthMonitor.ID)
	th.AssertEquals(t, "member-id", actual.LoadBalancer.Listeners[0].Pools[0].Members[0].ID)
	th.AssertEquals(t, "policy-id", actual.LoadBalancer.Listeners[0].L7Policies[0].ID)
	th.AssertEquals(t, "rule-id", actual.LoadBalancer.Listeners[0].L7Policies[0].Rules[0].ID)
	th.AssertEquals(t, "loadbalancer-id", actual.Loadbalancer.ID)
	th.AssertEquals(t, "listener-id", actual.Loadbalancer.Listeners[0].ID)
}

func TestGetStatusesInvalidID(t *testing.T) {
	actual, err := loadbalancers.GetStatuses(serviceClient(), "../listeners/listener-id")
	if err == nil || actual != nil {
		t.Fatalf("expected invalid ID error, got %#v, %v", actual, err)
	}
}
