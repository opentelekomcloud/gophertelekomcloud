package noauth

import (
	"fmt"
	"strings"

	"github.com/opentelekomcloud/gophertelekomcloud"
)

// EndpointOpts specifies a "noauth" Cinder Endpoint.
type EndpointOpts struct {
	// CinderEndpoint [required] is currently only used with "noauth" Cinder.
	// A cinder endpoint with "auth_strategy=noauth" is necessary, for example:
	// http://example.com:8776/v2.
	CinderEndpoint string
}

// NewClient prepares an unauthenticated ProviderClient instance.
func NewClient(options golangsdk.AuthOptions) (*golangsdk.ProviderClient, error) {
	if options.Username == "" {
		options.Username = "admin"
	}
	if options.TenantName == "" {
		options.TenantName = "admin"
	}

	client := &golangsdk.ProviderClient{
		TokenID: fmt.Sprintf("%s:%s", options.Username, options.TenantName),
	}

	return client, nil
}

func NewBlockStorageNoAuth(client *golangsdk.ProviderClient, opts EndpointOpts) (*golangsdk.ServiceClient, error) {
	sc := new(golangsdk.ServiceClient)
	if opts.CinderEndpoint == "" {
		return nil, fmt.Errorf("CinderEndpoint is required")
	}

	token := strings.Split(client.TokenID, ":")
	if len(token) != 2 {
		return nil, fmt.Errorf("Malformed noauth token")
	}

	sc.Endpoint = golangsdk.NormalizeURL(fmt.Sprintf("%s%s", golangsdk.NormalizeURL(opts.CinderEndpoint), token[1]))
	sc.ProviderClient = client
	return sc, nil
}
