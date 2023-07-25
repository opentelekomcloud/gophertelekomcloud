package keypairs

import (
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

// KeyPair is an SSH key known to the OpenStack Cloud that is available to be injected into bms servers.
type KeyPair struct {
	// Name is used to refer to this keypair from other services within this region.
	Name string `json:"name"`
	// Fingerprint is a short sequence of bytes that can be used to authenticate or validate a longer public key.
	Fingerprint string `json:"fingerprint"`
	// PublicKey is the public key from this pair, in OpenSSH format. "ssh-rsa AAAAB3Nz..."
	PublicKey string `json:"public_key"`
}

// KeyPairPage stores a single page of all KeyPair results from a List call.
// Use the ExtractKeyPairs function to convert the results to a slice of KeyPairs.
type KeyPairPage struct {
	pagination.LinkedPageBase
}

// IsEmpty determines whether a KeyPairPage is empty.
func (page KeyPairPage) IsEmpty() (bool, error) {
	ks, err := ExtractKeyPairs(page)
	return len(ks) == 0, err
}

// ExtractKeyPairs interprets a page of results as a slice of KeyPairs.
func ExtractKeyPairs(r pagination.Page) ([]KeyPair, error) {
	var res []struct {
		KeyPair KeyPair `json:"keypair"`
	}

	err := extract.IntoSlicePtr(r.(KeyPairPage).Body, &res, "keypairs")
	results := make([]KeyPair, len(res))

	for i, pair := range res {
		results[i] = pair.KeyPair
	}
	return results, err
}
