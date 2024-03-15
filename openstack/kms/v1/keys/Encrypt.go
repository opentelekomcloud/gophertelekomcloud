package keys

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type EncryptDataOpts struct {
	KeyID               string `json:"key_id" required:"true"`
	PlainText           string `json:"plain_text" required:"true"`
	EncryptionAlgorithm string `json:"encryption_algorithm,omitempty"`
	Sequence            string `json:"sequence,omitempty"`
}

func EncryptData(client *golangsdk.ServiceClient, opts EncryptDataOpts) (*EncryptResp, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL("kms", "encrypt-data"), b, nil,
		&golangsdk.RequestOpts{
			OkCodes: []int{200},
		})
	if err != nil {
		return nil, err
	}
	var res EncryptResp
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type EncryptResp struct {
	KeyID      string `json:"key_id"`
	CipherText string `json:"cipher_text"`
}
