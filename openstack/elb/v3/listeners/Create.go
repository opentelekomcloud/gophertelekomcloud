package listeners

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
)

type CreateOpts struct {
	AdminStateUp                     *bool                           `json:"admin_state_up,omitempty"`
	CAContainerRef                   string                          `json:"client_ca_tls_container_ref,omitempty"`
	DefaultPoolID                    string                          `json:"default_pool_id,omitempty"`
	DefaultTlsContainerRef           string                          `json:"default_tls_container_ref,omitempty"`
	Description                      string                          `json:"description,omitempty"`
	Http2Enable                      *bool                           `json:"http2_enable,omitempty"`
	LoadbalancerID                   string                          `json:"loadbalancer_id" required:"true"`
	Name                             string                          `json:"name,omitempty"`
	ProjectID                        string                          `json:"project_id,omitempty"`
	Protocol                         Protocol                        `json:"protocol" required:"true"`
	ProtocolPort                     int                             `json:"protocol_port" required:"true"`
	SniContainerRefs                 []string                        `json:"sni_container_refs,omitempty"`
	SniMatchAlgo                     string                          `json:"sni_match_algo,omitempty"`
	Tags                             []tags.ResourceTag              `json:"tags,omitempty"`
	TlsCiphersPolicy                 string                          `json:"tls_ciphers_policy,omitempty"`
	SecurityPolicy                   string                          `json:"security_policy_id,omitempty"`
	EnableMemberRetry                *bool                           `json:"enable_member_retry,omitempty"`
	KeepAliveTimeout                 int                             `json:"keepalive_timeout,omitempty"`
	ClientTimeout                    int                             `json:"client_timeout,omitempty"`
	MemberTimeout                    int                             `json:"member_timeout,omitempty"`
	IpGroup                          *IpGroup                        `json:"ipgroup,omitempty"`
	InsertHeaders                    *InsertHeaders                  `json:"insert_headers,omitempty"`
	TransparentClientIP              *bool                           `json:"transparent_client_ip_enable,omitempty"`
	EnhanceL7policy                  *bool                           `json:"enhance_l7policy_enable,omitempty"`
	ProtectionStatus                 string                          `json:"protection_status,omitempty"`
	ProtectionReason                 string                          `json:"protection_reason,omitempty"`
	AccessLogCustomizedHeadersConfig *AccessLogCustomizedHeadersOpts `json:"access_log_customized_headers_config,omitempty"`
}

type AccessLogCustomizedHeadersOpts struct {
	Enable         *bool    `json:"enable,omitempty"`
	IncludeHeaders []string `json:"include_headers,omitempty"`
	ExcludeHeaders []string `json:"exclude_headers,omitempty"`
}

func Create(client *golangsdk.ServiceClient, opts CreateOpts) (*Listener, error) {
	b, err := build.RequestBody(opts, "listener")
	if err != nil {
		return nil, err
	}
	raw, err := client.Post(client.ServiceURL("listeners"), b, nil, &golangsdk.RequestOpts{OkCodes: []int{201}})
	if err != nil {
		return nil, err
	}
	var res struct {
		Listener Listener `json:"listener"`
	}
	if err = extract.Into(raw.Body, &res); err != nil {
		return nil, err
	}
	return &res.Listener, nil
}
