package addons

import (
	"fmt"
	"strings"

	"github.com/opentelekomcloud/gophertelekomcloud"
)

// Delete will permanently delete a particular addon based on its unique ID.
func Delete(client *golangsdk.ServiceClient, id, clusterId string) (err error) {
	url := fmt.Sprintf("https://%s.%s", clusterId, client.ResourceBaseURL()[8:]) +
		strings.Join([]string{"addons", id + "?cluster_id=" + clusterId}, "/")

	_, err = client.Delete(url, &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: map[string]string{"Content-Type": "application/json"}, JSONBody: nil,
	})

	return
}
