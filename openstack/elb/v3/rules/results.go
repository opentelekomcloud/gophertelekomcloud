package rules

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
)

type commonResult struct {
	golangsdk.Result
}

func (r commonResult) Extract() (*ForwardingRule, error) {
	var rule ForwardingRule
	err := r.ExtractIntoStructPtr(&rule, "rule")
	if err != nil {
		return nil, err
	}
	return &rule, nil
}

type CreateResult struct {
	commonResult
}

type GetResult struct {
	commonResult
}

type UpdateResult struct {
	commonResult
}

type DeleteResult struct {
	golangsdk.ErrResult
}
