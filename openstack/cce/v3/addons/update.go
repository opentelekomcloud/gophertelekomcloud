package addons

import (
	"fmt"
	"strings"

	"github.com/opentelekomcloud/gophertelekomcloud"
)

func Update(client *golangsdk.ServiceClient, id, clusterId string, opts UpdateOpts) (r UpdateResult) {
	b, err := golangsdk.BuildRequestBody(opts, "")
	if err != nil {
		r.Err = err
		return
	}

	raw, err := client.Put(
		fmt.Sprintf("https://%s.%s", clusterId, client.ResourceBaseURL()[8:])+
			strings.Join([]string{"addons", id + "?cluster_id=" + clusterId}, "/"),
		b, nil, &golangsdk.RequestOpts{
			OkCodes: []int{200},
		})

	return
}
