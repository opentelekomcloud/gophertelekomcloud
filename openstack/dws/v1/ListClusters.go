package v1

type ListClustersRequest struct {
}

// GET /v1.0/{project_id}/clusters

type ListClustersResponse struct {
	//
	Clusters []ClusterInfo `json:"clusters,omitempty"`
	//
	Count int32 `json:"count,omitempty"`
}

type ClusterInfo struct {
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
	//
	PublicIp PublicIp `json:"public_ip"`
	//
	PublicEndpoints []PublicEndpoints `json:"public_endpoints"`
	//
	ActionProgress map[string]string `json:"action_progress"`
	//
	SubStatus string `json:"sub_status"`
	//
	TaskStatus string `json:"task_status"`
	//
	SecurityGroupId string `json:"security_group_id"`

	FailedReasons FailedReason `json:"failed_reasons,omitempty"`
}
