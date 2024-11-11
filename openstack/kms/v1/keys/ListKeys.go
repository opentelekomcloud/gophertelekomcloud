package keys

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ListOpts struct {
	// State of a CMK
	KeyState string `json:"key_state,omitempty"`
	Limit    string `json:"limit,omitempty"`
	Marker   string `json:"marker,omitempty"`
}

type ListKey struct {
	Keys       []string `json:"keys"`
	KeyDetails []Key    `json:"key_details"`
	NextMarker string   `json:"next_marker"`
	Truncated  string   `json:"truncated"`
}

func List(client *golangsdk.ServiceClient, opts ListOpts) (*ListKey, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL("kms", "list-keys"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res ListKey
	err = extract.Into(raw.Body, &res)
	return &res, err
}
