package loadbalancers

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// GetStatuses will return the status of a particular LoadBalancer.
func GetStatuses(client *golangsdk.ServiceClient, id string) (*LoadBalancerStatus, error) {
	// GET /v3/{project_id}/elb/loadbalancers/{loadbalancer_id}/statuses
	raw, err := client.Get(client.ServiceURL("loadbalancers", id, "statuses"), nil, nil)
	if err != nil {
		return nil, err
	}

	var res struct {
		LoadBalancer LoadBalancerStatus `json:"loadbalancer"`
	}
	err = extract.IntoStructPtr(raw.Body, &res, "statuses")
	return &res.LoadBalancer, err
}

type LoadBalancerStatus struct {
	// Specifies the load balancer name.
	//
	// Minimum: 1
	//
	// Maximum: 255
	Name string `json:"name"`
	// Specifies the provisioning status of the load balancer. The value can be ACTIVE or PENDING_DELETE.
	//
	// ACTIVE: The load balancer is successfully provisioned.
	//
	// PENDING_DELETE: The load balancer is being deleted.
	ProvisioningStatus string `json:"provisioning_status"`
	// Lists the listeners added to the load balancer.
	Listeners []LoadBalancerStatusListener `json:"listeners"`
	// Lists the backend server groups associated with the load balancer.
	Pools []LoadBalancerStatusPool `json:"pools"`
	// Specifies the load balancer ID.
	Id string `json:"id"`
	// Specifies the operating status of the load balancer.
	//
	// The value can only be one of the following:
	//
	// ONLINE (default): The load balancer is running normally.
	//
	// FROZEN: The load balancer has been frozen.
	//
	// DEGRADED: This status is displayed only when operating_status is set to OFFLINE for a backend server associated with the load balancer and the API for querying the load balancer status tree is called.
	//
	// DISABLED: This status is displayed only when admin_state_up of the load balancer is set to false.
	//
	// DEGRADED and DISABLED are returned only when the API for querying the load balancer status tree is called.
	OperatingStatus string `json:"operating_status"`
}

type LoadBalancerStatusListener struct {
	// Specifies the name of the listener added to the load balancer.
	//
	// Minimum: 1
	//
	// Maximum: 255
	Name string `json:"name"`
	// Specifies the provisioning status of the listener. The value can only be ACTIVE, indicating that the listener is successfully provisioned.
	ProvisioningStatus string `json:"provisioning_status"`
	// Specifies the operating status of the backend server group associated with the listener.
	Pools []LoadBalancerStatusPool `json:"pools"`
	// Specifies the operating status of the forwarding policy added to the listener.
	L7policies []LoadBalancerStatusPolicy `json:"l7policies"`
	// Specifies the listener ID.
	Id string `json:"id"`
	// Specifies the operating status of the listener.
	//
	// The value can only be one of the following:
	//
	// ONLINE (default): The listener is running normally.
	//
	// DEGRADED: This status is displayed only when provisioning_status of a forwarding policy or a forwarding rule added to the listener is set to ERROR or operating_status is set to OFFLINE for a backend server associated with the listener.
	//
	// DISABLED: This status is displayed only when admin_state_up of the load balancer or of the listener is set to false. Note: DEGRADED and DISABLED are returned only when the API for querying the load balancer status tree is called.
	OperatingStatus string `json:"operating_status"`
}

type LoadBalancerStatusPolicy struct {
	// Specifies whether requests are forwarded to another backend server group or redirected to an HTTPS listener. The value can be one of the following:
	//
	// REDIRECT_TO_POOL: Requests are forwarded to another backend server group.
	//
	// REDIRECT_TO_LISTENER: Requests are redirected to an HTTPS listener.
	Action string `json:"action"`
	// Specifies the forwarding policy ID.
	Id string `json:"id"`
	// Specifies the provisioning status of the forwarding policy.
	//
	// ACTIVE (default): The forwarding policy is provisioned successfully.
	//
	// ERROR: Another forwarding policy of the same listener has the same forwarding rule.
	ProvisioningStatus string `json:"provisioning_status"`
	// Specifies the policy name.
	//
	// Minimum: 1
	//
	// Maximum: 255
	Name string `json:"name"`
	// Specifies the forwarding rule.
	Rules []LoadBalancerStatusL7Rule `json:"rules"`
}

type LoadBalancerStatusL7Rule struct {
	// Specifies the ID of the forwarding rule.
	Id string `json:"id"`
	// Specifies the type of the match content. The value can be HOST_NAME or PATH.
	//
	// HOST_NAME: A domain name will be used for matching.
	//
	// PATH: A URL will be used for matching.
	//
	// The value must be unique for each forwarding rule in a forwarding policy.
	Type string `json:"type"`
	// Specifies the provisioning status of the forwarding rule.
	//
	// ACTIVE (default): The forwarding rule is successfully provisioned.
	//
	// ERROR: Another forwarding policy of the same listener has the same forwarding rule.
	ProvisioningStatus string `json:"provisioning_status"`
}

type LoadBalancerStatusPool struct {
	// Specifies the provisioning status of the backend server group. The value can only be ACTIVE, indicating that the backend server group is successfully provisioned.
	ProvisioningStatus string `json:"provisioning_status"`
	// Specifies the name of the backend server group.
	//
	// Minimum: 1
	//
	// Maximum: 255
	Name string `json:"name"`
	// Specifies the health check results of backend servers in the load balancer status tree.
	HealthMonitor LoadBalancerStatusHealthMonitor `json:"healthmonitor"`
	// Specifies the backend server.
	Members []LoadBalancerStatusMember `json:"members"`
	// Specifies the ID of the backend server group.
	Id string `json:"id"`
	// Specifies the operating status of the backend server group.
	//
	// The value can be one of the following:
	//
	// ONLINE: The backend server group is running normally.
	//
	// DEGRADED: This status is displayed only when operating_status of a backend server in the backend server group is set to OFFLINE.
	//
	// DISABLED: This status is displayed only when admin_state_up of the backend server group or of the associated load balancer is set to false.
	//
	// Note: DEGRADED and DISABLED are returned only when the API for querying the load balancer status tree is called.
	OperatingStatus string `json:"operating_status"`
}

type LoadBalancerStatusHealthMonitor struct {
	// Specifies the health check protocol. The value can be TCP, UDP_CONNECT, or HTTP.
	Type string `json:"type"`
	// Specifies the health check ID.
	Id string `json:"id"`
	// Specifies the health check name.
	//
	// Minimum: 1
	//
	// Maximum: 255
	Name string `json:"name"`
	// Specifies the provisioning status of the health check. The value can only be ACTIVE, indicating that the health check is successfully provisioned.
	ProvisioningStatus string `json:"provisioning_status"`
}

type LoadBalancerStatusMember struct {
	// Specifies the provisioning status of the backend server. The value can only be ACTIVE, indicating that the backend server is successfully provisioned.
	ProvisioningStatus string `json:"provisioning_status"`
	// Specifies the private IP address bound to the backend server.
	Address string `json:"address"`
	// Specifies the port used by the backend server to receive requests. The port number ranges from 1 to 65535.
	ProtocolPort *int `json:"protocol_port"`
	// Specifies the backend server ID.
	Id string `json:"id"`
	// Specifies the operating status of the backend server.
	//
	// The value can be one of the following:
	//
	// ONLINE: The backend server is running normally.
	//
	// NO_MONITOR: No health check is configured for the backend server group to which the backend server belongs.
	//
	// DISABLED: The backend server is not available. This status is displayed only when admin_state_up of the backend server, or the backend server group to which it belongs, or the associated load balancer is set to false and the API for querying the load balancer status tree is called.
	//
	// OFFLINE: The cloud server used as the backend server is stopped or does not exist.
	OperatingStatus string `json:"operating_status"`
}
