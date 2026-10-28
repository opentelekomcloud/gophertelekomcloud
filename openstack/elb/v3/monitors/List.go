package monitors

import (
	"bytes"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type ListOpts struct {
	ID                  string   `q:"id,omitempty"`
	Name                string   `q:"name,omitempty"`
	Type                string   `q:"type,omitempty"`
	Delay               int      `q:"delay,omitempty"`
	Timeout             int      `q:"timeout,omitempty"`
	MaxRetries          int      `q:"max_retries,omitempty"`
	HTTPMethod          string   `q:"http_method,omitempty"`
	URLPath             string   `q:"url_path,omitempty"`
	ExpectedCodes       string   `q:"expected_codes,omitempty"`
	AdminStateUp        *bool    `q:"admin_state_up,omitempty"`
	EnterpriseProjectID []string `q:"enterprise_project_id,omitempty"`
	Limit               int      `q:"limit,omitempty"`
	Marker              string   `q:"marker,omitempty"`
	PageReverse         bool     `q:"page_reverse,omitempty"`
	MonitorPort         []int    `q:"monitor_port,omitempty"`
	DomainName          []string `q:"domain_name,omitempty"`
	MaxRetriesDown      []int    `q:"max_retries_down,omitempty"`

	// These filters are documented but marked unsupported by the reviewed OTC API.
	HTTPVersion        []string `q:"http_version,omitempty"`
	ContentCheckEnable *bool    `q:"content_check_enable,omitempty"`
	ContentFormat      []string `q:"content_format,omitempty"`

	// TenantID is not documented in the reviewed OTC documentation.
	TenantID string `q:"tenant_id,omitempty"`
	// ProjectID is not documented in the reviewed OTC documentation.
	ProjectID string `q:"project_id,omitempty"`
	// PoolID is not documented in the reviewed OTC documentation.
	PoolID string `q:"pool_id,omitempty"`
	// Status is not documented in the reviewed OTC documentation.
	Status string `q:"status,omitempty"`
	// SortKey is not documented in the reviewed OTC documentation.
	SortKey string `q:"sort_key,omitempty"`
	// SortDir is not documented in the reviewed OTC documentation.
	SortDir string `q:"sort_dir,omitempty"`
}

func List(client *golangsdk.ServiceClient, opts ListOpts) ([]Monitor, error) {
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints("healthmonitors").
		WithQueryParams(&opts).
		Build()
	if err != nil {
		return nil, err
	}
	pages, err := pagination.Pager{
		Client:     client,
		InitialURL: client.ServiceURL(url.String()),
		CreatePage: func(r pagination.NewPageResult) pagination.NewPage {
			return MonitorPage{NewPageResult: r}
		},
	}.NewAllPages()
	if err != nil {
		return nil, err
	}
	return ExtractMonitors(pages)
}

type MonitorPage struct {
	pagination.NewPageResult
}

func (p MonitorPage) NewNextPageURL() (string, error) {
	var res struct {
		PageInfo PageInfo `json:"page_info"`
	}
	if err := extract.Into(bytes.NewReader(p.Body), &res); err != nil {
		return "", err
	}
	marker := res.PageInfo.NextMarker
	if p.URL.Query().Get("page_reverse") == "true" {
		marker = res.PageInfo.PreviousMarker
	}
	if marker == "" {
		return "", nil
	}
	next := p.URL
	query := next.Query()
	query.Set("marker", marker)
	next.RawQuery = query.Encode()
	return next.String(), nil
}

func (p MonitorPage) NewIsEmpty() (bool, error) {
	monitors, err := ExtractMonitors(p)
	return len(monitors) == 0, err
}

func ExtractMonitors(page pagination.NewPage) ([]Monitor, error) {
	var res struct {
		Monitors []Monitor `json:"healthmonitors"`
	}
	err := extract.Into(bytes.NewReader(page.(MonitorPage).Body), &res)
	return res.Monitors, err
}

type PageInfo struct {
	PreviousMarker string `json:"previous_marker"`
	NextMarker     string `json:"next_marker"`
	CurrentCount   int    `json:"current_count"`
}
