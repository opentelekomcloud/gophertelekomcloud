package nameservers

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

type commonResult struct {
	golangsdk.Result
}

type GetResult struct {
	commonResult
}

type Nameserver struct {
	Hostname string `json:"hostname"`
	Priority int    `json:"priority"`
}

// Extract is a function that accepts a result and extracts a nameserver.
func (r GetResult) Extract() ([]Nameserver, error) {
	var s []Nameserver
	err := r.ExtractIntoSlicePtr(&s, "nameservers")
	if err != nil {
		return nil, err
	}
	return s, nil
}
