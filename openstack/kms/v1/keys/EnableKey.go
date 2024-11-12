package keys

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func EnableKey(client *golangsdk.ServiceClient, keyID string) (*UpdateKeyState, error) {
	opts := struct {
		KeyID string `json:"key_id"`
	}{
		KeyID: keyID,
	}

	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL("kms", "enable-key"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res UpdateKeyState

	err = extract.IntoStructPtr(raw.Body, &res, "key_info")
	return &res, err
}
