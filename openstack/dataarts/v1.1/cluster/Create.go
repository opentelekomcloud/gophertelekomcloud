package cluster

import (
	"net/http"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	tag "github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
)

type CreateOpts struct {
	Cluster    Cluster `json:"cluster" required:"true"`
	AutoRemind *bool   `json:"auto_remind,omitempty"`
	PhoneNum   string  `json:"phone_num,omitempty"`
	Email      string  `json:"email,omitempty"`
}

type Cluster struct {
	ScheduleBootTime   string            `json:"scheduleBootTime,omitempty"`
	IsScheduleBootOff  *bool             `json:"isScheduleBootOff,omitempty"`
	Instances          []Instance        `json:"instances,omitempty"`
	DataStore          *Datastore        `json:"datastore,omitempty"`
	ExtendedProperties *ExtendedProp     `json:"extended_properties,omitempty"`
	ScheduleOffTime    string            `json:"scheduleOffTime,omitempty"`
	VpcId              string            `json:"vpcId,omitempty"`
	Name               string            `json:"name,omitempty"`
	SysTags            []tag.ResourceTag `json:"sys_tags,omitempty"`
	IsAutoOff          *bool             `json:"isAutoOff"`
}

type Instance struct {
	AZ        string `json:"availability_zone" required:"true"`
	Nics      []Nic  `json:"nics" required:"true"`
	FlavorRef string `json:"flavorRef" required:"true"`
	Type      string `json:"type" required:"true"`
}

type Datastore struct {
	Type    string `json:"type,omitempty"`
	Version string `json:"version,omitempty"`
}

type ExtendedProp struct {
	WorkSpaceId string `json:"workSpaceId,omitempty"`
	ResourceId  string `json:"resourceId,omitempty"`
	Trial       string `json:"trial,omitempty"`
}

type Nic struct {
	SecurityGroupId string `json:"securityGroupId" required:"true"`
	NetId           string `json:"net-id" required:"true"`
}

func Create(client *golangsdk.ServiceClient, opts CreateOpts) (*ClusterResponse, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// POST /v1.1/{project_id}/clusters
	raw, err := client.Post(client.ServiceURL("clusters"), b, nil, &golangsdk.RequestOpts{
		MoreHeaders: map[string]string{"Content-Type": "application/json", "X-Language": "en"},
	})
	return extra(err, raw)
}

type ClusterResponse struct {
	Name      string         `json:"name"`
	Id        string         `json:"id"`
	Task      Task           `json:"task"`
	Datastore Datastore      `json:"datastore"`
	Instances []InstanceResp `json:"instances"`
}

type Task struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type InstanceResp struct {
	Name    string `json:"name"`
	Id      string `json:"id"`
	Type    string `json:"type"`
	ShardId string `json:"shard_id"`
}

func extra(err error, raw *http.Response) (*ClusterResponse, error) {
	if err != nil {
		return nil, err
	}

	var res ClusterResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}
