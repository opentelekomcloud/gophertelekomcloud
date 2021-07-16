package nodes

import (
	"fmt"
	"strings"

	"github.com/opentelekomcloud/gophertelekomcloud"
)

const (
	rootPath = "nodes"
)

func listURL(client *golangsdk.ServiceClient, clusterID string) string {
	return CCEServiceURL(client, clusterID, rootPath)
}

func nodeURL(client *golangsdk.ServiceClient, clusterID string, k8sName string) string {
	return CCEServiceURL(client, clusterID, rootPath, k8sName)
}

func CCEServiceURL(client *golangsdk.ServiceClient, clusterID string, parts ...string) string {
	rbUrl := fmt.Sprintf("https://%s.%s", clusterID, client.ResourceBaseURL()[8:])
	return rbUrl + strings.Join(parts, "/")
}
