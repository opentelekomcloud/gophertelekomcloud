package nodes

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

func rootURL(c *golangsdk.ServiceClient, clusterid string) string {
	return c.ServiceURL("clusters", clusterid, "nodes")
}

func getJobURL(c *golangsdk.ServiceClient, jobid string) string {
	return c.ServiceURL("jobs", jobid)
}
