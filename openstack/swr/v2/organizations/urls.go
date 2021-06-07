package organizations

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	v2 "github.com/opentelekomcloud/gophertelekomcloud/openstack/swr/v2"
)

func listURL(client *golangsdk.ServiceClient) string {
	return client.ServiceURL(v2.Base, v2.Namespaces)
}

func organizationURL(client *golangsdk.ServiceClient, id string) string {
	return client.ServiceURL(v2.Base, v2.Namespaces, id)
}
