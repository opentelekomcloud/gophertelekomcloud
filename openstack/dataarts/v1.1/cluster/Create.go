package cluster

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	tag "github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
)

type CreateOpts struct {
	Cluster Cluster `json:"cluster" required:"true"`
	// AutoRemind Whether to enable message notification.
	// If you enable this function, you can configure a maximum of five mobile numbers or email addresses.
	// You will be notified of table/file migration job failures and EIP exceptions by SMS message or email.
	AutoRemind bool `json:"auto_remind,omitempty"`
	// PhoneNum Mobile number for receiving notifications.
	PhoneNum string `json:"phone_num,omitempty"`
	// Email address for receiving notifications.
	Email string `json:"email,omitempty"`
	// Request language.
	XLang string `json:"-"`
}

type Cluster struct {
	// ScheduleBootTime Time for scheduled startup of a CDM cluster. The CDM cluster starts at this time every day.
	ScheduleBootTime string `json:"scheduleBootTime,omitempty"`
	// IsScheduleBootOff Whether to enable scheduled startup/shutdown. The scheduled startup/shutdown and auto shutdown functions cannot be enabled at the same time.
	IsScheduleBootOff *bool `json:"isScheduleBootOff,omitempty"`
	// Instances Node list.
	Instances []Instance `json:"instances,omitempty"`
	// DataStore Cluster information.
	DataStore *Datastore `json:"datastore,omitempty"`
	// ExtendedProperties Extended attribute.
	ExtendedProperties *ExtendedProp `json:"extended_properties,omitempty"`
	// ScheduleOffTime Time for scheduled shutdown of a CDM cluster. The CDM cluster shuts down directly at this time every day without waiting for unfinished jobs to complete.
	ScheduleOffTime string `json:"scheduleOffTime,omitempty"`
	// VpcId VPC ID, which is used for configuring a network for the cluster.
	VpcId string `json:"vpcId,omitempty"`
	// Name Cluster name.
	Name string `json:"name,omitempty"`
	// SysTags Enterprise project information. For details, see the descriptions of sys_tags parameters.
	SysTags []tag.ResourceTag `json:"sys_tags,omitempty"`
	// IsAutoOff Whether to enable auto shutdown. The auto shutdown and scheduled startup/shutdown functions cannot be enabled at the same time.
	// When auto shutdown is enabled, if no job is running in the cluster and no scheduled job is available, a cluster will be automatically shut down 15 minutes after it starts running, which reduces costs for you.
	IsAutoOff bool `json:"isAutoOff,omitempty"`
}

type Instance struct {
	// AZ availability zone where a cluster is located.
	AZ string `json:"availability_zone" required:"true"`
	// Nics NIC list. A maximum of two NICs are supported. For details, see the descriptions of nics parameters.
	Nics []Nic `json:"nics" required:"true"`
	// FlavorRef Instance flavor.
	FlavorRef string `json:"flavorRef" required:"true"`
	// Type Node type. Currently, only cdm is available.
	Type string `json:"type" required:"true"`
}

type Nic struct {
	// SecurityGroupId Security group ID.
	SecurityGroupId string `json:"securityGroupId" required:"true"`
	// NetId Subnet ID.
	NetId string `json:"net-id" required:"true"`
}

type Datastore struct {
	// Type Generally, the value is cdm.
	Type string `json:"type,omitempty"`
	// Version Cluster version.
	Version string `json:"version,omitempty"`
}

type ExtendedProp struct {
	// WorkSpaceId Workspace ID.
	WorkSpaceId string `json:"workSpaceId,omitempty"`
	// ResourceId Resource ID.
	ResourceId string `json:"resourceId,omitempty"`
	// Trial Whether the cluster is a trial cluster.
	Trial string `json:"trial,omitempty"`
}

// Create is used to create a cluster.
// Send request POST /v1.1/{project_id}/clusters
func Create(client *golangsdk.ServiceClient, reqOpts CreateOpts) (*ClusterResp, error) {
	b, err := build.RequestBody(reqOpts, "")
	if err != nil {
		return nil, err
	}

	opts := &golangsdk.RequestOpts{
		MoreHeaders: map[string]string{HeaderContentType: ApplicationJson, HeaderXLanguage: reqOpts.XLang},
		OkCodes:     []int{202},
	}

	raw, err := client.Post(client.ServiceURL(clustersEndpoint), b, nil, opts)

	if err != nil {
		return nil, err
	}

	var res ClusterResp
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type ClusterResp struct {
	// Cluster name
	Name string `json:"name"`
	// Cluster ID
	Id string `json:"id"`
	// Task information
	Task Task `json:"task"`
	// Cluster information
	Datastore Datastore `json:"datastore"`
	// Cluster node information
	Instances []ClusterInstance `json:"instances"`
}

type Task struct {
	// Task ID
	Id string `json:"id"`
	// Task name
	Name string `json:"name"`
}

type ClusterInstance struct {
	// Node VM ID
	Id string `json:"id"`
	// Name of the VM on the node
	Name string `json:"name"`
	// Node type. Currently, only cdm is available.
	Type string `json:"type"`
	// Shard ID
	ShardId string `json:"shard_id"`
}
