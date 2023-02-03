package flavors

import (
	"encoding/json"
	"strconv"

	"github.com/opentelekomcloud/gophertelekomcloud"
)

// Get retrieves details of a single flavor. Use ExtractFlavor to convert its result into a Flavor.
func Get(client *golangsdk.ServiceClient, id string) (*Flavor, error) {
	raw, err := client.Get(client.ServiceURL("flavors", id), nil, nil)
	return extraFla(err, raw)
}

// Flavor represent (virtual) hardware configurations for server resources in a region.
type Flavor struct {
	// ID is the flavor's unique ID.
	ID string `json:"id"`
	// Disk is the amount of root disk, measured in GB.
	Disk int `json:"disk"`
	// RAM is the amount of memory, measured in MB.
	RAM int `json:"ram"`
	// Name is the name of the flavor.
	Name string `json:"name"`
	// RxTxFactor describes bandwidth alterations of the flavor.
	RxTxFactor float64 `json:"rxtx_factor"`
	// Swap is the amount of swap space, measured in MB.
	Swap int `json:"-"`
	// VCPUs indicates how many (virtual) CPUs are available for this flavor.
	VCPUs int `json:"vcpus"`
	// IsPublic indicates whether the flavor is public.
	IsPublic bool `json:"os-flavor-access:is_public"`
	// Ephemeral is the amount of ephemeral disk space, measured in GB.
	Ephemeral int `json:"OS-FLV-EXT-DATA:ephemeral"`
}

func (r *Flavor) UnmarshalJSON(b []byte) error {
	type tmp Flavor
	var s struct {
		tmp
		Swap interface{} `json:"swap"`
	}
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}

	*r = Flavor(s.tmp)

	switch t := s.Swap.(type) {
	case float64:
		r.Swap = int(t)
	case string:
		switch t {
		case "":
			r.Swap = 0
		default:
			swap, err := strconv.ParseFloat(t, 64)
			if err != nil {
				return err
			}
			r.Swap = int(swap)
		}
	}

	return nil
}
