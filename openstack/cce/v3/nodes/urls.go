package nodes

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

func rootURL(c *golangsdk.ServiceClient, clusterid string) string {
	return c.ServiceURL("clusters", clusterid, "nodes")
}
