package grants

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

type CreateGrant struct {
	GrantID string `json:"grant_id"`
}

type ListGrant struct {
	Grants     []Grant `json:"grants"`
	NextMarker string  `json:"next_marker"`
	Truncated  string  `json:"truncated"`
	Total      int     `json:"total"`
}

type Grant struct {
	KeyID             string   `json:"key_id"`
	GrantID           string   `json:"grant_id"`
	GranteePrincipal  string   `json:"grantee_principal"`
	Operations        []string `json:"operations"`
	IssuingPrincipal  string   `json:"issuing_principal"`
	CreationDate      string   `json:"creation_date"`
	Name              string   `json:"name"`
	RetiringPrincipal string   `json:"retiring_principal"`
}

// CreateResult contains the response body and error from a Create request.
type CreateResult struct {
	golangsdk.Result
}

func (r CreateResult) Extract() (*CreateGrant, error) {
	s := new(CreateGrant)
	err := r.ExtractInto(s)
	if err != nil {
		return nil, err
	}
	return s, nil
}

type DeleteResult struct {
	golangsdk.Result
}

type ListResult struct {
	golangsdk.Result
}

func (r ListResult) Extract() (*ListGrant, error) {
	s := new(ListGrant)
	err := r.ExtractInto(&s)
	if err != nil {
		return nil, err
	}
	return s, nil
}
