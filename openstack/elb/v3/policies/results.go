package policies

import (
	"bytes"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/structs"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/elb/v3/rules"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type Policy struct {
	ID                  string                `json:"id"`
	Action              Action                `json:"action"`
	Description         string                `json:"description"`
	ListenerID          string                `json:"listener_id"`
	Name                string                `json:"name"`
	Position            int                   `json:"position"`
	Priority            int                   `json:"priority"`
	ProjectID           string                `json:"project_id"`
	Status              string                `json:"provisioning_status"`
	RedirectListenerID  string                `json:"redirect_listener_id"`
	RedirectPoolID      string                `json:"redirect_pool_id"`
	Rules               []structs.ResourceRef `json:"rules"`
	Conditions          []rules.Condition     `json:"conditions"`
	FixedResponseConfig FixedResponseOptions  `json:"fixed_response_config"`
	RedirectUrl         string                `json:"redirect_url"`
	RedirectUrlConfig   RedirectUrlOptions    `json:"redirect_url_config"`
	RedirectPoolsConfig []RedirectPoolOptions `json:"redirect_pools_config"`
}

type commonResult struct {
	golangsdk.Result
}

func (r commonResult) Extract() (*Policy, error) {
	p := &Policy{}
	err := r.ExtractIntoStructPtr(p, "l7policy")
	if err != nil {
		return nil, err
	}
	return p, nil
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

	err := extract.IntoSlicePtr(bytes.NewReader(p.(PolicyPage).Body), &policies, "l7policies")
	if err != nil {
		return nil, err
	}
	return policies, nil
}
