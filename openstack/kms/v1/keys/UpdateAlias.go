package keys

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type UpdateAliasOpts struct {
	// ID of a CMK
	KeyID string `json:"key_id" required:"true"`
	// CMK alias
	KeyAlias string `json:"key_alias" required:"true"`
}

func UpdateAlias(client *golangsdk.ServiceClient, opts UpdateAliasOpts) error {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return err
	}

	_, err = client.Post(client.ServiceURL("kms", "update-key-alias"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return err
}
