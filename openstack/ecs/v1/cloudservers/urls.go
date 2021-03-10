package cloudservers

import "github.com/opentelekomcloud/gophertelekomcloud"

const (
	rootPath = "cloudservers"
	jobsPath = "jobs"
)

func createURL(sc *golangsdk.ServiceClient) string {
	return sc.ServiceURL(rootPath)
}

func deleteURL(sc *golangsdk.ServiceClient) string {
	return sc.ServiceURL(rootPath, "delete")
}

func getURL(sc *golangsdk.ServiceClient, serverID string) string {
	return sc.ServiceURL(rootPath, serverID)
}

func jobURL(sc *golangsdk.ServiceClient, jobID string) string {
	return sc.ServiceURL(jobsPath, jobID)
}
