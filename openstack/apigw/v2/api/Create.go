package api

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type CreateOpts struct {
	GatewayID           string            `json:"-"`
	GroupID             string            `json:"group_id" required:"true"`
	Name                string            `json:"name" required:"true"`
	Type                int               `json:"type" required:"true"`
	Version             string            `json:"version,omitempty"`
	ReqProtocol         string            `json:"req_protocol" required:"true"`
	ReqMethod           string            `json:"req_method" required:"true"`
	ReqUri              string            `json:"req_uri" required:"true"`
	AuthType            string            `json:"auth_type" required:"true"`
	AuthOpt             *AuthOpt          `json:"auth_opt,omitempty"`
	Cors                bool              `json:"cors,omitempty"`
	MatchMode           string            `json:"match_mode,omitempty"`
	BackendType         string            `json:"backend_type" required:"true"`
	Description         string            `json:"remark,omitempty"`
	BodyDescription     string            `json:"body_remark,omitempty"`
	ResultNormalSample  string            `json:"result_normal_sample,omitempty"`
	ResultFailureSample string            `json:"result_failure_sample,omitempty"`
	AuthorizerID        string            `json:"authorizer_id,omitempty"`
	Tags                []string          `json:"tags,omitempty"`
	RomaAppId           string            `json:"roma_app_id,omitempty"`
	DomainName          string            `json:"domain_name,omitempty"`
	ResponseID          string            `json:"response_id,omitempty"`
	ContentType         string            `json:"content_type,omitempty"`
	MockInfo            *MockInfo         `json:"mock_info,omitempty"`
	FuncInfo            *FuncInfo         `json:"func_info,omitempty"`
	ReqParams           []ReqParams       `json:"req_params,omitempty"`
	BackendParams       []BackendParams   `json:"backend_params,omitempty"`
	PolicyMocks         []PolicyMocks     `json:"policy_mocks,omitempty"`
	PolicyFunctions     []PolicyFunctions `json:"policy_functions,omitempty"`
	BackendApi          *BackendApi       `json:"backend_api,omitempty"`
	PolicyHttps         []PolicyHttps     `json:"policy_https,omitempty"`
}

type AuthOpt struct {
	AppCodeAuthType string `json:"app_code_auth_type" required:"true"`
}

type MockInfo struct {
	Description  string `json:"remark,omitempty"`
	Response     string `json:"result_content,omitempty"`
	Version      string `json:"version,omitempty"`
	AuthorizerID string `json:"authorizer_id,omitempty"`
}

type FuncInfo struct {
	FunctionUrn    string `json:"function_urn" required:"true"`
	Description    string `json:"remark,omitempty"`
	InvocationType string `json:"invocation_type" required:"true"`
	NetworkType    string `json:"network_type" required:"true"`
	Version        string `json:"version,omitempty"`
	AliasUrn       string `json:"alias_urn,omitempty"`
	Timeout        int    `json:"timeout" required:"true"`
	AuthorizerID   string `json:"authorizer_id,omitempty"`
}

type ReqParams struct {
	Name         string `json:"name" required:"true"`
	Type         string `json:"type" required:"true"`
	Location     string `json:"location" required:"true"`
	DefaultValue string `json:"default_value,omitempty"`
	SampleValue  string `json:"sample_value,omitempty"`
	Required     *int   `json:"required,omitempty"`
	ValidEnable  *int   `json:"valid_enable,omitempty"`
	Description  string `json:"remark,omitempty"`
	Enumerations string `json:"enumerations,omitempty"`
	MinNum       *int   `json:"min_num,omitempty"`
	MaxNum       *int   `json:"max_num,omitempty"`
	MinSize      *int   `json:"min_size,omitempty"`
	MaxSize      *int   `json:"max_size,omitempty"`
	PassThrough  *int   `json:"pass_through,omitempty"`
	Regular      string `json:"regular,omitempty"`
	JsonSchema   string `json:"json_schema,omitempty"`
}

type PolicyMocks struct {
	Response      string          `json:"result_content,omitempty"`
	EffectMode    string          `json:"effect_mode" required:"true"`
	Name          string          `json:"name" required:"true"`
	BackendParams []BackendParams `json:"backend_params,omitempty"`
	Conditions    []Conditions    `json:"conditions" required:"true"`
	AuthorizerID  string          `json:"authorizer_id,omitempty"`
}

type BackendParams struct {
	Origin      string `json:"origin" required:"true"`
	Name        string `json:"name" required:"true"`
	Description string `json:"remark,omitempty"`
	Location    string `json:"location" required:"true"`
	Value       string `json:"value" required:"true"`
}

type Conditions struct {
	ReqParamName    string `json:"req_param_name,omitempty"`
	ConditionType   string `json:"condition_type,omitempty"`
	ConditionOrigin string `json:"condition_origin" required:"true"`
	ConditionValue  string `json:"condition_value" required:"true"`
}

type PolicyFunctions struct {
	FunctionUrn    string          `json:"function_urn" required:"true"`
	InvocationType string          `json:"invocation_type" required:"true"`
	NetworkType    string          `json:"network_type" required:"true"`
	Version        string          `json:"version,omitempty"`
	AliasUrn       string          `json:"alias_urn,omitempty"`
	Timeout        int             `json:"timeout,omitempty"`
	EffectMode     string          `json:"effect_mode" required:"true"`
	Name           string          `json:"name" required:"true"`
	BackendParams  []BackendParams `json:"backend_params,omitempty"`
	Conditions     []Conditions    `json:"conditions" required:"true"`
	AuthorizerID   string          `json:"authorizer_id,omitempty"`
}

type BackendApi struct {
	AuthorizerID     string          `json:"authorizer_id,omitempty"`
	UrlDomain        string          `json:"url_domain,omitempty"`
	ReqProtocol      string          `json:"req_protocol" required:"true"`
	Description      string          `json:"remark,omitempty"`
	ReqMethod        string          `json:"req_method" required:"true"`
	Version          string          `json:"version,omitempty"`
	ReqUri           string          `json:"req_uri" required:"true"`
	Timeout          int             `json:"timeout" required:"true"`
	EnableClientSSL  *bool           `json:"enable_client_ssl,omitempty"`
	RetryCount       string          `json:"retry_count,omitempty"`
	VpcChannelInfo   *VpcChannelInfo `json:"vpc_channel_info,omitempty"`
	VpcChannelStatus *int            `json:"vpc_channel_status,omitempty"`
}

type VpcChannelInfo struct {
	VpcChannelProxyHost string `json:"vpc_channel_proxy_host,omitempty"`
	VpcChannelID        string `json:"vpc_channel_id" required:"true"`
}

type PolicyHttps struct {
	UrlDomain        string          `json:"url_domain,omitempty"`
	ReqProtocol      string          `json:"req_protocol" required:"true"`
	ReqMethod        string          `json:"req_method" required:"true"`
	ReqUri           string          `json:"req_uri" required:"true"`
	Timeout          *int            `json:"timeout,omitempty"`
	RetryCount       string          `json:"retry_count,omitempty"`
	EffectMode       string          `json:"effect_mode" required:"true"`
	Name             string          `json:"name" required:"true"`
	BackendParams    []BackendParams `json:"backend_params,omitempty"`
	Conditions       []Conditions    `json:"conditions" required:"true"`
	AuthorizerID     string          `json:"authorizer_id,omitempty"`
	VpcChannelInfo   *VpcChannelInfo `json:"vpc_channel_info,omitempty"`
	VpcChannelStatus *int            `json:"vpc_channel_status,omitempty"`
}

func Create(client *golangsdk.ServiceClient, opts CreateOpts) (*ApiResp, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL("apigw", "instances", opts.GatewayID, "apis"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{201},
	})
	if err != nil {
		return nil, err
	}

	var res ApiResp

	err = extract.Into(raw.Body, &res)
	return &res, err
}

type ApiResp struct {
	GroupID             string            `json:"group_id"`
	Name                string            `json:"name"`
	Type                int               `json:"type"`
	Version             string            `json:"version"`
	ReqProtocol         string            `json:"req_protocol"`
	ReqMethod           string            `json:"req_method"`
	ReqUri              string            `json:"req_uri"`
	AuthType            string            `json:"auth_type"`
	AuthOpt             *AuthOpt          `json:"auth_opt"`
	Cors                bool              `json:"cors"`
	MatchMode           string            `json:"match_mode"`
	BackendType         string            `json:"backend_type"`
	Description         string            `json:"remark"`
	BodyDescription     string            `json:"body_remark"`
	ResultNormalSample  string            `json:"result_normal_sample"`
	ResultFailureSample string            `json:"result_failure_sample"`
	AuthorizerID        string            `json:"authorizer_id"`
	Tags                []string          `json:"tags"`
	ResponseID          string            `json:"response_id"`
	RomaAppId           string            `json:"roma_app_id"`
	RomaAppName         string            `json:"roma_app_name"`
	DomainName          string            `json:"domain_name"`
	ContentType         string            `json:"content_type"`
	ID                  string            `json:"id"`
	Status              int               `json:"status"`
	ArrangeNecessary    int               `json:"arrange_necessary"`
	RegisterTime        string            `json:"register_time"`
	UpdateTime          string            `json:"update_time"`
	GroupName           string            `json:"group_name"`
	GroupVersion        string            `json:"group_version"`
	RunEnvId            string            `json:"run_env_id"`
	RunEnvName          string            `json:"run_env_name"`
	PublishID           string            `json:"publish_id"`
	PublishTime         string            `json:"publish_time"`
	CustomApiId         string            `json:"ld_api_id"`
	BackendApi          *BackendApi       `json:"backend_api"`
	ApiGroupInfo        *ApiGroupInfo     `json:"api_group_info"`
	FuncInfo            *FuncInfo         `json:"func_info"`
	MockInfo            *MockInfo         `json:"mock_info"`
	ReqParams           []ReqParams       `json:"req_params"`
	BackendParams       []BackendParams   `json:"backend_params"`
	PolicyMocks         []PolicyMocks     `json:"policy_mocks"`
	PolicyFunctions     []PolicyFunctions `json:"policy_functions"`
	PolicyHttps         []PolicyHttps     `json:"policy_https"`
}

type ApiGroupInfo struct {
	ID           string       `json:"id"`
	Name         string       `json:"name"`
	Status       int          `json:"status"`
	SlDomain     string       `json:"sl_domain"`
	RegisterTime string       `json:"register_time"`
	UpdateTime   string       `json:"update_time"`
	OnSellStatus int          `json:"on_sell_status"`
	UrlDomains   []UrlDomains `json:"url_domains"`
}

type UrlDomains struct {
	ID                  string `json:"id"`
	Domain              string `json:"domain"`
	CnameStatus         int    `json:"cname_status"`
	SslID               string `json:"ssl_id"`
	SslName             string `json:"ssl_name"`
	MinSslVersion       string `json:"min_ssl_version"`
	VfClientCertEnabled bool   `json:"verified_client_certificate_enabled"`
	HasTrustedCa        bool   `json:"is_has_trusted_root_ca"`
}
