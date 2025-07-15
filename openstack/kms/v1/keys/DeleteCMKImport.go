package keys

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type DeleteCMKImportOpts struct {
	KeyId    string `json:"key_id" required:"true"`
	Sequence string `json:"sequence,omitempty"`
}

func DeleteCMKImport(client *golangsdk.ServiceClient, opts DeleteCMKImportOpts) error {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return err
	}

	_, err = client.Post(client.ServiceURL("kms", "delete-imported-key-material"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return err
}
