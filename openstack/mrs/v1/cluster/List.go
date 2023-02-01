package cluster

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ListOpts struct {
	// Maximum number of clusters displayed on a page
	// Value range: 1 to 2147483646
	PageSize string `q:"pageSize,omitempty"`
	// Current page number
	CurrentPage string `q:"currentPage,omitempty"`
	// You can query a cluster list by cluster status.
	// - starting: Query a list of clusters that are being started.
	// - running: Query a list of running clusters.
	// - terminated: Query a list of terminated clusters.
	// - failed: Query a list of failed clusters.
	// - abnormal: Query a list of abnormal clusters.
	// - terminating: Query a list of clusters that are being terminated.
	// - frozen: Query a list of frozen clusters.
	// - scaling-out: Query a list of clusters that are being scaled out.
	// - scaling-in: Query a list of clusters that are being scaled in.
	ClusterState string `q:"clusterState,omitempty"`
	// You can search for a cluster by its tag.
	// If you specify multiple tags, the relationship between them is AND.
	// - The format of the tags parameter is tags=k1*v1,k2*v2,k3*v3.
	// - When the values of some tags are null, the format is tags=k1,k2,k3*v3.
	Tags string `q:"tags,omitempty"`
}

func List(client *golangsdk.ServiceClient, opts ListOpts) (*ListResponse, error) {
	q, err := golangsdk.BuildQueryString(opts)
	if err != nil {
		return nil, err
	}

	// GET /v1.1/{project_id}/cluster_infos
	raw, err := client.Get(client.ServiceURL("cluster_infos")+q.String(), nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res ListResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type ListResponse struct {
	// Total number of clusters in a list
	ClusterTotal *int32 `json:"clusterTotal,omitempty"`
	// Cluster parameters.
	Clusters []Cluster `json:"clusters,omitempty"`
}

type Cluster struct {
	// Cluster ID
	ClusterId string `json:"clusterId"`
	// Cluster name
	ClusterName string `json:"clusterName"`
	// Number of Master nodes deployed in a cluster
	MasterNodeNum string `json:"masterNodeNum"`
	// Number of Core nodes deployed in a cluster
	CoreNodeNum string `json:"coreNodeNum"`
	// Total number of nodes deployed in a cluster
	TotalNodeNum string `json:"totalNodeNum"`
	// Cluster status.
	// Valid values include:
	// - starting: The cluster is being started.
	// - running: The cluster is running.
	// - terminated: The cluster has been terminated.
	// - failed: The cluster fails.
	// - abnormal: The cluster is abnormal.
	// - terminating: The cluster is being terminated.
	// - frozen: The cluster has been frozen.
	// - scaling-out: The cluster is being scaled out.
	// - scaling-in: The cluster is being scaled in.
	ClusterState string `json:"clusterState"`
	// Cluster creation time, which is a 10-bit timestamp
	CreateAt string `json:"createAt"`
	// Cluster update time, which is a 10-bit timestamp
	UpdateAt string `json:"updateAt"`
	// Cluster billing mode
	BillingType string `json:"billingType"`
	// Cluster work region
	DataCenter string `json:"dataCenter"`
	// VPC name
	Vpc string `json:"vpc"`
	// Cluster creation fee, which is automatically calculated
	Fee string `json:"fee"`
	// Hadoop version
	HadoopVersion string `json:"hadoopVersion"`
	// Instance specifications of a Master node
	MasterNodeSize string `json:"masterNodeSize"`
	// Instance specifications of a Core node
	CoreNodeSize string `json:"coreNodeSize"`
	// Component list
	ComponentList []Component `json:"componentList"`
	// External IP address
	ExternalIp string `json:"externalIp"`
	// Backup external IP address
	ExternalAlternateIp string `json:"externalAlternateIp"`
	// Internal IP address
	InternalIp string `json:"internalIp"`
	// Cluster deployment ID
	DeploymentId string `json:"deploymentId"`
	// Cluster remarks
	Remark string `json:"remark"`
	// Cluster creation order ID
	OrderId string `json:"orderId"`
	// AZ ID
	AzId string `json:"azId"`
	// Product ID of a Master node
	MasterNodeProductId string `json:"masterNodeProductId"`
	// Specification ID of a Master node
	MasterNodeSpecId string `json:"masterNodeSpecId"`
	// Product ID of a Core node
	CoreNodeProductId string `json:"coreNodeProductId"`
	// Specification ID of a Core node
	CoreNodeSpecId string `json:"coreNodeSpecId"`
	// AZ name
	AzName string `json:"azName"`
	// Instance ID
	InstanceId string `json:"instanceId"`
	// URI for remotely logging in to an ECS
	Vnc string `json:"vnc"`
	// Project ID
	TenantId string `json:"tenantId"`
	// Disk storage space
	VolumeSize int    `json:"volumeSize"`
	VolumeType string `json:"volumeType"`
	// Subnet ID
	SubnetId string `json:"subnetId"`
	// Cluster type
	ClusterType int `json:"clusterType"`
	// Subnet name
	SubnetName string `json:"subnetName"`
	// Security group ID
	SecurityGroupsId string `json:"securityGroupsId"`
	// Security group ID of a non-Master node.
	// Currently, one MRS cluster uses only one security group.
	// Therefore, this field has been discarded.
	// This field returns the same value as securityGroupsId does for compatibility consideration.
	SlaveSecurityGroupsId string `json:"slaveSecurityGroupsId"`
	// Cluster operation progress description.
	// The cluster installation progress includes:
	// - Verifying cluster parameters: Cluster parameters are being verified.
	// - Applying for cluster resources: Cluster resources are being applied for.
	// - Creating VMs: The VMs are being created.
	// - Initializing VMs: The VMs are being initialized.
	// - Installing MRS Manager: MRS Manager is being installed.
	// - Deploying the cluster: The cluster is being deployed.
	// - Cluster installation failed: Failed to install the cluster.
	// The cluster scale-out progress includes:
	// - Preparing for scale-out: Cluster scale-out is being prepared.
	// - Creating VMs: The VMs are being created.
	// - Initializing VMs: The VMs are being initialized.
	// - Adding nodes to the cluster: The nodes are being added to the  cluster.
	// - Scale-out failed: Failed to scale out the cluster.
	// The cluster scale-in progress includes:
	// - Preparing for scale-in: Cluster scale-in is being prepared.
	// - Decommissioning instance: The instance is being decommissioned.
	// - Deleting VMs: The VMs are being deleted.
	// - Deleting nodes from the cluster: The nodes are being deleted from the cluster.
	// - Scale-in failed: Failed to scale in the cluster.
	// If the cluster installation, scale-out, or scale-in fails, stageDesc will display the failure cause.
	StageDesc string `json:"stageDesc"`
	// Whether MRS Manager installation is finished during cluster creation.
	// - true: MRS Manager installation is finished.
	// - false: MRS Manager installation is not finished.
	MrsManagerFinish bool `json:"mrsManagerFinish"`
	// Running mode of an MRS cluster
	// - 0: Normal cluster
	// - 1: Security cluster
	SafeMode int `json:"safeMode"`
	// Cluster version
	ClusterVersion string `json:"clusterVersion"`
	// Name of the public key file
	NodePublicCertName string `json:"nodePublicCertName"`
	// IP address of a Master node
	MasterNodeIp string `json:"masterNodeIp"`
	// Preferred private IP address
	PrivateIpFirst string `json:"privateIpFirst"`
	// Error message
	ErrorInfo string `json:"errorInfo"`
	// Start time of billing
	ChargingStartTime string `json:"chargingStartTime"`
	// Whether to collect logs when cluster installation fails
	// - 0: Do not collect.
	// - 1: Collect.
	LogCollection int `json:"logCollection"`
	// List of Task nodes.
	TaskNodeGroups []NodeGroupV10 `json:"taskNodeGroups"`
	// List of Master, Core and Task nodes
	NodeGroups []NodeGroupV10 `json:"nodeGroups"`
	// Data disk storage type of the Master node.
	// Currently, SATA, SAS and SSD are supported.
	MasterDataVolumeType string `json:"masterDataVolumeType"`
	// Data disk storage space of the Master node.
	// To increase data storage capacity, you can add disks at the same time when creating a cluster.
	// Value range: 100 GB to 32,000 GB
	MasterDataVolumeSize int `json:"masterDataVolumeSize"`
	// Number of data disks of the Master node.
	// The value can be set to 1 only.
	MasterDataVolumeCount int `json:"masterDataVolumeCount"`
	// Data disk storage type of the Core node.
	// Currently, SATA, SAS and SSD are supported.
	CoreDataVolumeType string `json:"coreDataVolumeType"`
	// Data disk storage space of the Core node.
	// To increase data storage capacity, you can add disks at the same time when creating a cluster.
	// Value range: 100 GB to 32,000 GB
	CoreDataVolumeSize int `json:"coreDataVolumeSize"`
	// Number of data disks of the Core node.
	// Value range: 1 to 10
	CoreDataVolumeCount int `json:"coreDataVolumeCount"`
	// Bootstrap action script information.
	// MRS 1.7.2 or later supports this parameter.
	BootstrapScripts []ScriptResult `json:"bootstrap_scripts"`
	// Node change status.
	// If this parameter is left blank, the cluster nodes are not changed.
	// Possible values are as follows:
	// - scaling-out: The cluster is being scaled out.
	// - scaling-in: The cluster is being scaled in.
	// - scaling-error: The cluster is in the running state
	//   and fails to be scaled in or out or the specifications fail to be scaled up for the last time.
	// - scaling-up: The Master node specifications are being scaled up.
	// - scaling_up_first: The standby Master node specifications are being scaled up.
	// - scaled_up_first: The standby Master node specifications have been scaled up successfully.
	// - scaled-up-success: The Master node specifications have been scaled up successfully.
	Scale string `json:"scale"`
}
