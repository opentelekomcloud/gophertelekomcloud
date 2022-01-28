package cluster

import "github.com/opentelekomcloud/gophertelekomcloud"

type Cluster struct {
	ClusterID             string         `json:"clusterId"`
	ClusterName           string         `json:"clusterName"`
	MasterNodeNum         string         `json:"masterNodeNum"`
	CoreNodeNum           string         `json:"coreNodeNum"`
	TotalNodeNum          string         `json:"totalNodeNum"`
	ClusterState          string         `json:"clusterState"`
	CreateAt              string         `json:"createAt"`
	UpdateAt              string         `json:"updateAt"`
	BillingType           string         `json:"billingType"`
	DataCenter            string         `json:"dataCenter"`
	Vpc                   string         `json:"vpc"`
	Duration              string         `json:"duration"`
	Fee                   string         `json:"fee"`
	HadoopVersion         string         `json:"hadoopVersion"`
	MasterNodeSize        string         `json:"masterNodeSize"`
	CoreNodeSize          string         `json:"coreNodeSize"`
	ComponentList         []Component    `json:"componentList"`
	ExternalIp            string         `json:"externalIp"`
	ExternalAlternateIp   string         `json:"externalAlternateIp"`
	InternalIp            string         `json:"internalIp"`
	DeploymentID          string         `json:"deploymentId"`
	Remark                string         `json:"remark"`
	OrderID               string         `json:"orderId"`
	AzID                  string         `json:"azId"`
	MasterNodeProductID   string         `json:"masterNodeProductId"`
	MasterNodeSpecID      string         `json:"masterNodeSpecId"`
	CoreNodeProductID     string         `json:"coreNodeProductId"`
	CoreNodeSpecID        string         `json:"coreNodeSpecId"`
	AzName                string         `json:"azName"`
	InstanceID            string         `json:"instanceId"`
	Vnc                   string         `json:"vnc"`
	TenantID              string         `json:"tenantId"`
	VolumeSize            int            `json:"volumeSize"`
	SubnetName            string         `json:"subnetName"`
	SecurityGroupsID      string         `json:"securityGroupsId"`
	SlaveSecurityGroupsID string         `json:"slaveSecurityGroupsId"`
	StageDesc             string         `json:"stageDesc"`
	MrsManagerFinish      bool           `json:"mrsManagerFinish"`
	SafeMode              int            `json:"safeMode"`
	ClusterVersion        string         `json:"clusterVersion"`
	NodePublicCertName    string         `json:"nodePublicCertName"`
	MasterNodeIp          string         `json:"masterNodeIp"`
	PrivateIpFirst        string         `json:"privateIpFirst"`
	ErrorInfo             string         `json:"errorInfo"`
	ChargingStartTime     string         `json:"chargingStartTime"`
	ClusterType           int            `json:"clusterType"`
	LogCollection         int            `json:"logCollection"`
	MasterDataVolumeType  string         `json:"masterDataVolumeType"`
	MasterDataVolumeSize  int            `json:"masterDataVolumeSize"`
	MasterDataVolumeCount int            `json:"masterDataVolumeCount"`
	CoreDataVolumeType    string         `json:"coreDataVolumeType"`
	CoreDataVolumeSize    int            `json:"coreDataVolumeSize"`
	CoreDataVolumeCount   int            `json:"coreDataVolumeCount"`
	Scale                 string         `json:"scale"`
	BootstrapScripts      []ScriptResult `json:"bootstrapScripts"`
}

type Component struct {
	ComponentID      string `json:"componentId"`
	ComponentName    string `json:"componentName"`
	ComponentVersion string `json:"componentVersion"`
	ComponentDesc    string `json:"componentDesc"`
}

type ScriptResult struct {
	Name                 string   `json:"name"`
	Uri                  string   `json:"uri"`
	Parameters           string   `json:"parameters"`
	Nodes                []string `json:"nodes"`
	ActiveMaster         bool     `json:"active_master"`
	BeforeComponentStart bool     `json:"before_component_start"`
	FailAction           string   `json:"fail_action"`
	StartTime            int      `json:"start_time"`
	State                string   `json:"state"`
}

type CreateClusterResult struct {
	ClusterID string `json:"cluster_id"`
	Result    bool   `json:"result"`
	Msg       string `json:"msg"`
}

type CreateResult struct {
	golangsdk.Result
}

func (r CreateResult) Extract() (*CreateClusterResult, error) {
	s := new(CreateClusterResult)
	err := r.ExtractIntoStructPtr(s, "")
	if err != nil {
		return nil, err
	}
	return s, nil
}

type GetResult struct {
	golangsdk.Result
}

func (r GetResult) Extract() (*Cluster, error) {
	s := new(Cluster)
	err := r.ExtractIntoStructPtr(s, "cluster")
	if err != nil {
		return nil, err
	}
	return s, nil

}

type DeleteResult struct {
	golangsdk.ErrResult
}
