package management

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dms/v2.1/instances"
)

const (
	crossVPCPath = "crossvpc"
	modifyPath   = "modify"
)

type CrossVPCModifyOpts struct {
	AdvertisedIpContents map[string]string `json:"advertised_ip_contents" required:"true"`
}

// CrossVPCModify is used to modify the private IP address for cross-VPC access.
// Send POST to /v2/{project_id}/instances/{instance_id}/crossvpc/modify
func CrossVPCModify(client *golangsdk.ServiceClient, instanceId string, opts PasswordOpts) (*CrossVPCModifyResp, error) {
	body, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL(instances.ResourcePath, instanceId, crossVPCPath, modifyPath), body, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res CrossVPCModifyResp
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type CrossVPCModifyResp struct {
	// Result of the cross-VPC access modification.
	Success bool `json:"success"`
	// Details of the result of the cross-VPC access modification.
	Results []*Result `json:"results"`
}

type Result struct {
	// advertised.listeners IP address or domain name.
	AdvertisedIp string `json:"advertised_ip"`
	// Status of the cross-VPC access modification.
	Success bool `json:"success"`
	// Listeners IP address.
	Ip string `json:"ip"`
}
