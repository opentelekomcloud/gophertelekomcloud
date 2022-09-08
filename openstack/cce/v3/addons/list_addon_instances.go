package addons

import (
	"fmt"

	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func ListAddonInstances(client *golangsdk.ServiceClient, clusterID string) (*AddonInstanceList, error) {
	sprintf := fmt.Sprintf("%s?cluster_id=%s",
		fmt.Sprintf("https://%s.%s%s", clusterID, client.ResourceBaseURL()[8:], "addons"), clusterID)

	raw, err := client.Get(sprintf, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res AddonInstanceList
	err = extract.Into(raw.Body, &res)
	return &res, err
}
