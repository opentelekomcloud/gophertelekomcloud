package clusters

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

const (
	clusters = "clusters"
)

func listURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL(clusters)
}

func singleURL(c *golangsdk.ServiceClient, id string) string {
	return c.ServiceURL(clusters, id)
}

func certificateURL(c *golangsdk.ServiceClient, clusterID string) string {
	return c.ServiceURL(clusters, clusterID, "sslCert")
}

func extendCommonURL(c *golangsdk.ServiceClient, clusterID string) string {
	return c.ServiceURL(clusters, clusterID, "extend")
}

func extendSpecialURL(c *golangsdk.ServiceClient, clusterID string) string {
	return c.ServiceURL(clusters, clusterID, "role_extend")
}
