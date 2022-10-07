package v1

type CreateClusterRequestBody struct {
	Cluster CreateClusterInfo `json:"cluster"`
}

type CreateClusterInfo struct {
	//
	NodeType string `json:"node_type"`
	//
	NumberOfNode int32 `json:"number_of_node"`
	//
	SubnetId string `json:"subnet_id"`
	//
	SecurityGroupId string `json:"security_group_id"`
	//
	VpcId string `json:"vpc_id"`
	//
	AvailabilityZone string `json:"availability_zone,omitempty"`
	//
	Port int32 `json:"port,omitempty"`
	//
	Name string `json:"name"`
	//
	UserName string `json:"user_name"`
	//
	UserPwd string `json:"user_pwd"`
	//
	PublicIp PublicIp `json:"public_ip,omitempty"`
	//
	NumberOfCn int32 `json:"number_of_cn,omitempty"`
	//
	EnterpriseProjectId string `json:"enterprise_project_id,omitempty"`
}

type PublicIp struct {
	//
	PublicBindType string `json:"public_bind_type"`
	//
	EipId string `json:"eip_id,omitempty"`
}

// POST /v1.0/{project_id}/clusters

type CreateClusterResponse struct {
	Cluster Cluster `json:"cluster,omitempty"`
}

type Cluster struct {
	//
	Id string `json:"id"`
}
