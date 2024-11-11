package keys

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type EncryptDEKOpts struct {
	// ID of a CMK
	KeyID string `json:"key_id" required:"true"`
	// Encryption context
	EncryptionContext string `json:"encryption_context,omitempty"`
	// Plain text data key length
	DataKeyPlainLength string `json:"datakey_plain_length,omitempty"`
	// Plain text to encrypt
	PlainText string `json:"plain_text" required:"true"`
}

type EncryptDEK struct {
	KeyID         string `json:"key_id"`
	DataKeyLength string `json:"datakey_length"`
	CipherText    string `json:"cipher_text"`
}

func EncryptDataEncryptionKey(client *golangsdk.ServiceClient, opts EncryptDEKOpts) (*EncryptDEK, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL("kms", "encrypt-datakey"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res EncryptDEK
	err = extract.Into(raw.Body, &res)
	return &res, err
}
