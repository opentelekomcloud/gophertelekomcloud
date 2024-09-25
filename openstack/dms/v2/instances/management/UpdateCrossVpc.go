package management

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// CrossVpcUpdateOpts is the structure required by the UpdateCrossVpc method to update the internal IP address for
// cross-VPC access.
type CrossVpcUpdateOpts struct {
	// User-defined advertised IP contents key-value pair.
	// The key is the listeners IP.
	// The value is advertised.listeners IP, or domain name.
	Contents map[string]string `json:"advertised_ip_contents" required:"true"`
}

// UpdateCrossVpc is a method to update the internal IP address for cross-VPC access using given parameters.
func UpdateCrossVpc(client *golangsdk.ServiceClient, instanceId string, opts CrossVpcUpdateOpts) (*CrossVpc, error) {
	body, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL("instances", instanceId, "crossvpc", "modify"), body, nil, nil)
	if err != nil {
		return nil, err
	}

	var r CrossVpc
	err = extract.Into(raw.Body, &r)

	return &r, err
}

// CrossVpc is the structure that represents the API response of 'UpdateCrossVpc' method.
type CrossVpc struct {
	// The result of cross-VPC access modification.
	Success bool `json:"success"`
	// The result list of broker cross-VPC access modification.
	Connections []Connection `json:"results"`
}

// Connection is the structure that represents the detail of the cross-VPC access.
type Connection struct {
	// advertised.listeners IP/domain name.
	AdvertisedIp string `json:"advertised_ip"`
	// The status of broker cross-VPC access modification.
	Success bool `json:"success"`
	// Listeners IP.
	ListenersIp string `json:"ip"`
}
