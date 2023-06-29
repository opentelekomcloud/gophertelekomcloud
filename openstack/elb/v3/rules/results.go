package rules

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
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

type RulePage struct {
	pagination.PageWithInfo
}

func (p RulePage) IsEmpty() (bool, error) {
	rules, err := ExtractRules(p)
	return len(rules) == 0, err
}

func ExtractRules(p pagination.Page) ([]ForwardingRule, error) {
	var policies []ForwardingRule
	err := p.(RulePage).ExtractIntoSlicePtr(&policies, "rules")
	if err != nil {
		return nil, err
	}
	return policies, nil
}
