package keys

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type UpdateAliasOpts struct {
	// ID of a CMK
	KeyID string `json:"key_id" required:"true"`
	// CMK description
	KeyAlias string `json:"key_alias" required:"true"`
}

func UpdateAlias(client *golangsdk.ServiceClient, opts UpdateAliasOpts) (*UpdateKeyAlias, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL("kms", "update-key-alias"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res UpdateKeyAlias

	err = extract.IntoStructPtr(raw.Body, &res, "key_info")
	return &res, err
}

type UpdateKeyAlias struct {
	KeyID    string `json:"key_id"`
	KeyAlias string `json:"key_alias"`
}
