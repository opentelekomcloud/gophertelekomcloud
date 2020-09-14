package openstack

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/opentelekomcloud/gophertelekomcloud"
	tokens3 "github.com/opentelekomcloud/gophertelekomcloud/openstack/identity/v3/tokens"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/utils"
)

const (
	// v3 represents Keystone v3.
	// The version can be anything from v3 to v3.x.
	v3 = "v3"
)

/*
NewClient prepares an unauthenticated ProviderClient instance.
Most users will probably prefer using the AuthenticatedClient function
instead.

This is useful if you wish to explicitly control the version of the identity
service that's used for authentication explicitly, for example.

A basic example of using this would be:

	ao, err := openstack.AuthOptionsFromEnv()
	provider, err := openstack.NewClient(ao.IdentityEndpoint)
	client, err := openstack.NewIdentityV3(provider, golangsdk.EndpointOpts{})
*/
func NewClient(endpoint string) (*golangsdk.ProviderClient, error) {
	base, err := utils.BaseEndpoint(endpoint)
	if err != nil {
		return nil, err
	}

	endpoint = golangsdk.NormalizeURL(endpoint)
	base = golangsdk.NormalizeURL(base)

	p := new(golangsdk.ProviderClient)
	p.IdentityBase = base
	p.IdentityEndpoint = endpoint
	p.UseTokenLock()

	return p, nil
}

/*
AuthenticatedClient logs in to an OpenStack cloud found at the identity endpoint
specified by the options, acquires a token, and returns a Provider Client
instance that's ready to operate.

If the full path to a versioned identity endpoint was specified  (example:
http://example.com:5000/v3), that path will be used as the endpoint to query.

If a versionless endpoint was specified (example: http://example.com:5000/),
the endpoint will be queried to determine which versions of the identity service
are available, then chooses the most recent or most supported version.

Example:

	ao, err := openstack.AuthOptionsFromEnv()
	provider, err := openstack.AuthenticatedClient(ao)
	client, err := openstack.NewNetworkV2(provider, golangsdk.EndpointOpts{
		Region: os.Getenv("OS_REGION_NAME"),
	})
*/
func AuthenticatedClient(options golangsdk.AuthOptions) (*golangsdk.ProviderClient, error) {
	client, err := NewClient(options.IdentityEndpoint)
	if err != nil {
		return nil, err
	}

	err = Authenticate(client, options)
	if err != nil {
		return nil, err
	}
	return client, nil
}

// Authenticate or re-authenticate against the most recent identity service
// supported at the provided endpoint.
func Authenticate(client *golangsdk.ProviderClient, options golangsdk.AuthOptions) error {
	versions := []*utils.Version{
		{ID: v3, Priority: 30, Suffix: "/v3/"},
	}

	chosen, endpoint, err := utils.ChooseVersion(client, versions)
	if err != nil {
		return err
	}

	if chosen.ID != v3 {
		return fmt.Errorf("unrecognized identity version: %s", chosen.ID)
	}

	return v3auth(client, endpoint, &options, golangsdk.EndpointOpts{})
}

// AuthenticateV3 explicitly authenticates against the identity v3 service.
func AuthenticateV3(client *golangsdk.ProviderClient, options tokens3.AuthOptionsBuilder, eo golangsdk.EndpointOpts) error {
	return v3auth(client, "", options, eo)
}

func v3auth(client *golangsdk.ProviderClient, endpoint string, opts tokens3.AuthOptionsBuilder, eo golangsdk.EndpointOpts) error {
	// Override the generated service endpoint with the one returned by the version endpoint.
	v3Client, err := NewIdentityV3(client, eo)
	if err != nil {
		return err
	}

	if endpoint != "" {
		v3Client.Endpoint = endpoint
	}

	var catalog *tokens3.ServiceCatalog

	var tokenID string
	// passthroughToken allows to passthrough the token without a scope
	var passthroughToken bool
	switch v := opts.(type) {
	case *golangsdk.AuthOptions:
		tokenID = v.TokenID
		passthroughToken = v.Scope == nil || *v.Scope == golangsdk.AuthScope{}
	case *tokens3.AuthOptions:
		tokenID = v.TokenID
		passthroughToken = v.Scope == tokens3.Scope{}
	}

	if tokenID != "" && passthroughToken {
		// passing through the token ID without requesting a new scope
		if opts.CanReauth() {
			return fmt.Errorf("cannot use AllowReauth, when the token ID is defined and auth scope is not set")
		}

		v3Client.SetToken(tokenID)
		result := tokens3.Get(v3Client, tokenID)
		if result.Err != nil {
			return result.Err
		}

		err = client.SetTokenAndAuthResult(result)
		if err != nil {
			return err
		}

		catalog, err = result.ExtractServiceCatalog()
		if err != nil {
			return err
		}
	} else {
		result := tokens3.Create(v3Client, opts)

		err = client.SetTokenAndAuthResult(result)
		if err != nil {
			return err
		}

		catalog, err = result.ExtractServiceCatalog()
		if err != nil {
			return err
		}
	}

	if opts.CanReauth() {
		// here we're creating a throw-away client (tac). it's a copy of the user's provider client, but
		// with the token and reauth func zeroed out. combined with setting `AllowReauth` to `false`,
		// this should retry authentication only once
		tac := *client
		tac.SetThrowaway(true)
		tac.ReauthFunc = nil
		tac.SetTokenAndAuthResult(nil)
		var tao tokens3.AuthOptionsBuilder
		switch ot := opts.(type) {
		case *golangsdk.AuthOptions:
			o := *ot
			o.AllowReauth = false
			tao = &o
		case *tokens3.AuthOptions:
			o := *ot
			o.AllowReauth = false
			tao = &o
		default:
			tao = opts
		}
		client.ReauthFunc = func() error {
			err := v3auth(&tac, endpoint, tao, eo)
			if err != nil {
				return err
			}
			client.CopyTokenFrom(&tac)
			return nil
		}
	}
	client.EndpointLocator = func(opts golangsdk.EndpointOpts) (string, error) {
		return V3EndpointURL(catalog, opts)
	}

	return nil
}

// NewIdentityV3 creates a ServiceClient that may be used to access the v3
// identity service.
func NewIdentityV3(client *golangsdk.ProviderClient, eo golangsdk.EndpointOpts) (*golangsdk.ServiceClient, error) {
	endpoint := client.IdentityBase + "v3/"
	clientType := "identity"
	var err error
	if !reflect.DeepEqual(eo, golangsdk.EndpointOpts{}) {
		eo.ApplyDefaults(clientType)
		endpoint, err = client.EndpointLocator(eo)
		if err != nil {
			return nil, err
		}
	}

	// Ensure endpoint still has a suffix of v3.
	// This is because EndpointLocator might have found a versionless
	// endpoint or the published endpoint is still /v2.0. In both
	// cases, we need to fix the endpoint to point to /v3.
	base, err := utils.BaseEndpoint(endpoint)
	if err != nil {
		return nil, err
	}

	base = golangsdk.NormalizeURL(base)

	endpoint = base + "v3/"

	return &golangsdk.ServiceClient{
		ProviderClient: client,
		Endpoint:       endpoint,
		Type:           clientType,
	}, nil
}

func initClientOpts(client *golangsdk.ProviderClient, eo golangsdk.EndpointOpts, clientType string) (*golangsdk.ServiceClient, error) {
	sc := new(golangsdk.ServiceClient)
	eo.ApplyDefaults(clientType)
	url, err := client.EndpointLocator(eo)
	if err != nil {
		return sc, err
	}
	sc.ProviderClient = client
	sc.Endpoint = url
	sc.Type = clientType
	return sc, nil
}

// initcommonServiceClient create a ServiceClient which can not get from clientType directly.
// firstly, we initialize a service client by "volumev2" type, the endpoint likes https://evs.{region}.{xxx.com}/v2/{project_id}
// then we replace the endpoint with the specified srv and version.
func initcommonServiceClient(client *golangsdk.ProviderClient, eo golangsdk.EndpointOpts, srv string, version string) (*golangsdk.ServiceClient, error) {
	sc, err := initClientOpts(client, eo, "volumev2")
	if err != nil {
		return nil, err
	}

	e := strings.Replace(sc.Endpoint, "v2", version, 1)
	sc.Endpoint = strings.Replace(e, "evs", srv, 1)
	sc.ResourceBase = sc.Endpoint
	return sc, err
}

// TODO: Need to change to apig client type from apig once available
// ApiGateWayV1 creates a service client that is used for Huawei cloud for API gateway.
func ApiGateWayV1(client *golangsdk.ProviderClient, eo golangsdk.EndpointOpts) (*golangsdk.ServiceClient, error) {
	sc, err := initClientOpts(client, eo, "network")
	if err != nil {
		return nil, err
	}
	sc.Endpoint = strings.Replace(sc.Endpoint, "vpc", "apig", 1)
	sc.ResourceBase = sc.Endpoint + "v1.0/apigw/"
	return sc, err
}

// NewBareMetalV1 creates a ServiceClient that may be used with the v1
// bare metal package.
func NewBareMetalV1(client *golangsdk.ProviderClient, eo golangsdk.EndpointOpts) (*golangsdk.ServiceClient, error) {
	return initClientOpts(client, eo, "baremetal")
}

// NewBareMetalIntrospectionV1 creates a ServiceClient that may be used with the v1
// bare metal introspection package.
func NewBareMetalIntrospectionV1(client *golangsdk.ProviderClient, eo golangsdk.EndpointOpts) (*golangsdk.ServiceClient, error) {
	return initClientOpts(client, eo, "baremetal-inspector")
}

// NewObjectStorageV1 creates a ServiceClient that may be used with the v1
// object storage package.
func NewObjectStorageV1(client *golangsdk.ProviderClient, eo golangsdk.EndpointOpts) (*golangsdk.ServiceClient, error) {
	return initClientOpts(client, eo, "object-store")
}

// NewComputeV2 creates a ServiceClient that may be used with the v2 compute
// package.
func NewComputeV2(client *golangsdk.ProviderClient, eo golangsdk.EndpointOpts) (*golangsdk.ServiceClient, error) {
	return initClientOpts(client, eo, "compute")
}

// NewNetworkV2 creates a ServiceClient that may be used with the v2 network
// package.
func NewNetworkV2(client *golangsdk.ProviderClient, eo golangsdk.EndpointOpts) (*golangsdk.ServiceClient, error) {
	sc, err := initClientOpts(client, eo, "network")
	if err != nil {
		return nil, err
	}
	sc.ResourceBase = sc.Endpoint + "v2.0/"
	return sc, err
}

// NewBlockStorageV1 creates a ServiceClient that may be used to access the v1
// block storage service.
func NewBlockStorageV1(client *golangsdk.ProviderClient, eo golangsdk.EndpointOpts) (*golangsdk.ServiceClient, error) {
	return initClientOpts(client, eo, "volume")
}

// NewBlockStorageV2 creates a ServiceClient that may be used to access the v2
// block storage service.
func NewBlockStorageV2(client *golangsdk.ProviderClient, eo golangsdk.EndpointOpts) (*golangsdk.ServiceClient, error) {
	return initClientOpts(client, eo, "volumev2")
}

// NewBlockStorageV3 creates a ServiceClient that may be used to access the v3 block storage service.
func NewBlockStorageV3(client *golangsdk.ProviderClient, eo golangsdk.EndpointOpts) (*golangsdk.ServiceClient, error) {
	return initClientOpts(client, eo, "volumev3")
}

// NewSharedFileSystemV2 creates a ServiceClient that may be used to access the v2 shared file system service.
func NewSharedFileSystemV2(client *golangsdk.ProviderClient, eo golangsdk.EndpointOpts) (*golangsdk.ServiceClient, error) {
	return initClientOpts(client, eo, "sharev2")
}

// NewCDNV1 creates a ServiceClient that may be used to access the OpenStack v1
// CDN service.
func NewCDNV1(client *golangsdk.ProviderClient, eo golangsdk.EndpointOpts) (*golangsdk.ServiceClient, error) {
	return initClientOpts(client, eo, "cdn")
}

// NewOrchestrationV1 creates a ServiceClient that may be used to access the v1
// orchestration service.
func NewOrchestrationV1(client *golangsdk.ProviderClient, eo golangsdk.EndpointOpts) (*golangsdk.ServiceClient, error) {
	return initClientOpts(client, eo, "orchestration")
}

// NewDBV1 creates a ServiceClient that may be used to access the v1 DB service.
func NewDBV1(client *golangsdk.ProviderClient, eo golangsdk.EndpointOpts) (*golangsdk.ServiceClient, error) {
	return initClientOpts(client, eo, "database")
}

// NewDNSV2 creates a ServiceClient that may be used to access the v2 DNS
// service.
func NewDNSV2(client *golangsdk.ProviderClient, eo golangsdk.EndpointOpts) (*golangsdk.ServiceClient, error) {
	sc, err := initClientOpts(client, eo, "dns")
	if err != nil {
		return nil, err
	}
	sc.ResourceBase = sc.Endpoint + "v2/"
	return sc, err
}

// NewImageServiceV1 creates a ServiceClient that may be used to access the v1
// image service.
func NewImageServiceV1(client *golangsdk.ProviderClient, eo golangsdk.EndpointOpts) (*golangsdk.ServiceClient, error) {
	sc, err := initClientOpts(client, eo, "image")
	if err != nil {
		return nil, err
	}
	sc.ResourceBase = sc.Endpoint + "v1/"
	return sc, err
}

// NewImageServiceV2 creates a ServiceClient that may be used to access the v2
// image service.
func NewImageServiceV2(client *golangsdk.ProviderClient, eo golangsdk.EndpointOpts) (*golangsdk.ServiceClient, error) {
	sc, err := initClientOpts(client, eo, "image")
	if err != nil {
		return nil, err
	}
	sc.ResourceBase = sc.Endpoint + "v2/"
	return sc, err
}

// NewLoadBalancerV2 creates a ServiceClient that may be used to access the v2
// load balancer service.
func NewLoadBalancerV2(client *golangsdk.ProviderClient, eo golangsdk.EndpointOpts) (*golangsdk.ServiceClient, error) {
	sc, err := initClientOpts(client, eo, "load-balancer")
	if err != nil {
		return nil, err
	}
	sc.ResourceBase = sc.Endpoint + "v2.0/"
	return sc, err
}

// NewOtcV1 creates a ServiceClient that may be used with the v1 network package.
func NewElbV1(client *golangsdk.ProviderClient, eo golangsdk.EndpointOpts, otctype string) (*golangsdk.ServiceClient, error) {
	sc, err := initClientOpts(client, eo, "compute")
	if err != nil {
		return nil, err
	}
	sc.Endpoint = strings.Replace(strings.Replace(sc.Endpoint, "ecs", otctype, 1), "/v2/", "/v1.0/", 1)
	sc.ResourceBase = sc.Endpoint
	sc.Type = otctype
	return sc, err
}

// NewSmnServiceV2 creates a ServiceClient that may be used to access the v2 Simple Message Notification service.
func NewSmnServiceV2(client *golangsdk.ProviderClient, eo golangsdk.EndpointOpts) (*golangsdk.ServiceClient, error) {
	sc, err := initClientOpts(client, eo, "compute")
	if err != nil {
		return nil, err
	}
	sc.Endpoint = strings.Replace(sc.Endpoint, "ecs", "smn", 1)
	sc.ResourceBase = sc.Endpoint + "notifications/"
	sc.Type = "smn"
	return sc, err
}

// NewRdsServiceV1 creates the a ServiceClient that may be used to access the v1
// rds service which is a service of db instances management.
func NewRdsServiceV1(client *golangsdk.ProviderClient, eo golangsdk.EndpointOpts) (*golangsdk.ServiceClient, error) {
	newsc, err := initClientOpts(client, eo, "compute")
	if err != nil {
		return nil, err
	}
	rdsendpoint := strings.Replace(strings.Replace(newsc.Endpoint, "ecs", "rds", 1), "/v2/", "/rds/v1/", 1)
	newsc.Endpoint = rdsendpoint
	newsc.ResourceBase = rdsendpoint
	newsc.Type = "rds"
	return newsc, err
}

func NewCESClient(client *golangsdk.ProviderClient, eo golangsdk.EndpointOpts) (*golangsdk.ServiceClient, error) {
	sc, err := initClientOpts(client, eo, "volumev2")
	if err != nil {
		return nil, err
	}
	e := strings.Replace(sc.Endpoint, "v2", "V1.0", 1)
	sc.Endpoint = strings.Replace(e, "evs", "ces", 1)
	sc.ResourceBase = sc.Endpoint
	return sc, err
}

// NewDRSServiceV2 creates a ServiceClient that may be used to access the v2 Data Replication Service.
func NewDRSServiceV2(client *golangsdk.ProviderClient, eo golangsdk.EndpointOpts) (*golangsdk.ServiceClient, error) {
	sc, err := initClientOpts(client, eo, "volumev2")
	return sc, err
}

func NewComputeV1(client *golangsdk.ProviderClient, eo golangsdk.EndpointOpts) (*golangsdk.ServiceClient, error) {
	sc, err := initClientOpts(client, eo, "network")
	if err != nil {
		return nil, err
	}
	sc.Endpoint = strings.Replace(sc.Endpoint, "vpc", "ecs", 1)
	sc.Endpoint = sc.Endpoint + "v1/"
	return sc, err
}

func NewComputeV11(client *golangsdk.ProviderClient, eo golangsdk.EndpointOpts) (*golangsdk.ServiceClient, error) {
	sc, err := initClientOpts(client, eo, "ecsv1.1")
	return sc, err
}

func NewEcsV1(client *golangsdk.ProviderClient, eo golangsdk.EndpointOpts) (*golangsdk.ServiceClient, error) {
	sc, err := initClientOpts(client, eo, "ecs")
	return sc, err
}

func NewRdsTagV1(client *golangsdk.ProviderClient, eo golangsdk.EndpointOpts) (*golangsdk.ServiceClient, error) {
	sc, err := initClientOpts(client, eo, "network")
	if err != nil {
		return nil, err
	}
	sc.Endpoint = strings.Replace(sc.Endpoint, "vpc", "rds", 1)
	sc.Endpoint = sc.Endpoint + "v1/"
	sc.ResourceBase = sc.Endpoint + client.ProjectID + "/rds/"
	return sc, err
}

// NewAutoScalingService creates a ServiceClient that may be used to access the
// auto-scaling service of huawei public cloud
func NewAutoScalingService(client *golangsdk.ProviderClient, eo golangsdk.EndpointOpts) (*golangsdk.ServiceClient, error) {
	sc, err := initClientOpts(client, eo, "volumev2")
	if err != nil {
		return nil, err
	}
	e := strings.Replace(sc.Endpoint, "v2", "autoscaling-api/v1", 1)
	sc.Endpoint = strings.Replace(e, "evs", "as", 1)
	sc.ResourceBase = sc.Endpoint
	return sc, err
}

// NewAutoScalingV1 creates a ServiceClient that may be used to access the AS service
func NewAutoScalingV1(client *golangsdk.ProviderClient, eo golangsdk.EndpointOpts) (*golangsdk.ServiceClient, error) {
	sc, err := initClientOpts(client, eo, "asv1")
	return sc, err
}

// NewKmsKeyV1 creates a ServiceClient that may be used to access the v1
// kms key service.
func NewKmsKeyV1(client *golangsdk.ProviderClient, eo golangsdk.EndpointOpts) (*golangsdk.ServiceClient, error) {
	sc, err := initClientOpts(client, eo, "compute")
	if err != nil {
		return nil, err
	}
	sc.Endpoint = strings.Replace(sc.Endpoint, "ecs", "kms", 1)
	sc.Endpoint = sc.Endpoint[:strings.LastIndex(sc.Endpoint, "v2")+3]
	sc.Endpoint = strings.Replace(sc.Endpoint, "v2", "v1.0", 1)
	sc.ResourceBase = sc.Endpoint
	sc.Type = "kms"
	return sc, err
}

func NewElasticLoadBalancer(client *golangsdk.ProviderClient, eo golangsdk.EndpointOpts) (*golangsdk.ServiceClient, error) {
	sc, err := initClientOpts(client, eo, "network")
	if err != nil {
		return sc, err
	}
	sc.Endpoint = strings.Replace(sc.Endpoint, "vpc", "elb", 1)
	sc.ResourceBase = sc.Endpoint + "v1.0/"
	return sc, err
}

// NewNetworkV1 creates a ServiceClient that may be used with the v1 network
// package.
func NewNetworkV1(client *golangsdk.ProviderClient, eo golangsdk.EndpointOpts) (*golangsdk.ServiceClient, error) {
	sc, err := initClientOpts(client, eo, "network")
	if err != nil {
		return nil, err
	}
	sc.ResourceBase = sc.Endpoint + "v1/"
	return sc, err
}

// NewNatV2 creates a ServiceClient that may be used with the v2 nat package.
func NewNatV2(client *golangsdk.ProviderClient, eo golangsdk.EndpointOpts) (*golangsdk.ServiceClient, error) {
	sc, err := initClientOpts(client, eo, "network")
	if err != nil {
		return nil, err
	}
	sc.Endpoint = strings.Replace(sc.Endpoint, "vpc", "nat", 1)
	sc.ResourceBase = sc.Endpoint + "v2.0/"
	return sc, err
}

// MapReduceV1 creates a ServiceClient that may be used with the v1 MapReduce service.
func MapReduceV1(client *golangsdk.ProviderClient, eo golangsdk.EndpointOpts) (*golangsdk.ServiceClient, error) {
	sc, err := initClientOpts(client, eo, "network")
	if err != nil {
		return nil, err
	}
	sc.Endpoint = strings.Replace(sc.Endpoint, "vpc", "mrs", 1)
	sc.Endpoint = sc.Endpoint + "v1.1/"
	sc.ResourceBase = sc.Endpoint + client.ProjectID + "/"
	return sc, err
}

// NewMapReduceV1 creates a ServiceClient that may be used with the v1 MapReduce service.
func NewMapReduceV1(client *golangsdk.ProviderClient, eo golangsdk.EndpointOpts) (*golangsdk.ServiceClient, error) {
	sc, err := initClientOpts(client, eo, "mrs")
	if err != nil {
		return nil, err
	}
	sc.ResourceBase = sc.Endpoint + client.ProjectID + "/"
	return sc, err
}

// AntiDDoSV1 creates a ServiceClient that may be used with the v1 Anti DDoS service.
func AntiDDoSV1(client *golangsdk.ProviderClient, eo golangsdk.EndpointOpts) (*golangsdk.ServiceClient, error) {
	sc, err := initClientOpts(client, eo, "network")
	if err != nil {
		return nil, err
	}
	sc.Endpoint = strings.Replace(sc.Endpoint, "vpc", "antiddos", 1)
	sc.Endpoint = sc.Endpoint + "v1/"
	sc.ResourceBase = sc.Endpoint + client.ProjectID + "/"
	return sc, err
}

// NewAntiDDoSV1 creates a ServiceClient that may be used with the v1 Anti DDoS Service
// package.
func NewAntiDDoSV1(client *golangsdk.ProviderClient, eo golangsdk.EndpointOpts) (*golangsdk.ServiceClient, error) {
	return initClientOpts(client, eo, "antiddos")
}

// NewAntiDDoSV2 creates a ServiceClient that may be used with the v2 Anti DDoS Service
// package.
func NewAntiDDoSV2(client *golangsdk.ProviderClient, eo golangsdk.EndpointOpts) (*golangsdk.ServiceClient, error) {
	sc, err := initClientOpts(client, eo, "antiddos")
	if err != nil {
		return nil, err
	}
	sc.ResourceBase = sc.Endpoint + "v2/" + client.ProjectID + "/"
	return sc, err
}

func NewCCEV3(client *golangsdk.ProviderClient, eo golangsdk.EndpointOpts) (*golangsdk.ServiceClient, error) {
	sc, err := initClientOpts(client, eo, "network")
	if err != nil {
		return nil, err
	}
	sc.Endpoint = strings.Replace(sc.Endpoint, "vpc", "cce", 1)
	sc.ResourceBase = sc.Endpoint + "api/v3/projects/" + client.ProjectID + "/"
	return sc, err
}

func NewCCEAddonV3(client *golangsdk.ProviderClient, eo golangsdk.EndpointOpts) (*golangsdk.ServiceClient, error) {
	sc, err := initClientOpts(client, eo, "network")
	if err != nil {
		return nil, err
	}
	sc.Endpoint = strings.Replace(sc.Endpoint, "vpc", "cce", 1)
	sc.ResourceBase = sc.Endpoint + "api/v3/"
	return sc, err
}

// NewDMSServiceV1 creates a ServiceClient that may be used to access the v1 Distributed Message Service.
func NewDMSServiceV1(client *golangsdk.ProviderClient, eo golangsdk.EndpointOpts) (*golangsdk.ServiceClient, error) {
	sc, err := initClientOpts(client, eo, "network")
	if err != nil {
		return nil, err
	}
	sc.Endpoint = strings.Replace(sc.Endpoint, "vpc", "dms", 1)
	sc.ResourceBase = sc.Endpoint + "v1.0/" + client.ProjectID + "/"
	return sc, err
}

// NewDCSServiceV1 creates a ServiceClient that may be used to access the v1 Distributed Cache Service.
func NewDCSServiceV1(client *golangsdk.ProviderClient, eo golangsdk.EndpointOpts) (*golangsdk.ServiceClient, error) {
	sc, err := initClientOpts(client, eo, "network")
	if err != nil {
		return nil, err
	}
	sc.Endpoint = strings.Replace(sc.Endpoint, "vpc", "dcs", 1)
	sc.ResourceBase = sc.Endpoint + "v1.0/" + client.ProjectID + "/"
	return sc, err
}

// NewOBSService creates a ServiceClient that may be used to access the Object Storage Service.
func NewOBSService(client *golangsdk.ProviderClient, eo golangsdk.EndpointOpts) (*golangsdk.ServiceClient, error) {
	sc, err := initClientOpts(client, eo, "object")
	return sc, err
}

// TODO: Need to change to sfs client type from evs once available
// NewSFSV2 creates a service client that is used for Huawei cloud  for SFS , it replaces the EVS type.
func NewHwSFSV2(client *golangsdk.ProviderClient, eo golangsdk.EndpointOpts) (*golangsdk.ServiceClient, error) {
	sc, err := initClientOpts(client, eo, "network")
	if err != nil {
		return nil, err
	}
	sc.Endpoint = strings.Replace(sc.Endpoint, "vpc", "sfs", 1)
	sc.ResourceBase = sc.Endpoint + "v2/" + client.ProjectID + "/"
	return sc, err
}

func NewBMSV2(client *golangsdk.ProviderClient, eo golangsdk.EndpointOpts) (*golangsdk.ServiceClient, error) {
	sc, err := initClientOpts(client, eo, "compute")
	if err != nil {
		return nil, err
	}
	ep := strings.Replace(sc.Endpoint, "v2", "v2.1", 1)
	sc.Endpoint = ep
	sc.ResourceBase = ep
	return sc, err
}

// NewDeHServiceV1 creates a ServiceClient that may be used to access the v1 Dedicated Hosts service.
func NewDeHServiceV1(client *golangsdk.ProviderClient, eo golangsdk.EndpointOpts) (*golangsdk.ServiceClient, error) {
	sc, err := initClientOpts(client, eo, "deh")
	return sc, err
}

// NewCSBSService creates a ServiceClient that can be used to access the Cloud Server Backup service.
func NewCSBSService(client *golangsdk.ProviderClient, eo golangsdk.EndpointOpts) (*golangsdk.ServiceClient, error) {
	sc, err := initClientOpts(client, eo, "data-protect")
	return sc, err
}

// NewHwCSBSServiceV1 creates a ServiceClient that may be used to access the Huawei Cloud Server Backup service.
func NewHwCSBSServiceV1(client *golangsdk.ProviderClient, eo golangsdk.EndpointOpts) (*golangsdk.ServiceClient, error) {
	sc, err := initClientOpts(client, eo, "compute")
	if err != nil {
		return nil, err
	}
	sc.Endpoint = strings.Replace(sc.Endpoint, "ecs", "csbs", 1)
	e := strings.Replace(sc.Endpoint, "v2", "v1", 1)
	sc.Endpoint = e
	sc.ResourceBase = e
	return sc, err
}

func NewMLSV1(client *golangsdk.ProviderClient, eo golangsdk.EndpointOpts) (*golangsdk.ServiceClient, error) {
	sc, err := initClientOpts(client, eo, "network")
	if err != nil {
		return nil, err
	}
	sc.Endpoint = strings.Replace(sc.Endpoint, "vpc", "mls", 1)
	sc.ResourceBase = sc.Endpoint + "v1.0/" + client.ProjectID + "/"
	return sc, err
}

func NewDWSClient(client *golangsdk.ProviderClient, eo golangsdk.EndpointOpts) (*golangsdk.ServiceClient, error) {
	sc, err := initClientOpts(client, eo, "volumev2")
	if err != nil {
		return nil, err
	}
	e := strings.Replace(sc.Endpoint, "v2", "v1.0", 1)
	sc.Endpoint = strings.Replace(e, "evs", "dws", 1)
	sc.ResourceBase = sc.Endpoint
	return sc, err
}

// NewVBSV2 creates a ServiceClient that may be used to access the VBS service for Orange and Telefonica Cloud.
func NewVBSV2(client *golangsdk.ProviderClient, eo golangsdk.EndpointOpts) (*golangsdk.ServiceClient, error) {
	sc, err := initClientOpts(client, eo, "vbsv2")
	return sc, err
}

// NewVBS creates a service client that is used for VBS.
func NewVBS(client *golangsdk.ProviderClient, eo golangsdk.EndpointOpts) (*golangsdk.ServiceClient, error) {
	sc, err := initClientOpts(client, eo, "volumev2")
	if err != nil {
		return nil, err
	}
	sc.Endpoint = strings.Replace(sc.Endpoint, "evs", "vbs", 1)
	sc.ResourceBase = sc.Endpoint
	return sc, err
}

// NewMAASV1 creates a ServiceClient that may be used to access the MAAS service.
func NewMAASV1(client *golangsdk.ProviderClient, eo golangsdk.EndpointOpts) (*golangsdk.ServiceClient, error) {
	sc, err := initClientOpts(client, eo, "maasv1")
	return sc, err
}

// MAASV1 creates a ServiceClient that may be used with the v1 MAAS service.
func MAASV1(client *golangsdk.ProviderClient, eo golangsdk.EndpointOpts) (*golangsdk.ServiceClient, error) {
	sc, err := initClientOpts(client, eo, "network")
	if err != nil {
		return nil, err
	}
	sc.Endpoint = "https://oms.myhuaweicloud.com/v1/"
	sc.ResourceBase = sc.Endpoint + client.ProjectID + "/"
	return sc, err
}

func NewHwAntiDDoSV1(client *golangsdk.ProviderClient, eo golangsdk.EndpointOpts) (*golangsdk.ServiceClient, error) {
	sc, err := initClientOpts(client, eo, "volumev2")
	if err != nil {
		return nil, err
	}
	e := strings.Replace(sc.Endpoint, "v2", "v1", 1)
	sc.Endpoint = strings.Replace(e, "evs", "antiddos", 1)
	sc.ResourceBase = sc.Endpoint
	return sc, err
}

// NewCTSService creates a ServiceClient that can be used to access the Cloud Trace service.
func NewCTSService(client *golangsdk.ProviderClient, eo golangsdk.EndpointOpts) (*golangsdk.ServiceClient, error) {
	sc, err := initClientOpts(client, eo, "cts")
	return sc, err
}

// NewELBV1 creates a ServiceClient that may be used to access the ELB service.
func NewELBV1(client *golangsdk.ProviderClient, eo golangsdk.EndpointOpts) (*golangsdk.ServiceClient, error) {
	sc, err := initClientOpts(client, eo, "elbv1")
	return sc, err
}

// NewRDSV1 creates a ServiceClient that may be used to access the RDS service.
func NewRDSV1(client *golangsdk.ProviderClient, eo golangsdk.EndpointOpts) (*golangsdk.ServiceClient, error) {
	sc, err := initClientOpts(client, eo, "rdsv1")
	return sc, err
}

// NewKMSV1 creates a ServiceClient that may be used to access the KMS service.
func NewKMSV1(client *golangsdk.ProviderClient, eo golangsdk.EndpointOpts) (*golangsdk.ServiceClient, error) {
	sc, err := initClientOpts(client, eo, "kms")
	return sc, err
}

// NewSMNV2 creates a ServiceClient that may be used to access the SMN service.
func NewSMNV2(client *golangsdk.ProviderClient, eo golangsdk.EndpointOpts) (*golangsdk.ServiceClient, error) {
	sc, err := initClientOpts(client, eo, "smnv2")
	if err != nil {
		return nil, err
	}
	sc.ResourceBase = sc.Endpoint + "notifications/"
	return sc, err
}

// NewCCE creates a ServiceClient that may be used to access the CCE service.
func NewCCE(client *golangsdk.ProviderClient, eo golangsdk.EndpointOpts) (*golangsdk.ServiceClient, error) {
	sc, err := initClientOpts(client, eo, "ccev2.0")
	if err != nil {
		return nil, err
	}
	sc.ResourceBase = sc.Endpoint + "api/v3/projects/" + client.ProjectID + "/"
	return sc, err
}

// NewWAF creates a ServiceClient that may be used to access the WAF service.
func NewWAFV1(client *golangsdk.ProviderClient, eo golangsdk.EndpointOpts) (*golangsdk.ServiceClient, error) {
	sc, err := initClientOpts(client, eo, "waf")
	if err != nil {
		return nil, err
	}
	sc.ResourceBase = sc.Endpoint + "v1/" + client.ProjectID + "/waf/"
	return sc, err
}

// NewRDSV3 creates a ServiceClient that may be used to access the RDS service.
func NewRDSV3(client *golangsdk.ProviderClient, eo golangsdk.EndpointOpts) (*golangsdk.ServiceClient, error) {
	sc, err := initClientOpts(client, eo, "rdsv3")
	return sc, err
}

// SDRSV1 creates a ServiceClient that may be used with the v1 SDRS service.
func SDRSV1(client *golangsdk.ProviderClient, eo golangsdk.EndpointOpts) (*golangsdk.ServiceClient, error) {
	sc, err := initClientOpts(client, eo, "network")
	if err != nil {
		return nil, err
	}
	sc.Endpoint = strings.Replace(sc.Endpoint, "vpc", "sdrs", 1)
	sc.Endpoint = sc.Endpoint + "v1/" + client.ProjectID + "/"
	sc.ResourceBase = sc.Endpoint
	return sc, err
}

// NewSDRSV1 creates a ServiceClient that may be used to access the SDRS service.
func NewSDRSV1(client *golangsdk.ProviderClient, eo golangsdk.EndpointOpts) (*golangsdk.ServiceClient, error) {
	sc, err := initClientOpts(client, eo, "sdrs")
	return sc, err
}

// NewBSSV1 creates a ServiceClient that may be used to access the BSS service.
func NewBSSV1(client *golangsdk.ProviderClient, eo golangsdk.EndpointOpts) (*golangsdk.ServiceClient, error) {
	sc, err := initClientOpts(client, eo, "bssv1")
	return sc, err
}

func NewSDKClient(c *golangsdk.ProviderClient, eo golangsdk.EndpointOpts, serviceType string) (*golangsdk.ServiceClient, error) {
	switch serviceType {
	case "mls":
		return NewMLSV1(c, eo)
	case "dws":
		return NewDWSClient(c, eo)
	case "nat":
		return NewNatV2(c, eo)
	}

	return initClientOpts(c, eo, serviceType)
}

// NewCESV1 creates a ServiceClient that may be used with the v1 CES service.
func NewCESV1(client *golangsdk.ProviderClient, eo golangsdk.EndpointOpts) (*golangsdk.ServiceClient, error) {
	sc, err := initClientOpts(client, eo, "cesv1")
	return sc, err
}

// NewDDSV3 creates a ServiceClient that may be used to access the DDS service.
func NewDDSV3(client *golangsdk.ProviderClient, eo golangsdk.EndpointOpts) (*golangsdk.ServiceClient, error) {
	sc, err := initClientOpts(client, eo, "ddsv3")
	return sc, err
}

// NewLTSV2 creates a ServiceClient that may be used to access the LTS service.
func NewLTSV2(client *golangsdk.ProviderClient, eo golangsdk.EndpointOpts) (*golangsdk.ServiceClient, error) {
	sc, err := initcommonServiceClient(client, eo, "lts", "v2.0")
	return sc, err
}

// NewHuaweiLTSV2 creates a ServiceClient that may be used to access the Huawei Cloud LTS service.
func NewHuaweiLTSV2(client *golangsdk.ProviderClient, eo golangsdk.EndpointOpts) (*golangsdk.ServiceClient, error) {
	sc, err := initcommonServiceClient(client, eo, "lts", "v2")
	return sc, err
}

// NewFGSV2 creates a ServiceClient that may be used with the v2 as
// package.
func NewFGSV2(client *golangsdk.ProviderClient, eo golangsdk.EndpointOpts) (*golangsdk.ServiceClient, error) {
	sc, err := initClientOpts(client, eo, "fgsv2")
	return sc, err
}

// NewVPCV1 creates a ServiceClient that may be used with the v1 network
// package.
func NewVPCV1(client *golangsdk.ProviderClient, eo golangsdk.EndpointOpts) (*golangsdk.ServiceClient, error) {
	sc, err := initClientOpts(client, eo, "vpc")
	return sc, err
}
