package proxy

import (
	"bytes"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type ListOpts struct {
	Offset int `q:"offset"`
	Limit  int `q:"limit"`
}

func List(client *golangsdk.ServiceClient, instanceId string, opts ListOpts) ([]ProxyInstanceResponse, error) {
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints("instances", instanceId, "proxies").
		WithQueryParams(&opts).Build()
	if err != nil {
		return nil, err
	}

	pages, err := pagination.Pager{
		Client:     client,
		InitialURL: client.ServiceURL(url.String()),
		CreatePage: func(r pagination.NewPageResult) pagination.NewPage {
			return ProxyPage{NewSinglePageBase: pagination.NewSinglePageBase{NewPageResult: r}}
		},
	}.NewAllPages()

	if err != nil {
		return nil, err
	}
	return ExtractProxies(pages)
}

func ExtractProxies(r pagination.NewPage) ([]ProxyInstanceResponse, error) {
	var s struct {
		ProxyList []ProxyInstanceResponse `json:"proxy_list"`
	}
	err := extract.Into(bytes.NewReader((r.(ProxyPage)).Body), &s)
	if err != nil {
		return nil, err
	}

	return s.ProxyList, nil
}

type ProxyPage struct {
	pagination.NewSinglePageBase
}

type ProxyInstanceResponse struct {
	Proxy         MysqlProxyV3       `json:"proxy"`
	MasterNode    MysqlProxyNodeV3   `json:"master_node"`
	ReadonlyNodes []MysqlProxyNodeV3 `json:"readonly_nodes"`
}

type MysqlProxyV3 struct {
	PoolId                  string            `json:"pool_id"`
	Status                  string            `json:"status"`
	Address                 string            `json:"address"`
	Port                    int               `json:"port"`
	PoolStatus              string            `json:"pool_status"`
	DelayThresholdInSeconds int               `json:"delay_threshold_in_seconds"`
	ElbVip                  string            `json:"elb_vip"`
	Eip                     string            `json:"eip"`
	Vcpus                   string            `json:"vcpus"`
	Ram                     string            `json:"ram"`
	NodeNum                 int               `json:"node_num"`
	Mode                    string            `json:"mode"`
	Nodes                   []MysqlProxyNodes `json:"nodes"`
	FlavorRef               string            `json:"flavor_ref"`
	Name                    string            `json:"name"`
	TransactionSplit        string            `json:"transaction_split"`
}

type MysqlProxyNodes struct {
	Id         string `json:"id"`
	Status     string `json:"status"`
	Name       string `json:"name"`
	Role       string `json:"role"`
	AzCode     string `json:"az_code"`
	FrozenFlag int    `json:"frozen_flag"`
}

type MysqlProxyNodeV3 struct {
	Id               string                `json:"id"`
	InstanceId       string                `json:"instance_id"`
	Status           string                `json:"status"`
	Name             string                `json:"name"`
	Weight           int                   `json:"weight"`
	AvailabilityZone []MysqlProxyAvailable `json:"availability_zone"`
}

type MysqlProxyAvailable struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}
