package cluster

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type UpdateOpts struct {
	// Cluster ID
	ClusterId string
	// Service ID.
	// This parameter is reserved for extension.
	// You do not need to set this parameter.
	ServiceId string `json:"service_id,omitempty"`
	// Plan ID.
	// This parameter is reserved for extension.
	// You do not need to set this parameter.
	PlanId string `json:"plan_id,omitempty"`
	// Core parameters.
	Parameters Parameters `json:"parameters"`
}

func Update(client *golangsdk.ServiceClient, opts UpdateOpts) (string, error) {
	body, err := build.RequestBody(opts, "")
	if err != nil {
		return "", err
	}

	// PUT /v1.1/{project_id}/cluster_infos/{cluster_id}
	raw, err := client.Put(client.ServiceURL("cluster_infos", opts.ClusterId), body, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return "", err
	}

	var res struct {
		Result string `json:"result"`
	}
	err = extract.Into(raw.Body, &res)
	return res.Result, err
}

type Parameters struct {
	// Order ID obtained by the system during scale-out or scale-in.
	// You do not need to set the parameter
	OrderId string `json:"order_id,omitempty"`
	// - scale_in: cluster scale-in
	// - scale_out: cluster scale-out
	ScaleType string `json:"scale_type"`
	// ID of the newly added or removed node.
	// The parameter value is fixed to node_orderadd.
	// The ID of a newly added or removed node includes node_orderadd, for example, node-orderadd-TBvSr.com.
	NodeId string `json:"node_id"`
	// Node group to be scaled out or in
	// - If the value of node_group is  core_node_default_group, the ode group is a Core node group.
	// - If the value of node_group is task_node_default_group, the node group is a Task node group.
	// If it is left blank, the default value core_node_default_group is used.
	NodeGroup string `json:"node_group,omitempty"`
	// Task node specifications.
	// - When the number of Task nodes is 0, this parameter is used to specify Task node specifications.
	// - When the number of Task nodes is greater than 0, this parameter is unavailable.
	TaskNodeInfo TaskNodeInfo `json:"task_node_info,omitempty"`
	// Number of nodes to be added or removed
	// - The maximum number of nodes to be added is 500 minus the number of Core and Task nodes.
	//   For example, the current number of Core nodes is 3,
	//   the number of nodes to be added must be less than or equal to 497.
	//   A maximum of 500 Core and Task nodes are supported by default.
	//   If more than 500 Core and Task nodes are required,
	//   contact technical support engineers or call a background API to modify the database.
	// - Nodes can be deleted for cluster scale-out
	//   when the number of Core nodes is greater than 3 or the number of Task nodes is greater than 0.
	//   For example,
	//   if there are 5 Core nodes and 5 Task nodes in a cluster,
	//   only 2 (5 minus 3) Core nodes are available for deletion and 5 or fewer than 5 Task nodes can be deleted.
	Instances int `json:"instances"`
	// This parameter is valid only when a bootstrap action is configured during cluster creation and takes effect during scale-out.
	// It indicates whether the bootstrap action specified during cluster creation is performed on nodes added during scale-out.
	// The default value is false, indicating that the bootstrap  action is performed.
	// MRS 1.7.2 or later supports this parameter.
	SkipBootstrapScripts string `json:"skip_bootstrap_scripts,omitempty"`
	// Whether to start components on the added nodes after cluster scale-out
	// - true: Do not start components after scale-out.
	// - false: Start components after scale-out.
	// This parameter is valid only in MRS 1.7.2 or later.
	ScaleWithoutStart *bool `json:"scale_without_start,omitempty"`
	// ID list of Task nodes to be deleted during task node scale-in.
	// - This parameter does not take effect when scale_type is set to scale-out.
	// - If scale_type is set to scale-in and cannot be left blank,
	//   the system deletes the specified Task nodes.
	// - When scale_type is set to scale-in and server_ids is left blank,
	//   the system automatically deletes the Task nodes based on the system rules.
	ServerIds []string `json:"server_ids,omitempty"`
}

type TaskNodeInfo struct {
	// Instance specifications of a Task node, for example, c6.4xlarge.4linux.mrs
	NodeSize string `json:"node_size"`
	// Data disk storage type of the Task node, supporting SATA, SAS, and SSD currently
	// - SATA: Common I/O
	// - SAS: High I/O
	// - SSD: Ultra-high I/O
	DataVolumeType string `json:"data_volume_type,omitempty"`
	// Number of data disks of a Task node
	// Value range: 1 to 10
	DataVolumeCount int `json:"data_volume_count,omitempty"`
	// Data disk storage space of a Task node
	// Value range: 100 GB to 32,000 GB
	DataVolumeSize int `json:"data_volume_size,omitempty"`
}
