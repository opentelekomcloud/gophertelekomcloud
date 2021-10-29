package policies

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/elb/v3/rules"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type Action string

const (
	ActionRedirectToPool     Action = "REDIRECT_TO_POOL"
	ActionRedirectToListener Action = "REDIRECT_TO_LISTENER"
)

type Rule struct {
	// Specifies the match content. The value can be one of the following:
	//
	//    HOST_NAME: A domain name will be used for matching.
	//    PATH: A URL will be used for matching.
	// If type is set to HOST_NAME, PATH, METHOD, or SOURCE_IP, only one forwarding rule can be created for each type.
	Type rules.RuleType `json:"type" required:"true"`

	// Specifies how requests are matched with the domain name or URL.
	//
	// If type is set to HOST_NAME, this parameter can only be set to EQUAL_TO (exact match).
	//
	// If type is set to PATH, this parameter can be set to REGEX (regular expression match), STARTS_WITH (prefix match), or EQUAL_TO (exact match).
	CompareType rules.CompareType `json:"compare_type" required:"true"`

	// Specifies the value of the match item. For example, if a domain name is used for matching, value is the domain name.
	//
	//    If type is set to HOST_NAME, the value can contain letters, digits, hyphens (-), and periods (.) and must start with a letter or digit.
	//    If you want to use a wildcard domain name, enter an asterisk (*) as the leftmost label of the domain name.
	//
	//    If type is set to PATH and compare_type to STARTS_WITH or EQUAL_TO, the value must start with a slash (/)
	//    and can contain only letters, digits, and special characters _~';@^-%#&$.*+?,=!:|/()[]{}
	Value string `json:"value" required:"true"`
}

type CreateOpts struct {
	// Specifies where requests will be forwarded. The value can be one of the following:
	//
	//    REDIRECT_TO_POOL: Requests will be forwarded to another backend server group.
	//    REDIRECT_TO_LISTENER: Requests will be redirected to an HTTPS listener.
	Action Action `json:"action" required:"true"`

	// Provides supplementary information about the forwarding policy.
	Description string `json:"description,omitempty"`

	// Specifies the ID of the listener to which the forwarding policy is added.
	ListenerID string `json:"listener_id" required:"true"`

	// Specifies the forwarding policy name.
	Name string `json:"name,omitempty"`

	// Specifies the forwarding policy priority. The value cannot be updated.
	Position int `json:"position,omitempty"`

	// Specifies the ID of the project where the forwarding policy is used.
	ProjectID string `json:"project_id,omitempty"`

	// Specifies the ID of the listener to which requests are redirected.
	// This parameter is mandatory when action is set to REDIRECT_TO_LISTENER.
	RedirectListenerID string `json:"redirect_listener_id,omitempty"`

	// Specifies the ID of the backend server group that requests are forwarded to.
	//
	// This parameter is valid and mandatory only when action is set to REDIRECT_TO_POOL.
	// The specified backend server group cannot be the default one associated with the listener,
	// or any backend server group associated with the forwarding policies of other listeners.
	RedirectPoolID string `json:"redirect_pool_id,omitempty"`

	// Lists the forwarding rules in the forwarding policy.
	// The list can contain a maximum of 10 forwarding rules (if conditions is specified, a condition is considered as a rule).
	Rules []Rule `json:"rules,omitempty"`
}

type CreateOptsBuilder interface {
	ToPolicyCreateMap() (map[string]interface{}, error)
}

func (opts CreateOpts) ToPolicyCreateMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "l7policy")
}

func Create(client *golangsdk.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToPolicyCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(baseURL(client), b, &r.Body, nil)
	return
}

type ListOptsBuilder interface {
	ToPolicyListQuery() (string, error)
}

type ListOpts struct {
	Marker             string   `q:"marker"`
	Limit              int      `q:"limit"`
	PageReverse        bool     `q:"page_reverse"`
	ID                 []string `q:"id"`
	Name               []string `q:"name"`
	Description        []string `q:"description"`
	ListenerID         []string `q:"listener_id"`
	Action             []string `q:"action"`
	RedirectPoolID     []string `q:"redirect_pool_id"`
	RedirectListenerID []string `q:"redirect_listener_id"`
	ProvisioningStatus []string `q:"provisioning_status"`
	DisplayAllRules    bool     `q:"display_all_rules"`
}

func (opts ListOpts) ToPolicyListQuery() (string, error) {
	q, err := golangsdk.BuildQueryString(opts)
	if err != nil {
		return "", err
	}
	return q.String(), nil
}

func List(client *golangsdk.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := baseURL(client)
	if opts != nil {
		q, err := opts.ToPolicyListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += q
	}

	return pagination.NewPager(client, url, func(r pagination.PageResult) pagination.Page {
		return PolicyPage{PageWithInfo: pagination.NewPageWithInfo(r)}
	})
}

func Get(client *golangsdk.ServiceClient, id string) (r GetResult) {
	_, r.Err = client.Get(resourceURL(client, id), &r.Body, nil)
	return
}

type UpdateOpts struct {
	Name               *string `json:"name,omitempty"`
	Description        *string `json:"description,omitempty"`
	RedirectListenerID string  `json:"redirect_listener_id,omitempty"`
	RedirectPoolID     string  `json:"redirect_pool_id,omitempty"`
	Rules              []Rule  `json:"rules,omitempty"`
}

type UpdateOptsBuilder interface {
	ToPolicyUpdateMap() (map[string]interface{}, error)
}

func (opts UpdateOpts) ToPolicyUpdateMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "l7policy")
}

func Update(client *golangsdk.ServiceClient, id string, opts UpdateOptsBuilder) (r UpdateResult) {
	b, err := opts.ToPolicyUpdateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Put(resourceURL(client, id), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200, 201},
	})
	return
}

func Delete(client *golangsdk.ServiceClient, id string) (r DeleteResult) {
	_, r.Err = client.Delete(resourceURL(client, id), nil)
	return
}
