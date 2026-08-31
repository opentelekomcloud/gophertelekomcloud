package subnets_test

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/testhelper/client"
)

func serviceClient() *golangsdk.ServiceClient {
	c := client.ServiceClient()
	c.ProjectID = "project-id"
	return c
}

const subnetJSON = `{
	"id": "subnet-id",
	"name": "subnet",
	"description": "description",
	"cidr": "192.168.20.0/24",
	"gateway_ip": "192.168.20.1",
	"ipv6_enable": true,
	"cidr_v6": "2407:c080:802:be7::/64",
	"gateway_ip_v6": "2407:c080:802:be7::1",
	"dhcp_enable": true,
	"primary_dns": "100.125.4.25",
	"secondary_dns": "100.125.129.199",
	"dnsList": [
		"100.125.4.25",
		"100.125.129.199"
	],
	"availability_zone": "eu-de-01",
	"vpc_id": "vpc-id",
	"status": "ACTIVE",
	"neutron_network_id": "network-id",
	"neutron_subnet_id": "neutron-subnet-id",
	"neutron_subnet_id_v6": "neutron-subnet-id-v6",
	"extra_dhcp_opts": [
		{
			"opt_value": "10.100.0.33,10.100.0.34",
			"opt_name": "ntp"
		}
	],
	"scope": "center",
	"tenant_id": "project-id"
}`
