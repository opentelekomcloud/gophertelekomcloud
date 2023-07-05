package monitor

import "github.com/opentelekomcloud/gophertelekomcloud"

type ListMetricItemsOpts struct {
	// Metric namespace.
	//
	// PAAS.CONTAINER: application metric.
	//
	// PAAS.NODE: node metric.
	//
	// PAAS.SLA: Service Level Agreement (SLA) metric.
	//
	// PAAS.AGGR: cluster metric.
	//
	// CUSTOMMETRICS: custom metric.
	//
	// Value Range
	//
	// PAAS.CONTAINER, PAAS.NODE, PAAS.SLA, PAAS.AGGR, CUSTOMMETRICS, and so on.
	Namespace string `json:"namespace" required:"true"`
	// Metric dimension.
	//
	// dimensions.name: dimension name, such as clusterName, clusterId, appName, appID, deploymentName, podName, podID, containerName, or containerID.
	//
	// dimensions.value: dimension value, such as a specific application instance ID.
	Dimensions []Dimension `json:"dimensions,omitempty"`
	// Metric name.
	//
	// Value Range
	//
	// 1-1000 characters.
	MetricName string `json:"metricName,omitempty"`
}

type Dimension struct {
	Name  string `json:"name" required:"true"`
	Value string `json:"value" required:"true"`
}

func ListMetricItems(client *golangsdk.ServiceClient, opts ListMetricItemsOpts) {
	// POST /v2/{project_id}/ams/metrics
	type RequestParameter struct {
		// If type (a URI parameter) is not inventory, the information carried by the array is used to query metrics.
		MetricItems []ListMetricItemsOpts `json:"metricItems,omitempty"`
	}
}

type ListMetricItemsResponse struct {
	// Response code. Example: AOM.0200, which indicates a success response.
	ErrorCode string `json:"errorCode"`
	// Response message.
	ErrorMessage string `json:"errorMessage"`
	// List of metrics.
	Metrics []Metric `json:"metrics"`
}

type Metric struct {
	// Namespace.
	Namespace string `json:"namespace"`
	// Metric name.
	MetricName string `json:"metricName"`
	// Metric unit.
	Unit string `json:"unit"`
	// List of metric dimensions.
	Dimensions []Dimension `json:"dimensions"`
}
