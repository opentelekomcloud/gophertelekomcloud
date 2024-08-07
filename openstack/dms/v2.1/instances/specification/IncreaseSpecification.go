package specification

import (
	"fmt"
	"strings"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dms/v2.1/instances"
)

const extendPath = "extend"

type IncreaseSpecOpts struct {
	// Message engine.
	Engine string `json:"-" required:"true"`
	// Type of the change.
	// Value:
	//    storage: Expand the storage space without adding brokers.
	//    horizontal: Add brokers without resizing the storage space of each broker.
	//    vertical: Modify the underlying flavor of brokers without adding brokers or storage space.
	OperType string `json:"oper_type" required:"true"`
	// New storage space.
	// This parameter is valid and mandatory when oper_type is set to storage or horizontal.
	// Instance storage space = Number of brokers x Storage space of each broker.
	// If oper_type is set to storage, the number of brokers remains unchanged, and the storage space of each broker must be expanded by at least 100 GB.
	// If oper_type is set to horizontal, the storage space of each broker remains unchanged.
	NewStorageSpace int `json:"new_storage_space"`
	// This parameter is valid only when oper_type is set to horizontal.
	// A maximum of 30 brokers are supported.
	NewBrokerNumber int `json:"new_broker_number"`
	// New product ID for scale-up.
	// This parameter is valid and mandatory when oper_type is set to vertical.
	// Obtain the product ID from Querying Product Specifications List.
	NewProductId string `json:"new_product_id"`
	// ID of the EIP bound to the instance.
	// Use commas (,) to separate multiple EIP IDs.
	// This parameter is mandatory when oper_type is set to horizontal.
	PublicIpId string `json:"publicip_id"`
	// Specified IPv4 private IP addresses.
	// The number of specified IP addresses must be less than or equal to the number of new brokers.
	// If the number of specified IP addresses is less than the number of brokers, the unspecified brokers are randomly assigned private IP addresses.
	TenantIps []string `json:"tenant_ips"`
	// ID of the standby subnet used by new brokers in instance expansion.
	// This value is transferred when a standby subnet is used in instance expansion.
	// Contact customer service to use the value.
	SecondTenantSubnetId string `json:"second_tenant_subnet_id"`
}

// IncreaseSpec is used to modify instance specifications.
// Send POST /v2/{engine}/{project_id}/instances/{instance_id}/extend
func IncreaseSpec(client *golangsdk.ServiceClient, instanceId string, opts IncreaseSpecOpts) (*IncreaseSpecResp, error) {
	body, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// Here we should patch a client, because for an increasing url path is different.
	// For all requests we use schema	/v2/{project_id}/instances
	// But for these					/v2/{engine}/{project_id}/instances
	paths := strings.SplitN(client.Endpoint, "v2", 2)
	url := fmt.Sprintf("%sv2/%s%s%s/%s/%s", paths[0], opts.Engine, paths[1], instances.ResourcePath, instanceId, extendPath)

	raw, err := client.Post(url, body, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res IncreaseSpecResp
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type IncreaseSpecResp struct {
	// ID of the specification modification task.
	JobId string `json:"job_id"`
}
