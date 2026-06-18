package policy

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type CreateOpts struct {
	// Central network ID.
	CentralNetworkId string `json:"-" required:"true"`
	// Name of the default central network plane.
	DefaultPlane string `json:"default_plane" required:"true"`
	// List of the central network planes.
	Planes []PlaneDocument `json:"planes" required:"true"`
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

func Create(client *golangsdk.ServiceClient, opts CreateOpts) (*CentralNetworkPolicy, error) {
	b, err := build.RequestBody(opts, "central_network_policy_document")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL(client.DomainID, "gcn", "central-network", opts.CentralNetworkId, "policies"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200, 201, 202},
	})
	if err != nil {
		return nil, err
	}

	var res struct {
		CentralNetworkPolicy CentralNetworkPolicy `json:"central_network_policy"`
	}
	err = extract.Into(raw.Body, &res)
	return &res.CentralNetworkPolicy, err
}

// CentralNetworkPolicy is a central network policy.
type CentralNetworkPolicy struct {
	// Instance ID.
	ID string `json:"id"`
	// Time when the resource was created.
	CreatedAt string `json:"created_at"`
	// Account ID.
	DomainId string `json:"domain_id"`
	// Policy status. Value options: AVAILABLE, CANCELING, APPLYING, FAILED, DELETED.
	State string `json:"state"`
	// Central network ID.
	CentralNetworkId string `json:"central_network_id"`
	// Policy document template version.
	DocumentTemplateVersion string `json:"document_template_version"`
	// Whether the policy is applied.
	IsApplied bool `json:"is_applied"`
	// Policy version.
	Version int `json:"version"`
	// Central network policy document.
	Document PolicyDocument `json:"document"`
}

// PolicyDocument is the central network policy document.
type PolicyDocument struct {
	// Name of the default central network plane.
	DefaultPlane string `json:"default_plane"`
	// List of the central network planes.
	Planes []PlaneDocument `json:"planes"`
	// List of the enterprise routers on a central network.
	ErInstances []AssociateErInstance `json:"er_instances"`
}
