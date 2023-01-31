package cluster

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
)

func Get(client *golangsdk.ServiceClient, id string) (*GetResponse, error) {
	// GET /v1.1/{project_id}/cluster_infos/{cluster_id}
	raw, err := client.Get(client.ServiceURL("cluster_infos", id), nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res GetResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type GetResponse struct {
	// Cluster ID
	ClusterID string `json:"clusterId"`
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
	DeploymentID string `json:"deploymentId"`
	// Cluster remarks
	Remark string `json:"remark"`
	// Cluster creation order ID
	OrderID string `json:"orderId"`
	// AZ ID
	AzID string `json:"azId"`
	// Product ID of a Master node
	MasterNodeProductID string `json:"masterNodeProductId"`
	// Specification ID of a Master node
	MasterNodeSpecID string `json:"masterNodeSpecId"`
	// Product ID of a Core node
	CoreNodeProductID string `json:"coreNodeProductId"`
	// Specification ID of a Core node
	CoreNodeSpecID string `json:"coreNodeSpecId"`
	// AZ name
	AzName string `json:"azName"`
	// Instance ID
	InstanceID string `json:"instanceId"`
	// URI for remotely logging in to an ECS
	Vnc string `json:"vnc"`
	// Project ID
	TenantID string `json:"tenantId"`
	// Disk storage space
	VolumeSize int `json:"volumeSize"`
	// Subnet ID
	SubnetID string `json:"subnetId"`
	// Subnet name
	SubnetName string `json:"subnetName"`
	// Security group ID
	SecurityGroupsID string `json:"securityGroupsId"`
	// Security group ID of a non-Master node.
	// Currently, one MRS cluster uses only one security group.
	// Therefore, this field has been discarded.
	// This field returns the same value as securityGroupsId does for compatibility consideration.
	SlaveSecurityGroupsID string `json:"slaveSecurityGroupsId"`
	// Bootstrap action script information.
	// MRS 1.7.2 or later supports this parameter.
	BootstrapScripts []ScriptResult `json:"bootstrap_scripts"`
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
	// Tag information
	Tags []tags.ResourceTag `json:"tags"`
	// Start time of billing
	ChargingStartTime string `json:"chargingStartTime"`
	// Cluster type
	ClusterType int `json:"clusterType"`
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

type Component struct {
	// Component ID
	// - Component IDs of MRS 3.1.2-LTS.3 are as follows:
	//	 – MRS 3.1.2-LTS.3_001: Hadoop
	//	 – MRS 3.1.2-LTS.3_002: Spark2x
	//	 – MRS 3.1.2-LTS.3_003: HBase
	//	 – MRS 3.1.2-LTS.3_004: Hive
	//	 – MRS 3.1.2-LTS.3_005: Hue
	//	 – MRS 3.1.2-LTS.3_006: Loader
	//	 – MRS 3.1.2-LTS.3_007: Kafka
	//	 – MRS 3.1.2-LTS.3_008: Flume
	//	 – MRS 3.1.2-LTS.3_009: FTP-Server
	//	 – MRS 3.1.2-LTS.3_010: Solr
	//	 – MRS 3.1.2-LTS.3_010: Redis
	//	 – MRS 3.1.2-LTS.3_011: Elasticsearch
	//	 – MRS 3.1.2-LTS.3_012: Flink
	//	 – MRS 3.1.2-LTS.3_013: Oozie
	//	 – MRS 3.1.2-LTS.3_014: GraphBase
	//	 – MRS 3.1.2-LTS.3_015: ZooKeeper
	//	 – MRS 3.1.2-LTS.3_016: HetuEngine
	//	 – MRS 3.1.2-LTS.3_017: Ranger
	//	 – MRS 3.1.2-LTS.3_018: Tez
	//	 – MRS 3.1.2-LTS.3_019: ClickHouse
	//	 – MRS 3.1.2-LTS.3_020: Metadata
	//	 – MRS 3.1.2-LTS.3_021: KMS
	// - Component IDs of MRS 3.1.0-LTS.1 are as follows:
	//	 – MRS 3.1.0-LTS.1_001: Hadoop
	//	 – MRS 3.1.0-LTS.1_002: Spark2x
	//	 – MRS 3.1.0-LTS.1_003: HBase
	//	 – MRS 3.1.0-LTS.1_004: Hive
	//	 – MRS 3.1.0-LTS.1_005: Hue
	//	 – MRS 3.1.0-LTS.1_006: Loader
	//	 – MRS 3.1.0-LTS.1_007: Kafka
	//	 – MRS 3.1.0-LTS.1_008: Flume
	//	 – MRS 3.1.0-LTS.1_009: Flink
	//	 – MRS 3.1.0-LTS.1_010: Oozie
	//	 – MRS 3.1.0-LTS.1_011: ZooKeeper
	//	 – MRS 3.1.0-LTS.1_012: HetuEngine
	//	 – MRS 3.1.0-LTS.1_013: Ranger
	//	 – MRS 3.1.0-LTS.1_014: Tez
	//	 – MRS 3.1.0-LTS.1_015: ClickHouse
	// - Component IDs of MRS 2.1.0 are as follows:
	//	 – MRS 2.1.0_001: Hadoop
	//	 – MRS 2.1.0_002: Spark
	//	 – MRS 2.1.0_003: HBase
	//	 – MRS 2.1.0_004: Hive
	//	 – MRS 2.1.0_005: Hue
	//	 – MRS 2.1.0_006: Kafka
	//	 – MRS 2.1.0_007: Storm
	//	 – MRS 2.1.0_008: Loader
	//	 – MRS 2.1.0_009: Flume
	//	 – MRS 2.1.0_010: Tez
	//	 – MRS 2.1.0_011: Presto
	//	 – MRS 2.1.0_014: Flink
	// - Component IDs of MRS 1.9.2 are as follows:
	//	 – MRS 1.9.2_001: Hadoop
	//	 – MRS 1.9.2_002: Spark
	//	 – MRS 1.9.2_003: HBase
	//	 – MRS 1.9.2_004: Hive
	//	 – MRS 1.9.2_005: Hue
	//	 – MRS 1.9.2_006: Kafka
	//	 – MRS 1.9.2_007: Storm
	//	 – MRS 1.9.2_008: Loader
	//	 – MRS 1.9.2_009: Flume
	//	 – MRS 1.9.2_010: Presto
	//	 – MRS 1.9.2_011: KafkaManager
	//	 – MRS 1.9.2_012: Flink
	//	 – MRS 1.9.2_013: OpenTSDB
	//	 – MRS 1.9.2_015: Alluxio
	//	 – MRS 1.9.2_16: Ranger
	//	 – MRS 1.9.2_17: Tez
	// - Component IDs of MRS 1.7.2 and MRS 1.6.3 are as follows:
	//	 – MRS 1.7.2_001: Hadoop
	//	 – MRS 1.7.2_002: Spark
	//	 – MRS 1.7.2_003: HBase
	//	 – MRS 1.7.2_004: Hive
	//	 – MRS 1.7.2_005: Hue
	//	 – MRS 1.7.2_006: Kafka
	//	 – MRS 1.7.2_007: Storm
	//	 – MRS 1.7.2_008: Loader
	//	 – MRS 1.7.2_009: Flume
	// For example, the component_id of Hadoop is MRS 2.1.0_001, MRS 1.9.2_001, MRS 1.7.2_001, MRS 1.6.3_001.
	ComponentID string `json:"componentId"`
	// Component name
	ComponentName string `json:"componentName"`
	// Component version
	ComponentVersion string `json:"componentVersion"`
	// Component description
	ComponentDesc string `json:"componentDesc"`
}

type NodeGroupV10 struct {
	// Node group name
	GroupName string `json:"groupName,omitempty"`
	// Number of nodes.
	// The value ranges from 0 to 500.
	// The minimum number of Master and Core nodes is 1 and the total number of Core and Task nodes cannot exceed 500.
	NodeNum *int32 `json:"nodeNum,omitempty"`
	// Instance specifications of a node
	NodeSize string `json:"nodeSize,omitempty"`
	// Instance specification ID of a node
	NodeSpecId string `json:"nodeSpecId,omitempty"`
	// Instance product ID of a node
	NodeProductId string `json:"nodeProductId,omitempty"`
	// VM product ID of a node
	VmProductId string `json:"vmProductId,omitempty"`
	// VM specifications of a node
	VmSpecCode string `json:"vmSpecCode,omitempty"`
	// System disk size of a node.
	// This parameter is not configurable and its default value is 40 GB.
	RootVolumeSize *int32 `json:"rootVolumeSize,omitempty"`
	// System disk product ID of a node
	RootVolumeProductId string `json:"rootVolumeProductId,omitempty"`
	// System disk type of a node
	RootVolumeType string `json:"rootVolumeType,omitempty"`
	// System disk product specifications of a node
	RootVolumeResourceSpecCode string `json:"rootVolumeResourceSpecCode,omitempty"`
	// System disk product type of a node
	RootVolumeResourceType string `json:"rootVolumeResourceType,omitempty"`
	// Data disk storage type of a node.
	// Currently, SATA, SAS and SSD are supported.
	// - SATA: Common I/O
	// - SAS: High I/O
	// - SSD: Ultra-high I/O
	DataVolumeType string `json:"dataVolumeType,omitempty"`
	// Number of data disks of a node
	DataVolumeCount *int32 `json:"dataVolumeCount,omitempty"`
	// Data disk storage space of a node
	DataVolumeSize *int32 `json:"dataVolumeSize,omitempty"`
	// Data disk product ID of a node
	DataVolumeProductId string `json:"dataVolumeProductId,omitempty"`
	// Data disk product specifications of a node
	DataVolumeResourceSpecCode string `json:"dataVolumeResourceSpecCode,omitempty"`
	// Data disk product type of a node
	DataVolumeResourceType string `json:"dataVolumeResourceType,omitempty"`
}

type ScriptResult struct {
	// Name of a bootstrap action script.
	// It must be unique in a cluster.
	// The value can contain only digits, letters, spaces, hyphens (-), and underscores (_) and cannot start with a space.
	// The value can contain 1 to 64 characters.
	Name string `json:"name"`
	// Path of the shell script.
	// Set this parameter to an OBS bucket path or a local VM path.
	// - OBS bucket path: Enter a script path manually.
	//   For example, enter the path of the public sample script provided by MRS.
	//   Example: s3a://bootstrap/presto/presto-install.sh.
	//   If dualroles is installed, the parameter of the presto-install.sh script is dualroles.
	//   If worker is installed, the parameter of the presto-install.sh script is worker.
	//   Based on the Presto usage habit,
	//   you are advised to install dualroles  on the active Master nodes and worker on the Core nodes.
	// - Local VM path: Enter a script path.
	//   The script path must start with a slash (/) and end with .sh.
	Uri string `json:"uri"`
	// Bootstrap action script parameters
	Parameters string `json:"parameters"`
	// Type of a node where the bootstrap action script is executed.
	// The value can be Master, Core, or Task.
	Nodes []string `json:"nodes"`
	// Whether the bootstrap action script runs only on active Master nodes.
	// The default value is false, indicating that the bootstrap action script can run on all Master nodes.
	ActiveMaster bool `json:"active_master"`
	// Time when the bootstrap action script is executed.
	// Currently, the following two options are available:
	// Before component start and After component start
	// The default value is false,
	// indicating that the bootstrap action script is executed after the component is started.
	BeforeComponentStart bool `json:"before_component_start"`
	// Whether to continue executing subsequent scripts
	// and creating a cluster after the bootstrap action script fails to be executed.
	// - continue: Continue to execute subsequent scripts.
	// - errorout: Stop the action.
	//   The default value is errorout, indicating that the action is stopped.
	// NOTE
	// You are advised to set this parameter to continue in the commissioning phase
	// so that the cluster can continue to be installed
	// and started no matter whether the bootstrap action is successful.
	FailAction string `json:"fail_action"`
	// Execution time of one boot operation script.
	StartTime int `json:"start_time"`
	// Running state of one bootstrap action script
	// - PENDING
	// - IN_PROGRESS
	// - SUCCESS
	// - FAILURE
	State string `json:"state"`
}
