package metrics

import (
	"bytes"
	"strconv"

	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type ListMetricsResponse struct {
	// Specifies the list of metric objects.
	Metrics []MetricInfoList `json:"metrics"`
	// Specifies the metadata of query results, including the pagination information.
	MetaData MetaData `json:"meta_data"`
}

type MetricInfoList struct {
	// Specifies the metric namespace.
	Namespace string `json:"namespace"`
	// Specifies the metric name, such as cpu_util.
	MetricName string `json:"metric_name"`
	// Specifies the metric unit.
	Unit string `json:"unit"`
	// Specifies the list of metric dimensions.
	Dimensions []MetricsDimension `json:"dimensions"`
}

type MetaData struct {
	// Specifies the number of returned results.
	Count int `json:"count"`
	// Specifies the pagination marker.
	Marker string `json:"marker"`
	// Specifies the total number of metrics.
	Total int `json:"total"`
}

type MetricsDimension struct {
	// Specifies the dimension.
	Name string `json:"name"`
	// Specifies the dimension value, for example, an ECS ID.
	Value string `json:"value"`
}

type ListResult struct {
	golangsdk.Result
}

func ExtractMetrics(r pagination.Page) (ListMetricsResponse, error) {
	var s ListMetricsResponse

	err := extract.Into(bytes.NewReader(r.(MetricsPage).Body), &s)
	return s, err
}

func ExtractAllPagesMetrics(r pagination.Page) (ListMetricsResponse, error) {
	var s ListMetricsResponse
	s.Metrics = make([]MetricInfoList, 0)

	err := extract.Into(bytes.NewReader(r.(MetricsPage).Body), &s)
	if len(s.Metrics) == 0 {
		return s, nil
	}
	s.MetaData.Count = len(s.Metrics)
	s.MetaData.Total = len(s.Metrics)
	var buf bytes.Buffer
	buf.WriteString(s.Metrics[len(s.Metrics)-1].Namespace)
	buf.WriteString(".")
	buf.WriteString(s.Metrics[len(s.Metrics)-1].MetricName)
	for _, dimension := range s.Metrics[len(s.Metrics)-1].Dimensions {
		buf.WriteString(".")
		buf.WriteString(dimension.Name)
		buf.WriteString(":")
		buf.WriteString(dimension.Value)
	}
	s.MetaData.Marker = buf.String()
	return s, err
}

type MetricsPage struct {
	pagination.LinkedPageBase
}

func (r MetricsPage) NextPageURL() (string, error) {
	metrics, err := ExtractMetrics(r)
	if err != nil {
		return "", err
	}

	if len(metrics.Metrics) < 1 {
		return "", nil
	}

	limit := r.URL.Query().Get("limit")
	num, _ := strconv.Atoi(limit)
	if num > len(metrics.Metrics) {
		return "", nil
	}

	metricslen := len(metrics.Metrics) - 1

	var buf bytes.Buffer
	buf.WriteString(metrics.Metrics[metricslen].Namespace)
	buf.WriteString(".")
	buf.WriteString(metrics.Metrics[metricslen].MetricName)
	for _, dimension := range metrics.Metrics[metricslen].Dimensions {
		buf.WriteString(".")
		buf.WriteString(dimension.Name)
		buf.WriteString(":")
		buf.WriteString(dimension.Value)
	}
	return r.WrapNextPageURL(buf.String())
}

func (r MetricsPage) IsEmpty() (bool, error) {
	s, err := ExtractMetrics(r)
	if err != nil {
		return false, err
	}
	return len(s.Metrics) == 0, err
}

func (r MetricsPage) WrapNextPageURL(start string) (string, error) {
	limit := r.URL.Query().Get("limit")
	if limit == "" {
		return "", nil
	}
	uq := r.URL.Query()
	uq.Set("start", start)
	r.URL.RawQuery = uq.Encode()
	return r.URL.String(), nil
}
