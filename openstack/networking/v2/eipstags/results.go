package eipstags

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

// CreateResult is a struct which contains the result of creation
type CreateResult struct {
	golangsdk.ErrResult
}

// ListResult contains the body of getting detailed EIP tags request
type ListResult struct {
	golangsdk.Result
}

// Extract method will parse the result body into Tag struct
func (r ListResult) Extract() ([]Tag, error) {
	var responseTags struct {
		Tags []Tag `json:"tags"`
	}
	err := r.Result.ExtractInto(&responseTags)
	return responseTags.Tags, err
}

// DeleteResult is a struct which contains the result of deletion
type DeleteResult struct {
	golangsdk.ErrResult
}

// ActionResult is a struct which contains the result of action
type ActionResult struct {
	golangsdk.ErrResult
}
