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
	Err error
}

// GetResult contains the body of getting detailed tags request
type GetResult struct {
	golangsdk.Result
}

// Extract method will parse the result body into ResourceTag struct
func (r GetResult) Extract() ([]ResourceTag, error) {
	var s []ResourceTag
	err := r.ExtractIntoSlicePtr(&s, "tags")
	if err != nil {
		return nil, err
	}
	return s, nil
}

// ListResult contains the body of getting all tags request
type ListResult struct {
	golangsdk.Result
}

// Extract method will parse the result body into ListedTag struct
func (r ListResult) Extract() ([]ListedTag, error) {
	var s []ListedTag
	err := r.ExtractIntoSlicePtr(&s, "tags")
	if err != nil {
		return nil, err
	}
	return s, nil
}
