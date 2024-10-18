package proxy

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func ListFlavors(client *golangsdk.ServiceClient, instanceID string) ([]ProxyFlavor, error) {
	raw, err := client.Get(client.ServiceURL("instances", instanceID, "proxy", "flavors"), nil, nil)
	if err != nil {
		return nil, err
	}

	var res []ProxyFlavor
	err = extract.IntoSlicePtr(raw.Body, &res, "proxy_flavor_groups")
	return res, err
}

type ProxyFlavor struct {
	GroupType    string        `json:"group_type"`
	ProxyFlavors []MysqlFlavor `json:"proxy_flavors"`
}

type MysqlFlavor struct {
	Vcpus    string `json:"vcpus"`
	Ram      string `json:"ram"`
	DbType   string `json:"db_type"`
	ID       string `json:"id"`
	SpecCode string `json:"spec_code"`
}
