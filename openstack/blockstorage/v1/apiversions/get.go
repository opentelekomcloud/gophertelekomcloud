package apiversions

import (
	"strings"

	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func Get(client *golangsdk.ServiceClient, v string) (*APIVersion, error) {
	raw, err := client.Get(client.ServiceURL(strings.TrimRight(v, "/")+"/"), nil, nil)
	if err != nil {
		return nil, err
	}

	var res struct {
		Version APIVersion `json:"version"`
	}
	err = extract.Into(raw.Body, &res)
	return &res.Version, err
}
