package policies

import "github.com/opentelekomcloud/gophertelekomcloud"

type UpdateOpts struct {
	Name                *string               `json:"name,omitempty"`
	Description         *string               `json:"description,omitempty"`
	RedirectListenerID  string                `json:"redirect_listener_id,omitempty"`
	RedirectPoolID      string                `json:"redirect_pool_id,omitempty"`
	Rules               []Rule                `json:"rules,omitempty"`
	RedirectUrlConfig   *RedirectUrlOptions   `json:"redirect_url_config,omitempty"`
	FixedResponseConfig *FixedResponseOptions `json:"fixed_response_config,omitempty"`
	Priority            int                   `json:"priority,omitempty"`
	RedirectPoolsConfig []RedirectPoolOptions `json:"redirect_pools_config,omitempty"`
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
	_, r.Err = client.Put(client.ServiceURL("l7policies", id), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200, 201},
	})
	return
}
