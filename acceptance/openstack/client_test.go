package openstack

import (
	"fmt"
	"net/http"
	"os"
	"sync"
	"testing"
	"time"

	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
	fake "github.com/opentelekomcloud/gophertelekomcloud/testhelper/client"
)

func TestAuthenticatedClient(t *testing.T) {
	cc, err := clients.CloudAndClient()
	th.AssertNoErr(t, err)

	if cc.TokenID == "" {
		t.Errorf("No token ID assigned to the client")
	}

	if cc.ProjectID == "" {
		t.Errorf("Project ID is not set for the client")
	}
	if cc.UserID == "" {
		t.Errorf("User ID is not set for the client")
	}
	if cc.DomainID == "" {
		t.Errorf("Domain ID is not set for the client")
	}
	if cc.RegionID == "" {
		t.Errorf("Region ID is not set for the client")
	}

	// Find the storage service in the service catalog.
	storage, err := openstack.NewObjectStorageV1(cc.ProviderClient, golangsdk.EndpointOpts{
		Region: cc.RegionName,
	})
	th.AssertNoErr(t, err)
	t.Logf("Located a storage service at endpoint: [%s]", storage.Endpoint)
}

func TestAuthTokenNoRegion(t *testing.T) {
	cc, err := clients.CloudAndClient()
	th.AssertNoErr(t, err)

	envPrefix := tools.RandomString("", 5)
	th.AssertNoErr(t, os.Setenv(envPrefix+"_TOKEN", cc.TokenID))
	th.AssertNoErr(t, os.Setenv(envPrefix+"_AUTH_URL", cc.IdentityEndpoint))

	env := openstack.NewEnv(envPrefix)
	client, err := env.AuthenticatedClient()
	th.AssertNoErr(t, err)
	_, err = openstack.NewComputeV2(client, golangsdk.EndpointOpts{})
	th.AssertNoErr(t, err)
}

func TestReAuth(t *testing.T) {
	cloud, err := clients.CloudAndClient()
	th.AssertNoErr(t, err)

	opts, err := openstack.AuthOptionsFromInfo(&cloud.AuthInfo, cloud.AuthType)
	th.AssertNoErr(t, err)

	ao := opts.(golangsdk.AuthOptions)
	ao.AllowReauth = true

	scl, err := openstack.AuthenticatedClient(ao)
	th.AssertNoErr(t, err)

	t.Logf("Creating a compute client")
	_, err = openstack.NewComputeV2(scl, golangsdk.EndpointOpts{
		Region: cloud.RegionName,
	})
	th.AssertNoErr(t, err)

	t.Logf("Sleeping for 1 second")
	time.Sleep(1 * time.Second)
	t.Logf("Attempting to reauthenticate")

	th.AssertNoErr(t, scl.ReauthFunc())

	t.Logf("Creating a compute client")
	_, err = openstack.NewComputeV2(scl, golangsdk.EndpointOpts{
		Region: cloud.RegionName,
	})
	th.AssertNoErr(t, err)
}

type failHandler struct {
	ExpectedFailures int
	ErrorCode        int
	FailCount        int
	mut              *sync.RWMutex
}

func (f *failHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if f.mut == nil {
		f.mut = new(sync.RWMutex)
	}

	defer func() { _ = r.Body.Close() }()
	if f.FailCount < f.ExpectedFailures {
		f.mut.Lock()
		f.FailCount += 1
		f.mut.Unlock()
		w.WriteHeader(f.ErrorCode)
	} else {
		w.WriteHeader(200)
	}
}

func TestGatewayRetry(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	for _, code := range []int{http.StatusBadGateway, http.StatusGatewayTimeout} {

		failHandler := &failHandler{ExpectedFailures: 1, ErrorCode: code}
		th.Mux.Handle(fmt.Sprintf("/%d", code), failHandler)

		codeURL := fmt.Sprintf("%d", code)
		t.Run(codeURL, func(sub *testing.T) {
			client := fake.ServiceClient()
			_, err := client.Delete(client.ServiceURL(codeURL), &golangsdk.RequestOpts{
				OkCodes: []int{200},
			})
			th.AssertNoErr(sub, err)
			th.AssertEquals(sub, failHandler.ExpectedFailures, failHandler.FailCount)
		})
	}
}

func TestTooManyRequestsRetry(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	code := http.StatusTooManyRequests
	retries := 2
	timeout := 20 * time.Second
	failHandler := &failHandler{ExpectedFailures: 1, ErrorCode: code}
	th.Mux.Handle(fmt.Sprintf("/%d", code), failHandler)

	codeURL := fmt.Sprintf("%d", code)
	t.Run(codeURL, func(sub *testing.T) {
		client := fake.ServiceClient()
		client.MaxBackoffRetries = &retries
		client.BackoffRetryTimeout = &timeout

		_, err := client.Delete(client.ServiceURL(codeURL), &golangsdk.RequestOpts{
			OkCodes: []int{200},
		})
		th.AssertNoErr(sub, err)
		th.AssertEquals(sub, failHandler.ExpectedFailures, failHandler.FailCount)
	})
}

func TestAuthTempAKSK(t *testing.T) {
	securityToken := os.Getenv("OS_SECURITY_TOKEN")
	if securityToken == "" {
		t.Skip("OS_SECURITY_TOKEN env var is missing but client_test requires")
	}
	cc, err := clients.CloudAndClient()
	th.AssertNoErr(t, err)

	if cc.ProjectID == "" {
		t.Errorf("Project ID is not set for the client")
	}
	if cc.AuthInfo.AuthURL == "" {
		t.Errorf("Auth URL is not set for the client")
	}
	if cc.AKSKAuthOptions.AccessKey == "" {
		t.Errorf("Access Key is not set for the client")
	}
	if cc.AKSKAuthOptions.SecretKey == "" {
		t.Errorf("Secret Key is not set for the client")
	}
	if cc.AKSKAuthOptions.SecurityToken == "" {
		t.Errorf("Security Token is not set for the client")
	}

	// Find several services in the service catalog.
	storage, err := openstack.NewObjectStorageV1(cc.ProviderClient, golangsdk.EndpointOpts{
		Region: cc.RegionName,
	})
	th.AssertNoErr(t, err)
	t.Logf("Located a storage service at endpoint: [%s]", storage.Endpoint)

	compute, err := openstack.NewComputeV2(cc.ProviderClient, golangsdk.EndpointOpts{
		Region: cc.RegionName,
	})
	th.AssertNoErr(t, err)
	t.Logf("Located a compute service at endpoint: [%s]", compute.Endpoint)
}
