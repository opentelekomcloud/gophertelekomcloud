package gateway

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func Get(client *golangsdk.ServiceClient, id string) (*Gateway, error) {
	raw, err := client.Get(client.ServiceURL("apigw/instances", id), nil, nil)
	if err != nil {
		return nil, err
	}

	var res Gateway
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type Gateway struct {
	ID                    string `json:"id"`
	ProjectID             string `json:"project_id"`
	InstanceName          string `json:"instance_name"`
	Status                string `json:"status"`
	InstanceStatus        int    `json:"instance_status"`
	Type                  string `json:"type"`
	Spec                  string `json:"spec"`
	CreateTime            int64  `json:"create_time"`
	EipAddress            string `json:"eip_address"`
	ChargingMode          int    `json:"charging_mode"`
	LoadbalancerProvider  string `json:"loadbalancer_provider"`
	Description           string `json:"description"`
	VpcID                 string `json:"vpc_id"`
	SubnetID              string `json:"subnet_id"`
	SecurityGroupID       string `json:"security_group_id"`
	MaintainBegin         string `json:"maintain_begin"`
	MaintainEnd           string `json:"maintain_end"`
	IngressIp             string `json:"ingress_ip"`
	UserID                string `json:"user_id"`
	NatEipAddress         string `json:"nat_eip_address"`
	BandwidthSize         int    `json:"bandwidth_size"`
	BandwidthChargingMode string `json:"bandwidth_charging_mode"`
	AvailableZoneIDs      string `json:"available_zone_ids"`
	InstanceVersion       string `json:"instance_version"`
	VirsubnetID           string `json:"virsubnet_id"`
	RomaEipAddress        string `json:"roma_eip_address"`
	// Listeners                    *Listeners        `json:"listeners"`
	SupportedFeatures            []string          `json:"supported_features"`
	EndpointService              *EndpointService  `json:"endpoint_service"`
	EndpointServices             []EndpointService `json:"endpoint_services"`
	NodeIps                      *NodeIps          `json:"node_ips"`
	PublicIps                    []IpDetail        `json:"publicips"`
	PrivateIps                   []IpDetail        `json:"privateips"`
	IsReleasable                 *bool             `json:"is_releasable"`
	IngressBandwidthChargingMode string            `json:"ingress_bandwidth_charging_mode"`
}

type EndpointService struct {
	ServiceName string `json:"service_name"`
	CreatedAt   string `json:"created_at"`
}

type NodeIps struct {
	LiveData []string `json:"livedata"`
	Shubao   []string `json:"shubao"`
}

type IpDetail struct {
	IpAddress     string `json:"ip_address"`
	BandwidthSize int    `json:"bandwidth_size"`
}
