package routes

import (
	"log"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/networking/v2/extensions/security/rules"
)

func CreateMultipleSgsRules(clientV2 *golangsdk.ServiceClient, sgID string, count int) ([]string, error) {
	i := 0
	createdSgs := make([]string, count)
	for i < count {
		opts := rules.CreateOpts{
			Description:  "description",
			SecGroupID:   sgID,
			PortRangeMin: 1000 + i,
			PortRangeMax: 5000 + i,
			Direction:    "ingress",
			EtherType:    "IPv4",
			Protocol:     "TCP",
		}
		log.Printf("[DEBUG] Create OpenTelekomCloud Neutron security group: %#v", opts)
		securityGroupRule, err := rules.Create(clientV2, opts).Extract()
		if err != nil {
			return createdSgs, err
		}
		createdSgs[i] = securityGroupRule.ID
		i += 1
	}
	return createdSgs, nil
}
