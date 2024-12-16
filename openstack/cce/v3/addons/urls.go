package addons

import (
	"fmt"
	"strings"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
)

func CCEServiceURL(client *golangsdk.ServiceClient, clusterID string, parts ...string) string {
	rbUrl := fmt.Sprintf("https://%s.%s", clusterID, client.ResourceBaseURL()[8:])
	return rbUrl + strings.Join(parts, "/")
}

func instanceURL(client *golangsdk.ServiceClient, clusterID string) string {
	return fmt.Sprintf("%s?cluster_id=%s", CCEServiceURL(client, clusterID, "addons"), clusterID)
}
