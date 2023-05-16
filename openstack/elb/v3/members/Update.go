package members

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

// UpdateOpts is the common options' struct used in this package's Update operation.
type UpdateOpts struct {
	// Specifies the administrative status of the backend server.
	//
	// Although this parameter can be used in the APIs for creating and updating backend servers, its actual value depends on whether cloud servers exist. If cloud servers exist, the value is true. Otherwise, the value is false.
	//
	// This parameter is unsupported. Please do not use it.
	AdminStateUp *bool `json:"admin_state_up,omitempty"`
	// Specifies the backend server name.
	//
	// Minimum: 0
	//
	// Maximum: 255
	Name string `json:"name,omitempty"`
	// Specifies the weight of the backend server. Requests are routed to backend servers in the same backend server group based on their weights.
	//
	// The value ranges from 0 to 100, and the default value is 1. The larger the weight is, the higher proportion of requests the backend server receives. If the weight is set to 0, the backend server will not accept new requests.
	//
	// If lb_algorithm is set to SOURCE_IP, this parameter will not take effect.
	//
	// Minimum: 0
	//
	// Maximum: 100
	Weight *int `json:"weight,omitempty"`
}

// Update allows Member to be updated.
func Update(client *golangsdk.ServiceClient, poolID string, memberID string, opts UpdateOpts) (*Member, error) {
	b, err := build.RequestBody(opts, "member")
	if err != nil {
		return nil, err
	}

	// PUT /v3/{project_id}/elb/pools/{pool_id}/members/{member_id}
	raw, err := client.Put(client.ServiceURL("pools", poolID, "members", memberID), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return extra(err, raw)
}
