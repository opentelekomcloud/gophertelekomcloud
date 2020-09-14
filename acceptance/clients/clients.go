// Package clients contains functions for creating OpenStack service clients
// for use in acceptance tests. It also manages the required environment
// variables to run the tests.
package clients

import (
	"fmt"
	"os"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack"
)

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
		Region: os.Getenv("OS_REGION_NAME"),
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
		Region: os.Getenv("OS_REGION_NAME"),
	})
}

// NewNetworkV2Client returns a *ServiceClient for making calls to the
// OpenStack Networking v2 API. An error will be returned if authentication
// or client creation was not possible.
func NewNetworkV2Client() (*golangsdk.ServiceClient, error) {
	ao, err := openstack.AuthOptionsFromEnv()
	if err != nil {
		return nil, err
	}

	client, err := openstack.AuthenticatedClient(ao)
	if err != nil {
		return nil, err
	}

	return openstack.NewNetworkV2(client, golangsdk.EndpointOpts{
		Region: os.Getenv("OS_REGION_NAME"),
	})
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
		Region: os.Getenv("OS_REGION_NAME"),
	})
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
