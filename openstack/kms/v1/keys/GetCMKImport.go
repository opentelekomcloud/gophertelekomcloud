package keys

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type GetCMKImportOpts struct {
	KeyId             string `json:"key_id" required:"true"`
	WrappingAlgorithm string `json:"wrapping_algorithm" required:"true"`
	Sequence          string `json:"sequence,omitempty"`
}

func GetCMKImport(client *golangsdk.ServiceClient, opts GetCMKImportOpts) (*CMKImport, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL("kms", "get-parameters-for-import"), &b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res CMKImport

	err = extract.Into(raw.Body, &res)
	return &res, err
}

type CMKImport struct {
	KeyId          string `json:"key_id"`
	ImportToken    string `json:"import_token"`
	ExpirationTime int    `json:"expiration_time"`
	PublicKey      string `json:"public_key"`
}
