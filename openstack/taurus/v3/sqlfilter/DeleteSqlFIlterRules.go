package sqlfilter

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type DeleteSqlFilterRulesOpts struct {
	SqlFilterRules []DeleteNodeSqlFilterRuleInfo `json:"sql_filter_rules" required:"true"`
}

type DeleteNodeSqlFilterRuleInfo struct {
	NodeId string                    `json:"node_id" required:"true"`
	Rules  []DeleteNodeSqlFilterRule `json:"rules" required:"true"`
}

type DeleteNodeSqlFilterRule struct {
	SqlType  string   `json:"sql_type" required:"true"`
	Patterns []string `json:"patterns" required:"true"`
}

func DeleteSqlFilterRules(client *golangsdk.ServiceClient, instanceId string, opts DeleteSqlFilterRulesOpts) (*string, error) {
	raw, err := client.Delete(client.ServiceURL("instances", instanceId, "sql-filter", "rules"), &golangsdk.RequestOpts{
		OkCodes: []int{200},
		MoreHeaders: map[string]string{
			"Content-Type": "application/json",
		},
		JSONBody: opts,
	})
	if err != nil {
		return nil, err
	}

	var res jobResponse
	return &res.JobId, extract.Into(raw.Body, &res)
}
