package v1

type ListClusterDetailsRequest struct {
	//
	ClusterId string `json:"cluster_id"`
}

// GET /v1.0/{project_id}/clusters/{cluster_id}

type ListClusterDetailsResponse struct {
	Cluster ClusterDetail `json:"cluster,omitempty"`
}

type ClusterDetail struct {
	//
	Id string `json:"id"`
	//
	Name string `json:"name"`
	//
	Status string `json:"status"`
	//
	Version string `json:"version"`
	//
	Updated string `json:"updated"`
	//
	Created string `json:"created"`
	//
	Port int32 `json:"port"`
	//
	Endpoints []Endpoints `json:"endpoints"`
	//
	Nodes []Nodes `json:"nodes"`
	//
	Tags []Tags `json:"tags"`
	//
	UserName string `json:"user_name"`
	//
	NumberOfNode int32 `json:"number_of_node"`
	//
	RecentEvent int32 `json:"recent_event"`
	//
	AvailabilityZone string `json:"availability_zone"`
	//
	EnterpriseProjectId string `json:"enterprise_project_id"`
	//
	NodeType string `json:"node_type"`
	//
	VpcId string `json:"vpc_id"`
	//
	SubnetId string `json:"subnet_id"`

	PublicIp PublicIp `json:"public_ip"`
	//
	PublicEndpoints []PublicEndpoints `json:"public_endpoints"`
	//
	ActionProgress map[string]string `json:"action_progress"`
	//
	SubStatus string `json:"sub_status"`
	//
	TaskStatus string `json:"task_status"`

	ParameterGroup ParameterGroup `json:"parameter_group,omitempty"`
	//
	NodeTypeId string `json:"node_type_id"`
	//
	SecurityGroupId string `json:"security_group_id"`
	//
	PrivateIp []string `json:"private_ip"`

	MaintainWindow MaintainWindow `json:"maintain_window"`

	ResizeInfo ResizeInfo `json:"resize_info,omitempty"`

	FailedReasons FailedReason `json:"failed_reasons,omitempty"`
}

type Endpoints struct {
	//
	ConnectInfo string `json:"connect_info,omitempty"`
	//
	JdbcUrl string `json:"jdbc_url,omitempty"`
}

type Nodes struct {
	//
	Id string `json:"id"`
	//
	Status string `json:"status"`
}

type Tags struct {
	//
	Key string `json:"key"`
	//
	Value string `json:"value"`
}

type PublicEndpoints struct {
	//
	PublicConnectInfo string `json:"public_connect_info,omitempty"`
	//
	JdbcUrl string `json:"jdbc_url,omitempty"`
}

type ParameterGroup struct {
	//
	Id string `json:"id"`
	//
	Name   string `json:"name"`
	Status string `json:"status"`
}

type MaintainWindow struct {
	//
	Day string `json:"day,omitempty"`
	//
	StartTime string `json:"start_time,omitempty"`
	//
	EndTime string `json:"end_time,omitempty"`
}

type ResizeInfo struct {
	//
	TargetNodeNum int32 `json:"target_node_num,omitempty"`
	//
	OriginNodeNum int32 `json:"origin_node_num,omitempty"`
	//
	ResizeStatus string `json:"resize_status,omitempty"`
	//
	StartTime string `json:"start_time,omitempty"`
}

type FailedReason struct {
	//
	ErrorCode string `json:"error_code,omitempty"`
	//
	ErrorMsg string `json:"error_msg,omitempty"`
}
