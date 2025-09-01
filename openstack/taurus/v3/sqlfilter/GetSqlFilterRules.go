package sqlfilter

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type GetSqlFilterRulesOpts struct {
	NodeId string `q:"node_id" required:"true"`
	Type   string `q:"type,omitempty"`
}

func GetSqlFilterRules(client *golangsdk.ServiceClient, instanceId string, opts GetSqlFilterRulesOpts) (*SqlFilterRulesResponse, error) {
	url := client.ServiceURL("instances", instanceId, "sql-filter", "rules")

	query, err := golangsdk.BuildQueryString(opts)
	if err != nil {
		return nil, err
	}
	url += query.String()

	raw, err := client.Get(url, nil, nil)
	if err != nil {
		return nil, err
	}

	var res SqlFilterRulesResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type SqlFilterRulesResponse struct {
	NodeId         string          `json:"node_id"`
	SqlFilterRules []SqlFilterRule `json:"sql_filter_rules"`
}

type SqlFilterRule struct {
	SqlType  string                 `json:"sql_type"`
	Patterns []SqlFilterRulePattern `json:"patterns"`
}

type SqlFilterRulePattern struct {
	Pattern        string `json:"pattern"`
	MaxConcurrency int    `json:"max_concurrency"`
}
