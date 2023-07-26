package mappings

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

// List enumerates the Groups to which the current token has access.
func List(client *golangsdk.ServiceClient) pagination.Pager {
	url := listURL(client)

	return pagination.Pager{
		Client:     client,
		InitialURL: url,
		CreatePage: func(r pagination.PageResult) pagination.Page {
			return MappingPage{pagination.LinkedPageBase{PageResult: r}}
		},
	}
}

// Get retrieves details on a single Mapping, by ID.
func Get(client *golangsdk.ServiceClient, id string) (r GetResult) {
	_, r.Err = client.Get(mappingURL(client, id), &r.Body, nil)
	return
}

// CreateOptsBuilder allows extensions to add additional parameters to
// the Create request.
type CreateOptsBuilder interface {
	ToMappingCreateMap() (map[string]any, error)
}

// CreateOpts provides options used to create a mapping.
type CreateOpts struct {
	// Rules used to map federated users to local users.
	Rules []RuleOpts `json:"rules" required:"true"`
}

type RuleOpts struct {
	Local  []LocalRuleOpts  `json:"local" required:"true"`
	Remote []RemoteRuleOpts `json:"remote" required:"true"`
}

type LocalRuleOpts struct {
	User   *UserOpts  `json:"user,omitempty"`
	Group  *GroupOpts `json:"group,omitempty"`
	Groups string     `json:"groups,omitempty"`
}

type UserOpts struct {
	Name string `json:"name" required:"true"`
}

type GroupOpts struct {
	Name   string  `json:"name" required:"true"`
	Domain *Domain `json:"domain,omitempty"`
}

type Domain struct {
	Name string `json:"name,omitempty"`
	ID   string `json:"id,omitempty"`
}

type RemoteRuleOpts struct {
	Type     string   `json:"type" required:"true"`
	NotAnyOf []string `json:"not_any_of,omitempty"`
	AnyOneOf []string `json:"any_one_of,omitempty"`
	Regex    *bool    `json:"regex,omitempty"`
}

// ToMappingCreateMap formats a CreateOpts into a create request.
func (opts CreateOpts) ToMappingCreateMap() (map[string]any, error) {
	b, err := golangsdk.BuildRequestBody(opts, "mapping")
	if err != nil {
		return nil, err
	}

	return b, nil
}

// Create creates a new Mapping, by ID.
func Create(client *golangsdk.ServiceClient, mappingID string, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToMappingCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Put(mappingURL(client, mappingID), &b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{201},
	})
	return
}

// UpdateOptsBuilder allows extensions to add additional parameters to
// the Update request.
type UpdateOptsBuilder interface {
	ToMappingUpdateMap() (map[string]any, error)
}

// UpdateOpts provides options for updating a mapping.
type UpdateOpts struct {
	// Rules used to map federated users to local users.
	Rules []RuleOpts `json:"rules" required:"true"`
}

// ToMappingUpdateMap formats a UpdateOpts into an update request.
func (opts UpdateOpts) ToMappingUpdateMap() (map[string]any, error) {
	b, err := golangsdk.BuildRequestBody(opts, "mapping")
	if err != nil {
		return nil, err
	}

	return b, nil
}

// Update updates an existing Mapping, by ID.
func Update(client *golangsdk.ServiceClient, mappingID string, opts UpdateOptsBuilder) (r UpdateResult) {
	b, err := opts.ToMappingUpdateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Patch(mappingURL(client, mappingID), &b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

// Delete deletes a mapping, by ID.
func Delete(client *golangsdk.ServiceClient, mappingID string) (r DeleteResult) {
	_, r.Err = client.Delete(mappingURL(client, mappingID), nil)
	return
}
