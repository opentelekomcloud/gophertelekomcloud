package testing

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/blockstorage/noauth"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestNoAuth(t *testing.T) {
	ao := golangsdk.AuthOptions{
		Username:   "user",
		TenantName: "test",
	}
	provider, err := noauth.NewClient(ao)
	th.AssertNoErr(t, err)
	noauthClient, err := noauth.NewBlockStorageNoAuth(provider, noauth.EndpointOpts{
		CinderEndpoint: "http://cinder:8776/v2",
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, naTestResult.Endpoint, noauthClient.Endpoint)
	th.AssertEquals(t, naTestResult.TokenID, noauthClient.TokenID)

	ao2 := golangsdk.AuthOptions{}
	provider2, err := noauth.NewClient(ao2)
	th.AssertNoErr(t, err)
	noauthClient2, err := noauth.NewBlockStorageNoAuth(provider2, noauth.EndpointOpts{
		CinderEndpoint: "http://cinder:8776/v2/",
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, naResult.Endpoint, noauthClient2.Endpoint)
	th.AssertEquals(t, naResult.TokenID, noauthClient2.TokenID)

	errTest, err := noauth.NewBlockStorageNoAuth(provider2, noauth.EndpointOpts{})
	_ = errTest
	if err == nil {
		t.Fatalf("Expected to receive error")
	}
	th.AssertEquals(t, errorResult, err.Error())
}
