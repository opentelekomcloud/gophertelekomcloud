package metricdata

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"net/url"
)

type ShowEventDataOpts struct {
	// Query the namespace of a service.
	Namespace string `q:"namespace"`
	// Specifies the dimension. For example, the ECS dimension is instance_id.
	// For details about the dimensions corresponding to the monitoring metrics of each service,
	// see the monitoring metrics description of the corresponding service in Services Interconnected with Cloud Eye.
	//
	// Specifies the dimension. A maximum of three dimensions are supported,
	// and the dimensions are numbered from 0 in dim.{i}=key,value format.
	// The key cannot exceed 32 characters and the value cannot exceed 256 characters.
	Dim0 string `q:"dim.0"`
	Dim1 string `q:"dim.1"`
	Dim2 string `q:"dim.2"`
	// Specifies the event type.
	Type string `q:"type"`
	// Specifies the start time of the query.
	From string `q:"from"`
	// Specifies the end time of the query.
	To string `q:"to"`
}

func ListEventData(client *golangsdk.ServiceClient, opts ShowEventDataOpts) ([]EventDataInfo, error) {
	var opts2 interface{} = &opts
	q, err := build.QueryString(opts2)
	if err != nil {
		return nil, err
	}

	// GET /V1.0/{project_id}/event-data
	raw, err := client.Get(client.ServiceURL("event-data")+q.String(), nil, nil)
	if err != nil {
		return nil, err
	}

	var res []EventDataInfo
	err = extract.IntoSlicePtr(raw.Body, &res, "datapoints")

	return res, err
}

type EventDataInfo struct {
	// Specifies the event type, for example, instance_host_info.
	Type string `json:"type"`
	// Specifies when the event is reported. It is a UNIX timestamp and the unit is ms.
	Timestamp int64 `json:"timestamp"`
	// Specifies the host configuration information.
	Value string `json:"value"`
}
