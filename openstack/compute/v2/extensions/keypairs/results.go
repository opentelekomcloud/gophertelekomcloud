package keypairs

import (
	"net/http"

	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// KeyPair is an SSH key known to the OpenStack Cloud that is available to be injected into servers.
type KeyPair struct {
	// Name is used to refer to this keypair from other services within this
	// region.
	Name string `json:"name"`
	// Fingerprint is a short sequence of bytes that can be used to authenticate
	// or validate a longer public key.
	Fingerprint string `json:"fingerprint"`
	// PublicKey is the public key from this pair, in OpenSSH format.
	// "ssh-rsa AAAAB3Nz..."
	PublicKey string `json:"public_key"`
	// PrivateKey is the private key from this pair, in PEM format.
	// "-----BEGIN RSA PRIVATE KEY-----\nMIICXA..."
	// It is only present if this KeyPair was just returned from a Create call.
	PrivateKey string `json:"private_key"`
	// UserID is the user who owns this KeyPair.
	UserID string `json:"user_id"`
}

func extra(err error, raw *http.Response) (*KeyPair, error) {
	if err != nil {
		return nil, err
	}

	var res KeyPair
	err = extract.IntoStructPtr(raw.Body, &res, "keypair")
	return &res, err
}
