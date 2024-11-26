package clusters

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

const (
	rootPath     = "clusters"
	certPath     = "clustercert"
	masterIpPath = "mastereip"
)

func certificateURL(c *golangsdk.ServiceClient, id string) string {
	return c.ServiceURL(rootPath, id, certPath)
}

func masterIpURL(c *golangsdk.ServiceClient, id string) string {
	return c.ServiceURL(rootPath, id, masterIpPath)
}
