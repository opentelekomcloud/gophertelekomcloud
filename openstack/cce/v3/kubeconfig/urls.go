package kubeconfig

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
)

const (
	clusterPath = "clusters"
	rootPath    = "clustercert"
)

func rootURL(client *golangsdk.ServiceClient, clusterID string) string {
	return client.ServiceURL(clusterPath, clusterID, rootPath)
}
