package flow_logs

import (
	"bytes"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type ListOpts struct {
	ID           string `q:"id"`
	Name         string `q:"name"`
	TenantID     string `q:"tenant_id"`
	Description  string `q:"description"`
	ResourceType string `q:"resource_type"`
	ResourceID   string `q:"resource_id"`
	TrafficType  string `q:"traffic_type"`
	LogGroupID   string `q:"log_group_id"`
	LogTopicID   string `q:"log_topic_id"`
	Status       string `q:"status"`
	Limit        *int   `q:"limit"`
	Marker       string `q:"marker"`
}

func List(client *golangsdk.ServiceClient, opts ListOpts) ([]FlowLog, error) {
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints(client.ProjectID, "fl", "flow_logs").
		WithQueryParams(&opts).
		Build()
	if err != nil {
		return nil, err
	}

	pages, err := pagination.Pager{
		Client:     client,
		InitialURL: client.ServiceURL(url.String()),
		CreatePage: func(r pagination.NewPageResult) pagination.NewPage {
			return FlowLogPage{NewPageResult: r}
		},
	}.NewAllPages()
	if err != nil {
		return nil, err
	}
	return ExtractFlowLogs(pages)
}

type FlowLogPage struct {
	pagination.NewPageResult
}

func (p FlowLogPage) NewNextPageURL() (string, error) {
	flowLogs, err := ExtractFlowLogs(p)
	if err != nil {
		return "", err
	}
	if len(flowLogs) == 0 {
		return "", nil
	}

	q := p.URL.Query()
	q.Set("marker", flowLogs[len(flowLogs)-1].ID)
	nextURL := p.URL
	nextURL.RawQuery = q.Encode()
	return nextURL.String(), nil
}

func (p FlowLogPage) NewIsEmpty() (bool, error) {
	flowLogs, err := ExtractFlowLogs(p)
	return len(flowLogs) == 0, err
}

func ExtractFlowLogs(page pagination.NewPage) ([]FlowLog, error) {
	var res struct {
		FlowLogs []FlowLog `json:"flow_logs"`
	}
	err := extract.Into(bytes.NewReader(page.(FlowLogPage).Body), &res)
	return res.FlowLogs, err
}
