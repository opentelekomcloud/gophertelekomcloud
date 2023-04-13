package rules

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type ListOptsBuilder interface {
	ToRuleListQuery() (string, error)
}

type ListOpts struct {
	ID          []string      `q:"id"`
	CompareType []CompareType `q:"compare_type"`
	Value       []string      `q:"value"`
	Type        []RuleType    `q:"type"`

	Limit       int    `q:"limit"`
	Marker      string `q:"marker"`
	PageReverse bool   `q:"page_reverse"`
}

func List(client *golangsdk.ServiceClient, policyID string, opts ListOptsBuilder) pagination.Pager {
	url := client.ServiceURL("l7policies", policyID, "rules")

	if opts != nil {
		q, err := opts.ToRuleListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += q
	}

	return pagination.NewPager(client, url, func(r pagination.PageResult) pagination.Page {
		return RulePage{PageWithInfo: pagination.NewPageWithInfo(r)}
	})
}
