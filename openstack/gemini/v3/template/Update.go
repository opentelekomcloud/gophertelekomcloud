package template

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type UpdateOpts struct {
	ConfigId    string            `json:"-"`
	Name        string            `json:"name,omitempty"`
	Description string            `json:"description,omitempty"`
	Values      map[string]string `json:"values,omitempty"`
}

func Update(client *golangsdk.ServiceClient, opts UpdateOpts) error {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return err
	}

	_, err = client.Put(client.ServiceURL("configurations", opts.ConfigId), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})

	return err
}
