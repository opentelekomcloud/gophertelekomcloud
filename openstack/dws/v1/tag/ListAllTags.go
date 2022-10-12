package tag

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func ListAllTags(client *golangsdk.ServiceClient) ([]TagWithMultiValue, error) {
	// GET /v1.0/{project_id}/clusters/tags
	raw, err := client.Get(client.ServiceURL("clusters", "tags"), nil, nil)
	if err != nil {
		return nil, err
	}

	var res []TagWithMultiValue
	err = extract.IntoSlicePtr(raw.Body, &res, "tags")
	return res, err
}

type TagWithMultiValue struct {
	// Tag key. A tag key can contain a maximum of 127 Unicode characters, which cannot be null. The first and last characters cannot be spaces.
	// It can contain uppercase letters (A to Z), lowercase letters (a to z), digits (0-9), hyphens (-), and underscores (_).
	Key string `json:"key"`
	// Tag value. A tag value can contain a maximum of 255 Unicode characters, which can be null.
	// The first and last characters cannot be spaces.
	// It can contain uppercase letters (A to Z), lowercase letters (a to z), digits (0-9), hyphens (-), and underscores (_).
	Values []string `json:"values,omitempty"`
}
