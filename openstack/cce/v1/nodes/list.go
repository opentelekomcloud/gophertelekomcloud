package nodes

import (
	"fmt"
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack"
	"strings"
)

// List returns collection of nodes.
func List(client *golangsdk.ServiceClient, clusterID string) (r ListResult) {
	raw, err := client.Get(
		fmt.Sprintf("https://%s.%s", clusterID, client.ResourceBaseURL()[8:])+strings.Join([]string{"nodes"}, "/"),
		nil, openstack.StdRequestOpts())
	return
}
