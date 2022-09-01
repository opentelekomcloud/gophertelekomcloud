package addons

import (
	"fmt"
	"strings"

	"github.com/opentelekomcloud/gophertelekomcloud"
)

func ListAddonInstances(client *golangsdk.ServiceClient, clusterID string) (r ListInstanceResult) {
	raw, err := client.Get(fmt.Sprintf("%s?cluster_id=%s",
		fmt.Sprintf("https://%s.%s", clusterID, client.ResourceBaseURL()[8:])+
			strings.Join([]string{"addons"}, "/"), clusterID),
		nil, &golangsdk.RequestOpts{
			OkCodes: []int{200},
		})

	return
}
