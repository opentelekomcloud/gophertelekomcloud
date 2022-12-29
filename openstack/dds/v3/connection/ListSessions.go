package connection

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ListSessionOpts struct {
	PlanSummary string `q:"plan_summary"`
	Type        string `q:"type"`
	NameSpace   string `q:"namespace"`
	CostTime    int    `q:"cost_time"`
}

func ListSessions(client *golangsdk.ServiceClient, nodeId string, opts ListSessionOpts) (*ListSessionsResponse, error) {
	q, err := golangsdk.BuildQueryString(opts)
	if err != nil {
		return nil, err
	}

	raw, err := client.Get(client.ServiceURL("nodes", nodeId, "sessions")+q.String(), nil, nil)
	if err != nil {
		return nil, err
	}

	var res ListSessionsResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type ListSessionsResponse struct {
	Sessions   []SessionResponse `json:"sessions"`
	TotalCount int               `json:"total_count"`
}

type SessionResponse struct {
	Id          string `json:"id"`
	Active      bool   `json:"active"`
	Operation   string `json:"operation"`
	Type        string `json:"type"`
	CostTime    string `json:"cost_time"`
	PlanSummary string `json:"plan_summary"`
	Host        string `json:"engine_versions"`
	Client      string `json:"client"`
	Description string `json:"description"`
	Namespace   string `json:"namespace"`
}
