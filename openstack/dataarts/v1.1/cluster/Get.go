package cluster

import (
	"net/http"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func Get(client *golangsdk.ServiceClient, clusterId string) (*ClusterQuery, error) {
	// GET /v1.1/{project_id}/clusters/{cluster_id}
	raw, err := client.Get(client.ServiceURL("clusters", clusterId), nil, nil)
	return extraResp(err, raw)
}

func extraResp(err error, raw *http.Response) (*ClusterQuery, error) {
	if err != nil {
		return nil, err
	}

	var res ClusterQuery
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type ClusterQuery struct {
	PublicEndpoint           string               `json:"public_endpoint"`
	Instances                []DetailedInstances  `json:"instances"`
	SecurityGroupId          string               `json:"security_group_id"`
	SubnetId                 string               `json:"subnet_id"`
	VpcId                    string               `json:"vpc_id"`
	CustomerConfig           CustomerConfig       `json:"customerConfig"`
	Datastore                Datastore            `json:"datastore"`
	IsAutoOff                bool                 `json:"isAutoOff"`
	PublicEndpointDomainName string               `json:"publicEndpointDomainName"`
	BakExpectedStartTime     string               `json:"bakExpectedStartTime"`
	BakKeepDay               string               `json:"bakKeepDay"`
	MaintainWindow           MaintainWindow       `json:"maintainWindow"`
	RecentEvent              int                  `json:"recentEvent"`
	FlavorName               string               `json:"flavorName"`
	AzName                   string               `json:"azName"`
	EndpointDomainName       string               `json:"endpointDomainName"`
	PublicEndpointStatus     PublicEndpointStatus `json:"publicEndpointStatus"`
	IsScheduleBootOff        bool                 `json:"isScheduleBootOff"`
	Namespace                string               `json:"namespace"`
	EipId                    string               `json:"eipId"`
	FailedReasons            FailedReasons        `json:"failedReasons"`
	DbUser                   string               `json:"dbuser"`
	Links                    []ClusterLinks       `json:"links"`
	ClusterMode              string               `json:"clusterMode"`
	Task                     ClusterTask          `json:"task"`
	Created                  string               `json:"created"`
	StatusDetail             string               `json:"statusDetail"`
	ConfigStatus             string               `json:"config_status"`
	ActionProgress           ActionProgress       `json:"actionProgress"`
	Name                     string               `json:"name"`
	Id                       string               `json:"id"`
	IsFrozen                 string               `json:"isFrozen"`
	Actions                  []string             `json:"actions"`
	Updated                  string               `json:"updated"`
	Status                   string               `json:"status"`
}

type DetailedInstances struct {
	Flavor        Flavor         `json:"flavor"`
	Volume        Volume         `json:"volume"`
	Status        string         `json:"status"`
	Actions       []string       `json:"actions"`
	Type          string         `json:"string"`
	Name          string         `json:"name"`
	Id            string         `json:"id"`
	IsFrozen      string         `json:"isFrozen"`
	Components    string         `json:"components"`
	ConfigStatus  string         `json:"config_status"`
	Role          string         `json:"role"`
	Group         string         `json:"group"`
	Links         []ClusterLinks `json:"links"`
	ParamsGroupId string         `json:"paramsGroupId"`
	PublicIp      string         `json:"publicIp"`
	ManageIp      string         `json:"manageIp"`
	TrafficIp     string         `json:"trafficIp"`
	ShardId       string         `json:"shard_id"`
	ManageFixIp   string         `json:"manage_fix_ip"`
	PrivateIp     string         `json:"private_ip"`
	InternalIp    string         `json:"internal_ip"`
	Resource      []Resource     `json:"resource"`
}

type Flavor struct {
	Id    string         `json:"id"`
	Links []ClusterLinks `json:"links"`
}

type Volume struct {
	Type string `json:"type"`
	Size int64  `json:"size"`
}

type Resource struct {
	ResourceId   string `json:"resource_id"`
	ResourceType string `json:"resource_type"`
}

type CustomerConfig struct {
	FailureRemind   string `json:"failureRemind"`
	ClusterName     string `json:"clusterName"`
	ServiceProvider string `json:"serviceProvider"`
	LocalDisk       string `json:"localDisk"`
	Ssl             string `json:"ssl"`
	CreateFrom      string `json:"createFrom"`
	ResourceId      string `json:"resourceId"`
	FlavorType      string `json:"flavorType"`
	WorkSpaceId     string `json:"workSpaceId"`
	Trial           string `json:"trial"`
}

type MaintainWindow struct {
	Dat       string `json:"day"`
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
}

type PublicEndpointStatus struct {
	Status       string `json:"status"`
	ErrorMessage string `json:"errorMessage"`
}

type FailedReasons struct {
	CreateFailed CreateFailed `json:"CREATE_FAILED"`
}

type CreateFailed struct {
	ErrorMsg  string `json:"errorMsg"`
	ErrorCode string `json:"errorCode"`
}

type ClusterLinks struct {
	Rel  string `json:"rel"`
	Href string `json:"href"`
}

type ClusterTask struct {
	Description string `json:"description"`
	Name        string `json:"name"`
	Id          string `json:"id"`
}

type ActionProgress struct {
	Creating     string `json:"creating"`
	Growing      string `json:"growing"`
	Restoring    string `json:"restoring"`
	Snapshotting string `json:"snapshotting"`
	Repairing    string `json:"repairing"`
}
