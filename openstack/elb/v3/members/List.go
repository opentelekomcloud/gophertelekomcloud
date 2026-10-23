package members

import (
	"bytes"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type ListOpts struct {
	Name                []string `q:"name,omitempty"`
	Weight              []int    `q:"weight,omitempty"`
	AdminStateUp        *bool    `q:"admin_state_up,omitempty"`
	Address             []string `q:"address,omitempty"`
	ProtocolPort        []int    `q:"protocol_port,omitempty"`
	ID                  []string `q:"id,omitempty"`
	OperatingStatus     []string `q:"operating_status,omitempty"`
	EnterpriseProjectID []string `q:"enterprise_project_id,omitempty"`
	Limit               int      `q:"limit,omitempty"`
	Marker              string   `q:"marker,omitempty"`
	PageReverse         bool     `q:"page_reverse,omitempty"`
	IPVersion           []string `q:"ip_version,omitempty"`
	MemberType          []string `q:"member_type,omitempty"`
	InstanceID          []string `q:"instance_id,omitempty"`
	AvailabilityZone    []string `q:"availability_zone,omitempty"`
	SubnetCIDRID        []string `q:"subnet_cidr_id,omitempty"`

	// SubnetID is not documented in the reviewed OTC documentation.
	SubnetID []string `q:"subnet_sidr_id,omitempty"`
	// SortKey is not documented in the reviewed OTC documentation.
	SortKey string `q:"sort_key,omitempty"`
	// SortDir is not documented in the reviewed OTC documentation.
	SortDir string `q:"sort_dir,omitempty"`
}

func List(client *golangsdk.ServiceClient, poolID string, opts ListOpts) ([]Member, error) {
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints("pools", poolID, "members").
		WithQueryParams(&opts).
		Build()
	if err != nil {
		return nil, err
	}
	pages, err := pagination.Pager{
		Client:     client,
		InitialURL: client.ServiceURL(url.String()),
		CreatePage: func(r pagination.NewPageResult) pagination.NewPage {
			return MemberPage{NewPageResult: r}
		},
	}.NewAllPages()
	if err != nil {
		return nil, err
	}
	return ExtractMembers(pages)
}

type MemberPage struct {
	pagination.NewPageResult
}

func (p MemberPage) NewNextPageURL() (string, error) {
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

func (p MemberPage) NewIsEmpty() (bool, error) {
	members, err := ExtractMembers(p)
	return len(members) == 0, err
}

func ExtractMembers(page pagination.NewPage) ([]Member, error) {
	var res struct {
		Members []Member `json:"members"`
	}
	err := extract.Into(bytes.NewReader(page.(MemberPage).Body), &res)
	return res.Members, err
}
