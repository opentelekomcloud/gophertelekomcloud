package monitors

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type GetStreamMonitorOpts struct {
	// Name of the stream to be queried.
	// Maximum: 60
	StreamName string
	// Stream monitoring metric.
	// (Either label or label_list must be specified.
	// If both label_list and label are specified, label_list prevails.)
	// - total_put_bytes_per_stream: total input traffic (byte)
	// - total_get_bytes_per_stream: total output traffic (byte)
	// - total_put_records_per_stream: total number of input records
	// - total_get_records_per_stream: total number of output records
	// - total_put_req_latency: average processing time of upload requests (millisecond)
	// - total_get_req_latency: average processing time of download requests (millisecond)
	// - total_put_req_suc_per_stream: number of successful upload requests
	// - total_get_req_suc_per_stream: number of successful download requests
	// - traffic_control_put: number of rejected upload requests due to flow control
	// - traffic_control_get: number of rejected download requests due to flow control
	// Enumeration values:
	//  total_put_bytes_per_stream
	//  total_get_bytes_per_stream
	//  total_put_records_per_stream
	//  total_get_records_per_stream
	//  total_put_req_latency
	//  total_get_req_latency
	//  total_put_req_suc_per_stream
	//  total_get_req_suc_per_stream
	//  traffic_control_put
	//  traffic_control_get
	Label string `q:"label,omitempty"`
	// List of labels separated by commas (,) to query multiple labels in batches.
	// (Either label or label_list must be set.
	// If both label_list and label exist, label_list prevails.)
	LabelList string `q:"label_list,omitempty"`
	// Monitoring start time, which is a 10-digit timestamp.
	StartTime int64 `q:"start_time"`
	// Monitoring end time, which is a 10-digit timestamp.
	EndTime int64 `q:"end_time"`
}

func GetStreamMonitor(client *golangsdk.ServiceClient, opts GetStreamMonitorOpts) (*GetStreamMonitorResponse, error) {
	url, err := golangsdk.NewURLBuilder().WithEndpoints("streams", opts.StreamName, "metrics").WithQueryParams(&opts).Build()
	if err != nil {
		return nil, err
	}

	// GET /v2/{project_id}/streams/{stream_name}/metrics
	raw, err := client.Get(client.ServiceURL(url.String()), nil, nil)
	if err != nil {
		return nil, err
	}

	var res GetStreamMonitorResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type GetStreamMonitorResponse struct {
	// Data object.
	Metrics Metrics `json:"metrics,omitempty"`
	// List of monitored data objects.
	MetricsList []Metrics `json:"metrics_list,omitempty"`
}

type Metrics struct {
	// Monitoring data.
	DataPoints []DataPoint `json:"data_points,omitempty"`
	// Metric
	Label string `json:"label,omitempty"`
}

type DataPoint struct {
	// Timestamp.
	Timestamp int64 `json:"timestamp,omitempty"`
	// Monitoring value corresponding to the timestamp.
	Value int64 `json:"value,omitempty"`
}
