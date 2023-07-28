package metadata

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type ImportOptsBuilder interface {
	ToMetadataImportMap() (map[string]interface{}, error)
}

type ImportOpts struct {
	XAccountType string `json:"xaccount_type"`
	DomainID     string `json:"domain_id" required:"true"`
	Metadata     string `json:"metadata" required:"true"`
}

func (opts ImportOpts) ToMetadataImportMap() (map[string]interface{}, error) {
	return build.RequestBodyMap(opts, "")
}

func Import(client *golangsdk.ServiceClient, provider, protocol string, opts ImportOptsBuilder) (r ImportResult) {
	b, err := opts.ToMetadataImportMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(metadataURL(client, provider, protocol), b, &r.Body, nil)
	return
}

func Get(client *golangsdk.ServiceClient, provider, protocol string) (r GetResult) {
	_, r.Err = client.Get(metadataURL(client, provider, protocol), &r.Body, &golangsdk.RequestOpts{
		MoreHeaders: map[string]string{"Content-Type": "application/json"},
	})
	return
}
