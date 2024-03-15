package keys

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type DecryptDataOpts struct {
	CipherText          string `json:"cipher_text" required:"true"`
	KeyID               string `json:"key_id,omitempty"`
	EncryptionAlgorithm string `json:"encryption_algorithm,omitempty"`
	Sequence            string `json:"sequence,omitempty"`
}

func DecryptData(client *golangsdk.ServiceClient, opts DecryptDataOpts) (*DecryptResp, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL("kms", "decrypt-data"), b, nil,
		&golangsdk.RequestOpts{OkCodes: []int{200}})
	if err != nil {
		return nil, err
	}
	var res DecryptResp
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type DecryptResp struct {
	KeyID           string `json:"key_id"`
	PlainText       string `json:"plain_text"`
	PlainTextBase64 string `json:"plain_text_base64"`
}
