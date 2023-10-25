package routetables

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

// UpdateOpts contains the values used when updating a route table
type UpdateOpts struct {
	Name        string                 `json:"name,omitempty"`
	Description *string                `json:"description,omitempty"`
	Routes      map[string][]RouteOpts `json:"routes,omitempty"`
}

func Update(client *golangsdk.ServiceClient, id string, opts UpdateOpts) error {
	body, err := build.RequestBody(opts, "routetable")
	if err != nil {
		return err
	}
	_, err = client.Put(client.ServiceURL(client.ProjectID, "routetables", id), body, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return err
}
