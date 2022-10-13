package floatingips

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

// CreateOpts specifies a Floating IP allocation request.
type CreateOpts struct {
	// Pool is the pool of Floating IPs to allocate one from.
	Pool string `json:"pool" required:"true"`
}

// Create requests the creation of a new Floating IP.
func Create(client *golangsdk.ServiceClient, opts CreateOpts) (*FloatingIP, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL("os-floating-ips"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return extra(err, raw)
}
