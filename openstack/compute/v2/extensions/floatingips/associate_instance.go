package floatingips

import "github.com/opentelekomcloud/gophertelekomcloud"

// AssociateOpts specifies the required information to associate a Floating IP with an instance
type AssociateOpts struct {
	// FloatingIP is the Floating IP to associate with an instance.
	FloatingIP string `json:"address" required:"true"`
	// FixedIP is an optional fixed IP address of the server.
	FixedIP string `json:"fixed_address,omitempty"`
}

// AssociateInstance pairs an allocated Floating IP with a server.
func AssociateInstance(client *golangsdk.ServiceClient, serverID string, opts AssociateOpts) (err error) {
	b, err := golangsdk.BuildRequestBody(opts, "addFloatingIp")
	if err != nil {
		return
	}

	_, err = client.Post(client.ServiceURL("servers/"+serverID+"/action"), b, nil, nil)
	return
}
