package policies

import (
	"fmt"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"net/url"
	"reflect"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
)

var RequestOpts = golangsdk.RequestOpts{
	MoreHeaders: map[string]string{"Content-Type": "application/json;charset=utf8"},
}

// ListOpts Page&PerPage options currently don't have any effect
type ListOpts struct {
	DisplayName string
	Type        string
	ID          string
	Page        int `q:"page"`
	PerPage     int `q:"per_page"`
}

func (opts ListOpts) ToPolicyListQuery() (string, error) {
	var opts2 interface{} = opts
	q, err := build.QueryString(opts2)
	if err != nil {
		return "", err
	}
	return q.String(), err
}

func List(client *golangsdk.ServiceClient, opts ListOpts) (ListPolicy, error) {
	var r ListResult
	_, r.Err = client.Get(listURL(client), &r.Body, &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: RequestOpts.MoreHeaders, JSONBody: nil,
	})

	allPolicies, err := r.ExtractPolicies()
	if err != nil {
		return allPolicies, err
	}

	return FilterRoles(allPolicies, opts)
}

func FilterRoles(roles ListPolicy, opts ListOpts) (ListPolicy, error) {
	refinedListPolicy := ListPolicy{
		Links: roles.Links,
	}
	var matched bool
	m := map[string]interface{}{}

	if opts.DisplayName != "" {
		m["DisplayName"] = opts.DisplayName
	}

	if opts.Type != "" {
		m["Type"] = opts.Type
	}

	if opts.Type != "" {
		m["ID"] = opts.ID
	}

	if len(m) > 0 && len(roles.Roles) > 0 {

		for _, role := range roles.Roles {
			matched = true

			for key, value := range m {
				if sVal := getStructField(&role, key); !(sVal == value) {
					matched = false
				}
			}

			if matched {
				refinedListPolicy.Roles = append(refinedListPolicy.Roles, role)
			}
		}
	} else {
		refinedListPolicy.Roles = roles.Roles
	}

	refinedListPolicy.TotalNumber = len(refinedListPolicy.Roles)

	return refinedListPolicy, nil

}

func getStructField(v *Policy, field string) string {
	r := reflect.ValueOf(v)
	f := reflect.Indirect(r).FieldByName(field)
	return f.String()
}

func Get(client *golangsdk.ServiceClient, id string) (r GetResult) {
	_, r.Err = client.Get(getURL(client, id), &r.Body, &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: RequestOpts.MoreHeaders, JSONBody: nil,
	})
	return
}

type CreateOpts struct {
	// Display name of the custom policy
	DisplayName string `json:"display_name" required:"true"`
	// Display mode which can only be AX or XA
	// AX: Account level.
	// XA: Project level.
	Type string `json:"type" required:"true"`
	// Description of the custom policy
	Description string `json:"description" required:"true"`
	// DescriptionCn string `json:"description_cn"`
	// Content of custom policy.
	Policy CreatePolicy `json:"policy" required:"true"`
}

type CreatePolicy struct {
	// Policy version. When creating a custom policy, set this parameter to 1.1
	Version string `json:"Version" required:"true"`
	// Statement of the policy
	// A policy can contain a maximum of eight statements.
	Statement []CreateStatement `json:"Statement" required:"true"`
}

type CreateStatement struct {
	// Specific operation permission on a resource. A maximum of 100 actions are allowed.
	// The value format is `Service name:Resource type:Operation`, for example, `vpc:ports:create`
	// `Service name`: indicates the product name, such as ecs, evs, or vpc. Only lowercase letters are allowed.
	// Resource types and operations are not case-sensitive. You can use an asterisk (*) to represent all operations.
	Action []string `json:"Action" required:"true"`
	// Effect of the permission. The value can be `Allow` or `Deny`.
	// If both `Allow` and `Deny` statements are found in a policy,
	// the authentication starts from the `Deny` statements.
	Effect string `json:"Effect" required:"true"`
	// Conditions for the permission to take effect. A maximum of 10 conditions are allowed.
	// Conditions can't be used if `Resource` is selected using delegating agencies.
	Condition Condition `json:"Condition,omitempty"`
	// Cloud resource. The array can contain a maximum of 10 resource strings - []string,
	// and each string cannot exceed 128 characters.
	// Format: ::::. For example, `obs:::bucket:*`. Asterisks are allowed.
	// In the case of a custom policy for agencies, the type of this parameter is map[string][]string
	// and the value should be set to
	// "Resource": {"uri": ["/iam/agencies/07805acaba800fdd4fbdc00b8f888c7c"]}.
	Resource interface{} `json:"Resource,omitempty"`
}

type Condition map[string]map[string][]string

func (opts CreateOpts) ToPolicyCreateMap() (map[string]interface{}, error) {
	b, err := golangsdk.BuildRequestBody(opts, "role")
	if err != nil {
		return nil, err
	}
	return b, nil
}

type CreateOptsBuilder interface {
	ToPolicyCreateMap() (map[string]interface{}, error)
}

func Create(client *golangsdk.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	reqBody, err := opts.ToPolicyCreateMap()
	if err != nil {
		r.Err = fmt.Errorf("failed to create policy create map: %s", err)
		return
	}
	_, err = client.Post(createURL(client), reqBody, &r.Body, &golangsdk.RequestOpts{
		OkCodes:     []int{201},
		MoreHeaders: RequestOpts.MoreHeaders, JSONBody: nil,
	})
	r.Err = err
	return
}

func (opts CreateOpts) ToPolicyUpdateMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "role")
}

func Update(c *golangsdk.ServiceClient, id string, opts CreateOpts) (r UpdateResult) {
	b, err := opts.ToPolicyUpdateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = c.Patch(updateURL(c, id), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: RequestOpts.MoreHeaders, JSONBody: nil,
	})
	return
}

func Delete(client *golangsdk.ServiceClient, id string) (r DeleteResult) {
	_, r.Err = client.Delete(deleteURL(client, id), &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: RequestOpts.MoreHeaders, JSONBody: nil,
	})
	return
}
