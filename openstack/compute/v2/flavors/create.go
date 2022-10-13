package flavors

import (
	"net/http"

	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

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

// Create requests the creation of a new flavor.
func Create(client *golangsdk.ServiceClient, opts CreateOpts) (*Flavor, error) {
	b, err := build.RequestBody(opts, "flavor")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL("flavors"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200, 201},
	})
	return extraFla(err, raw)
}

func extraFla(err error, raw *http.Response) (*Flavor, error) {
	if err != nil {
		return nil, err
	}

	var res Flavor
	err = extract.IntoStructPtr(raw.Body, &res, "flavor")
	return &res, err
}
