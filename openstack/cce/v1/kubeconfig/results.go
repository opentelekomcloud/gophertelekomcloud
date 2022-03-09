package nodes

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
)

type GetResult struct {
	golangsdk.Result
}

// Extract interprets any extraSpecResult as an ExtraSpec, if possible.
func (r GetResult) Extract() (map[string]string, error) {
	var s map[string]string
	err := r.ExtractInto(&s)
	if err != nil {
		return nil, err
	}
	return s, nil
}
