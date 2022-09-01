package nodes

import (
	"fmt"
	"strings"

	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack"
)

// Get retrieves a particular nodes based on its unique ID and cluster ID.
func Get(client *golangsdk.ServiceClient, clusterID, k8sName string) (r GetResult) {
	raw, err := client.Get(
		fmt.Sprintf("https://%s.%s", clusterID, client.ResourceBaseURL()[8:])+strings.Join([]string{"nodes", k8sName}, "/"),
		nil, openstack.StdRequestOpts())
	return
}
