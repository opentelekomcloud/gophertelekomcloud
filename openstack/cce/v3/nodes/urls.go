package nodes

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

const (
	rootPath     = "clusters"
	resourcePath = "nodes"
)

func rootURL(c *golangsdk.ServiceClient, clusterid string) string {
	return c.ServiceURL("clusters", clusterid, "nodes")
}

func resourceURL(c *golangsdk.ServiceClient, clusterid, nodeid string) string {
	return c.ServiceURL("clusters", clusterid, "nodes", nodeid)
}

func getJobURL(c *golangsdk.ServiceClient, jobid string) string {
	return c.ServiceURL("jobs", jobid)
}
