package instances

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type GetPrivateDomainNameParams struct {
	DnsType string `q:"dns_type" required:"true"`
}

// This function is used to query the domain name of a DB instance.
func GetPrivateDomainName(client *golangsdk.ServiceClient, instanceId string, opts GetPrivateDomainNameParams) (*PrivateDomain, error) {
	// GET /v3/{project_id}/instances/{instance_id}/dns
	url, err := golangsdk.NewURLBuilder().WithEndpoints("instances", instanceId, "dns").WithQueryParams(&opts).Build()
	if err != nil {
		return nil, err
	}

	raw, err := client.Get(client.ServiceURL(url.String()), nil, nil)
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
