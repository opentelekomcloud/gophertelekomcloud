package addons

import (
	"fmt"
	"strings"

	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// Get retrieves a particular addon based on its unique ID.
func Get(client *golangsdk.ServiceClient, id, clusterId string) (*Addon, error) {
	raw, err := client.Get(fmt.Sprintf("https://%s.%s", clusterId, client.ResourceBaseURL()[8:])+
		strings.Join([]string{"addons", id + "?cluster_id=" + clusterId}, "/"), nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res Addon
	err = extract.Into(raw, &res)
	return &res, err
}
