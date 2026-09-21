package pools_test

import (
	"net/http"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/pointerto"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/elb/v3/pools"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestCreate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	th.Mux.HandleFunc("/pools", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodPost)
		th.TestJSONRequest(t, r, `{"pool":{"lb_algorithm":"LEAST_CONNECTIONS","protocol":"HTTP","loadbalancer_id":"loadbalancer-id","listener_id":"listener-id","project_id":"project-id","name":"pool-test","description":"pool","session_persistence":{"type":"APP_COOKIE","cookie_name":"session","persistence_timeout":10},"slow_start":{"enable":true,"duration":30},"admin_state_up":true,"member_deletion_protection_enable":false,"vpc_id":"vpc-id","type":"instance","protection_status":"PROTECTED","protection_reason":"managed"}}`)
		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write([]byte(`{"pool":` + poolJSON + `}`))
	})
	actual, err := pools.Create(serviceClient(), pools.CreateOpts{
		LBMethod: "LEAST_CONNECTIONS", Protocol: "HTTP", LoadbalancerID: "loadbalancer-id",
		ListenerID: "listener-id", ProjectID: "project-id", Name: "pool-test", Description: "pool",
		Persistence: &pools.SessionPersistence{
			Type: "APP_COOKIE", CookieName: "session", PersistenceTimeout: 10,
		},
		SlowStart: &pools.SlowStart{Enable: true, Duration: 30}, AdminStateUp: pointerto.Bool(true),
		DeletionProtectionEnable: pointerto.Bool(false), VpcId: "vpc-id", Type: "instance",
		ProtectionStatus: "PROTECTED", ProtectionReason: "managed",
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "pool-id", actual.ID)
	th.AssertEquals(t, "2026-09-21T08:01:00Z", actual.UpdatedAt)
}

func TestCreateMissingRequiredFields(t *testing.T) {
	actual, err := pools.Create(serviceClient(), pools.CreateOpts{})
	if err == nil || actual != nil {
		t.Fatalf("expected validation error, got %#v, %v", actual, err)
	}
}
