package addons

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func GetTemplates(client *golangsdk.ServiceClient) (*AddonTemplateList, error) {
	raw, err := client.Get(client.ServiceURL("addontemplates"), nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res AddonTemplateList
	return &res, extract.Into(raw.Body, &res)
}
