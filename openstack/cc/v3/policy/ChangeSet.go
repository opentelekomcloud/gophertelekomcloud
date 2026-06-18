package policy

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// ListChangeSet queries the changes between the current policy and an applied policy.
func ListChangeSet(client *golangsdk.ServiceClient, centralNetworkId, policyId string) (*ChangeSetResp, error) {
	raw, err := client.Get(client.ServiceURL(client.DomainID, "gcn", "central-network", centralNetworkId, "policies", policyId, "change-set"), nil, nil)
	if err != nil {
		return nil, err
	}

	var res ChangeSetResp
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type ChangeSetResp struct {
	// Request ID.
	RequestId string `json:"request_id"`
	// Pagination query information.
	PageInfo PageInfo `json:"page_info"`
	// List of central network policy changes.
	CentralNetworkPolicyChangeSet []ElementChangeEntry `json:"central_network_policy_change_set"`
}

// ElementChangeEntry describes a single change between policies.
type ElementChangeEntry struct {
	// Change operation. Value options: CreateCentralNetworkPlane, DeleteCentralNetworkPlane,
	// UpdateCentralNetworkPlane, CreateCentralNetworkErInstance, DeleteCentralNetworkErInstance,
	// CreateCentralNetworkErConnection, DeleteCentralNetworkErConnection, CreateCentralNetworkErTable,
	// DeleteCentralNetworkErTable, SwitchCentralNetworkErTable.
	OperationId string `json:"operation_id"`
	// Plane to be created.
	CreateCentralNetworkPlane *PlaneChangeDocument `json:"create_central_network_plane"`
	// Original plane.
	OriginalCentralNetworkPlane *PlaneChangeDocument `json:"original_central_network_plane"`
	// Newest plane.
	NewestCentralNetworkPlane *PlaneChangeDocument `json:"newest_central_network_plane"`
	// Plane to be deleted.
	DeleteCentralNetworkPlane *PlaneChangeDocument `json:"delete_central_network_plane"`
	// Enterprise router instance to be created.
	CreateCentralNetworkErInstance *AssociateErInstance `json:"create_central_network_er_instance"`
	// Enterprise router instance to be deleted.
	DeleteCentralNetworkErInstance *AssociateErInstance `json:"delete_central_network_er_instance"`
	// Central network plane name.
	CentralNetworkPlaneName string `json:"central_network_plane_name"`
	// Enterprise router connections to be created.
	CreateCentralNetworkErConnection []AssociateErTable `json:"create_central_network_er_connection"`
	// Enterprise router connections to be deleted.
	DeleteCentralNetworkErConnection []AssociateErTable `json:"delete_central_network_er_connection"`
	// Enterprise router route table to be created.
	CreateCentralNetworkErTable *AssociateErTable `json:"create_central_network_er_table"`
	// Enterprise router route table to be deleted.
	DeleteCentralNetworkErTable *AssociateErTable `json:"delete_central_network_er_table"`
	// Enterprise router route table to be switched.
	SwitchCentralNetworkErTable *SwitchErTableDocument `json:"switch_central_network_er_table"`
}

// PlaneChangeDocument describes a plane in a change set entry.
type PlaneChangeDocument struct {
	// Plane name.
	Name string `json:"name"`
	// Whether the plane is the default one.
	IsDefault bool `json:"is_default"`
	// Enterprise router route tables associated with the plane.
	AssociateErTables []AssociateErTable `json:"associate_er_tables"`
	// Connections between enterprise routers excluded from the plane.
	ExcludeErConnections [][]AssociateErInstance `json:"exclude_er_connections"`
}

// SwitchErTableDocument describes a route table switch.
type SwitchErTableDocument struct {
	// Project ID.
	ProjectId string `json:"project_id"`
	// Region ID.
	RegionId string `json:"region_id"`
	// Enterprise router ID.
	EnterpriseRouterId string `json:"enterprise_router_id"`
	// Route table ID of the original enterprise router.
	OriginalEnterpriseRouterTableId string `json:"original_enterprise_router_table_id"`
	// Route table ID of the new enterprise router.
	NewEnterpriseRouterTableId string `json:"new_enterprise_router_table_id"`
}
