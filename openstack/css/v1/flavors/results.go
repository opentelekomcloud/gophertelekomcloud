package flavors

import (
	"encoding/json"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Flavor struct {
	// RAM - memory size of an instance.
	// Unit: GB
	RAM int `json:"ram"`
	// CPU - number of vCPUs of an instance.
	CPU int `json:"cpu"`
	// Name - flavor name.
	Name   string `json:"name"`
	Region string `json:"region"`
	// DiskMin - minimal disk capacity of an instance.
	DiskMin int `json:"-"`
	// DiskMax - maximum disk capacity of an instance.
	DiskMax int `json:"-"`
	// FlavorID - ID of a flavor.
	FlavorID string `json:"flavor_id"`
}

func (f *Flavor) UnmarshalJSON(b []byte) error {
	type tmp Flavor
	var flavor struct {
		tmp
		DiskRange string `json:"diskrange"`
	}
	if err := json.Unmarshal(b, &flavor); err != nil {
		return err
	}
	*f = Flavor(flavor.tmp)
	dRange := strings.Split(flavor.DiskRange, ",")
	if len(dRange) != 2 {
		return fmt.Errorf("disk range format is invalid: %s", flavor.DiskRange)
	}
	min, errMin := strconv.Atoi(dRange[0])
	max, errMax := strconv.Atoi(dRange[1])
	if errMin != nil || errMax != nil {
		return fmt.Errorf("failed to convert disk range format: %s, %s", errMin, errMax)
	}
	f.DiskMin = min
	f.DiskMax = max
	return nil
}

type Version struct {
	// Version - engine version
	Version string `json:"version"`
	// Type - instance type.
	// The options are `ess`, `ess-cold`, `ess-master`, and `ess-client`.
	Type string `json:"type"`
	// Flavors - list of flavors
	Flavors []Flavor `json:"flavors"`
}

type Limit struct {
	Min int
	Max int
}

func (r Limit) Matches(value int) bool {
	if r.Max == 0 {
		r.Max = math.MaxInt32
	}
	return value >= r.Min && value <= r.Max
}

// FilterOpts to filter version list by
type FilterOpts struct {
	Version string
	// One of ess, ess-master, ess-client, ess-cloud
	Type string
	// Name of the searched flavor
	FlavorName string
	DiskMin    *Limit
	DiskMax    *Limit
	// Region - region the flavor is available for
	Region string
	CPU    *Limit
	// RAM - memory size, GB
	RAM *Limit
}

func matches(value int, limits *Limit) bool {
	if limits == nil {
		return true
	}
	return (*limits).Matches(value)
}

func filterFlavors(flavors []Flavor, opts FilterOpts) []Flavor {
	var resFlavors []Flavor
	for _, flv := range flavors {
		if opts.FlavorName != "" && flv.Name != opts.FlavorName {
			continue
		}
		if opts.Region != "" && flv.Region != opts.Region {
			continue
		}

		if !matches(flv.CPU, opts.CPU) {
			continue
		}
		if !matches(flv.RAM, opts.RAM) {
			continue
		}
		if !matches(flv.DiskMin, opts.DiskMin) {
			continue
		}
		if !matches(flv.DiskMax, opts.DiskMax) {
			continue
		}

		resFlavors = append(resFlavors, flv)
	}
	return resFlavors
}

func findSingleFlavor(flavors []Flavor, opts FilterOpts) *Flavor {
	for _, flv := range flavors {
		if (opts.FlavorName == "" || flv.Name == opts.FlavorName) &&
			(opts.Region == "" || flv.Region == opts.Region) &&
			matches(flv.CPU, opts.CPU) &&
			matches(flv.RAM, opts.RAM) &&
			matches(flv.DiskMin, opts.DiskMin) &&
			matches(flv.DiskMax, opts.DiskMax) {
			return &flv
		}
	}
	return nil
}

// FilterVersions - filters flavors in version list by given options (with AND operator)
func FilterVersions(versions []Version, opts FilterOpts) []Version {
	var resVersions []Version
	for _, version := range versions {

		if opts.Version != "" && version.Version != opts.Version {
			continue
		}
		if opts.Type != "" && version.Type != opts.Type {
			continue
		}
		version.Flavors = filterFlavors(version.Flavors, opts)
		resVersions = append(resVersions, version)
	}
	return resVersions
}

// FindFlavor - finds first flavor matching options
func FindFlavor(versions []Version, opts FilterOpts) *Flavor {
	for _, version := range versions {
		if opts.Version != "" && version.Version != opts.Version {
			continue
		}
		if opts.Type != "" && version.Type != opts.Type {
			continue
		}
		flavor := findSingleFlavor(version.Flavors, opts)
		if flavor != nil {
			return flavor
		}
	}
	return nil
}
