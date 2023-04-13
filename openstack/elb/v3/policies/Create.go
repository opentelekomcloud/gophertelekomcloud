package policies

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/elb/v3/rules"
)

type CreateOpts struct {
	// Specifies where requests will be forwarded. The value can be one of the following:
	//
	//    REDIRECT_TO_POOL: Requests will be forwarded to another backend server group.
	//    REDIRECT_TO_LISTENER: Requests will be redirected to an HTTPS listener.
	Action Action `json:"action" required:"true"`

	// Specifies the conditions contained in a forwarding rule.
	// This parameter will take effect when enhance_l7policy_enable is set to true.
	// If conditions is specified, key and value will not take effect,
	// and the value of this parameter will contain all conditions configured for the forwarding rule.
	// The keys in the list must be the same, whereas each value must be unique.
	Conditions []rules.Condition `json:"conditions,omitempty"`

	// Provides supplementary information about the forwarding policy.
	Description string `json:"description,omitempty"`

	// Specifies the configuration of the page that will be returned.
	// This parameter will take effect when
	FixedResponseConfig *FixedResponseOptions `json:"fixed_response_config,omitempty"`

	// Specifies the ID of the listener to which the forwarding policy is added.
	ListenerID string `json:"listener_id" required:"true"`

	// Specifies the forwarding policy name.
	Name string `json:"name,omitempty"`

	// Specifies the forwarding policy priority. The value cannot be updated.
	Position int `json:"position,omitempty"`

	// Specifies the forwarding policy priority. The value cannot be updated.
	Priority int `json:"priority,omitempty"`

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

	// Specifies the URL to which requests are forwarded.
	//
	// Format: protocol://host:port/path?query
	RedirectUrl string `json:"redirect_url,omitempty"`

	// Specifies the URL to which requests are forwarded.
	// This parameter is mandatory when action is set to REDIRECT_TO_URL.
	// It cannot be specified if the value of action is not REDIRECT_TO_URL.
	RedirectUrlConfig *RedirectUrlOptions `json:"redirect_url_config,omitempty"`

	// Lists the forwarding rules in the forwarding policy.
	// The list can contain a maximum of 10 forwarding rules (if conditions is specified, a condition is considered as a rule).
	Rules []Rule `json:"rules,omitempty"`

	// Specifies the configuration of the backend server group that the requests are forwarded to. This parameter is valid only when action is set to REDIRECT_TO_POOL.
	RedirectPoolsConfig []RedirectPoolOptions `json:"redirect_pools_config,omitempty"`
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
	_, r.Err = client.Post(client.ServiceURL("l7policies"), b, &r.Body, nil)
	return
}
