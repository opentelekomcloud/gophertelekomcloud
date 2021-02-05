package tags

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
)

type ListedTag struct {
	Key    string   `json:"key"`
	Values []string `json:"values"`
}

// ActionResult is the action result which is the result of create or delete operations
type ActionResult struct {
	golangsdk.ErrResult
}

// GetResult contains the body of getting detailed tags request
type GetResult struct {
	golangsdk.Result
}

// Extract method will parse the result body into ResourceTag struct
func (r GetResult) Extract() ([]ResourceTag, error) {
	var responseTags struct {
		Tags []ResourceTag `json:"tags"`
	}
	err := r.Result.ExtractInto(&responseTags)
	return responseTags.Tags, err
}

// ListResult contains the body of getting all tags request
type ListResult struct {
	golangsdk.Result
}

// Extract method will parse the result body into ListedTag struct
func (r ListResult) Extract() ([]ListedTag, error) {
	var responseTags struct {
		ListedTags []ListedTag `json:"tags"`
	}
	err := r.Result.ExtractInto(&responseTags)
	return responseTags.ListedTags, err
}
