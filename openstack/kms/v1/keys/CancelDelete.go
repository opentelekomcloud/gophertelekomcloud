package keys

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func CancelDelete(client *golangsdk.ServiceClient, keyID string) (*Key, error) {
	opts := struct {
		KeyID string `json:"key_id"`
	}{
		KeyID: keyID,
	}

	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL("kms", "cancel-key-deletion"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res Key

	err = extract.Into(raw.Body, &res)
	return &res, err
}
