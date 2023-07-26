package ruletypes

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

// ListRuleTypes returns the list of rule types from the server
func ListRuleTypes(c *golangsdk.ServiceClient) (result pagination.Pager) {
	return pagination.Pager{
		Client:     c,
		InitialURL: listRuleTypesURL(c),
		CreatePage: func(r pagination.PageResult) pagination.Page {
			return ListRuleTypesPage{SinglePageBase: pagination.SinglePageBase{PageResult: r}}
		},
	}
}
