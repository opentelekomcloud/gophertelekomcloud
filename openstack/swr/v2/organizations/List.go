package organizations

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ListOpts struct {
	// Organization name
	Namespace string `q:"namespace"`
}

func List(client *golangsdk.ServiceClient, opts ListOpts) ([]Organization, error) {
	url, err := golangsdk.NewURLBuilder().WithEndpoints("manage", "namespaces").WithQueryParams(&opts).Build()
	if err != nil {
		return nil, err
	}

	// GET /v2/manage/namespaces
	raw, err := client.Get(client.ServiceURL(url.String()), nil, nil)
	if err != nil {
		return nil, err
	}

	var res []Organization
	err = extract.IntoSlicePtr(raw.Body, &res, "namespaces")
	return res, err
}
