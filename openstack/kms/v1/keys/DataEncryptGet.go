package keys

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type DataEncryptOpts struct {
	// ID of a CMK
	KeyID string `json:"key_id" required:"true"`
	// CMK description
	EncryptionContext string `json:"encryption_context,omitempty"`
	// 36-byte serial number of a request message
	DatakeyLength string `json:"datakey_length,omitempty"`
}

func DataEncryptGet(client *golangsdk.ServiceClient, opts DataEncryptOpts) (*DataKey, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL("kms", "create-datakey"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})

	if err != nil {
		return nil, err
	}

	var res DataKey

	err = extract.Into(raw.Body, &res)
	return &res, err
}

type DataKey struct {
	// Current ID of a CMK
	KeyID      string `json:"key_id"`
	PlainText  string `json:"plain_text"`
	CipherText string `json:"cipher_text"`
}
