package cc

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/cc/v3/central_network"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/cc/v3/policy"
)

func WaitForCentralNetworkAvailable(client *golangsdk.ServiceClient, secs int, id string) error {
	return golangsdk.WaitFor(secs, func() (bool, error) {
		cn, err := central_network.Get(client, id)
		if err != nil {
			return false, err
		}
		return cn.State == "AVAILABLE", nil
	})
}

func WaitForPolicyAvailable(client *golangsdk.ServiceClient, secs int, centralNetworkId, policyId string) error {
	return golangsdk.WaitFor(secs, func() (bool, error) {
		list, err := policy.List(client, policy.ListOpts{CentralNetworkId: centralNetworkId})
		if err != nil {
			return false, err
		}
		for _, p := range list.CentralNetworkPolicies {
			if p.ID == policyId {
				return p.State == "AVAILABLE", nil
			}
		}

		return true, nil
	})
}
