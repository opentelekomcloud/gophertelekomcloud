package proxy

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func GetFlavors(client *golangsdk.ServiceClient, instanceId string) ([]ProxyFlavorGroups, error) {
	raw, err := client.Get(client.ServiceURL("instances", instanceId, "proxy", "flavors"), nil, nil)
	if err != nil {
		return nil, err
	}

	var res []ProxyFlavorGroups
	err = extract.IntoSlicePtr(raw.Body, &res, "proxy_flavor_groups")
	return res, err
}

type ProxyFlavorGroups struct {
	GroupType    string                    `json:"group_type"`
	ProxyFlavors []MysqlProxyComputeFlavor `json:"proxy_flavors"`
}

type MysqlProxyComputeFlavor struct {
	Vcpus    string            `json:"vcpus"`
	Ram      string            `json:"ram"`
	DbType   string            `json:"db_type"`
	Id       string            `json:"id"`
	SpecCode string            `json:"spec_code"`
	AzStatus map[string]string `json:"az_status"`
}
