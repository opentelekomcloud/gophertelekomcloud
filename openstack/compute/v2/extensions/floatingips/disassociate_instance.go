package floatingips

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

// DisassociateOpts specifies the required information to disassociate a
// Floating IP with a server.
type DisassociateOpts struct {
	FloatingIP string `json:"address" required:"true"`
}

// DisassociateInstance decouples an allocated Floating IP from an instance
func DisassociateInstance(client *golangsdk.ServiceClient, serverID string, opts DisassociateOpts) (err error) {
	b, err := build.RequestBody(opts, "removeFloatingIp")
	if err != nil {
		return
	}

	_, err = client.Post(client.ServiceURL("servers/"+serverID+"/action"), b, nil, nil)
	return
}
