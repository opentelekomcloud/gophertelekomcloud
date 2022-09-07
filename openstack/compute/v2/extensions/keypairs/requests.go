package keypairs

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/compute/v2/servers"
)

// CreateOptsExt adds a KeyPair option to the base CreateOpts.
type CreateOptsExt struct {
	servers.CreateOptsBuilder

	// KeyName is the name of the key pair.
	KeyName string `json:"key_name,omitempty"`
}

// ToServerCreateMap adds the key_name to the base server creation options.
func (opts CreateOptsExt) ToServerCreateMap() (map[string]interface{}, error) {
	base, err := opts.CreateOptsBuilder.ToServerCreateMap()
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

// CreateOptsBuilder allows extensions to add additional parameters to the
// Create request.
type CreateOptsBuilder interface {
	ToKeyPairCreateMap() (map[string]interface{}, error)
}

// CreateOpts specifies KeyPair creation or import parameters.
type CreateOpts struct {
	// Name is a friendly name to refer to this KeyPair in other services.
	Name string `json:"name" required:"true"`

	// PublicKey [optional] is a pregenerated OpenSSH-formatted public key.
	// If provided, this key will be imported and no new key will be created.
	PublicKey string `json:"public_key,omitempty"`
}

// ToKeyPairCreateMap constructs a request body from CreateOpts.
func (opts CreateOpts) ToKeyPairCreateMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "keypair")
}
