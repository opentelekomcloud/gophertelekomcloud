package keys

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type EncryptDEKOpts struct {
	// ID of a CMK
	KeyID string `json:"key_id" required:"true"`
	// CMK description
	EncryptionContext string `json:"encryption_context,omitempty"`
	// 36-byte serial number of a request message
	DataKeyPlainLength string `json:"datakey_plain_length,omitempty"`
	// Both the plaintext (64 bytes) of a DEK and the SHA-256 hash value (32 bytes)
	// of the plaintext are expressed as a hexadecimal character string.
	PlainText string `json:"plain_text" required:"true"`
}

func EncryptDEKGet(client *golangsdk.ServiceClient, opts EncryptDEKOpts) (*EncryptDEK, error) {
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

type EncryptDEK struct {
	KeyID         string `json:"key_id"`
	DataKeyLength string `json:"datakey_length"`
	CipherText    string `json:"cipher_text"`
}
