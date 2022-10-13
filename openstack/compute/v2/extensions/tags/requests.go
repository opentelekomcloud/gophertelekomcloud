package tags

import (
	"net/http"

	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type Tags struct {
	// Tags is a list of any tags. Tags are arbitrarily defined strings attached to a resource.
	Tags []string `json:"tags"`
}

// Create implements create tags request
func Create(client *golangsdk.ServiceClient, serverId string, opts Tags) (*Tags, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Put(client.ServiceURL("servers", serverId, "tags"), b, nil, &golangsdk.RequestOpts{OkCodes: []int{200}})
	return extra(err, raw)
}

// Get implements tags get request
func Get(client *golangsdk.ServiceClient, serverId string) (*Tags, error) {
	raw, err := client.Get(client.ServiceURL("servers", serverId, "tags"), nil, nil)
	return extra(err, raw)
}

// Delete implements image delete request
func Delete(client *golangsdk.ServiceClient, serverId string) (err error) {
	_, err = client.Delete(client.ServiceURL("servers", serverId, "tags"), nil)
	return
}

func extra(err error, raw *http.Response) (*Tags, error) {
	if err != nil {
		return nil, err
	}

	var res Tags
	err = extract.Into(raw.Body, &res)
	return &res, err
}
