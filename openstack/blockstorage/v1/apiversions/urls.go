package apiversions

import (
	"net/url"
	"strings"

	"github.com/opentelekomcloud/gophertelekomcloud"
)

func getURL(c *golangsdk.ServiceClient, version string) string {
	return c.ServiceURL(strings.TrimRight(version, "/") + "/")
}

func listURL(c *golangsdk.ServiceClient) string {
	u, _ := url.Parse(c.ServiceURL(""))
	u.Path = "/"
	return u.String()
}
