package sqlfilter

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type UpdateSqlFilterRulesOpts struct {
	SqlFilterRules []NodeSqlFilterRuleInfo `json:"sql_filter_rules" required:"true"`
}

type NodeSqlFilterRuleInfo struct {
	NodeId string              `json:"node_id" required:"true"`
	Rules  []NodeSqlFilterRule `json:"rules" required:"true"`
}

type NodeSqlFilterRule struct {
	SqlType  string                     `json:"sql_type" required:"true"`
	Patterns []NodeSqlFilterRulePattern `json:"patterns" required:"true"`
}

type NodeSqlFilterRulePattern struct {
	Pattern        string `json:"pattern" required:"true"`
	MaxConcurrency int    `json:"max_concurrency" required:"true"`
}

func UpdateSqlFilterRules(client *golangsdk.ServiceClient, instanceId string, opts UpdateSqlFilterRulesOpts) (*string, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Put(client.ServiceURL("instances", instanceId, "sql-filter", "rules"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res jobResponse
	return &res.JobId, extract.Into(raw.Body, &res)
}
