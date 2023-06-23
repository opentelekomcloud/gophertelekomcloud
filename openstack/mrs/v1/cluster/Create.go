package cluster

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
)

type CreateOpts struct {
	// Cluster billing mode.
	// Set this parameter to 12.
	BillingType int `json:"billing_type"`
	// Region of the cluster.
	DataCenter string `json:"data_center"`
	// AZ ID.
	// - AZ1(eude-01):bf84aba586ce4e948da0b97d9a7d62fb
	// - AZ2(eude-02):bf84aba586ce4e948da0b97d9a7d62fc
	AvailableZoneId string `json:"available_zone_id"`
	// Cluster name.
	// It must be unique.
	// It contains only 1 to 64 characters.
	// Only letters, digits, hyphens (-), and underscores (_) are allowed.
	ClusterName string `json:"cluster_name"`
	// Name of the VPC where the subnet locates.
	// Perform the following operations to obtain the VPC name from the VPC management console:
	// 1. Log in to the management console.
	// 2. Click Virtual Private Cloud and select Virtual Private Cloud from the left list.
	// On the Virtual Private Cloud page, obtain the VPC name from the list.
	Vpc string `json:"vpc"`
	// ID of the VPC where the subnet locates
	// Perform the following operations to obtain the VPC ID from the VPC management console:
	// 1. Log in to the management console.
	// 2. Click Virtual Private Cloud and select Virtual Private Cloud from the left list.
	// On the Virtual Private Cloud page, obtain the VPC ID from the list.
	VpcId string `json:"vpc_id"`
	// Network ID
	// Perform the following operations to obtain the network ID of the  VPC from the VPC management console:
	// 1. Log in to the management console.
	// 2. Click Virtual Private Cloud and select Virtual Private Cloud from the left list.
	// On the Virtual Private Cloud page, obtain the network ID of the VPC from the list.
	SubnetId string `json:"subnet_id"`
	// Subnet name
	// Perform the following operations to obtain the subnet name from the VPC management console:
	// 1. Log in to the management console.
	// 2. Click Virtual Private Cloud and select Virtual Private Cloud from the left list.
	// On the Virtual Private Cloud page, obtain the subnet name of the VPC from the list.
	SubnetName string `json:"subnet_name"`
	// Security group ID of the cluster
	// - If this parameter is left blank,
	// MRS automatically creates a security group, whose name starts with mrs_{cluster_name}.
	// - If this parameter is not left blank, a fixed security group is used to create a cluster.
	// The transferred ID must be the security group ID owned by the current tenant.
	// The security group must include an inbound rule in
	// which all protocols and all ports are allowed and
	// the source is the IP address of the specified node on the management plane.
	SecurityGroupsId string `json:"security_groups_id,omitempty"`
	// Cluster tag
	// - A cluster allows a maximum of 10 tags. A tag name (key) must be unique in a cluster.
	// - A tag key or value cannot contain the following special characters: =*<>\,|/
	Tags []tags.ResourceTag `json:"tags,omitempty"`
	// Cluster version
	// Possible values are as follows:
	// - MRS 1.6.3
	// - MRS 1.7.2
	// - MRS 1.9.2
	// - MRS 2.1.0
	// - MRS 3.1.0-LTS.1
	// - MRS 3.1.2-LTS.3
	ClusterVersion string `json:"cluster_version"`
	// Cluster type
	// - 0: analysis cluster
	// - 1: streaming cluster
	// The default value is 0.
	// Note: Currently, hybrid clusters cannot be created using APIs.
	ClusterType *int `json:"cluster_type,omitempty"`
	// Running mode of an MRS cluster
	// - 0: normal cluster.
	// In a normal cluster, Kerberos authentication is disabled,
	// and users can use all functions provided by the cluster.
	// - 1: security cluster.
	// In a security cluster, Kerberos authentication is enabled,
	// and common users cannot use the file management
	// and job management functions of an MRS cluster or view cluster resource usage
	// and the job records of Hadoop and Spark.
	// To use these functions, the users must obtain the relevant permissions from the MRS Manager administrator.
	// NOTE
	// For MRS 1.7.2 or earlier,
	// the request body contains the cluster_admin_secret field only when safe_mode is set to 1.
	SafeMode int `json:"safe_mode"`
	// Password of the MRS Manager administrator
	// - Must contain 8 to 32 characters.
	// - Must contain at least three of
	// the following:
	// 	– Lowercase letters
	// 	– Uppercase letters
	// 	– Digits
	// 	– Special characters: `~!@#$ %^&*()-_=+\|[{}];:'",<.>/? and space
	// - Cannot be the username or the username spelled backwards.
	// NOTE
	// For MRS 1.7.2 or earlier, this parameter is mandatory only when safe_mode is set to 1.
	ClusterAdminSecret string `json:"cluster_admin_secret,omitempty"`
	// Cluster login mode
	// - 0: password
	// - 1: key pair
	// The default value is 1.
	// - If login_mode is set to 0, the request body contains the cluster_master_secret field.
	// - If login_mode is set to 1, the request body contains the node_public_cert_name field.
	// NOTE
	// This parameter is valid only for clusters of MRS 1.6.3
	// or later instead of clusters of versions earlier than MRS 1.6.3.
	LoginMode *int `json:"login_mode,omitempty"`
	// Password of user root for logging in to a cluster node
	// If login_mode is set to 0, the request body contains the cluster_master_secret field.
	// A password must meet the following requirements:
	// - Must be 8 to 26 characters long.
	// - Must contain at least three of the following:
	// uppercase letters, lowercase letters, digits,
	// and special characters (!@$%^-_=+[{}]:,./?), but must not contain spaces.
	// - Cannot be the username or the username spelled backwards
	ClusterMasterSecret string `json:"cluster_master_secret"`
	// Name of a key pair You can use a key pair to log in to the Master node in the cluster.
	// If login_mode is set to 1, the request body contains the node_public_cert_name field.
	NodePublicCertName string `json:"node_public_cert_name,omitempty"`
	// Whether to collect logs when cluster creation fails
	// - 0: Do not collect.
	// - 1: Collect.
	// The default value is 1, indicating that OBS buckets will be created
	// and only used to collect logs that record MRS cluster creation failures.
	LogCollection *int `json:"log_collection,omitempty"`
	// List of nodes.
	NodeGroups []NodeGroup `json:"node_groups,omitempty"`
	// List of service components to be installed.
	ComponentList []ComponentList `json:"component_list"`
	// Jobs can be submitted when a cluster is created. Currently, only one job can be created.
	AddJobs []AddJobs `json:"add_jobs,omitempty"`
	// Bootstrap action script information.
	BootstrapScripts []BootstrapScript `json:"bootstrap_scripts,omitempty"`
	// Number of Master nodes. If cluster HA is enabled, set this parameter to 2. If cluster HA is disabled, set this parameter to 1.
	MasterNodeNum int `json:"master_node_num" required:"true"`
	// Instance specifications of the Master node, for example, c6.4xlarge.4linux.mrs.
	// MRS supports host specifications determined by CPU, memory, and disk space.
	MasterNodeSize string `json:"master_node_size" required:"true"`
	// Number of Core nodes
	// Value range: 1 to 500
	// A maximum of 500 Core nodes are supported by default. If more than 500 Core nodes are required, contact technical support.
	CoreNodeNum int `json:"core_node_num" required:"true"`
	// Instance specifications of the Core node, for example, c6.4xlarge.4linux.mrs.
	CoreNodeSize string `json:"core_node_size" required:"true"`
	// This parameter is a multi-disk parameter, indicating the data disk storage type of the Master node. Currently, SATA, SAS and SSD are supported.
	MasterDataVolumeType string `json:"master_data_volume_type,omitempty"`
	// This parameter is a multi-disk parameter, indicating the data disk storage space of the Master node.
	// To increase data storage capacity, you can add disks at the same time when creating a cluster.
	// Value range: 100 GB to 32,000 GB
	MasterDataVolumeSize int `json:"master_data_volume_size,omitempty"`
	// This parameter is a multi-disk parameter, indicating the number of data disks of the Master node.
	// The value can be set to 1 only.
	MasterDataVolumeCount int `json:"master_data_volume_count,omitempty"`
	// This parameter is a multi-disk parameter, indicating the data disk storage type of the Core node. Currently, SATA, SAS and SSD are supported.
	CoreDataVolumeType string `json:"core_data_volume_type,omitempty"`
	// This parameter is a multi-disk parameter, indicating the data disk storage space of the Core node.
	// To increase data storage capacity, you can add disks at the same time when creating a cluster.
	// Value range: 100 GB to 32,000 GB
	CoreDataVolumeSize int `json:"core_data_volume_size,omitempty"`
	// This parameter is a multi-disk parameter, indicating the number of data disks of the Core node.
	// Value range: 1 to 10
	CoreDataVolumeCount int `json:"core_data_volume_count,omitempty"`
	// Data disk storage type of the Master and Core nodes. Currently, SATA, SAS and SSD are supported.
	// Disk parameters can be represented by volume_type and volume_size, or multi-disk parameters.
	// If the volume_type and volume_size parameters coexist with the multi-disk parameters,
	// the system reads the volume_type and volume_size parameters first. You are advised to use the multi-disk parameters.
	// SATA: Common I/O
	// SAS: High I/O
	// SSD: Ultra-high I/O
	VolumeType string `json:"volume_type,omitempty"`
	// Data disk storage space of the Master and Core nodes. To increase data storage capacity,
	// you can add disks at the same time when creating a cluster. Select a proper disk storage space based on the following application scenarios:
	// Separation of data storage and computing: Data is stored in the OBS system.
	// Costs of clusters are relatively low but computing performance is poor. The clusters can be deleted at any time.
	// It is recommended when data computing is infrequently performed.
	// Integration of data storage and computing: Data is stored in the HDFS system.
	// Costs of clusters are relatively high but computing performance is good. The clusters cannot be deleted in a short term.
	// It is recommended when data computing is frequently performed.
	// Value range: 100 GB to 32,000 GB
	// This parameter is not recommended. For details, see the description of the volume_type parameter.
	VolumeSize int `json:"volume_size,omitempty"`
}

func Create(client *golangsdk.ServiceClient, opts CreateOpts) (*CreateResponse, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// POST /v1.1/{project_id}/run-job-flow
	raw, err := client.Post(client.ServiceURL("run-job-flow"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res CreateResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type NodeGroup struct {
	// Node group name.
	// - master_node_default_group
	// - core_node_analysis_group
	// - core_node_streaming_group
	// - task_node_analysis_group
	// - task_node_streaming_group
	GroupName string `json:"group_name"`
	// Number of nodes.
	// The value ranges from 0 to 500 and the default value is 0.
	// The total number of Core and Task nodes cannot exceed 500.
	NodeNum int `json:"node_num"`
	// Instance specifications of a node.
	// For details about the configuration method, see the remarks of master_node_size.
	NodeSize string `json:"node_size"`
	// Data disk storage space of a node.
	RootVolumeSize string `json:"root_volume_size,omitempty"`
	// Data disk storage type of a node.
	// Currently, SATA, SAS and SSD are supported.
	// - SATA: Common I/O
	// - SAS: High I/O
	// - SSD: Ultra-high I/O
	RootVolumeType string `json:"root_volume_type,omitempty"`
	// Data disk storage type of a node.
	// Currently, SATA, SAS and SSD are supported.
	// - SATA: Common I/O
	// - SAS: High I/O
	// - SSD: Ultra-high I/O
	DataVolumeType string `json:"data_volume_type,omitempty"`
	// Number of data disks of a node.
	// Value range: 0 to 10
	DataVolumeCount *int `json:"data_volume_count,omitempty"`
	// Data disk storage space of a node.
	// Value range: 100 GB to 32,000 GB
	DataVolumeSize *int `json:"data_volume_size,omitempty"`
	// Auto scaling rule information.
	// This parameter is valid only when group_name is set to task_node_analysis_group or task_node_streaming_group.
	AutoScalingPolicy *AutoScalingPolicy `json:"auto_scaling_policy,omitempty"`
}

type AutoScalingPolicy struct {
	// Whether to enable the auto scaling rule.
	AutoScalingEnable bool `json:"auto_scaling_enable"`
	// Minimum number of nodes left in the node group.
	// Value range: 0 to 500
	MinCapacity int `json:"min_capacity"`
	// Maximum number of nodes in the node group.
	// Value range: 0 to 500
	MaxCapacity int `json:"max_capacity"`
	// Resource plan list.
	// If this parameter is left blank, the resource plan is disabled.
	// When auto scaling is enabled, either a resource plan or an auto scaling rule must be configured.
	// MRS 1.6.3 or later supports this parameter.
	ResourcesPlans []ResourcesPlan `json:"resources_plans,omitempty"`
	// List of custom scaling automation scripts.
	// If this parameter is left blank, a hook script is disabled.
	// MRS 1.7.2 or later supports this parameter.
	ExecScripts []ExecScript `json:"exec_scripts,omitempty"`
	// List of auto scaling rules.
	// When auto scaling is enabled, either a resource plan or an auto scaling rule must be configured.
	Rules []Rules `json:"rules,omitempty"`
}

type ResourcesPlan struct {
	// Cycle type of a resource plan.
	// Currently, only the following cycle type is supported:
	// - daily
	PeriodType string `json:"period_type"`
	// Start time of a resource plan.
	// The value is in the format of hour:minute, indicating that the time ranges from 0:00 to 23:59.
	StartTime string `json:"start_time"`
	// End time of a resource plan.
	// The value is in the same format as that  of start_time.
	// The interval between end_time and start_time must be greater than or equal to 30 minutes.
	EndTime string `json:"end_time"`
	// Minimum number of the preserved nodes in a node group in a resource plan.
	// Value range: 0 to 500
	MinCapacity int `json:"min_capacity"`
	// Maximum number of the preserved nodes in a node group in a resource plan.
	// Value range: 0 to 500
	MaxCapacity int `json:"max_capacity"`
}

type Rules struct {
	// Name of an auto scaling rule.
	// It contains only 1 to 64 characters.
	// Only letters, digits, hyphens (-), and underscores (_) are allowed.
	// Rule names must be unique in a node group.
	Name string `json:"name"`
	// Description about an auto scaling rule.
	// It contains a maximum of 1,024 characters.
	Description string `json:"description,omitempty"`
	// Auto scaling rule adjustment type.
	// The options are as follows:
	// - scale_out: cluster scale-out
	// - scale_in: cluster scale-in
	AdjustmentType string `json:"adjustment_type"`
	// Cluster cooling time after an auto scaling rule is triggered, when no auto scaling operation is performed.
	// The unit is minute.
	// Value range: 0 to 10,080.
	// One week is equal to 10,080 minutes.
	CoolDownMinutes int `json:"cool_down_minutes"`
	// Number of nodes that can be adjusted once.
	// Value range: 1 to 100
	ScalingAdjustment int `json:"scaling_adjustment"`
	// Condition for triggering a rule.
	Trigger *Trigger `json:"trigger"`
}

type Trigger struct {
	// Metric name.
	// This triggering condition makes a judgment according to the value of the metric.
	// A metric name contains a maximum of 64 characters.
	MetricName string `json:"metric_name"`
	// Metric threshold to trigger a rule
	// The parameter value must be an integer or number with two decimal places only.
	MetricValue string `json:"metric_value"`
	// Metric judgment logic operator.
	// The options are as follows:
	// - LT: less than
	// - GT: greater than
	// - LTOE: less than or equal to
	// - GTOE: greater than or equal to
	ComparisonOperator string `json:"comparison_operator,omitempty"`
	// Number of consecutive five-minute periods, during which a metric threshold is reached
	// Value range: 1 to 288
	EvaluationPeriods int `json:"evaluation_periods"`
}

type ExecScript struct {
	// Name of a custom automation script.
	// It must be unique in a same cluster.
	// The value can contain only digits, letters, spaces, hyphens (-),
	// and underscores (_) and cannot start with a space.
	// The value can contain 1 to 64 characters.
	Name string `json:"name"`
	// Path of a custom automation script.
	// Set this parameter to an OBS bucket path or a local VM path.
	// - OBS bucket path: Enter a script path manually, for example, s3a://XXX/scale.sh.
	// - Local VM path: Enter a script path.
	// 	 The script path must start with a slash (/) and end with .sh.
	Uri string `json:"uri"`
	// Parameters of a custom automation script.
	// - Multiple parameters are separated by space.
	// - The following predefined system parameters can be transferred:
	//	 – ${mrs_scale_node_num}: Number of the nodes to be added or removed
	//	 – ${mrs_scale_type}: Scaling type.
	//     The value can be scale_out or scale_in.
	//	 – ${mrs_scale_node_hostnames}: Host names of the nodes to be added or removed
	//	 – ${mrs_scale_node_ips}: IP addresses of the nodes to be added or removed
	//	 – ${mrs_scale_rule_name}: Name of the rule that triggers auto scaling
	// - Other user-defined parameters are used in the same way as those of common shell scripts.
	//   Parameters are separated by space.
	Parameters string `json:"parameters,omitempty"`
	// Type of a node where the custom automation script is executed.
	// The node type can be Master, Core, or Task.
	Nodes []string `json:"nodes"`
	// Whether the custom automation script runs only on the active Master node.
	// The default value is false, indicating that the custom automation script can run on all Master nodes.
	ActiveMaster *bool `json:"active_master,omitempty"`
	// Time when a script is executed.
	// The following four options are supported:
	// - before_scale_out: before scale-out
	// - before_scale_in: before scale-in
	// - after_scale_out: after scale-out
	// - after_scale_in: after scale-in
	ActionStage string `json:"action_stage"`
	// Whether to continue to execute subsequent scripts
	// and create a cluster after the custom automation script fails to be executed.
	// - continue: Continue to execute subsequent scripts.
	// - errorout: Stop the action.
	// NOTE
	// - You are advised to set this parameter to continue in the commissioning phase
	//   so that the cluster can continue to be installed
	//   and started no matter whether the custom automation script is executed successfully.
	// - The scale-in operation cannot be undone.
	// Therefore, fail_action must be set to continue for the scripts that are executed after scale-in.
	FailAction string `json:"fail_action"`
}

type ComponentList struct {
	// Component name
	ComponentName string `json:"component_name"`
}

type AddJobs struct {
	// Job type code
	// - 1: MapReduce
	// - 2: Spark
	// - 3: Hive Script
	// - 4: HiveQL (not supported currently)
	// - 5: DistCp, importing and exporting data (not supported currently)
	// - 6: Spark Script
	// - 7: Spark SQL, submitting Spark SQL statements (not supported currently).
	// NOTE
	// Spark and Hive jobs can be added to only clusters that include Spark and Hive components.
	JobType int `json:"job_type" required:"true"`
	// Job name.
	// It contains 1 to 64 characters.
	// Only letters, digits, hyphens (-), and underscores (_) are allowed.
	// NOTE
	// Identical job names are allowed but not recommended.
	JobName string `json:"job_name" required:"true"`
	// Path of the JAR or SQL file for program execution.
	// The parameter must meet the following requirements:
	// - Contains a maximum of 1,023 characters,
	// excluding special characters such as ;|&><'$. The parameter value cannot be empty or full of spaces.
	// - Files can be stored in HDFS or OBS.
	// The path varies depending on the file system.
	// 	– OBS: The path must start with s3a://. Files or programs encrypted by KMS are not supported.
	// 	– HDFS: The path starts with a slash (/).
	// - Spark Script must end with .sql while MapReduce and Spark Jar must end with .jar. sql and jar are case-insensitive.
	JarPath string `json:"jar_path,omitempty"`
	// Key parameter for program execution.
	// The parameter is specified by the function of the user's program.
	// MRS is only responsible for loading the parameter.
	// The parameter contains a maximum of 2,047 characters,
	// excluding special characters such as ;|&>'<$, and can be left blank.
	Arguments string `json:"arguments,omitempty"`
	// Address for inputting data.
	// Files can be stored in HDFS or OBS.
	// The path varies depending on the file system.
	// - OBS: The path must start with s3a://.
	// Files or programs encrypted by KMS are not supported.
	// - HDFS: The path starts with a slash (/).
	// The parameter contains a maximum of 1,023 characters,
	// excluding special characters such as ;|&>'<$, and can be left blank.
	Input string `json:"input,omitempty"`
	// Address for outputting data.
	// Files can be stored in HDFS or OBS.
	// The path varies depending on the file system.
	// - OBS: The path must start with s3a://.
	// - HDFS: The path starts with a slash (/).
	// If the specified path does not exist,
	// the system will automatically create it.
	// The parameter contains a maximum of 1,023 characters,
	// excluding special characters such as ;|&>'<$, and can be left blank.
	Output string `json:"output,omitempty"`
	// Path for storing job logs that record job running status.
	// Files can be stored in HDFS or OBS.
	// The path varies depending on the file system.
	// - OBS: The path must start with s3a://.
	// - HDFS: The path starts with a slash (/).
	// The parameter contains a maximum of 1,023 characters,
	// excluding special characters such as ;|&>'<$, and can be left blank.
	JobLog string `json:"job_log,omitempty"`
	// Whether to delete the cluster after the job execution is complete
	// - true: Yes
	// - false: No
	ShutdownCluster *bool `json:"shutdown_cluster,omitempty"`
	// Data import and export
	// - import
	// - export
	FileAction string `json:"file_action,omitempty"`
	// - true: Submit a job during cluster creation.
	// - false: Submit a job after the cluster is created.
	// Set this parameter to true in this example.
	SubmitJobOnceClusterRun *bool `json:"submit_job_once_cluster_run" required:"true"`
	// HiveQL statement
	Hql string `json:"hql,omitempty"`
	// SQL program path.
	// This parameter is needed by Spark Script and Hive Script jobs only, and must meet the following requirements:
	// - Contains a maximum of 1,023 characters, excluding special characters such as ;|&><'$.
	//   The parameter value cannot be empty or full of spaces.
	// - Files can be stored in HDFS or OBS. The path varies depending on the file system.
	// 	– OBS: The path must start with s3a://.
	// 	 	   Files or programs encrypted by KMS are not supported.
	// 	– HDFS: The path starts with a slash (/).
	// - Ends with .sql. sql is case-insensitive.
	HiveScriptPath string `json:"hive_script_path" required:"true"`
}

type BootstrapScript struct {
	// Name of a bootstrap action script.
	// It must be unique in a cluster.
	// The value can contain only digits, letters, spaces, hyphens (-),
	// and underscores (_) and cannot start with a space.
	// The value can contain 1 to 64 characters.
	Name string `json:"name" required:"true"`
	// Path of a bootstrap action script.
	// Set this parameter to an OBS bucket path or a local VM path.
	// - OBS bucket path: Enter a script path manually.
	//   For example, enter the path of the public sample script provided by MRS.
	//   Example: s3a://bootstrap/presto/presto-install.sh.
	//   If dualroles is installed, the parameter of the presto-install. sh script is dualroles.
	//   If worker is installed, the parameter of the presto-install.sh script is worker.
	//   Based on the Presto usage habit,
	//   you are advised to install dualroles on the active Master nodes and worker on the Core nodes.
	// - Local VM path: Enter a script path.
	//   The script path must start with a slash (/) and end with .sh.
	Uri string `json:"uri" required:"true"`
	// Bootstrap action script parameters.
	Parameters string `json:"parameters,omitempty"`
	// Type of a node where the bootstrap action script is executed.
	// The value can be Master, Core, or Task.
	Nodes []string `json:"nodes" required:"true"`
	// Whether the bootstrap action script runs only on active Master nodes.
	// The default value is false, indicating that the bootstrap action script can run on all Master nodes.
	ActiveMaster *bool `json:"active_master,omitempty"`
	// Time when the bootstrap action script is executed.
	// Currently, the following two options are available:
	// Before component start and After component start
	// The default value is false,
	// indicating that the bootstrap action script is executed after the component is started.
	BeforeComponentStart *bool `json:"before_component_start,omitempty"`
	// Whether to continue executing subsequent scripts
	// and creating a cluster after the bootstrap action script fails to be executed.
	// - continue: Continue to execute subsequent scripts.
	// - errorout: Stop the action.
	//   The default value is errorout, indicating that the action is stopped.
	// NOTE
	// You are advised to set this parameter to continue in the commissioning phase
	// so that the cluster can continue to be installed
	// and started no matter whether the bootstrap action is successful.
	FailAction string `json:"fail_action" required:"true"`
}

type CreateResponse struct {
	// Cluster ID, which is returned by the system after the cluster is created.
	ClusterId string `json:"cluster_id"`
	// Operation result.
	// - true: The operation is successful.
	// - false: The operation failed.
	Result bool `json:"result"`
	// System message, which can be empty.
	Msg string `json:"msg"`
}
