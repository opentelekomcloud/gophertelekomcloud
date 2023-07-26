package stacktemplates

import (
	"encoding/json"

	"github.com/opentelekomcloud/gophertelekomcloud"
)

// GetResult represents the result of a Get operation.
type GetResult struct {
	golangsdk.Result
}

// Extract returns the JSON template and is called after a Get operation.
func (r GetResult) Extract() ([]byte, error) {
	if r.Err != nil {
		return nil, r.Err
	}
	mapBody := make(map[string]any)

	// make sure return pretty-printed body
	if err := json.Unmarshal(r.Body, &mapBody); err != nil {
		return nil, err
	}

	template, err := json.MarshalIndent(mapBody, "", "  ")
	if err != nil {
		return nil, err
	}
	return template, nil
}
