package central_network

import (
	"net/http"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type CreateOpts struct {
	// Central network name.
	Name string `json:"name" required:"true"`
	// Resource description. Angle brackets (<>) are not allowed.
	Description string `json:"description,omitempty"`
	// ID of the enterprise project that the central network belongs to.
	EnterpriseProjectId string `json:"enterprise_project_id,omitempty"`
	// Central network policy document.
	PolicyDocument *PolicyDocument `json:"policy_document,omitempty"`
	// Name of the default central network plane.
	DefaultPlane string `json:"default_plane,omitempty"`
	// List of the central network planes.
	Planes []PlaneDocument `json:"planes,omitempty"`
	// List of the enterprise routers on a central network.
	ErInstances []AssociateErInstance `json:"er_instances,omitempty"`
}

// PolicyDocument is the central network policy document.
type PolicyDocument struct {
	// Name of the default central network plane.
	DefaultPlane string `json:"default_plane,omitempty"`
	// List of the central network planes.
	Planes []PlaneDocument `json:"planes,omitempty"`
	// List of the enterprise routers on a central network.
	ErInstances []AssociateErInstance `json:"er_instances,omitempty"`
}

// PlaneDocument describes a central network plane.
type PlaneDocument struct {
	// Central network plane name.
	Name string `json:"name,omitempty"`
	// Enterprise router route tables associated with the central network plane.
	AssociateErTables []AssociateErTable `json:"associate_er_tables,omitempty"`
	// Connections between enterprise routers excluded from the central network plane.
	ExcludeErConnections [][]AssociateErInstance `json:"exclude_er_connections,omitempty"`
}

// AssociateErTable describes an associated enterprise router route table.
type AssociateErTable struct {
	// Project ID.
	ProjectId string `json:"project_id,omitempty"`
	// Region ID.
	RegionId string `json:"region_id,omitempty"`
	// Enterprise router ID.
	EnterpriseRouterId string `json:"enterprise_router_id,omitempty"`
	// Enterprise router route table ID.
	EnterpriseRouterTableId string `json:"enterprise_router_table_id,omitempty"`
}

// AssociateErInstance describes an enterprise router instance on a central network.
type AssociateErInstance struct {
	// Enterprise router ID.
	EnterpriseRouterId string `json:"enterprise_router_id,omitempty"`
	// Project ID.
	ProjectId string `json:"project_id,omitempty"`
	// Region ID.
	RegionId string `json:"region_id,omitempty"`
}

func Create(client *golangsdk.ServiceClient, opts CreateOpts) (*CentralNetwork, error) {
	b, err := build.RequestBody(opts, "central_network")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL(client.DomainID, "gcn", "central-networks"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200, 201, 202},
	})
	if err != nil {
		return nil, err
	}

	return extra(raw)
}

// CentralNetworkResp is the wrapper for a single central network response.
type CentralNetworkResp struct {
	// Request ID.
	RequestId string `json:"request_id"`
	// Central network.
	CentralNetwork CentralNetwork `json:"central_network"`
}

// CentralNetwork is the central network instance.
type CentralNetwork struct {
	// Central network ID.
	ID string `json:"id"`
	// Central network name.
	Name string `json:"name"`
	// Resource description.
	Description string `json:"description"`
	// Time when the resource was created.
	CreatedAt string `json:"created_at"`
	// Time when the resource was updated.
	UpdatedAt string `json:"updated_at"`
	// ID of the account that the central network belongs to.
	DomainId string `json:"domain_id"`
	// Central network status. Value options: AVAILABLE, CREATING, UPDATING, FAILED, DELETING, DELETED, RESTORING.
	State string `json:"state"`
	// ID of the enterprise project that the central network belongs to.
	EnterpriseProjectId string `json:"enterprise_project_id"`
	// Default plane ID.
	DefaultPlaneId string `json:"default_plane_id"`
	// Deprecated.
	AutoAssociateRouteEnabled bool `json:"auto_associate_route_enabled"`
	// Deprecated.
	AutoPropagateRouteEnabled bool `json:"auto_propagate_route_enabled"`
	// Central network planes.
	Planes []CentralNetworkPlane `json:"planes"`
	// Enterprise routers on the central network.
	ErInstances []CentralNetworkErInstance `json:"er_instances"`
	// Connections on the central network.
	Connections []CentralNetworkConnectionInfo `json:"connections"`
}

// CentralNetworkPlane describes a plane returned by the API.
type CentralNetworkPlane struct {
	// Plane ID.
	ID string `json:"id"`
	// Plane name.
	Name string `json:"name"`
	// Enterprise router route tables associated with the plane.
	AssociateErTables []AssociateErTable `json:"associate_er_tables"`
	// Connections between enterprise routers excluded from the plane.
	ExcludeErConnections [][]AssociateErInstance `json:"exclude_er_connections"`
}

// CentralNetworkErInstance describes an enterprise router on the central network.
type CentralNetworkErInstance struct {
	// Instance ID.
	ID string `json:"id"`
	// Enterprise router ID.
	EnterpriseRouterId string `json:"enterprise_router_id"`
	// Project ID.
	ProjectId string `json:"project_id"`
	// Region ID.
	RegionId string `json:"region_id"`
	// BGP autonomous system number.
	Asn int64 `json:"asn"`
	// Geographic site code.
	SiteCode string `json:"site_code"`
}

// CentralNetworkConnectionInfo describes a connection on the central network.
type CentralNetworkConnectionInfo struct {
	// Instance ID.
	ID string `json:"id"`
	// Plane ID.
	PlaneId string `json:"plane_id"`
	// Global connection bandwidth ID.
	GlobalConnectionBandwidthId string `json:"global_connection_bandwidth_id"`
	// Bandwidth size in Mbit/s.
	BandwidthSize int64 `json:"bandwidth_size"`
	// Connection type. Value options: ER-ER, ER-GDGW, ER-ER_ROUTE_TABLE.
	ConnectionType string `json:"connection_type"`
	// Two connection endpoints.
	ConnectionPointPair []ConnectionPoint `json:"connection_point_pair"`
	// Connection status.
	State string `json:"state"`
}

// ConnectionPoint describes a single connection endpoint.
type ConnectionPoint struct {
	// Instance ID.
	ID string `json:"id"`
	// Project ID.
	ProjectId string `json:"project_id"`
	// Region ID.
	RegionId string `json:"region_id"`
	// Geographic site code.
	SiteCode string `json:"site_code"`
	// Connection endpoint ID.
	InstanceId string `json:"instance_id"`
	// Parent resource ID.
	ParentInstanceId string `json:"parent_instance_id"`
	// Resource type. Value options: ER, GDGW, ER_ROUTE_TABLE.
	Type string `json:"type"`
}

func extra(raw *http.Response) (*CentralNetwork, error) {
	var res CentralNetworkResp
	if err := extract.Into(raw.Body, &res); err != nil {
		return nil, err
	}
	return &res.CentralNetwork, nil
}
