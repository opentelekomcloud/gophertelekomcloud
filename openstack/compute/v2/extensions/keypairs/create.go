package keypairs

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/compute/v2/servers"
)

// CreateOptsExt adds a KeyPair option to the base CreateOpts.
type CreateOptsExt struct {
	servers.CreateOpts
	// KeyName is the name of the key pair.
	KeyName string `json:"key_name,omitempty"`
}

// ToServerCreateMap adds the key_name to the base server creation options.
func (opts CreateOptsExt) ToServerCreateMap() (map[string]interface{}, error) {
	base, err := opts.CreateOpts.ToServerCreateMap()
	if err != nil {
		return nil, err
	}

	if opts.KeyName == "" {
		return base, nil
	}

	serverMap := base["server"].(map[string]interface{})
	serverMap["key_name"] = opts.KeyName

	return base, nil
}

// CreateOpts specifies KeyPair creation or import parameters.
type CreateOpts struct {
	// Name is a friendly name to refer to this KeyPair in other services.
	Name string `json:"name" required:"true"`
	// PublicKey [optional] is a pregenerated OpenSSH-formatted public key.
	// If provided, this key will be imported and no new key will be created.
	PublicKey string `json:"public_key,omitempty"`
}

// Create requests the creation of a new KeyPair on the server, or to import a pre-existing keypair.
func Create(client *golangsdk.ServiceClient, opts CreateOpts) (*KeyPair, error) {
	b, err := build.RequestBody(opts, "keypair")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL("os-keypairs"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return extra(err, raw)
}
