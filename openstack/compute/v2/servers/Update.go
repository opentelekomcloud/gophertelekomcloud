package servers

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

// UpdateOpts specifies the base attributes that may be updated on an existing server.
type UpdateOpts struct {
	// Name changes the displayed name of the server.
	// The server host name will *not* change.
	// Server names are not constrained to be unique, even within the same tenant.
	Name string `json:"name,omitempty"`
	// AccessIPv4 provides a new IPv4 address for the instance.
	AccessIPv4 string `json:"accessIPv4,omitempty"`
	// AccessIPv6 provides a new IPv6 address for the instance.
	AccessIPv6 string `json:"accessIPv6,omitempty"`
}

// Update requests that various attributes of the indicated server be changed.
func Update(client *golangsdk.ServiceClient, id string, opts UpdateOpts) (*Server, error) {
	b, err := build.RequestBody(opts, "server")
	if err != nil {
		return nil, err
	}

	raw, err := client.Put(client.ServiceURL("servers", id), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return ExtractSer(err, raw)
}
