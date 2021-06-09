package repositories

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	v2 "github.com/opentelekomcloud/gophertelekomcloud/openstack/swr/v2"
)

func createURL(client *golangsdk.ServiceClient, namespace string) string {
	return client.ServiceURL(v2.Base, v2.Namespaces, namespace, v2.Repos)
}

func listURL(client *golangsdk.ServiceClient) string {
	return client.ServiceURL(v2.Base, v2.Repos)
}

func repoURL(client *golangsdk.ServiceClient, namespace, repository string) string {
	return client.ServiceURL(v2.Base, v2.Namespaces, namespace, v2.Repos, repository)
}
