package routetables

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func Get(client *golangsdk.ServiceClient, id string) (*RouteTable, error) {
	raw, err := client.Get(client.ServiceURL(client.ProjectID, "routetables", id), nil, nil)
	if err != nil {
		return nil, err
	}

	var res RouteTable
	err = extract.IntoStructPtr(raw.Body, &res, "routetable")
	return &res, err
}
