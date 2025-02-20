package management

type CreateTags struct {
	// Resource tag key.
	Key string `json:"key,omitempty"`

	// Resource tag value.
	Value string `json:"value,omitempty"`
}

type Tags struct {
	// Resource tag key.
	Key string `json:"key"`
	// Resource tag value.
	Value string `json:"value"`
}
