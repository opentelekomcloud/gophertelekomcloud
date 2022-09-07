package addons

import (
	"fmt"
	"strings"

	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func ListAddonInstances(client *golangsdk.ServiceClient, clusterID string) (*AddonInstanceList, error) {
	raw, err := client.Get(fmt.Sprintf("%s?cluster_id=%s",
		fmt.Sprintf("https://%s.%s", clusterID, client.ResourceBaseURL()[8:])+
			strings.Join([]string{"addons"}, "/"), clusterID), nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res AddonInstanceList
	err = extract.Into(raw.Body, &res)
	return &res, err
}
