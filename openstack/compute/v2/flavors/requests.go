package flavors

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
)

// ListOptsBuilder allows extensions to add additional parameters to the List request.
type ListOptsBuilder interface {
	ToFlavorListQuery() (string, error)
}

/*
AccessType maps to OpenStack's Flavor.is_public field. Although the is_public
field is boolean, the request options are ternary, which is why AccessType is
a string. The following values are allowed:

The AccessType argument is optional, and if it is not supplied, OpenStack
returns the PublicAccess flavors.
*/
type AccessType string

const (
	// PublicAccess returns public flavors and private flavors associated with that project.
	PublicAccess AccessType = "true"
	// PrivateAccess (admin only) returns private flavors, across all projects.
	PrivateAccess AccessType = "false"
	// AllAccess (admin only) returns public and private flavors across all projects.
	AllAccess AccessType = "None"
)

/*
ListOpts filters the results returned by the List() function.
For example, a flavor with a minDisk field of 10 will not be returned if you
specify MinDisk set to 20.

Typically, software will use the last ID of the previous call to List to set
the Marker for the current call.
*/
type ListOpts struct {
	// ChangesSince, if provided, instructs List to return only those things which
	// have changed since the timestamp provided.
	ChangesSince string `q:"changes-since"`
	// MinDisk and MinRAM, if provided, elides flavors which do not meet your criteria.
	MinDisk int `q:"minDisk"`
	MinRAM  int `q:"minRam"`
	// SortDir allows to select sort direction. It can be "asc" or "desc" (default).
	SortDir string `q:"sort_dir"`
	// SortKey allows to sort by one of the flavors attributes. Default is flavorId.
	SortKey string `q:"sort_key"`
	// Marker and Limit control paging. Marker instructs List where to start listing from.
	Marker string `q:"marker"`
	// Limit instructs List to refrain from sending excessively large lists of flavors.
	Limit int `q:"limit"`
	// AccessType, if provided, instructs List which set of flavors to return.
	// If IsPublic not provided, flavors for the current project are returned.
	AccessType AccessType `q:"is_public"`
}

// ToFlavorListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToFlavorListQuery() (string, error) {
	q, err := golangsdk.BuildQueryString(opts)
	if err != nil {
		return "", err
	}
	return q.String(), err
}

type CreateOptsBuilder interface {
	ToFlavorCreateMap() (map[string]interface{}, error)
}

// CreateOpts specifies parameters used for creating a flavor.
type CreateOpts struct {
	// Name is the name of the flavor.
	Name string `json:"name" required:"true"`
	// RAM is the memory of the flavor, measured in MB.
	RAM int `json:"ram" required:"true"`
	// VCPUs is the number of vcpus for the flavor.
	VCPUs int `json:"vcpus" required:"true"`
	// Disk the amount of root disk space, measured in GB.
	Disk *int `json:"disk" required:"true"`
	// ID is a unique ID for the flavor.
	ID string `json:"id,omitempty"`
	// Swap is the amount of swap space for the flavor, measured in MB.
	Swap *int `json:"swap,omitempty"`
	// RxTxFactor alters the network bandwidth of a flavor.
	RxTxFactor float64 `json:"rxtx_factor,omitempty"`
	// IsPublic flags a flavor as being available to all projects or not.
	IsPublic *bool `json:"os-flavor-access:is_public,omitempty"`
	// Ephemeral is the amount of ephemeral disk space, measured in GB.
	Ephemeral *int `json:"OS-FLV-EXT-DATA:ephemeral,omitempty"`
}

// ToFlavorCreateMap constructs a request body from CreateOpts.
func (opts CreateOpts) ToFlavorCreateMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "flavor")
}

// AddAccessOptsBuilder allows extensions to add additional parameters to the
// AddAccess requests.
type AddAccessOptsBuilder interface {
	ToFlavorAddAccessMap() (map[string]interface{}, error)
}

// AddAccessOpts represents options for adding access to a flavor.
type AddAccessOpts struct {
	// Tenant is the project/tenant ID to grant access.
	Tenant string `json:"tenant"`
}

// ToFlavorAddAccessMap constructs a request body from AddAccessOpts.
func (opts AddAccessOpts) ToFlavorAddAccessMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "addTenantAccess")
}

// RemoveAccessOptsBuilder allows extensions to add additional parameters to the
// RemoveAccess requests.
type RemoveAccessOptsBuilder interface {
	ToFlavorRemoveAccessMap() (map[string]interface{}, error)
}

// RemoveAccessOpts represents options for removing access to a flavor.
type RemoveAccessOpts struct {
	// Tenant is the project/tenant ID to grant access.
	Tenant string `json:"tenant"`
}

// ToFlavorRemoveAccessMap constructs a request body from RemoveAccessOpts.
func (opts RemoveAccessOpts) ToFlavorRemoveAccessMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "removeTenantAccess")
}

// CreateExtraSpecsOptsBuilder allows extensions to add additional parameters to the
// CreateExtraSpecs requests.
type CreateExtraSpecsOptsBuilder interface {
	ToFlavorExtraSpecsCreateMap() (map[string]interface{}, error)
}

// ExtraSpecsOpts is a map that contains key-value pairs.
type ExtraSpecsOpts map[string]string

// ToFlavorExtraSpecsCreateMap assembles a body for a Create request based on
// the contents of ExtraSpecsOpts.
func (opts ExtraSpecsOpts) ToFlavorExtraSpecsCreateMap() (map[string]interface{}, error) {
	return map[string]interface{}{"extra_specs": opts}, nil
}

// UpdateExtraSpecOptsBuilder allows extensions to add additional parameters to
// the Update request.
type UpdateExtraSpecOptsBuilder interface {
	ToFlavorExtraSpecUpdateMap() (map[string]string, string, error)
}

// ToFlavorExtraSpecUpdateMap assembles a body for an Update request based on
// the contents of a ExtraSpecOpts.
func (opts ExtraSpecsOpts) ToFlavorExtraSpecUpdateMap() (map[string]string, string, error) {
	if len(opts) != 1 {
		err := golangsdk.ErrInvalidInput{}
		err.Argument = "flavors.ExtraSpecOpts"
		err.Info = "Must have 1 and only one key-value pair"
		return nil, "", err
	}

	var key string
	for k := range opts {
		key = k
	}

	return opts, key, nil
}

// IDFromName is a convienience function that returns a flavor's ID given its
// name.
func IDFromName(client *golangsdk.ServiceClient, name string) (string, error) {
	count := 0
	id := ""
	allPages, err := ListDetail(client, nil).AllPages()
	if err != nil {
		return "", err
	}

	all, err := ExtractFlavors(allPages)
	if err != nil {
		return "", err
	}

	for _, f := range all {
		if f.Name == name {
			count++
			id = f.ID
		}
	}

	switch count {
	case 0:
		err := &golangsdk.ErrResourceNotFound{}
		err.ResourceType = "flavor"
		err.Name = name
		return "", err
	case 1:
		return id, nil
	default:
		err := &golangsdk.ErrMultipleResourcesFound{}
		err.ResourceType = "flavor"
		err.Name = name
		err.Count = count
		return "", err
	}
}
