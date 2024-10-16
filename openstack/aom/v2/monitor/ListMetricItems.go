package monitor

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ListMetricItemsQuery struct {
	// Metric query mode. The information carried by metricItems in the request body is used to query metrics.
	Type string `q:"type"`
	// Maximum number of returned records. Value range: 1-1000. Default value: 1000.
	Limit int `q:"limit"`
	// Start position of a pagination query. The value is a non-negative integer.
	Start int `q:"start"`
}

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

func ListMetricItems(client *golangsdk.ServiceClient, opts []ListMetricItemsOpts, query ListMetricItemsQuery) (*ListMetricItemsResponse, error) {
	q, err := golangsdk.BuildQueryString(query)
	if err != nil {
		return nil, err
	}

	b, err := build.RequestBody(opts, "metricItems")
	if err != nil {
		return nil, err
	}

	// POST /v2/{project_id}/ams/metrics
	raw, err := client.Post(client.ServiceURL("ams", "metrics")+q.String(), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res ListMetricItemsResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
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
