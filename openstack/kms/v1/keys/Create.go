package keys

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type CreateOpts struct {
	// Alias of a CMK
	KeyAlias string `json:"key_alias" required:"true"`
	// CMK description
	KeyDescription string `json:"key_description,omitempty"`
	// Region where a CMK resides
	Realm string `json:"realm,omitempty"`
	// Purpose of a CMK (The default value is Encrypt_Decrypt)
	KeyUsage string `json:"key_usage,omitempty"`
	// Origin of a CMK. The default value is kms. Possible values: `kms` / `external`
	Origin string `json:"origin,omitempty"`
}

func Create(client *golangsdk.ServiceClient, opts CreateOpts) (*Key, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL("kms", "create-key"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})

	if err != nil {
		return nil, err
	}

	var res Key

	err = extract.IntoStructPtr(raw.Body, &res, "key_info")
	return &res, err
}
