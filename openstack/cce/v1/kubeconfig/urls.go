package nodes

import (
	"fmt"
	"strings"

	"github.com/opentelekomcloud/gophertelekomcloud"
)

const (
	rootPath = "clustercert"
)

func certURL(client *golangsdk.ServiceClient, clusterID string) string {
	return CCEServiceURL(client, clusterID, rootPath)
}

func CCEServiceURL(client *golangsdk.ServiceClient, clusterID string, parts ...string) string {
	rbUrl := fmt.Sprintf("https://%s.%s", clusterID, client.ResourceBaseURL()[8:])
	return rbUrl + strings.Join(parts, "/")
}
