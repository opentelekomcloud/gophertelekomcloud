package snapshots

import "github.com/opentelekomcloud/gophertelekomcloud"

// policyURL used to set or query the snapshot policy
func policyURL(c *golangsdk.ServiceClient, clusterId string) string {
	return c.ServiceURL("clusters", clusterId, "index_snapshot/policy")
}

// enableURL used to automatically perform basic configurations for a cluster snapshot
func enableURL(c *golangsdk.ServiceClient, clusterId string) string {
	return c.ServiceURL("clusters", clusterId, "index_snapshot/auto_setting")
}

// disableURL used to disable the snapshot function
func disableURL(c *golangsdk.ServiceClient, clusterId string) string {
	return c.ServiceURL("clusters", clusterId, "index_snapshots")
}

// createURL used to manually create a snapshot
func createURL(c *golangsdk.ServiceClient, clusterId string) string {
	return c.ServiceURL("clusters", clusterId, "index_snapshot")
}

func configURL(c *golangsdk.ServiceClient, clusterId string) string {
	return c.ServiceURL("clusters", clusterId, "index_snapshot", "setting")
}

// listURL used to query all snapshots of a cluster
func listURL(c *golangsdk.ServiceClient, clusterId string) string {
	return c.ServiceURL("clusters", clusterId, "index_snapshots")
}

// deleteURL used to delete a snapshot
func deleteURL(c *golangsdk.ServiceClient, clusterId, snapId string) string {
	return c.ServiceURL("clusters", clusterId, "index_snapshot", snapId)
}
