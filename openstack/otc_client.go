package openstack

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/utils"
)

type Client interface {
	LoadConfig() error
	AuthenticatedClient() (*golangsdk.ProviderClient, error)
}

// NewOSClient create new instance of Client with given env prefix
func NewOSClient(envPrefix string) Client {
	return &client{envPrefix: envPrefix}
}

// client represents openstack client
type client struct {
	envPrefix  string
	config     *Config
	maxRetries int
}

// LoadConfig load openstack configuration
func (c *client) LoadConfig() error {
	env := NewEnv(c.envPrefix)
	cfg, err := env.LoadOpenstackConfig()
	if err != nil {
		return err
	}
	c.config = cfg
	return nil
}

func info2opts(authInfo *AuthInfo, authType AuthType) (*golangsdk.AuthOptions, error) {
	scope := &golangsdk.AuthScope{
		DomainID:   authInfo.DomainID,
		DomainName: authInfo.DomainName,
	}
	// project scope
	if authInfo.ProjectID != "" || authInfo.ProjectName != "" {
		scope.ProjectID = authInfo.ProjectID
		scope.ProjectName = authInfo.ProjectName

		if authInfo.ProjectDomainName != "" {
			scope.DomainName = authInfo.ProjectDomainName
		}
		if authInfo.ProjectDomainID != "" {
			scope.ProjectID = authInfo.ProjectDomainID
		}
	}
	// user scope
	if authInfo.Username != "" || authInfo.UserID != "" {
		if authInfo.UserDomainName != "" {
			scope.DomainName = authInfo.UserDomainName
		}
		if authInfo.UserDomainID != "" {
			scope.ProjectID = authInfo.UserDomainID
		}
	}
	// agency
	if authInfo.AgencyName != "" && authInfo.AgencyDomainName != "" {
		scope.DomainName = authInfo.AgencyDomainName
	}
	if authInfo.AgencyName != "" && authInfo.DelegatedProject != "" {
		scope.ProjectName = authInfo.DelegatedProject
	}

	var akskOpts *golangsdk.AKSKAuthOptions
	if authInfo.AccessKey != "" && authInfo.SecretKey != "" {
		akskOpts = &golangsdk.AKSKAuthOptions{
			Access: authInfo.AccessKey,
			Secret: authInfo.SecretKey,
		}
	}
	ao := &golangsdk.AuthOptions{
		IdentityEndpoint: authInfo.AuthURL,
		TokenID:          authInfo.Token,
		Username:         authInfo.Username,
		UserID:           authInfo.UserID,
		Password:         authInfo.Password,
		Scope:            scope,
		AKSKAuthOptions:  akskOpts,
	}

	if authType == "" {
		authType = guessAuthType(ao)
	}

	// If an auth_type of "token" was specified, then make sure
	// Gophercloud properly authenticates with a token. This involves
	// unsetting a few other auth options. The reason this is done
	// here is to wait until all auth settings (both in clouds.yaml
	// and via environment variables) are set and then unset them.
	if strings.Contains(string(authType), "token") {
		if ao.TokenID == "" {
			return nil, fmt.Errorf("AuthType is `token`, but no token has been provided")
		}
		ao.Username = ""
		ao.Password = ""
		ao.UserID = ""
		ao.DomainID = ""
		ao.DomainName = ""
	}

	if authType == "aksk" {
		if !validAKSK(ao) {
			return nil, fmt.Errorf("AuthType is `aksk`, but no AK/SK has been provided")
		}
		ao.Username = ""
		ao.Password = ""
		ao.UserID = ""
		ao.DomainID = ""
	}

	// Check for absolute minimum requirements.
	if ao.IdentityEndpoint == "" {
		err := golangsdk.ErrMissingInput{Argument: "auth_url"}
		return nil, err
	}
	return ao, nil
}

func validAKSK(ao *golangsdk.AuthOptions) bool {
	return ao.AKSKAuthOptions != nil && ao.AKSKAuthOptions.Access != "" && ao.AKSKAuthOptions.Secret != ""
}

func guessAuthType(ao *golangsdk.AuthOptions) AuthType {
	if ao.TokenID != "" {
		return "token"
	}
	if validAKSK(ao) {
		return "aksk"
	}
	return "password"
}

func configHttpClient(client *golangsdk.ProviderClient) *golangsdk.ProviderClient {
	transport := &http.Transport{Proxy: http.ProxyFromEnvironment}

	// if OS_DEBUG is set, log the requests and responses
	var osDebug bool
	if os.Getenv("OS_DEBUG") != "" {
		osDebug = true
	}

	client.HTTPClient = http.Client{
		Transport: &LogRoundTripper{
			Rt:         transport,
			OsDebug:    osDebug,
			MaxRetries: 10,
		},
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			if client.AKSKAuthOptions.Access != "" {
				golangsdk.ReSign(req, golangsdk.SignOptions{
					AccessKey: client.AKSKAuthOptions.Access,
					SecretKey: client.AKSKAuthOptions.Secret,
				})
			}
			return nil
		},
	}
	return client
}

func (c *client) AuthenticatedClient() (*golangsdk.ProviderClient, error) {
	e := NewEnv(c.envPrefix)
	cloudConfig := e.CloudFromEnv()
	if c, ok := c.config.Clouds[c.config.DefaultCloud]; ok {
		merged, err := mergeClouds(cloudConfig, c)
		if err != nil {
			return nil, err
		}
		cloudConfig = merged
	}
	authInfo := cloudConfig.AuthInfo
	endpoint := authInfo.AuthURL

	base, err := utils.BaseEndpoint(endpoint)
	if err != nil {
		return nil, err
	}

	endpoint = golangsdk.NormalizeURL(endpoint)
	base = golangsdk.NormalizeURL(base)

	client := new(golangsdk.ProviderClient)
	client.IdentityBase = base
	client.IdentityEndpoint = endpoint
	client.Region = cloudConfig.RegionName
	client.UseTokenLock()

	opts, err := info2opts(&authInfo, cloudConfig.AuthType)
	if err != nil {
		return nil, err
	}
	client = configHttpClient(client)

	err = Authenticate(client, *opts)
	if err != nil {
		return nil, fmt.Errorf("failed to authorize client: %s", err)
	}
	return client, nil
}
