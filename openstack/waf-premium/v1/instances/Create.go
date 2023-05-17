package instances

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type CreateOpts struct {
	// Billing mode. Currently, only pay-per-use billing (30) is
	// supported. Make sure your account balance is enough, or
	// the dedicated WAF engine will forward requests directly to
	// the origin server without inspection.
	ChargeMode int `json:"chargemode"`
	// Region where a dedicated engine is to be created. Its
	// value is EU-DE.
	Region string `json:"region" required:"true"`
	// AZ where the dedicated engine is to be created.
	AvailabilityZone string `json:"available_zone" required:"true"`
	// Dedicated engine CPU architecture. Its value can be x86.
	Architecture string `json:"arch" required:"true"`
	// Prefix of the dedicated WAF engine name, which is user-defined
	InstanceName string `json:"instancename" required:"true"`
	// Specification of the dedicated engine version. The value can be
	// waf.instance.enterprise or waf.instance.professional. An
	// enterprise edition dedicated engine has more functions
	// than a professional edition one. For more details, see the
	// Web Application Firewall (WAF) User Guide
	Specification string `json:"specification" required:"true"`
	// ID of the specification of the ECS hosting the dedicated
	// engine. It can be obtained by calling the ECS ListFlavors API.
	// For the enterprise edition, ECS specification with 8 vCPUs
	// and 16 GB memory are used. For the professional edition,
	// ECS specification with 2 vCPUs and 4 GB memory are used.
	Flavor string `json:"cpu_flavor" required:"true"`
	// ID of the VPC where the dedicated engine is located. It
	// can be obtained by calling the ListVpcs API.
	VpcId string `json:"vpc_id" required:"true"`
	// ID of the VPC subnet where the dedicated engine is
	// located. It can be obtained by calling the ListSubnets API.
	// subnet_id has the same value as network_id obtained by
	// calling the OpenStack APIs.
	SubnetId string `json:"subnet_id" required:"true"`
	// ID of the security group where  the dedicated engine is
	// located. It can be obtained by calling the ListSecurityGroups API.
	SecurityGroupsId []string `json:"security_group" required:"true"`
	// Number of dedicated engines to be provisioned
	Count int `json:"count" required:"true"`
}

// Create will create a new instance on the values in CreateOpts. To extract
// the instance object from the response, call the Extract method on the CreateResult.
func Create(client *golangsdk.ServiceClient, opts CreateOpts) (*InstanceResponse, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// POST /v1/{project_id}/premium-waf/instance
	raw, err := client.Post(client.ServiceURL("instance"), b,
		nil, &golangsdk.RequestOpts{
			OkCodes: []int{201},
		})
	if err != nil {
		return nil, err
	}

	var res InstanceResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type InstanceResponse struct {
	Instances []Info `json:"instances"`
}

type Info struct {
	// Instance id
	Id string `json:"id"`
	// Instance name
	Name string `json:"name"`
}
