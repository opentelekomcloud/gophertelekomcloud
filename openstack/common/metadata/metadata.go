package metadata

import (
	"io"

	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// result is a struct wrapper for a metadata map
type result struct {
	Metadata map[string]interface{} `json:"metadata"`
}

func Extract(reader io.Reader) (map[string]interface{}, error) {
	metadata := new(result)
	err := extract.Into(reader, metadata)
	if err != nil {
		return nil, err
	}

	return metadata.Metadata, nil
}
