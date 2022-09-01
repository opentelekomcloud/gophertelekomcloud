package addons

import (
	"fmt"
	"strings"

	"github.com/opentelekomcloud/gophertelekomcloud"
)

func ListTemplates(client *golangsdk.ServiceClient, clusterID string, opts ListOpts) (r ListTemplateResult) {
	q, err := golangsdk.BuildQueryString(opts)
	if err != nil {
		r.Err = err
		return
	}

	raw, err := client.Get(fmt.Sprintf("https://%s.%s", clusterID, client.ResourceBaseURL()[8:])+
		strings.Join([]string{"addontemplates"}, "/")+q.String(),
		nil, &golangsdk.RequestOpts{
			OkCodes: []int{200},
		})
	return
}
