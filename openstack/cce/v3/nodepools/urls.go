package nodepools

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

const (
	rootPath     = "clusters"
	resourcePath = "nodepools"
)

func rootURL(c *golangsdk.ServiceClient, clusterid string) string {
	return c.ServiceURL(rootPath, clusterid, resourcePath)
}
