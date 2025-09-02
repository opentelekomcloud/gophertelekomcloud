package load_balancer

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type LoadBalancerResp struct {
	// The cluster agency
	Agency string `json:"agency"`
	// Server certificate name
	ServerCertName string `json:"serverCertName"`
	// Server certificate ID
	ServerCertId string `json:"serverCertId"`
	// CA certificate name
	CacertName string `json:"cacertName"`
	// CA certificate ID
	CacertId string `json:"cacertId"`
	// Indicates whether ELB is enabled
	Enabled bool `json:"elb_enable"`
	// Authentication mode
	AuthenticationType string `json:"authentication_type"`
	// Load balancer object information
	LoadBalancer LoadBalancer `json:"loadBalancer"`
	// Start time of automatic log backup.
	Listener Elbv3Listener `json:"listener"`
	// Indicates whether to enable the log function.
	Healthmonitors []*Member `json:"healthmonitors"`
}

type LoadBalancer struct {
	// Load balancer ID.
	Id string `json:"id"`
	// Load balancer name
	Name string `json:"name"`
	// Whether the LB is for dedicated use. The value can be false (shared) or true (dedicated).
	Guaranteed string `json:"guaranteed"`
	// Resource billing information. If the value is left blank, the resource will be billed in pay-per-use mode. If the value is not left blank,
	// the resource is billed on a yearly/monthly basis.
	BillingInfo string `json:"billing_info"`
	// Description.
	Description string `json:"description"`
	// ID of the VPC to which the load balancer belongs
	VpcId string `json:"vpc_id"`
	// Provisioning status of the load balancer
	ProvisioningStatus string `json:"provisioning_status"`
	// Associated listener list
	Listeners []*IdListWrapper `json:"listeners"`
	// IPv4 virtual IP address bound to the load balancer
	VipAddress string `json:"vip_address"`
	// Port ID bound to the private IPv4 IP address of the load balancer.
	VipPortId string `json:"vip_port_id"`
	// IPv6 address of the load balancer.
	Ipv6VipAddress string `json:"ipv6_vip_address"`
	// EIP bound to the load balancer
	Publicips []*PublicIpInfo `json:"publicips"`
}

type Elbv3Listener struct {
	// Listener id.
	Id string `json:"id"`
	// Listener name.
	Name string `json:"name"`
	// Protocol used by the listener
	Protocol string `json:"protocol"`
	// Port used by the listener.
	ProtocolPort int `json:"protocol_port"`
	// ipgroup information in the listener object.
	Ipgroup *ListenerIpGroup `json:"ipgroup"`
}

type Member struct {
	// Member Id
	Id string `json:"id"`
	// Specifies the backend server name.
	Name string `json:"name"`
	// Private IP address bound to the backend server.
	Address string `json:"address"`
	// Specifies the port used by the backend server.
	ProtocolPort int `json:"protocol_port"`
	// Specifies the operating status of the backend server. - ONLINE: The backend server is running normally.
	// NO_MONITOR: No health check is configured for the backend server group to which the backend server belongs.
	// OFFLINE: The cloud server used as the backend server is stopped or does not exist.
	OperatingStatus string `json:"operating_status"`
	// ID of the instance used as the backend server. If this parameter is left blank, the backend server is not an ECS.
	InstanceId string `json:"instance_id"`
}

type IdListWrapper struct {
	// Listener id.
	Id string `json:"id"`
}

type PublicIpInfo struct {
	// EIP configuration ID
	PublicIpId string `json:"publicip_id"`
	// Specifies the EIP.
	PublicIpAddress string `json:"publicip_address"`
	// IP address version. Value range: 4 and 6. 4 indicates IPv4, and 6 indicates IPv6.
	IpVersion int `json:"ip_version"`
}

type ListenerIpGroup struct {
	// ID of the IP address group associated with the listener This parameter is mandatory during creation and is optional during update.
	IpgroupId string `json:"ipgroup_id"`
	// Status of an access control group. True: Enable access control. False: Disable access control.
	EnableIpgroup bool `json:"enable_ipgroup"`
}

// This API is used to obtain information about the load balancers of a cluster.
func Get(client *golangsdk.ServiceClient, id string) (*LoadBalancerResp, error) {
	raw, err := client.Get(client.ServiceURL("clusters", id, "es-listeners"), nil, &golangsdk.RequestOpts{
		MoreHeaders: map[string]string{
			"Content-Type": "application/json",
		},
	})
	if err != nil {
		return nil, err
	}

	var res LoadBalancerResp
	err = extract.Into(raw.Body, &res)
	return &res, err
}
