package policies

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type RuleRef struct {
	ID string `json:"id"`
}

type Policy struct {
	ID                 string    `json:"id"`
	Action             Action    `json:"action"`
	Description        string    `json:"description"`
	ListenerID         string    `json:"listener_id"`
	Name               string    `json:"name"`
	Position           int       `json:"position"`
	ProjectID          string    `json:"project_id"`
	Status             string    `json:"provisioning_status"`
	RedirectListenerID string    `json:"redirect_listener_id"`
	RedirectPoolID     string    `json:"redirect_pool_id"`
	Rules              []RuleRef `json:"rules"`
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

type PageInfo struct {
	PreviousMarker string `json:"previous_marker"`
	NextMarker     string `json:"next_marker"`
	CurrentCount   int    `json:"current_count"`
}

type PolicyPage struct {
	pagination.MarkerPageBase
}

func (p PolicyPage) LastMarker() (string, error) {
	var info PageInfo
	err := p.ExtractIntoStructPtr(&info, "page_info")
	if err != nil {
		return "", err
	}
	return info.NextMarker, nil
}

// NextPageURL generates the URL for the page of results after this one.
func (p PolicyPage) NextPageURL() (string, error) {
	currentURL := p.URL

	mark, err := p.Owner.LastMarker()
	if err != nil {
		return "", err
	}
	if mark == "" {
		return "", nil
	}

	q := currentURL.Query()
	q.Set("marker", mark)
	currentURL.RawQuery = q.Encode()

	return currentURL.String(), nil
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
