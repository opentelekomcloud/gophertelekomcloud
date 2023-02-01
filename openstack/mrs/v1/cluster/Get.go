package cluster

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack"
)

func Get(client *golangsdk.ServiceClient, id string) (*Cluster, error) {
	// GET /v1.1/{project_id}/cluster_infos/{cluster_id}
	raw, err := client.Get(client.ServiceURL("cluster_infos", id), nil, openstack.StdRequestOpts())
	if err != nil {
		return nil, err
	}

	var res Cluster
	err = extract.IntoStructPtr(raw.Body, &res, "cluster")
	return &res, err
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
	ComponentId string `json:"componentId"`
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
	NodeNum int `json:"nodeNum,omitempty"`
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
	RootVolumeSize int `json:"rootVolumeSize,omitempty"`
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
	DataVolumeCount int `json:"dataVolumeCount,omitempty"`
	// Data disk storage space of a node
	DataVolumeSize int `json:"dataVolumeSize,omitempty"`
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
