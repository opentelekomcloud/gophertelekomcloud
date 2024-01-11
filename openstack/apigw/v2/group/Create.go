package group

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type CreateOpts struct {
	GatewayID   string `json:"-"`
	Name        string `json:"name" required:"true"`
	Description string `json:"remark,omitempty"`
}

func Create(client *golangsdk.ServiceClient, opts CreateOpts) (*GroupResp, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL("apigw", "instances", opts.GatewayID, "api-groups"), b,
		nil, &golangsdk.RequestOpts{
			OkCodes: []int{201},
		})
	if err != nil {
		return nil, err
	}

	var res GroupResp

	err = extract.Into(raw.Body, &res)
	return &res, err
}

type GroupResp struct {
	ID           string       `json:"id"`
	Name         string       `json:"name"`
	Status       int          `json:"status"`
	SlDomain     string       `json:"sl_domain"`
	RegisterTime string       `json:"register_time"`
	UpdateTime   string       `json:"update_time"`
	OnSellStatus int          `json:"on_sell_status"`
	UrlDomains   []UrlDomains `json:"url_domains"`
	SlDomains    []string     `json:"sl_domains"`
	Description  string       `json:"remark"`
}

type UrlDomains struct {
	DomainId            string `json:"id"`
	DomainName          string `json:"name"`
	CnameStatus         int    `json:"cname_status"`
	SslID               string `json:"ssl_id"`
	SslName             string `json:"ssl_name"`
	MinSslVersion       string `json:"min_ssl_version"`
	VfClientCertEnabled bool   `json:"verified_client_certificate_enabled"`
	HasTrustedCa        bool   `json:"is_has_trusted_root_ca"`
}
