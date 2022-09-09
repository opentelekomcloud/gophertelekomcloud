package hypervisors

import (
	"encoding/json"
	"fmt"

	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// List makes a request against the API to list hypervisors.
func List(client *golangsdk.ServiceClient) ([]Hypervisor, error) {
	raw, err := client.Get(client.ServiceURL("os-hypervisors", "detail"), nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res []Hypervisor
	err = extract.IntoSlicePtr(raw.Body, &res, "hypervisors")
	return res, err
}

func (r *Hypervisor) UnmarshalJSON(b []byte) error {
	type tmp Hypervisor
	var s struct {
		tmp
		CPUInfo           interface{} `json:"cpu_info"`
		HypervisorVersion interface{} `json:"hypervisor_version"`
		FreeDiskGB        interface{} `json:"free_disk_gb"`
		LocalGB           interface{} `json:"local_gb"`
	}

	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}

	*r = Hypervisor(s.tmp)

	// Newer versions return the CPU info as the correct type.
	// Older versions return the CPU info as a string and need to be
	// unmarshalled by the json parser.
	var tmpb []byte

	switch t := s.CPUInfo.(type) {
	case string:
		tmpb = []byte(t)
	case map[string]interface{}:
		tmpb, err = json.Marshal(t)
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("CPUInfo has unexpected type: %T", t)
	}

	err = json.Unmarshal(tmpb, &r.CPUInfo)
	if err != nil {
		return err
	}

	// These fields may be returned as a scientific notation, so they need
	// converted to int.
	switch t := s.HypervisorVersion.(type) {
	case int:
		r.HypervisorVersion = t
	case float64:
		r.HypervisorVersion = int(t)
	default:
		return fmt.Errorf("Hypervisor version of unexpected type")
	}

	switch t := s.FreeDiskGB.(type) {
	case int:
		r.FreeDiskGB = t
	case float64:
		r.FreeDiskGB = int(t)
	default:
		return fmt.Errorf("Free disk GB of unexpected type")
	}

	switch t := s.LocalGB.(type) {
	case int:
		r.LocalGB = t
	case float64:
		r.LocalGB = int(t)
	default:
		return fmt.Errorf("Local GB of unexpected type")
	}

	return nil
}
