package domains

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ListOpts struct {
	// Organization name
	Namespace string `json:"-" required:"true"`
	// Image repository name
	Repository string `json:"-" required:"true"`
}

func List(client *golangsdk.ServiceClient, opts ListOpts) ([]AccessDomain, error) {
	// GET /v2/manage/namespaces/{namespace}/repositories/{repository}/access-domains
	url := client.ServiceURL("manage", "namespaces", opts.Namespace, "repos", opts.Repository, "access-domains")
	raw, err := client.Get(url, nil, nil)
	if err != nil {
		return nil, err
	}

	var res []AccessDomain
	err = extract.IntoSlicePtr(raw.Body, &res, "")
	return res, err
}
