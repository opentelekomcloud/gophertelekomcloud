package datastores

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

func List(client *golangsdk.ServiceClient, databasesname string) pagination.Pager {
	url := client.ServiceURL("datastores", databasesname)

	pageRdsList := pagination.NewPager(client, url, func(r pagination.PageResult) pagination.Page {
		return DataStoresPage{pagination.SinglePageBase(r)}
	})

	rdsheader := map[string]string{"Content-Type": "application/json"}
	pageRdsList.Headers = rdsheader
	return pageRdsList
}

type DataStoresResult struct {
	golangsdk.Result
}

type DataStores struct {
	//
	DataStores []dataStores `json:"dataStores" `
}

type dataStores struct {
	//
	Id string `json:"id" `
	//
	Name string `json:"name"`
}

type DataStoresPage struct {
	pagination.SinglePageBase
}

func (r DataStoresPage) IsEmpty() (bool, error) {
	data, err := ExtractDataStores(r)
	if err != nil {
		return false, err
	}
	return len(data.DataStores) == 0, err
}

func ExtractDataStores(r pagination.Page) (DataStores, error) {
	var s DataStores
	err := (r.(DataStoresPage)).ExtractInto(&s)
	return s, err
}
