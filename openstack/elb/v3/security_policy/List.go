package security_policy

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type ListOpts struct {
	// Specifies the ID of the last record on the previous page.
	//
	// Note:
	//
	// This parameter must be used together with limit.
	//
	// If this parameter is not specified, the first page will be queried.
	//
	// This parameter cannot be left blank or set to an invalid ID.
	Marker string `q:"marker"`
	// Specifies the number of records on each page.
	//
	// Minimum: 0
	//
	// Maximum: 2000
	//
	// Default: 2000
	Limit *int `q:"limit"`
	// Specifies whether to use reverse query. Values:
	//
	// true: Query the previous page.
	//
	// false (default): Query the next page.
	//
	// Note:
	//
	// This parameter must be used together with limit.
	//
	// If page_reverse is set to true and you want to query the previous page, set the value of marker to the value of previous_marker.
	PageReverse *bool `q:"page_reverse"`
	// Specifies the ID of the custom security policy.
	//
	// Multiple IDs can be queried in the format of id=xxx&id=xxx.
	Id []string `q:"id"`
	// Specifies the name of the custom security policy.
	//
	// Multiple names can be queried in the format of name=xxx&name=xxx.
	Name []string `q:"name"`
	// Provides supplementary information about the custom security policy.
	//
	// Multiple descriptions can be queried in the format of description=xxx&description=xxx.
	Description []string `q:"description"`
	// Specifies the TLS protocols supported by the custom security policy. (Multiple protocols are separated using spaces.)
	//
	// Multiple protocols can be queried in the format of protocols=xxx&protocols=xxx.
	Protocols []string `q:"protocols"`
	// Specifies the cipher suites supported by the custom security policy. (Multiple cipher suites are separated using colons.)
	//
	// Multiple cipher suites can be queried in the format of ciphers=xxx&ciphers=xxx.
	Ciphers []string `q:"ciphers"`
}

func List(client *golangsdk.ServiceClient, opts ListOpts) pagination.Pager {
	q, err := golangsdk.BuildQueryString(&opts)
	if err != nil {
		return pagination.Pager{Err: err}
	}

	return pagination.NewPager(client, client.ServiceURL("security-policies")+q.String(), func(r pagination.PageResult) pagination.Page {
		return SecurityPage{PageWithInfo: pagination.NewPageWithInfo(r)}
	})
}

type SecurityPage struct {
	pagination.PageWithInfo
}

func (p SecurityPage) IsEmpty() (bool, error) {
	rules, err := ExtractSecurity(p)
	return len(rules) == 0, err
}

func ExtractSecurity(p pagination.Page) ([]SecurityPolicy, error) {
	var res []SecurityPolicy
	err := extract.IntoSlicePtr(p.(SecurityPage).BodyReader(), &res, "security_policies")
	return res, err
}
