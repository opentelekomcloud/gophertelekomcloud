// Package clients contains functions for creating OpenStack service clients
// for use in acceptance tests. It also manages the required environment
// variables to run the tests.
package clients

import (
	"encoding/json"
	"fmt"
	"strings"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/identity/v3/credentials"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/obs"
)

const envPrefix = "OS_"

var EnvOS = openstack.NewEnv(envPrefix)

// NewAutoscalingV1Client returns authenticated AutoScaling v1 client
func NewAutoscalingV1Client() (*golangsdk.ServiceClient, error) {
	cc, err := CloudAndClient()
	if err != nil {
		return nil, err
	}

	return openstack.NewAutoScalingV1(cc.ProviderClient, golangsdk.EndpointOpts{
		Region: cc.RegionName,
	})
}

// NewAutoscalingV2Client returns authenticated AutoScaling v2 client
func NewAutoscalingV2Client() (*golangsdk.ServiceClient, error) {
	cc, err := CloudAndClient()
	if err != nil {
		return nil, err
	}

	return openstack.NewAutoScalingV2(cc.ProviderClient, golangsdk.EndpointOpts{
		Region: cc.RegionName,
	})
}

// NewBlockStorageV1Client returns a *ServiceClient for making calls
// to the OpenStack Block Storage v1 API. An error will be returned
// if authentication or client creation was not possible.
func NewBlockStorageV1Client() (*golangsdk.ServiceClient, error) {
	cc, err := CloudAndClient()
	if err != nil {
		return nil, err
	}

	return openstack.NewBlockStorageV1(cc.ProviderClient, golangsdk.EndpointOpts{
		Region: cc.RegionName,
	})
}

// NewBlockStorageV2Client returns a *ServiceClient for making calls
// to the OpenStack Block Storage v2 API. An error will be returned
// if authentication or client creation was not possible.
func NewBlockStorageV2Client() (*golangsdk.ServiceClient, error) {
	cc, err := CloudAndClient()
	if err != nil {
		return nil, err
	}

	return openstack.NewBlockStorageV2(cc.ProviderClient, golangsdk.EndpointOpts{
		Region: cc.RegionName,
	})
}

// NewBlockStorageV3Client returns a *ServiceClient for making calls
// to the OpenStack Block Storage v3 API. An error will be returned
// if authentication or client creation was not possible.
func NewBlockStorageV3Client() (*golangsdk.ServiceClient, error) {
	cc, err := CloudAndClient()
	if err != nil {
		return nil, err
	}

	return openstack.NewBlockStorageV3(cc.ProviderClient, golangsdk.EndpointOpts{
		Region: cc.RegionName,
	})
}

// NewComputeV2Client returns a *ServiceClient for making calls
// to the OpenStack Compute v2 API. An error will be returned
// if authentication or client creation was not possible.
func NewComputeV2Client() (*golangsdk.ServiceClient, error) {
	cc, err := CloudAndClient()
	if err != nil {
		return nil, err
	}

	return openstack.NewComputeV2(cc.ProviderClient, golangsdk.EndpointOpts{
		Region: cc.RegionName,
	})
}

// NewComputeV1Client returns a *ServiceClient for making calls
// to the OpenStack Compute v1 API. An error will be returned
// if authentication or client creation was not possible.
func NewComputeV1Client() (*golangsdk.ServiceClient, error) {
	cc, err := CloudAndClient()
	if err != nil {
		return nil, err
	}

	return openstack.NewComputeV1(cc.ProviderClient, golangsdk.EndpointOpts{
		Region: cc.RegionName,
	})
}

func NewCTSV1Client() (*golangsdk.ServiceClient, error) {
	cc, err := CloudAndClient()
	if err != nil {
		return nil, err
	}

	return openstack.NewCTSV1(cc.ProviderClient, golangsdk.EndpointOpts{
		Region: cc.RegionName,
	})
}

func NewCTSV2Client() (*golangsdk.ServiceClient, error) {
	cc, err := CloudAndClient()
	if err != nil {
		return nil, err
	}

	return openstack.NewCTSV2(cc.ProviderClient, golangsdk.EndpointOpts{
		Region: cc.RegionName,
	})
}

func NewCTSV3Client() (*golangsdk.ServiceClient, error) {
	cc, err := CloudAndClient()
	if err != nil {
		return nil, err
	}

	return openstack.NewCTSV3(cc.ProviderClient, golangsdk.EndpointOpts{
		Region: cc.RegionName,
	})
}

// NewDCaaSV2Client returns a *ServiceClient for making calls
// to the OpenStack v2 API. An error will be returned
// if authentication or client creation was not possible.
func NewDCaaSV2Client() (*golangsdk.ServiceClient, error) {
	cc, err := CloudAndClient()
	if err != nil {
		return nil, err
	}

	return openstack.NewDCaaSV2(cc.ProviderClient, golangsdk.EndpointOpts{
		Region: cc.RegionName,
	})
}

// NewDNSV2Client returns a *ServiceClient for making calls
// to the OpenStack Compute v2 API. An error will be returned
// if authentication or client creation was not possible.
func NewDNSV2Client() (*golangsdk.ServiceClient, error) {
	cc, err := CloudAndClient()
	if err != nil {
		return nil, err
	}

	return openstack.NewDNSV2(cc.ProviderClient, golangsdk.EndpointOpts{
		Region: cc.RegionName,
	})
}

func NewDWSV1Client() (*golangsdk.ServiceClient, error) {
	cc, err := CloudAndClient()
	if err != nil {
		return nil, err
	}

	return openstack.NewDWSV1(cc.ProviderClient, golangsdk.EndpointOpts{
		Region: cc.RegionName,
	})
}

// NewIdentityV3Client returns a *ServiceClient for making calls
// to the OpenStack Identity v3 API on a `project` level. An error will be returned
// if authentication or client creation was not possible.
func NewIdentityV3Client() (*golangsdk.ServiceClient, error) {
	cc, err := CloudAndClient()
	if err != nil {
		return nil, err
	}

	return openstack.NewIdentityV3(cc.ProviderClient, golangsdk.EndpointOpts{
		Region: cc.RegionName,
	})
}

// NewIdentityV30AdminClient returns a *ServiceClient for making calls
// to the OpenStack Identity v3 API on a `domain` level with v3.0 IAM endpoint.
// An error will be returned if authentication or client creation was not possible.
func NewIdentityV30AdminClient() (*golangsdk.ServiceClient, error) {
	cloud, err := EnvOS.Cloud()
	if err != nil {
		return nil, fmt.Errorf("error constructing cloud configuration: %w", err)
	}

	opts := golangsdk.AuthOptions{
		IdentityEndpoint: cloud.AuthInfo.AuthURL,
		Username:         cloud.AuthInfo.Username,
		Password:         cloud.AuthInfo.Password,
		DomainName:       cloud.AuthInfo.UserDomainName,
		TenantID:         cloud.AuthInfo.ProjectID,
	}

	pClient, err := openstack.AuthenticatedClient(opts)
	if err != nil {
		return nil, fmt.Errorf("error creating provider client: %w", err)
	}
	client, err := openstack.NewIdentityV3(pClient, golangsdk.EndpointOpts{})

	if err != nil {
		return nil, err
	}

	client.Endpoint = strings.Replace(client.Endpoint, "v3", "v3.0", 1)
	return client, err
}

// NewIdentityV3AdminClient returns a *ServiceClient for making calls
// to the OpenStack Identity v3 API on a `domain` level with v3 IAM endpoint.
// An error will be returned if authentication or client creation was not possible.
func NewIdentityV3AdminClient() (*golangsdk.ServiceClient, error) {
	cloud, err := EnvOS.Cloud()
	if err != nil {
		return nil, fmt.Errorf("error constructing cloud configuration: %w", err)
	}

	opts := golangsdk.AuthOptions{
		IdentityEndpoint: cloud.AuthInfo.AuthURL,
		Username:         cloud.AuthInfo.Username,
		Password:         cloud.AuthInfo.Password,
		DomainName:       cloud.AuthInfo.UserDomainName,
		TenantID:         cloud.AuthInfo.ProjectID,
	}

	pClient, err := openstack.AuthenticatedClient(opts)
	if err != nil {
		return nil, fmt.Errorf("error creating provider client: %w", err)
	}
	client, err := openstack.NewIdentityV3(pClient, golangsdk.EndpointOpts{})

	if err != nil {
		return nil, err
	}

	return client, err
}

// NewIdentityV3UnauthenticatedClient returns an unauthenticated *ServiceClient
// for the OpenStack Identity v3 API. An error  will be returned if
// authentication or client creation was not possible.
func NewIdentityV3UnauthenticatedClient() (*golangsdk.ServiceClient, error) {
	cloud, err := EnvOS.Cloud()
	if err != nil {
		return nil, err
	}

	client, err := openstack.NewClient(cloud.AuthInfo.AuthURL)
	if err != nil {
		return nil, err
	}

	return openstack.NewIdentityV3(client, golangsdk.EndpointOpts{})
}

func NewIMSV1Client() (*golangsdk.ServiceClient, error) {
	cc, err := CloudAndClient()
	if err != nil {
		return nil, err
	}

	return openstack.NewIMSV1(cc.ProviderClient, golangsdk.EndpointOpts{})
}

func NewIMSV2Client() (*golangsdk.ServiceClient, error) {
	cc, err := CloudAndClient()
	if err != nil {
		return nil, err
	}

	return openstack.NewIMSV2(cc.ProviderClient, golangsdk.EndpointOpts{})
}

// NewNetworkV1Client returns a *ServiceClient for making calls to the
// OpenStack Networking v1 API. An error will be returned if authentication
// or client creation was not possible.
func NewNetworkV1Client() (*golangsdk.ServiceClient, error) {
	cc, err := CloudAndClient()
	if err != nil {
		return nil, err
	}

	return openstack.NewNetworkV1(cc.ProviderClient, golangsdk.EndpointOpts{
		Region: cc.RegionName,
	})
}

func NewVPCEndpointV1Client() (*golangsdk.ServiceClient, error) {
	cc, err := CloudAndClient()
	if err != nil {
		return nil, err
	}

	return openstack.NewVpcEpV1(cc.ProviderClient, golangsdk.EndpointOpts{
		Region: cc.RegionName,
	})
}

// NewPeerNetworkV1Client returns a *ServiceClient for making calls to the
// OpenStack Networking v1 API for VPC peer. An error will be returned if authentication
// or client creation was not possible.
func NewPeerNetworkV1Client() (*golangsdk.ServiceClient, error) {
	cc, err := CloudAndClient()
	if err != nil {
		return nil, err
	}

	err = UpdatePeerTenantDetails(cc.Cloud)
	if err != nil {
		return nil, err
	}

	return openstack.NewNetworkV1(cc.ProviderClient, golangsdk.EndpointOpts{
		Region: cc.RegionName,
	})
}

// NewNetworkV2Client returns a *ServiceClient for making calls to the
// OpenStack Networking v2 API. An error will be returned if authentication
// or client creation was not possible.
func NewNetworkV2Client() (*golangsdk.ServiceClient, error) {
	cc, err := CloudAndClient()
	if err != nil {
		return nil, err
	}

	return openstack.NewNetworkV2(cc.ProviderClient, golangsdk.EndpointOpts{
		Region: cc.RegionName,
	})
}

// NewElbV2Client returns authenticated ELB v2 client
func NewElbV2Client() (*golangsdk.ServiceClient, error) {
	cc, err := CloudAndClient()
	if err != nil {
		return nil, err
	}

	return openstack.NewELBV2(cc.ProviderClient, golangsdk.EndpointOpts{
		Region: cc.RegionName,
	})
}

// NewElbV3Client returns authenticated ELB v3 client
func NewElbV3Client() (*golangsdk.ServiceClient, error) {
	cc, err := CloudAndClient()
	if err != nil {
		return nil, err
	}

	return openstack.NewELBV3(cc.ProviderClient, golangsdk.EndpointOpts{
		Region: cc.RegionName,
	})
}

// NewNatV2Client returns authenticated NAT v2 client
func NewNatV2Client() (*golangsdk.ServiceClient, error) {
	cc, err := CloudAndClient()
	if err != nil {
		return nil, err
	}

	return openstack.NewNatV2(cc.ProviderClient, golangsdk.EndpointOpts{
		Region: cc.RegionName,
	})
}

// NewPeerNetworkV2Client returns a *ServiceClient for making calls to the
// OpenStack Networking v2 API for Peer. An error will be returned if authentication
// or client creation was not possible.
func NewPeerNetworkV2Client() (*golangsdk.ServiceClient, error) {
	cc, err := CloudAndClient()
	if err != nil {
		return nil, err
	}
	err = UpdatePeerTenantDetails(cc.Cloud)
	if err != nil {
		return nil, err
	}

	return openstack.NewNetworkV2(cc.ProviderClient, golangsdk.EndpointOpts{
		Region: cc.RegionName,
	})
}

// NewObjectStorageV1Client returns a *ServiceClient for making calls to the
// OpenStack Object Storage v1 API. An error will be returned if authentication
// or client creation was not possible.
func NewObjectStorageV1Client() (*golangsdk.ServiceClient, error) {
	cc, err := CloudAndClient()
	if err != nil {
		return nil, err
	}

	return openstack.NewObjectStorageV1(cc.ProviderClient, golangsdk.EndpointOpts{
		Region: cc.RegionName,
	})
}

func NewOBSClient() (*obs.ObsClient, error) {
	cc, err := CloudAndClient()
	if err != nil {
		return nil, err
	}

	if err := setupTemporaryAKSK(cc); err != nil {
		return nil, fmt.Errorf("failed to construct OBS client without AK/SK: %s", err)
	}

	client, err := openstack.NewOBSService(cc.ProviderClient, golangsdk.EndpointOpts{
		Region: cc.RegionName,
	})
	if err != nil {
		return nil, err
	}
	opts := cc.AKSKAuthOptions
	return obs.New(
		opts.AccessKey, opts.SecretKey, client.Endpoint,
		obs.WithSecurityToken(opts.SecurityToken), obs.WithSignature(obs.SignatureObs),
	)
}

func NewOBSClientWithoutHeader() (*obs.ObsClient, error) {
	cc, err := CloudAndClient()
	if err != nil {
		return nil, err
	}

	if err := setupTemporaryAKSK(cc); err != nil {
		return nil, fmt.Errorf("failed to construct OBS client without AK/SK: %s", err)
	}

	client, err := openstack.NewOBSService(cc.ProviderClient, golangsdk.EndpointOpts{
		Region: cc.RegionName,
	})
	if err != nil {
		return nil, err
	}
	opts := cc.AKSKAuthOptions
	return obs.New(
		opts.AccessKey, opts.SecretKey, client.Endpoint,
		obs.WithSecurityToken(opts.SecurityToken),
	)
}

// NewSharedFileSystemV2Client returns a *ServiceClient for making calls
// to the OpenStack Shared File System v2 API. An error will be returned
// if authentication or client creation was not possible.
func NewSharedFileSystemV2Client() (*golangsdk.ServiceClient, error) {
	cc, err := CloudAndClient()
	if err != nil {
		return nil, err
	}
	return openstack.NewSharedFileSystemV2(cc.ProviderClient, golangsdk.EndpointOpts{
		Region: cc.RegionName,
	})
}

func NewKMSV1Client() (*golangsdk.ServiceClient, error) {
	cc, err := CloudAndClient()
	if err != nil {
		return nil, err
	}
	return openstack.NewKMSV1(cc.ProviderClient, golangsdk.EndpointOpts{
		Region: cc.RegionName,
	})
}

// NewSharedFileSystemTurboV1Client returns a *ServiceClient for making calls
// to the OpenStack Shared File System Turbo v1 API. An error will be returned
// if authentication or client creation was not possible.
func NewSharedFileSystemTurboV1Client() (*golangsdk.ServiceClient, error) {
	cc, err := CloudAndClient()
	if err != nil {
		return nil, err
	}
	return openstack.NewSharedFileSystemTurboV1(cc.ProviderClient, golangsdk.EndpointOpts{
		Region: cc.RegionName,
	})
}

// NewRdsV3 returns authenticated RDS v3 client
func NewRdsV3() (*golangsdk.ServiceClient, error) {
	cc, err := CloudAndClient()
	if err != nil {
		return nil, err
	}
	return openstack.NewRDSV3(cc.ProviderClient, golangsdk.EndpointOpts{
		Region: cc.RegionName,
	})
}

// NewMrsV1 returns authenticated MRS v1 client
func NewMrsV1() (*golangsdk.ServiceClient, error) {
	cc, err := CloudAndClient()
	if err != nil {
		return nil, err
	}
	return openstack.NewMapReduceV1(cc.ProviderClient, golangsdk.EndpointOpts{
		Region: cc.RegionName,
	})
}

// NewSDRSV1 returns authenticated SDRS v3 client
func NewSDRSV1() (*golangsdk.ServiceClient, error) {
	cc, err := CloudAndClient()
	if err != nil {
		return nil, err
	}
	return openstack.NewSDRSV1(cc.ProviderClient, golangsdk.EndpointOpts{
		Region: cc.RegionName,
	})
}

// NewWafV1Client returns authenticated WAF v1 client
func NewWafV1Client() (*golangsdk.ServiceClient, error) {
	cc, err := CloudAndClient()
	if err != nil {
		return nil, err
	}
	return openstack.NewWAFV1(cc.ProviderClient, golangsdk.EndpointOpts{Region: cc.RegionName})
}

// NewWafdV1Client returns authenticated WAF premium v1 client
func NewWafdV1Client() (*golangsdk.ServiceClient, error) {
	cc, err := CloudAndClient()
	if err != nil {
		return nil, err
	}
	if cc.RegionName != "eu-ch2" {
		return openstack.NewWAFDV1(cc.ProviderClient, golangsdk.EndpointOpts{Region: cc.RegionName})
	} else {
		return openstack.NewWAFDSwissV1(cc.ProviderClient, golangsdk.EndpointOpts{Region: cc.RegionName})
	}
}

// NewCsbsV1Client returns authenticated CSBS v1 client
func NewCsbsV1Client() (*golangsdk.ServiceClient, error) {
	cc, err := CloudAndClient()
	if err != nil {
		return nil, err
	}
	return openstack.NewCSBSService(cc.ProviderClient, golangsdk.EndpointOpts{
		Region: cc.RegionName,
	})
}

// NewCssV1Client returns authenticated CSS v1 client
func NewCssV1Client() (*golangsdk.ServiceClient, error) {
	cc, err := CloudAndClient()
	if err != nil {
		return nil, err
	}
	return openstack.NewCSSService(cc.ProviderClient, golangsdk.EndpointOpts{
		Region: cc.RegionName,
	})
}

func NewCceV1Client() (*golangsdk.ServiceClient, error) {
	cc, err := CloudAndClient()
	if err != nil {
		return nil, err
	}
	return openstack.NewCCEv1(cc.ProviderClient, golangsdk.EndpointOpts{
		Region: cc.RegionName,
	})
}

func NewCceV3Client() (*golangsdk.ServiceClient, error) {
	cc, err := CloudAndClient()
	if err != nil {
		return nil, err
	}
	return openstack.NewCCE(cc.ProviderClient, golangsdk.EndpointOpts{Region: cc.RegionName})
}

func NewCceV3AddonClient() (*golangsdk.ServiceClient, error) {
	cc, err := CloudAndClient()
	if err != nil {
		return nil, err
	}
	client, err := openstack.NewCCE(cc.ProviderClient, golangsdk.EndpointOpts{Region: cc.RegionName})
	if err != nil {
		return nil, err
	}
	client.ResourceBase = fmt.Sprintf("%sapi/v3/", client.Endpoint)
	return client, nil
}

func NewCbrV3Client() (*golangsdk.ServiceClient, error) {
	cc, err := CloudAndClient()
	if err != nil {
		return nil, err
	}
	return openstack.NewCBRService(cc.ProviderClient, golangsdk.EndpointOpts{
		Region: cc.RegionName,
	})
}

func NewVbsV2Client() (*golangsdk.ServiceClient, error) {
	cc, err := CloudAndClient()
	if err != nil {
		return nil, err
	}
	return openstack.NewVBSServiceV2(cc.ProviderClient, golangsdk.EndpointOpts{
		Region: cc.RegionName,
	})
}

// NewDataArtsV11Client returns authenticated DataArts v1.1 client
func NewDataArtsV11Client() (*golangsdk.ServiceClient, error) {
	cc, err := CloudAndClient()
	if err != nil {
		return nil, err
	}
	return openstack.NewDataArtsV11(cc.ProviderClient, golangsdk.EndpointOpts{
		Region: cc.RegionName,
	})
}

// NewDdsV3Client returns authenticated DDS v3 client
func NewDdsV3Client() (*golangsdk.ServiceClient, error) {
	cc, err := CloudAndClient()
	if err != nil {
		return nil, err
	}
	return openstack.NewDDSServiceV3(cc.ProviderClient, golangsdk.EndpointOpts{
		Region: cc.RegionName,
	})
}

// NewDrsV3Client returns authenticated DDS v3 client
func NewDrsV3Client() (*golangsdk.ServiceClient, error) {
	cc, err := CloudAndClient()
	if err != nil {
		return nil, err
	}
	return openstack.NewDRSServiceV3(cc.ProviderClient, golangsdk.EndpointOpts{
		Region: cc.RegionName,
	})
}

// NewDcsV1Client returns authenticated DCS v1 client
func NewDcsV1Client() (*golangsdk.ServiceClient, error) {
	cc, err := CloudAndClient()
	if err != nil {
		return nil, err
	}
	return openstack.NewDCSServiceV1(cc.ProviderClient, golangsdk.EndpointOpts{
		Region: cc.RegionName,
	})
}

// NewDmsV1Client returns authenticated DMS v1 client
func NewDmsV1Client() (*golangsdk.ServiceClient, error) {
	cc, err := CloudAndClient()
	if err != nil {
		return nil, err
	}
	return openstack.NewDMSServiceV1(cc.ProviderClient, golangsdk.EndpointOpts{
		Region: cc.RegionName,
	})
}

func NewDisV2Client() (*golangsdk.ServiceClient, error) {
	cc, err := CloudAndClient()
	if err != nil {
		return nil, err
	}
	return openstack.NewDISServiceV2(cc.ProviderClient, golangsdk.EndpointOpts{
		Region: cc.RegionName,
	})
}

func NewDmsV11Client() (*golangsdk.ServiceClient, error) {
	cc, err := CloudAndClient()
	if err != nil {
		return nil, err
	}
	return openstack.NewDMSServiceV11(cc.ProviderClient, golangsdk.EndpointOpts{
		Region: cc.RegionName,
	})
}

// NewDmsV2Client returns authenticated DMS v2 client
func NewDmsV2Client() (*golangsdk.ServiceClient, error) {
	cc, err := CloudAndClient()
	if err != nil {
		return nil, err
	}
	return openstack.NewDMSServiceV2(cc.ProviderClient, golangsdk.EndpointOpts{
		Region: cc.RegionName,
	})
}

// NewSwrV2Client returns authenticated SWR v2 client
func NewSwrV2Client() (client *golangsdk.ServiceClient, err error) {
	cc, err := CloudAndClient()
	if err != nil {
		return nil, err
	}
	return openstack.NewSWRV2(cc.ProviderClient, golangsdk.EndpointOpts{
		Region: cc.RegionName,
	})
}

// NewSmnV2Client returns authenticated SMN v2 client
func NewSmnV2Client() (client *golangsdk.ServiceClient, err error) {
	cc, err := CloudAndClient()
	if err != nil {
		return nil, err
	}
	return openstack.NewSMNV2(cc.ProviderClient, golangsdk.EndpointOpts{
		Region: cc.RegionName,
	})
}

// NewTmsV1Client returns authenticated TMS v1.0 client
func NewTmsV1Client() (client *golangsdk.ServiceClient, err error) {
	iamClient, err := NewIdentityV3AdminClient()
	if err != nil {
		return nil, err
	}

	iamClient.Endpoint = strings.Replace(iamClient.Endpoint, "v3", "v1.0", 1)
	iamClient.Endpoint = strings.Replace(iamClient.Endpoint, "iam", "tms", 1)
	return iamClient, err
}

// NewCesV1Client returns authenticated CES v1 client
func NewCesV1Client() (client *golangsdk.ServiceClient, err error) {
	cc, err := CloudAndClient()
	if err != nil {
		return nil, err
	}
	return openstack.NewCESClient(cc.ProviderClient, golangsdk.EndpointOpts{
		Region: cc.RegionName,
	})
}

// NewLtsV2Client returns authenticated LTS v2 client
func NewLtsV2Client() (client *golangsdk.ServiceClient, err error) {
	cc, err := CloudAndClient()
	if err != nil {
		return nil, err
	}
	return openstack.NewLTSV2(cc.ProviderClient, golangsdk.EndpointOpts{
		Region: cc.RegionName,
	})
}

func NewGaussDBClient() (client *golangsdk.ServiceClient, err error) {
	cc, err := CloudAndClient()
	if err != nil {
		return nil, err
	}

	return openstack.NewGaussDBV3(cc.ProviderClient, golangsdk.EndpointOpts{})
}

func UpdatePeerTenantDetails(cloud *openstack.Cloud) error {
	if id := EnvOS.GetEnv("Peer_Tenant_ID"); id != "" {
		cloud.AuthInfo.ProjectID = id
		cloud.AuthInfo.ProjectName = ""
		return nil

	}
	if name := EnvOS.GetEnv("Peer_Tenant_Name"); name != "" {
		cloud.AuthInfo.ProjectID = ""
		cloud.AuthInfo.ProjectName = name
		return nil
	}
	return fmt.Errorf("you're missing some important setup:\n OS_Peer_Tenant_ID or OS_Peer_Tenant_Name env variables must be provided")
}

// copyCloud makes a deep copy of cloud
func copyCloud(src *openstack.Cloud) (*openstack.Cloud, error) {
	srcJson, err := json.Marshal(src)
	if err != nil {
		return nil, fmt.Errorf("error marshalling cloud: %s", err)
	}
	res := new(openstack.Cloud)
	if err := json.Unmarshal(srcJson, res); err != nil {
		return nil, fmt.Errorf("error unmarshalling cloud: %s", err)
	}
	return res, nil
}

// cc stands for `cloud` & `client`
type cc struct {
	*openstack.Cloud
	*golangsdk.ProviderClient
}

// CloudAndClient returns copy of cloud configuration and authenticated client for OS_ environment
func CloudAndClient() (*cc, error) {
	cloud, err := EnvOS.Cloud()
	if err != nil {
		return nil, fmt.Errorf("error constructing cloud configuration: %w", err)
	}
	cloud, err = copyCloud(cloud)
	if err != nil {
		return nil, fmt.Errorf("error copying cloud: %w", err)
	}
	client, err := EnvOS.AuthenticatedClient()
	if err != nil {
		return nil, err
	}
	return &cc{cloud, client}, nil
}

func setupTemporaryAKSK(config *cc) error {
	if config.AKSKAuthOptions.AccessKey != "" {
		return nil
	}

	client, err := NewIdentityV3Client()
	if err != nil {
		return fmt.Errorf("error creating identity v3 domain client: %s", err)
	}
	credential, err := credentials.CreateTemporary(client, credentials.CreateTemporaryOpts{
		Methods: []string{"token"},
		Token:   client.Token(),
	}).Extract()
	if err != nil {
		return fmt.Errorf("error creating temporary AK/SK: %s", err)
	}

	config.AKSKAuthOptions.AccessKey = credential.AccessKey
	config.AKSKAuthOptions.SecretKey = credential.SecretKey
	config.AKSKAuthOptions.SecurityToken = credential.SecurityToken
	return nil
}
