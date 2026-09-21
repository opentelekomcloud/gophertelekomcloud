package pools_test

import (
	"net/http"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/pointerto"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/elb/v3/pools"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestUpdate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	th.Mux.HandleFunc("/pools/pool-id", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodPut)
		th.TestJSONRequest(t, r, `{"pool":{"name":"","description":"","lb_algorithm":"ROUND_ROBIN","session_persistence":{"type":"SOURCE_IP","persistence_timeout":5},"admin_state_up":true,"slow_start":{"enable":false,"duration":0},"member_deletion_protection_enable":false,"vpc_id":"vpc-id","type":"instance"}}`)
		_, _ = w.Write([]byte(`{"pool":` + poolJSON + `}`))
	})
	empty := ""
	actual, err := pools.Update(serviceClient(), "pool-id", pools.UpdateOpts{
		Name: &empty, Description: &empty, LBMethod: "ROUND_ROBIN",
		Persistence:  &pools.SessionPersistence{Type: "SOURCE_IP", PersistenceTimeout: 5},
		AdminStateUp: pointerto.Bool(true), SlowStart: &pools.SlowStart{},
		DeletionProtectionEnable: pointerto.Bool(false), VpcId: "vpc-id", Type: "instance",
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "pool-id", actual.ID)
	th.AssertEquals(t, "managed", actual.ProtectionReason)
}

func TestUpdateInvalidID(t *testing.T) {
	actual, err := pools.Update(serviceClient(), "../listeners/listener-id", pools.UpdateOpts{})
	if err == nil || actual != nil {
		t.Fatalf("expected invalid ID error, got %#v, %v", actual, err)
	}
}
