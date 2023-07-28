package topicattributes

import "github.com/opentelekomcloud/gophertelekomcloud"

type GetResult struct {
	golangsdk.Result
}

func (r GetResult) Extract() (map[string]string, error) {
	var attributes struct {
		Attributes map[string]string `json:"attributes"`
	}
	err := r.ExtractIntoStructPtr(&attributes, "")
	if err != nil {
		return nil, err
	}
	return attributes.Attributes, nil
}

type UpdateResult struct {
	golangsdk.Result
}

type DeleteResult struct {
	golangsdk.ErrResult
}
