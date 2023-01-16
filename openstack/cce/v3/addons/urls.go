package addons

import (
	"fmt"
	"strings"

	"github.com/opentelekomcloud/gophertelekomcloud"
)

const (
	rootPath      = "addons"
	templatesPath = "addontemplates"
)

func rootURL(client *golangsdk.ServiceClient, clusterID string) string {
	return CCEServiceURL(client, clusterID, rootPath)
}

func resourceURL(client *golangsdk.ServiceClient, id, clusterID string) string {
	return CCEServiceURL(client, clusterID, rootPath, id+"?cluster_id="+clusterID)
}

func CCEServiceURL(client *golangsdk.ServiceClient, clusterID string, parts ...string) string {
	rbUrl := fmt.Sprintf("https://%s.%s", clusterID, client.ResourceBaseURL()[8:])
	return rbUrl + strings.Join(parts, "/")
}

func templatesURL(client *golangsdk.ServiceClient, clusterID string) string {
	return CCEServiceURL(client, clusterID, templatesPath)
}

func instanceURL(client *golangsdk.ServiceClient, clusterID string) string {
	return fmt.Sprintf("%s?cluster_id=%s", CCEServiceURL(client, clusterID, rootPath), clusterID)
}

// GET /api/v3/addontemplates
func addonTemplatesURL(client *golangsdk.ServiceClient) string {
	return client.ServiceURL(templatesPath)
}
