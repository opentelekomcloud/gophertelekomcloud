package secgroups

import (
	"net/http"

	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// List will return a collection of all the security groups for a particular tenant.
func List(client *golangsdk.ServiceClient) ([]SecurityGroup, error) {
	raw, err := client.Get(client.ServiceURL("os-security-groups"), nil, nil)
	return extra2(err, raw)
}

// ListByServer will return a collection of all the security groups which are associated with a particular server.
func ListByServer(client *golangsdk.ServiceClient, serverID string) ([]SecurityGroup, error) {
	raw, err := client.Get(client.ServiceURL("servers", serverID, "os-security-groups"), nil, nil)
	return extra2(err, raw)
}

func extra2(err error, raw *http.Response) ([]SecurityGroup, error) {
	if err != nil {
		return nil, err
	}

	var res []SecurityGroup
	err = extract.IntoSlicePtr(raw.Body, &res, "security_groups")
	return res, err
}
