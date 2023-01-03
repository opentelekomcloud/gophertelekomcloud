package repositories

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
)

func Get(client *golangsdk.ServiceClient, organization, repository string) (r GetResult) {
	_, r.Err = client.Get(client.ServiceURL("manage", "namespaces", organization, "repos", repository), &r.Body, nil)
	return
}

type UpdateOptsBuilder interface {
	ToRepositoryUpdateMap() (map[string]interface{}, error)
}

type UpdateOpts struct {
	Category    string `json:"category,omitempty"`
	Description string `json:"description,omitempty"`
	IsPublic    bool   `json:"is_public"` // this is mandatory field, so no pointers
}

func (opts UpdateOpts) ToRepositoryUpdateMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "")
}

func Update(client *golangsdk.ServiceClient, organization, repository string, opts UpdateOptsBuilder) (r UpdateResult) {
	b, err := opts.ToRepositoryUpdateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Patch(client.ServiceURL("manage", "namespaces", organization, "repos", repository), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{201},
	})
	return
}
