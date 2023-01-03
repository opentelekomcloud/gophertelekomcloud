package organizations

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type ListOptsBuilder interface {
	ToNamespaceListQuery() (string, error)
}

type ListOpts struct {
	Namespace string `q:"namespace"`
}

func (opts ListOpts) ToNamespaceListQuery() (string, error) {
	q, err := golangsdk.BuildQueryString(opts)
	if err != nil {
		return "", err
	}
	return q.String(), err
}

func List(client *golangsdk.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := client.ServiceURL("manage", "namespaces")
	if opts != nil {
		q, err := opts.ToNamespaceListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += q
	}
	return pagination.NewPager(client, url, func(r pagination.PageResult) pagination.Page {
		return OrganizationPage{pagination.SinglePageBase(r)}
	})
}

func Get(client *golangsdk.ServiceClient, id string) (r GetResult) {
	_, r.Err = client.Get(client.ServiceURL("manage", "namespaces", id), &r.Body, nil)
	return
}

type CreatePermissionsOptsBuilder interface {
	ToPermissionCreateMap() (map[string]interface{}, error)
}

type CreatePermissionsOpts Auth

func (opts CreatePermissionsOpts) ToPermissionCreateMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "")
}

func CreatePermissions(client *golangsdk.ServiceClient, organization string, opts CreatePermissionsOptsBuilder) (r CreatePermissionsResult) {
	b, err := opts.ToPermissionCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	realBody := []interface{}{b}
	_, r.Err = client.Post(client.ServiceURL("manage", "namespaces", organization, "access"), realBody, &r.Body, nil)
	return
}

func DeletePermissions(client *golangsdk.ServiceClient, organization string, userID string) (r DeletePermissionsResult) {
	_, r.Err = client.Request("DELETE", client.ServiceURL("manage", "namespaces", organization, "access"), &golangsdk.RequestOpts{
		JSONBody: []interface{}{userID},
	})
	return
}

type UpdatePermissionsOptsBuilder interface {
	ToPermissionUpdateMap() (map[string]interface{}, error)
}

type UpdatePermissionsOpts Auth

func (opts UpdatePermissionsOpts) ToPermissionUpdateMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "")
}

func UpdatePermissions(client *golangsdk.ServiceClient, organization string, opts UpdatePermissionsOptsBuilder) (r UpdatePermissionsResult) {
	b, err := opts.ToPermissionUpdateMap()
	if err != nil {
		r.Err = err
		return
	}
	realBody := []interface{}{b}
	_, r.Err = client.Patch(client.ServiceURL("manage", "namespaces", organization, "access"), realBody, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{201},
	})
	return
}

func GetPermissions(client *golangsdk.ServiceClient, organization string) (r GetPermissionsResult) {
	_, r.Err = client.Get(client.ServiceURL("manage", "namespaces", organization, "access"), &r.Body, nil)
	return
}
