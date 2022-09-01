package addons

import (
	"fmt"
	"strings"

	"github.com/opentelekomcloud/gophertelekomcloud"
)

// Get retrieves a particular addon based on its unique ID.
func Get(client *golangsdk.ServiceClient, id, clusterId string) (r GetResult) {
	raw, err := client.Get(
		fmt.Sprintf("https://%s.%s", clusterId, client.ResourceBaseURL()[8:])+
			strings.Join([]string{"addons", id + "?cluster_id=" + clusterId}, "/"),
		nil, &golangsdk.RequestOpts{
			OkCodes: []int{200},
		})
	return
}
