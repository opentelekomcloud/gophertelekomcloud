package keypairs

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// List returns a Pager that allows you to iterate over a collection of KeyPairs.
func List(client *golangsdk.ServiceClient) ([]KeyPair, error) {
	raw, err := client.Get(client.ServiceURL("os-keypairs"), nil, nil)
	if err != nil {
		return nil, err
	}

	var res []struct {
		KeyPair KeyPair `json:"keypair"`
	}
	err = extract.IntoSlicePtr(raw.Body, &res, "keypairs")

	results := make([]KeyPair, len(res))
	for i, pair := range res {
		results[i] = pair.KeyPair
	}
	return results, err
}
