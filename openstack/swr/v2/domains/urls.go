package domains

import (
	"fmt"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	v2 "github.com/opentelekomcloud/gophertelekomcloud/openstack/swr/v2"
)

const domains = "access-domains"

func listURL(client *golangsdk.ServiceClient, org, repo string) string {
	return client.ServiceURL(v2.Base, v2.Namespaces, org, v2.Repos, repo, domains)
}

func singleURL(client *golangsdk.ServiceClient, org, repo, domain string) string {
	return fmt.Sprintf("%s/%s", listURL(client, org, repo), domain)
}
