package repositories

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	v2 "github.com/opentelekomcloud/gophertelekomcloud/openstack/swr/v2"
)

const (
	repos = "repos"
)

func createURL(client *golangsdk.ServiceClient, namespace string) string {
	return client.ServiceURL(v2.Base, v2.Namespaces, namespace, repos)
}

func listURL(client *golangsdk.ServiceClient) string {
	return client.ServiceURL(v2.Base, repos)
}

func repoURL(client *golangsdk.ServiceClient, namespace, repository string) string {
	return client.ServiceURL(v2.Base, v2.Namespaces, namespace, repos, repository)
}
