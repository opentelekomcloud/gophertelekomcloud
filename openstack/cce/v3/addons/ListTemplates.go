package addons

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ListOpts struct {
	Name string `q:"addon_template_name"`
}

func ListTemplates(client *golangsdk.ServiceClient, clusterID string, opts ListOpts) (*AddonTemplateList, error) {
	url, err := golangsdk.NewURLBuilder().WithEndpoints("addontemplates").WithQueryParams(&opts).Build()
	if err != nil {
		return nil, err
	}

	raw, err := client.Get(CCEServiceURL(client, clusterID, url.String()), err, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res AddonTemplateList
	return &res, extract.Into(raw.Body, &res)
}
