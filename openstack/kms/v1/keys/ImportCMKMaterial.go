package keys

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type ImportCMKOpts struct {
	KeyId                string `json:"key_id" required:"true"`
	ImportToken          string `json:"import_token" required:"true"`
	EncryptedKeyMaterial string `json:"encrypted_key_material" required:"true"`
	ExpirationTime       string `json:"expiration_time,omitempty"`
	Sequence             string `json:"sequence,omitempty"`
}

func ImportCMKMaterial(client *golangsdk.ServiceClient, opts ImportCMKOpts) error {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return err
	}

	_, err = client.Post(client.ServiceURL("kms", "import-key-material"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})

	return err
}
