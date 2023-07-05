package log

import "github.com/opentelekomcloud/gophertelekomcloud"

type ListLogItemsOpts struct {
	// Log type:
	//
	// app_log: application log
	//
	// node_log: host log
	//
	// custom_log: log in a custom path
	Category string `json:"category" required:"true"`
	// Log filter criteria, which vary according to log sources.
	SearchKey SearchKey `json:"searchKey" required:"true"`
	// Cloud Container Engine (CCE) cluster ID.
	ClusterId string `json:"clusterId" required:"true"`
	// CCE cluster namespace.
	NameSpace string `json:"nameSpace,omitempty"`
	// Service name.
	AppName string `json:"appName,omitempty"`
	// Container pod name.
	PodName string `json:"podName,omitempty"`
	// Log file name.
	PathFile string `json:"pathFile,omitempty"`
	// IP address of the VM where logs are located.
	HostIP string `json:"hostIP,omitempty"`
	// Enter a keyword between two adjacent delimiters for exact search.
	//
	// Enter a keyword for fuzzy search. Example: RROR, ERRO?, *ROR*, ERR*, or ER*OR.
	//
	// Enter a phrase for exact search. Example: Start to refresh alm Statistic.
	//
	// Enter contents containing AND (&&) or OR (||) for search. Example: query&&logs or query||logs.
	//
	// Default delimiters:
	//
	// , '";=()[]{}@&<>/:\n\t\r
	KeyWord string `json:"keyWord,omitempty"`
	// Start time of the query (UTC, in ms).
	StartTime *int64 `json:"startTime" required:"true"`
	// End time of the query (UTC, in ms).
	EndTime *int64 `json:"endTime" required:"true"`
	// Whether to hide the system log (icagent\kubectl) during the query. 0 (default): Hide. 1: Not hide.
	//
	// Value Range
	//
	// 0 or 1
	HideSyslog *int `json:"hideSyslog,omitempty"`
}

type SearchKey struct {
	// CCE cluster ID.
	ClusterId string `json:"clusterId" required:"true"`
	// CCE cluster namespace.
	NameSpace string `json:"nameSpace,omitempty"`
	// Application name.
	AppName string `json:"appName,omitempty"`
	// Container instance name.
	PodName string `json:"podName,omitempty"`
	// Log file name.
	PathFile string `json:"pathFile,omitempty"`
	// IP address of the VM where logs are located.
	HostIP string `json:"hostIP,omitempty"`
}

func ListLogItems(client *golangsdk.ServiceClient, opts ListLogItemsOpts) {
	// POST /v2/{project_id}/als/action
}

type LogItemResponse struct {
	// Response code. Example: AOM.0200, which indicates a success response.
	ErrorCode string `json:"errorCode"`
	// Response message.
	ErrorMessage string `json:"errorMessage"`
	// Metadata, including total number of returned records and results.
	Result LogResult `json:"result"`
}

type LogResult struct {
	// Number of returned records.
	Total *int `json:"total"`
	// Data array.
	Data []ItemData `json:"data"`
}

type ItemData struct {
	// Log type.
	Category string `json:"category"`
	// Hash value of the log source.
	LogHash string `json:"loghash"`
	// Cloud Container Engine (CCE) cluster ID.
	ClusterId string `json:"clusterId"`
	// CCE cluster name.
	ClusterName string `json:"clusterName"`
	// CCE cluster namespace.
	NameSpace string `json:"nameSpace"`
	// CCE container pod name.
	PodName string `json:"podName"`
	// Service name.
	AppName string `json:"appName"`
	// Service ID of an AOM resource.
	ServiceID string `json:"serviceID"`
	// CCE container name.
	ContainerName string `json:"containerName"`
	// Source log data.
	LogContent string `json:"logContent"`
	// Absolute path of a log file.
	PathFile string `json:"pathFile"`
	// IP address of the VM where log files are located.
	HostIP string `json:"hostIP"`
	// ID of a host in a cluster.
	HostId string `json:"hostId"`
	// Name of the VM where log files are located.
	HostName string `json:"hostName"`
	// Log collection time (UTC time, in ms).
	CollectTime string `json:"collectTime"`
	// Sequence number of a log line.
	LineNum string `json:"lineNum"`
	// Size of a single-line log.
	LogContentSize string `json:"logContentSize"`
}
