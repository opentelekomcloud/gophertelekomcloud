package v1

type RestoreClusterRequest struct {
	//
	SnapshotId string `json:"snapshot_id"`

	Body RestoreClusterRequestBody `json:"body,omitempty"`
}

type RestoreClusterRequestBody struct {
	Restore Restore `json:"restore"`
}

type Restore struct {
	//
	Name string `json:"name"`
	//
	SubnetId string `json:"subnet_id,omitempty"`
	//
	SecurityGroupId string `json:"security_group_id,omitempty"`
	//
	VpcId string `json:"vpc_id,omitempty"`
	//
	AvailabilityZone string `json:"availability_zone,omitempty"`
	//
	Port int32 `json:"port,omitempty"`
	//
	PublicIp PublicIp `json:"public_ip,omitempty"`
	//
	EnterpriseProjectId string `json:"enterprise_project_id,omitempty"`
}

// POST /v1.0/{project_id}/snapshots/{snapshot_id}/actions

type RestoreClusterResponse struct {
	Cluster Cluster `json:"cluster,omitempty"`
}
