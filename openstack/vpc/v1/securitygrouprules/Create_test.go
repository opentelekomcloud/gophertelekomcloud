package securitygrouprules_test

import (
	"net/http"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/pointerto"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/vpc/v1/securitygrouprules"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestCreate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	th.Mux.HandleFunc("/project-id/security-group-rules", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodPost)
		th.TestJSONRequest(t, r, `{"security_group_rule":{"security_group_id":"group-id","description":"allow web","direction":"ingress","ethertype":"IPv4","protocol":"tcp","port_range_min":0,"port_range_max":80,"remote_ip_prefix":"10.0.0.0/24"}}`)
		_, _ = w.Write([]byte(`{"security_group_rule":` + ruleJSON + `}`))
	})
	actual, err := securitygrouprules.Create(serviceClient(), securitygrouprules.CreateOpts{
		SecurityGroupID: "group-id", Description: "allow web", Direction: "ingress",
		EtherType: "IPv4", Protocol: "tcp", PortRangeMin: pointerto.Int(0),
		PortRangeMax: pointerto.Int(80), RemoteIPPrefix: "10.0.0.0/24",
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "rule-id", actual.ID)
	th.AssertEquals(t, 80, *actual.PortRangeMin)
}

func TestCreateMissingRequiredOpts(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	actual, err := securitygrouprules.Create(serviceClient(), securitygrouprules.CreateOpts{})
	if err == nil || actual != nil {
		t.Fatalf("expected nil response and validation error, got %#v, %v", actual, err)
	}
}
