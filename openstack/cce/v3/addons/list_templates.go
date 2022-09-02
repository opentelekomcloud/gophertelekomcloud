package addons

import (
	"fmt"
	"strings"

	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ListOpts struct {
	Name string `q:"addon_template_name"`
}

func ListTemplates(client *golangsdk.ServiceClient, clusterID string, opts ListOpts) (*AddonTemplateList, error) {
	q, err := golangsdk.BuildQueryString(opts)
	if err != nil {
		return nil, err
	}

	raw, err := client.Get(fmt.Sprintf("https://%s.%s", clusterID, client.ResourceBaseURL()[8:])+
		strings.Join([]string{"addontemplates"}, "/")+q.String(), nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res AddonTemplateList
	err = extract.Into(raw, &res)
	return &res, err
}
