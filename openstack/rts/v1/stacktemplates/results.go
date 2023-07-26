package stacktemplates

import (
	"encoding/json"
	"io"

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
	mapBody := make(map[string]interface{})

	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	// make sure return pretty-printed body
	if err := json.Unmarshal(bytes, &mapBody); err != nil {
		return nil, err
	}

	template, err := json.MarshalIndent(mapBody, "", "  ")
	if err != nil {
		return nil, err
	}
	return template, nil
}
