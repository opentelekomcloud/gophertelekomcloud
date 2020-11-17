// Package clients contains functions for creating OpenStack service clients
// for use in acceptance tests. It also manages the required environment
// variables to run the tests.
package clients

import (
	"fmt"
	"os"
	"strings"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/utils"
)

var (
	OS_FLAVOR_ID         = os.Getenv("OS_FLAVOR_ID")
	OS_FLAVOR_NAME       = os.Getenv("OS_FLAVOR_NAME")
	OS_IMAGE_ID          = os.Getenv("OS_IMAGE_ID")
	OS_IMAGE_NAME        = os.Getenv("OS_IMAGE_NAME")
	OS_NETWORK_ID        = os.Getenv("OS_NETWORK_ID")
	OS_POOL_NAME         = os.Getenv("OS_POOL_NAME")
	OS_REGION_NAME       = os.Getenv("OS_REGION_NAME")
	OS_ACCESS_KEY        = os.Getenv("OS_ACCESS_KEY")
	OS_SECRET_KEY        = os.Getenv("OS_SECRET_KEY")
	OS_AVAILABILITY_ZONE = os.Getenv("OS_AVAILABILITY_ZONE")
	OS_VPC_ID            = os.Getenv("OS_VPC_ID")
	OS_SUBNET_ID         = os.Getenv("OS_SUBNET_ID")
	OS_TENANT_ID         = os.Getenv("OS_TENANT_ID")
	OS_KEYPAIR_NAME      = os.Getenv("OS_KEYPAIR_NAME")
	OS_TENANT_NAME       = getTenantName()
	OS_USER_ID           = os.Getenv("OS_USER_ID")

	osEnv = openstack.NewEnv("OS_")
)

func getTenantName() string {
	tn := os.Getenv("OS_TENANT_NAME")
	if tn == "" {
		tn = os.Getenv("OS_PROJECT_NAME")
	}
	return tn
}

// AcceptanceTestChoices contains image and flavor selections for use by the acceptance tests.
type AcceptanceTestChoices struct {
	// ImageID contains the ID of a valid image.
	ImageID string

	// FlavorID contains the ID of a valid flavor.
	FlavorID string

	// FlavorIDResize contains the ID of a different flavor available on the same OpenStack installation, that is distinct
	// from FlavorID.
	FlavorIDResize string

	// FloatingIPPool contains the name of the pool from where to obtain floating IPs.
	FloatingIPPoolName string

	// NetworkName is the name of a network to launch the instance on.
	NetworkName string

	// ExternalNetworkID is the network ID of the external network.
	ExternalNetworkID string

	// ShareNetworkID is the Manila Share network ID
	ShareNetworkID string

	// DBDatastoreType is the datastore type for DB tests.
	DBDatastoreType string

	// DBDatastoreTypeID is the datastore type version for DB tests.
	DBDatastoreVersion string
}

// AcceptanceTestChoicesFromEnv populates a ComputeChoices struct from environment variables.
// If any required state is missing, an `error` will be returned that enumerates the missing properties.
func AcceptanceTestChoicesFromEnv() (*AcceptanceTestChoices, error) {
	imageID := os.Getenv("OS_IMAGE_ID")
	flavorID := os.Getenv("OS_FLAVOR_ID")
	flavorIDResize := os.Getenv("OS_FLAVOR_ID_RESIZE")
	networkName := os.Getenv("OS_NETWORK_NAME")
	floatingIPPoolName := os.Getenv("OS_POOL_NAME")
	externalNetworkID := os.Getenv("OS_EXTGW_ID")
	shareNetworkID := os.Getenv("OS_SHARE_NETWORK_ID")
	dbDatastoreType := os.Getenv("OS_DB_DATASTORE_TYPE")
	dbDatastoreVersion := os.Getenv("OS_DB_DATASTORE_VERSION")

	missing := make([]string, 0, 3)
	if imageID == "" {
		missing = append(missing, "OS_IMAGE_ID")
	}
	if flavorID == "" {
		missing = append(missing, "OS_FLAVOR_ID")
	}
	if flavorIDResize == "" {
		missing = append(missing, "OS_FLAVOR_ID_RESIZE")
	}
	if floatingIPPoolName == "" {
		missing = append(missing, "OS_POOL_NAME")
	}
	if externalNetworkID == "" {
		missing = append(missing, "OS_EXTGW_ID")
	}
	if networkName == "" {
		networkName = "private"
	}
	if shareNetworkID == "" {
		missing = append(missing, "OS_SHARE_NETWORK_ID")
	}
	notDistinct := ""
	if flavorID == flavorIDResize {
		notDistinct = "OS_FLAVOR_ID and OS_FLAVOR_ID_RESIZE must be distinct."
	}

	if len(missing) > 0 || notDistinct != "" {
		text := "You're missing some important setup:\n"
		if len(missing) > 0 {
			text += " * These environment variables must be provided: " + strings.Join(missing, ", ") + "\n"
		}
		if notDistinct != "" {
			text += " * " + notDistinct + "\n"
		}

		return nil, fmt.Errorf(text)
	}

	return &AcceptanceTestChoices{
		ImageID:            imageID,
		FlavorID:           flavorID,
		FlavorIDResize:     flavorIDResize,
		FloatingIPPoolName: floatingIPPoolName,
		NetworkName:        networkName,
		ExternalNetworkID:  externalNetworkID,
		ShareNetworkID:     shareNetworkID,
		DBDatastoreType:    dbDatastoreType,
		DBDatastoreVersion: dbDatastoreVersion,
	}, nil
}

// NewBlockStorageV1Client returns a *ServiceClient for making calls
// to the OpenStack Block Storage v1 API. An error will be returned
// if authentication or client creation was not possible.
func NewBlockStorageV1Client() (*golangsdk.ServiceClient, error) {
	ao, err := openstack.AuthOptionsFromEnv()
	if err != nil {
		return nil, err
	}

	client, err := openstack.AuthenticatedClient(ao)
	if err != nil {
		return nil, err
	}

	return openstack.NewBlockStorageV1(client, golangsdk.EndpointOpts{
		Region: utils.GetRegion(ao),
	})
}

// NewBlockStorageV2Client returns a *ServiceClient for making calls
// to the OpenStack Block Storage v2 API. An error will be returned
// if authentication or client creation was not possible.
func NewBlockStorageV2Client() (*golangsdk.ServiceClient, error) {
	ao, err := openstack.AuthOptionsFromEnv()
	if err != nil {
		return nil, err
	}

	client, err := openstack.AuthenticatedClient(ao)
	if err != nil {
		return nil, err
	}

	return openstack.NewBlockStorageV2(client, golangsdk.EndpointOpts{
		Region: utils.GetRegion(ao)})
}

// NewBlockStorageV3Client returns a *ServiceClient for making calls
// to the OpenStack Block Storage v3 API. An error will be returned
// if authentication or client creation was not possible.
func NewBlockStorageV3Client() (*golangsdk.ServiceClient, error) {
	ao, err := openstack.AuthOptionsFromEnv()
	if err != nil {
		return nil, err
	}

	client, err := openstack.AuthenticatedClient(ao)
	if err != nil {
		return nil, err
	}

	return openstack.NewBlockStorageV3(client, golangsdk.EndpointOpts{
		Region: utils.GetRegion(ao),
	})
}

// NewComputeV2Client returns a *ServiceClient for making calls
// to the OpenStack Compute v2 API. An error will be returned
// if authentication or client creation was not possible.
func NewComputeV2Client() (*golangsdk.ServiceClient, error) {
	ao, err := openstack.AuthOptionsFromEnv()
	if err != nil {
		return nil, err
	}

	client, err := openstack.AuthenticatedClient(ao)
	if err != nil {
		return nil, err
	}

	return openstack.NewComputeV2(client, golangsdk.EndpointOpts{
		Region: utils.GetRegion(ao),
	})
}

// NewDNSV2Client returns a *ServiceClient for making calls
// to the OpenStack Compute v2 API. An error will be returned
// if authentication or client creation was not possible.
func NewDNSV2Client() (*golangsdk.ServiceClient, error) {
	ao, err := openstack.AuthOptionsFromEnv()
	if err != nil {
		return nil, err
	}

	client, err := openstack.AuthenticatedClient(ao)
	if err != nil {
		return nil, err
	}

	return openstack.NewDNSV2(client, golangsdk.EndpointOpts{
		Region: utils.GetRegion(ao),
	})
}

// NewIdentityV3Client returns a *ServiceClient for making calls
// to the OpenStack Identity v3 API. An error will be returned
// if authentication or client creation was not possible.
func NewIdentityV3Client() (*golangsdk.ServiceClient, error) {
	ao, err := openstack.AuthOptionsFromEnv()
	if err != nil {
		return nil, err
	}

	client, err := openstack.AuthenticatedClient(ao)
	if err != nil {
		return nil, err
	}

	return openstack.NewIdentityV3(client, golangsdk.EndpointOpts{
		Region: utils.GetRegion(ao),
	})
}

// NewIdentityV3UnauthenticatedClient returns an unauthenticated *ServiceClient
// for the OpenStack Identity v3 API. An error  will be returned if
// authentication or client creation was not possible.
func NewIdentityV3UnauthenticatedClient() (*golangsdk.ServiceClient, error) {
	ao, err := openstack.AuthOptionsFromEnv()
	if err != nil {
		return nil, err
	}

	client, err := openstack.NewClient(ao.IdentityEndpoint)
	if err != nil {
		return nil, err
	}

	return openstack.NewIdentityV3(client, golangsdk.EndpointOpts{})
}

// NewImageServiceV2Client returns a *ServiceClient for making calls to the
// OpenStack Image v2 API. An error will be returned if authentication or
// client creation was not possible.
func NewImageServiceV2Client() (*golangsdk.ServiceClient, error) {
	ao, err := openstack.AuthOptionsFromEnv()
	if err != nil {
		return nil, err
	}

	client, err := openstack.AuthenticatedClient(ao)
	if err != nil {
		return nil, err
	}

	return openstack.NewImageServiceV2(client, golangsdk.EndpointOpts{
		Region: utils.GetRegion(ao),
	})
}

// NewNetworkV1Client returns a *ServiceClient for making calls to the
// OpenStack Networking v1 API. An error will be returned if authentication
// or client creation was not possible.
func NewNetworkV1Client() (*golangsdk.ServiceClient, error) {
	ao, err := openstack.AuthOptionsFromEnv()
	if err != nil {
		return nil, err
	}

	client, err := openstack.AuthenticatedClient(ao)
	if err != nil {
		return nil, err
	}

	return openstack.NewNetworkV1(client, golangsdk.EndpointOpts{
		Region: utils.GetRegion(ao),
	})
}

// NewPeerNetworkV1Client returns a *ServiceClient for making calls to the
// OpenStack Networking v1 API for VPC peer. An error will be returned if authentication
// or client creation was not possible.
func NewPeerNetworkV1Client() (*golangsdk.ServiceClient, error) {
	ao, err := openstack.AuthOptionsFromEnv()
	if err != nil {
		return nil, err
	}

	err = UpdatePeerTenantDetails(&ao)
	if err != nil {
		return nil, err
	}

	client, err := openstack.AuthenticatedClient(ao)
	if err != nil {
		return nil, err
	}

	return openstack.NewNetworkV1(client, golangsdk.EndpointOpts{
		Region: utils.GetRegion(ao),
	})
}

// NewNetworkV2Client returns a *ServiceClient for making calls to the
// OpenStack Networking v2 API. An error will be returned if authentication
// or client creation was not possible.
func NewNetworkV2Client() (*golangsdk.ServiceClient, error) {
	cloud, err := osEnv.Cloud()
	if err != nil {
		return nil, err
	}
	client, err := osEnv.AuthenticatedClient()
	if err != nil {
		return nil, err
	}

	return openstack.NewNetworkV2(client, golangsdk.EndpointOpts{Region: cloud.RegionName})
}

// NewPeerNetworkV2Client returns a *ServiceClient for making calls to the
// OpenStack Networking v2 API for Peer. An error will be returned if authentication
// or client creation was not possible.
func NewPeerNetworkV2Client() (*golangsdk.ServiceClient, error) {
	ao, err := openstack.AuthOptionsFromEnv()
	if err != nil {
		return nil, err
	}

	err = UpdatePeerTenantDetails(&ao)
	if err != nil {
		return nil, err
	}

	client, err := openstack.AuthenticatedClient(ao)
	if err != nil {
		return nil, err
	}

	return openstack.NewNetworkV2(client, golangsdk.EndpointOpts{
		Region: utils.GetRegion(ao),
	})
}

// NewObjectStorageV1Client returns a *ServiceClient for making calls to the
// OpenStack Object Storage v1 API. An error will be returned if authentication
// or client creation was not possible.
func NewObjectStorageV1Client() (*golangsdk.ServiceClient, error) {
	ao, err := openstack.AuthOptionsFromEnv()
	if err != nil {
		return nil, err
	}

	client, err := openstack.AuthenticatedClient(ao)
	if err != nil {
		return nil, err
	}

	return openstack.NewObjectStorageV1(client, golangsdk.EndpointOpts{
		Region: utils.GetRegion(ao),
	})
}

// NewSharedFileSystemV2Client returns a *ServiceClient for making calls
// to the OpenStack Shared File System v2 API. An error will be returned
// if authentication or client creation was not possible.
func NewSharedFileSystemV2Client() (*golangsdk.ServiceClient, error) {
	ao, err := openstack.AuthOptionsFromEnv()
	if err != nil {
		return nil, err
	}

	client, err := openstack.AuthenticatedClient(ao)
	if err != nil {
		return nil, err
	}

	return openstack.NewSharedFileSystemV2(client, golangsdk.EndpointOpts{
		Region: utils.GetRegion(ao),
	})
}

// NewRdsV3 returns authenticated RDS v3 client
func NewRdsV3() (*golangsdk.ServiceClient, error) {
	ao, err := openstack.AuthOptionsFromEnv()
	if err != nil {
		return nil, err
	}
	client, err := openstack.AuthenticatedClient(ao)
	if err != nil {
		return nil, err
	}

	return openstack.NewRDSV3(client, golangsdk.EndpointOpts{
		Region: utils.GetRegion(ao),
	})
}

// NewWafV1Client returns authenticated WAF v1 client
func NewWafV1Client() (*golangsdk.ServiceClient, error) {
	cloud, err := osEnv.Cloud()
	if err != nil {
		return nil, err
	}
	client, err := osEnv.AuthenticatedClient()
	if err != nil {
		return nil, err
	}
	return openstack.NewWAFV1(client, golangsdk.EndpointOpts{Region: cloud.RegionName})
}

// NewCsbsV1Client returns authenticated CSBS v1 client
func NewCsbsV1Client() (*golangsdk.ServiceClient, error) {
	cloud, err := osEnv.Cloud()
	if err != nil {
		return nil, err
	}
	client, err := osEnv.AuthenticatedClient()
	if err != nil {
		return nil, err
	}
	return openstack.NewCSBSService(client, golangsdk.EndpointOpts{Region: cloud.RegionName})
}

func UpdatePeerTenantDetails(ao *golangsdk.AuthOptions) error {

	if peerTenantID := os.Getenv("OS_Peer_Tenant_ID"); peerTenantID != "" {
		ao.TenantID = peerTenantID
		ao.TenantName = ""
		return nil

	} else if peerTenantName := os.Getenv("OS_Peer_Tenant_Name"); peerTenantName != "" {
		ao.TenantID = ""
		ao.TenantName = peerTenantName
		return nil

	} else {
		return fmt.Errorf("You're missing some important setup:\n OS_Peer_Tenant_ID or OS_Peer_Tenant_Name env variables must be provided.")
	}
}
