package nics

import (
	"reflect"

	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

// ListOpts allows the filtering and sorting of paginated collections through
// the API.
type ListOpts struct {
	// ID is the unique identifier for the nic.
	ID string `json:"port_id"`

	// Status indicates whether a nic is currently operational.
	Status string `json:"port_state"`
}

// List returns collection of nics. It accepts a ListOpts struct, which allows you to filter and sort
// the returned collection for greater efficiency.
func List(client *golangsdk.ServiceClient, serverId string, opts ListOpts) ([]Nic, error) {
	pages, err := pagination.NewPager(client, client.ServiceURL("servers", serverId, "os-interface"),
		func(r pagination.PageResult) pagination.Page {
			return NicPage{pagination.LinkedPageBase{PageResult: r}}
		}).AllPages()
	if err != nil {
		return nil, err
	}

	allNICs, err := ExtractNics(pages)
	if err != nil {
		return nil, err
	}

	return filterNICs(allNICs, opts)
}

// filterNICs used to filter nics using id and status.
func filterNICs(nics []Nic, opts ListOpts) ([]Nic, error) {

	var refinedNICs []Nic
	var matched bool
	m := map[string]interface{}{}

	if opts.ID != "" {
		m["ID"] = opts.ID
	}
	if opts.Status != "" {
		m["Status"] = opts.Status
	}
	if len(m) > 0 && len(nics) > 0 {
		for _, nic := range nics {
			matched = true

			for key, value := range m {
				if sVal := getStructField(&nic, key); !(sVal == value) {
					matched = false
				}
			}

			if matched {
				refinedNICs = append(refinedNICs, nic)
			}
		}

	} else {
		refinedNICs = nics
	}

	return refinedNICs, nil
}

func getStructField(v *Nic, field string) string {
	r := reflect.ValueOf(v)
	f := reflect.Indirect(r).FieldByName(field)
	return f.String()
}
