package metrics

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type ListMetricsRequest struct {
	// Specifies the dimension. For example, the ECS dimension is instance_id.
	// For details about each service dimension, see Services Interconnected with Cloud Eye.
	// A maximum of three dimensions are supported,
	// and the dimensions are numbered from 0 in dim.{i}=key,value format.
	// The key cannot exceed 32 characters and the value cannot exceed 256 characters.
	// Single dimension: dim.0=instance_id,6f3c6f91-4b24-4e1b-b7d1-a94ac1cb011d
	// Multiple dimensions: dim.0=key,value&dim.1=key,value。
	Dim string `q:"dim,omitempty"`
	// The value ranges from 1 to 1000, and is 1000 by default.
	// This parameter is used to limit the number of query results.
	Limit *int `q:"limit,omitempty"`
	// Specifies the metric ID. For example, if the monitoring metric of an ECS is CPU usage, metric_name is cpu_util.
	MetricName string `q:"metric_name,omitempty"`
	// Query the namespace of a service.
	Namespace string `q:"namespace,omitempty"`
	// Specifies the result sorting method, which is sorted by timestamp.
	// The default value is desc.
	// asc: The query results are displayed in the ascending order.
	// desc: The query results are displayed in the descending order.
	Order string `q:"order,omitempty"`
	// Specifies the paging start value.
	// The format is namespace.metric_name.key:value.
	Start string `q:"start,omitempty"`
}

type ListMetricsBuilder interface {
	ToMetricsListMap() (string, error)
}

func (opts ListMetricsRequest) ToMetricsListMap() (string, error) {
	s, err := build.QueryString(opts)
	if err != nil {
		return "", err
	}
	return s.String(), err
}

func ListMetrics(client *golangsdk.ServiceClient, opts ListMetricsBuilder) pagination.Pager {
	url := metricsURL(client)
	if opts != nil {
		query, err := opts.ToMetricsListMap()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}

	return pagination.Pager{
		Client:     client,
		InitialURL: url,
		CreatePage: func(r pagination.PageResult) pagination.Page {
			return MetricsPage{pagination.LinkedPageBase{PageResult: r}}
		},
	}
}
