package template

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type CreateOpts struct {
	Name        string            `json:"name" required:"true"`
	Description string            `json:"description,omitempty"`
	Values      map[string]string `json:"values,omitempty"`
	DataStore   DataStoreOpt      `json:"datastore" required:"true"`
}

type DataStoreOpt struct {
	Type    string `json:"type" required:"true"`
	Version string `json:"version" required:"true"`
}

func Create(client *golangsdk.ServiceClient, opts CreateOpts) (*ConfigurationResult, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL("configurations"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res ConfigurationResult
	return &res, extract.IntoStructPtr(raw.Body, &res, "configuration")
}

type ConfigurationResult struct {
	Id                   string `json:"id"`
	Name                 string `json:"name"`
	DataStoreVersionName string `json:"datastore_version_name"`
	DataStoreName        string `json:"datastore_name"`
	Description          string `json:"description"`
	Created              string `json:"created"`
	Updated              string `json:"updated"`
}
