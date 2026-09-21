package listeners

import (
	"bytes"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type ListOpts struct {
	Limit                   int        `q:"limit,omitempty"`
	Marker                  string     `q:"marker,omitempty"`
	PageReverse             bool       `q:"page_reverse,omitempty"`
	ProtocolPort            []int      `q:"protocol_port,omitempty"`
	Protocol                []Protocol `q:"protocol,omitempty"`
	Description             []string   `q:"description,omitempty"`
	DefaultTLSContainerRef  []string   `q:"default_tls_container_ref,omitempty"`
	ClientCATLSContainerRef []string   `q:"client_ca_tls_container_ref,omitempty"`
	AdminStateUp            *bool      `q:"admin_state_up,omitempty"`
	ConnectionLimit         []int      `q:"connection_limit,omitempty"`
	DefaultPoolID           []string   `q:"default_pool_id,omitempty"`
	ID                      []string   `q:"id,omitempty"`
	Name                    []string   `q:"name,omitempty"`
	Http2Enable             *bool      `q:"http2_enable,omitempty"`
	LoadBalancerID          []string   `q:"loadbalancer_id,omitempty"`
	TLSCiphersPolicy        []string   `q:"tls_ciphers_policy,omitempty"`
	MemberAddress           []string   `q:"member_address,omitempty"`
	MemberDeviceID          []string   `q:"member_device_id,omitempty"`
	EnterpriseProjectID     []string   `q:"enterprise_project_id,omitempty"`
	EnableMemberRetry       *bool      `q:"enable_member_retry,omitempty"`
	MemberTimeout           []int      `q:"member_timeout,omitempty"`
	ClientTimeout           []int      `q:"client_timeout,omitempty"`
	KeepAliveTimeout        []int      `q:"keepalive_timeout,omitempty"`
	TransparentClientIP     *bool      `q:"transparent_client_ip_enable,omitempty"`
	EnhanceL7policy         *bool      `q:"enhance_l7policy_enable,omitempty"`
	MemberInstanceID        []string   `q:"member_instance_id,omitempty"`
	ProtectionStatus        []string   `q:"protection_status,omitempty"`
}

func List(client *golangsdk.ServiceClient, opts ListOpts) ([]Listener, error) {
	url, err := golangsdk.NewURLBuilder().WithEndpoints("listeners").WithQueryParams(&opts).Build()
	if err != nil {
		return nil, err
	}
	pages, err := pagination.Pager{
		Client:     client,
		InitialURL: client.ServiceURL(url.String()),
		CreatePage: func(r pagination.NewPageResult) pagination.NewPage {
			return ListenerPage{NewPageResult: r}
		},
	}.NewAllPages()
	if err != nil {
		return nil, err
	}
	return ExtractListeners(pages)
}

type ListenerPage struct {
	pagination.NewPageResult
}

func (p ListenerPage) NewNextPageURL() (string, error) {
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

func (p ListenerPage) NewIsEmpty() (bool, error) {
	listeners, err := ExtractListeners(p)
	return len(listeners) == 0, err
}

func ExtractListeners(page pagination.NewPage) ([]Listener, error) {
	var res struct {
		Listeners []Listener `json:"listeners"`
	}
	err := extract.Into(bytes.NewReader(page.(ListenerPage).Body), &res)
	return res.Listeners, err
}

type PageInfo struct {
	PreviousMarker string `json:"previous_marker"`
	NextMarker     string `json:"next_marker"`
	CurrentCount   int    `json:"current_count"`
}
