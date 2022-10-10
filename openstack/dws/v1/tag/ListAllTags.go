package tag

type ListAllTagsRequest struct {
}

// GET /v1.0/{project_id}/clusters/tags

type ListAllTagsResponse struct {
	// Tag list.
	Tags []TagWithMultiValue `json:"tags,omitempty"`
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
