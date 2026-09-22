package pools

import (
	"bytes"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type ListOpts struct {
	Marker                         string   `q:"marker,omitempty"`
	Limit                          int      `q:"limit,omitempty"`
	PageReverse                    bool     `q:"page_reverse,omitempty"`
	Description                    []string `q:"description,omitempty"`
	HealthMonitorID                []string `q:"healthmonitor_id,omitempty"`
	LBMethod                       []string `q:"lb_algorithm,omitempty"`
	Protocol                       []string `q:"protocol,omitempty"`
	AdminStateUp                   *bool    `q:"admin_state_up,omitempty"`
	Name                           []string `q:"name,omitempty"`
	ID                             []string `q:"id,omitempty"`
	LoadbalancerID                 []string `q:"loadbalancer_id,omitempty"`
	EnterpriseProjectID            []string `q:"enterprise_project_id,omitempty"`
	IPVersion                      []string `q:"ip_version,omitempty"`
	MemberAddress                  []string `q:"member_address,omitempty"`
	MemberDeviceID                 []string `q:"member_device_id,omitempty"`
	MemberDeletionProtectionEnable *bool    `q:"member_deletion_protection_enable,omitempty"`
	ListenerID                     []string `q:"listener_id,omitempty"`
	MemberInstanceID               []string `q:"member_instance_id,omitempty"`
	VpcID                          []string `q:"vpc_id,omitempty"`
	Type                           []string `q:"type,omitempty"`
	ProtectionStatus               []string `q:"protection_status,omitempty"`

	// SortKey is not documented in the reviewed OTC documentation.
	SortKey string `q:"sort_key,omitempty"`
	// SortDir is not documented in the reviewed OTC documentation.
	SortDir string `q:"sort_dir,omitempty"`
}

func List(client *golangsdk.ServiceClient, opts ListOpts) ([]Pool, error) {
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints("pools").
		WithQueryParams(&opts).
		Build()
	if err != nil {
		return nil, err
	}
	pages, err := pagination.Pager{
		Client:     client,
		InitialURL: client.ServiceURL(url.String()),
		CreatePage: func(r pagination.NewPageResult) pagination.NewPage {
			return PoolPage{NewPageResult: r}
		},
	}.NewAllPages()
	if err != nil {
		return nil, err
	}
	return ExtractPools(pages)
}

type PoolPage struct {
	pagination.NewPageResult
}

func (p PoolPage) NewNextPageURL() (string, error) {
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

func (p PoolPage) NewIsEmpty() (bool, error) {
	pools, err := ExtractPools(p)
	return len(pools) == 0, err
}

func ExtractPools(page pagination.NewPage) ([]Pool, error) {
	var res struct {
		Pools []Pool `json:"pools"`
	}
	err := extract.Into(bytes.NewReader(page.(PoolPage).Body), &res)
	return res.Pools, err
}

type PageInfo struct {
	PreviousMarker string `json:"previous_marker"`
	NextMarker     string `json:"next_marker"`
	CurrentCount   int    `json:"current_count"`
}
