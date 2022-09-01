package addons

import (
	"fmt"
	"strings"

	"github.com/opentelekomcloud/gophertelekomcloud"
)

// Create accepts a CreateOpts struct and uses the values to create a new
// addon.
func Create(client *golangsdk.ServiceClient, opts CreateOpts, clusterId string) (r CreateResult) {
	b, err := golangsdk.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(
		fmt.Sprintf("https://%s.%s", clusterId, client.ResourceBaseURL()[8:])+
			strings.Join([]string{"addons"}, "/"),
		b, nil, &golangsdk.RequestOpts{OkCodes: []int{201}})
	return
}
