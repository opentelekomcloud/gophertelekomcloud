package policies

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type commonResult struct {
	golangsdk.Result
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

type PolicyPage struct {
	pagination.PageWithInfo
}

func (p PolicyPage) IsEmpty() (bool, error) {
	l, err := ExtractPolicies(p)
	if err != nil {
		return false, err
	}
	return len(l) == 0, nil
}

func ExtractPolicies(p pagination.Page) ([]Policy, error) {
	var policies []Policy
	err := p.(PolicyPage).ExtractIntoSlicePtr(&policies, "l7policies")
	if err != nil {
		return nil, err
	}
	return policies, nil
}
