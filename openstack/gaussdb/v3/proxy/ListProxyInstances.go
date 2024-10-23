package proxy

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ListProxyInstancesOpts struct {
	InstanceID string `json:"-"`
	Limit      int    `q:"limit"`
	Offset     int    `q:"offset"`
}

func ListProxyInstances(client *golangsdk.ServiceClient, opts ListProxyInstancesOpts) (*ListProxyInstanceResponse, error) {
	url, err := golangsdk.NewURLBuilder().WithEndpoints("instances", opts.InstanceID, "proxies").WithQueryParams(&opts).Build()
	if err != nil {
		return nil, err
	}

	raw, err := client.Get(client.ServiceURL(url.String()), nil, nil)
	if err != nil {
		return nil, err
	}

	var res ListProxyInstanceResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type ListProxyInstanceResponse struct {
	ProxyList []MysqlProxyResponse `json:"proxy_list"`
}

type MysqlProxyResponse struct {
	Proxy         *MysqlProxy        `json:"proxy"`
	MasterNode    *MysqlProxyNodeV3  `json:"master_node"`
	ReadonlyNodes []MysqlProxyNodeV3 `json:"readonly_nodes"`
}

type MysqlProxy struct {
	PoolID                  string           `json:"pool_id"`
	Status                  string           `json:"status"`
	Address                 string           `json:"address"`
	Port                    int              `json:"port"`
	PoolStatus              string           `json:"pool_status"`
	DelayThresholdInSeconds int              `json:"delay_threshold_in_seconds"`
	ElbVIP                  string           `json:"elb_vip"`
	EIP                     string           `json:"eip"`
	Vcpus                   string           `json:"vcpus"`
	Ram                     string           `json:"ram"`
	NodeNum                 int              `json:"node_num"`
	Mode                    string           `json:"mode"`
	Nodes                   []MysqlProxyNode `json:"nodes"`
	TransactionSplit        string           `json:"transaction_split"`
}

type MysqlProxyNode struct {
	ID         string `json:"id"`
	Status     string `json:"status"`
	Name       string `json:"name"`
	Role       string `json:"role"`
	AzCode     string `json:"az_code"`
	FrozenFlag int    `json:"frozen_flag"`
}

type MysqlProxyNodeV3 struct {
	ID               string               `json:"id"`
	InstanceId       string               `json:"instance_id"`
	Status           string               `json:"status"`
	Name             string               `json:"name"`
	Weight           int                  `json:"weight"`
	AvailabilityZone *MysqlProxyAvailable `json:"availability_zone"`
}

type MysqlProxyAvailable struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}
