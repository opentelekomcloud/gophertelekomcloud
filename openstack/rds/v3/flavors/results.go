package flavors

import (
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type Flavor struct {
	VCPUs        string            `json:"vcpus"`
	RAM          int               `json:"ram"`
	SpecCode     string            `json:"spec_code"`
	InstanceMode string            `json:"instance_mode"`
	AzStatus     map[string]string `json:"az_status"`
}

type DbFlavorsPage struct {
	pagination.SinglePageBase
}

func (r DbFlavorsPage) IsEmpty() (bool, error) {
	flavors, err := ExtractDbFlavors(r)
	if err != nil {
		return false, err
	}
	return len(flavors) == 0, err
}

func ExtractDbFlavors(r pagination.Page) ([]Flavor, error) {
	var s []Flavor
	err := (r.(DbFlavorsPage)).ExtractIntoSlicePtr(&s, "flavors")
	if err != nil {
		return nil, err
	}
	return s, nil
}
