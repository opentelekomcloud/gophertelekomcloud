package eipstags

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

// CreateResult
type CreateResult struct {
	golangsdk.ErrResult
}

// ListResult
type ListResult struct {
	golangsdk.Result
}

// Tags model
type Tags struct {
	// Tags is a list of any tags. Tags are arbitrarily defined strings
	// attached to a resource.
	Tags []string `json:"tags"`
}

func (r ListResult) Extract() ([]Tag, error) {
	var responseTags struct {
		Tags []Tag `json:"tags"`
	}
	err := r.Result.ExtractInto(&responseTags)
	return responseTags.Tags, err
}

// DeleteResult
type DeleteResult struct {
	golangsdk.ErrResult
}

// ActionResult
type ActionResult struct {
	golangsdk.ErrResult
}
