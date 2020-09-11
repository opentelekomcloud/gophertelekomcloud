package alarmreminding

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
)

func WarnAlertURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL("warnalert", "alertconfig", "query")
}
