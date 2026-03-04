package instances

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func GetPrivateDomainName(client *golangsdk.ServiceClient, instanceId string) (*PrivateDomain, error) {
	// GET /v3/{project_id}/instances/{instance_id}/dns
	raw, err := client.Get(client.ServiceURL("instances", instanceId, "dns"), nil, nil)
	if err != nil {
		return nil, err
	}
	var res PrivateDomain
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type PrivateDomain struct {
	InstanceId  string `json:"instance_id"`
	DnsName     string `json:"dns_name"`
	DnsType     string `json:"dns_type"`
	Ipv4Address string `json:"ipv4_address"`
	Status      string `json:"status"`
}
