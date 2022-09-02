package v3

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/testhelper/client"
)

const TokenID = client.TokenID

func ServiceClient() *golangsdk.ServiceClient {
	sc := client.ServiceClient()
	sc.ResourceBase = sc.Endpoint + "api/" + "v3/" + "projects/" + "c59fd21fd2a94963b822d8985b884673/"
	return sc
}
