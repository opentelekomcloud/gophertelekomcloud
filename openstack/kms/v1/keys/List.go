package keys

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ListOpts struct {
	KeyState string `q:"key_state"`
	Limit    string `q:"limit"`
	Marker   string `q:"marker"`
}

func List(client *golangsdk.ServiceClient, opts ListOpts) (*ListKey, error) {
	url, err := golangsdk.NewURLBuilder().WithEndpoints("kms", "list-keys").WithQueryParams(&opts).Build()
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL(url.String()), nil, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res ListKey
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type ListKey struct {
	Keys       []string `json:"keys"`
	KeyDetails []Key    `json:"key_details"`
	NextMarker string   `json:"next_marker"`
	Truncated  string   `json:"truncated"`
}
