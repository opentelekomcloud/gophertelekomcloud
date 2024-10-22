package instances

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ModifyParametersOpts struct {
	// Parameters that need to be modified
	Values Values `json:"values" required:"true"`
}

type Values struct {
	// Data association among multiple sharded tables. The optimizer processes JOIN operations at the MySQL layer based on these associations.
	// The format is [{tb.col1,tb2.col2},{tb.col2,tb3.col1},...]. (Optional)
	BindTable string `json:"bind_table,omitempty"`
	// DDM server's character set. To store emoticons, set both this parameter and the character set on RDS to utf8mb4.
	// Enumerated values: gbk, utf8, utf8mb4 (Optional)
	CharacterSetServer string `json:"character_set_server,omitempty"`
	// Collation on the DDM server.
	// Enumerated values: utf8_unicode_ci, utf8_bin, gbk_chinese_ci, gbk_bin, utf8mb4_unicode_ci, utf8mb4_bin (Optional)
	CollationServer string `json:"collation_server,omitempty"`
	// Concurrency level of scanning table shards in a logical table.
	// Enumerated values: RDS_INSTANCE, DATA_NODE, PHY_TABLE (Optional)
	ConcurrentExecutionLevel string `json:"concurrent_execution_level,omitempty"`
	// Number of seconds the server waits for activity on a connection before closing it.
	// Range: 60-28800. Default: 28800 (Optional)
	ConnectionIdleTimeout string `json:"connection_idle_timeout,omitempty"`
	// Whether the table recycle bin is enabled.
	// Enumerated values: OFF, ON (Optional)
	EnableTableRecycle string `json:"enable_table_recycle,omitempty"`
	// Whether constant values can be inserted by executing the LOAD DATA statement.
	// Enumerated values: OFF, ON (Optional)
	InsertToLoadData string `json:"insert_to_load_data,omitempty"`
	// Timeout limit of an in-transit transaction, in seconds.
	// Range: 0-100. Default: 1 (Optional)
	LiveTransactionTimeoutOnShutdown string `json:"live_transaction_timeout_on_shutdown,omitempty"`
	// Minimum duration of a query to be logged as slow, in seconds.
	// Range: 0.01-10. Default: 1 (Optional)
	LongQueryTime string `json:"long_query_time,omitempty"`
	// Maximum size of a packet or any generated intermediate string.
	// Range: 1024-1073741824. Default: 16777216 (Optional)
	MaxAllowedPacket string `json:"max_allowed_packet,omitempty"`
	// Maximum of concurrent RDS client connections allowed per DDM instance. Default: 0 (Optional)
	MaxBackendConnections string `json:"max_backend_connections,omitempty"`
	// Concurrent connections allowed per DDM instance, depending on the class and quantity of associated RDS instances.
	// Range: 10-40000. Default: 20000 (Optional)
	MaxConnections string `json:"max_connections,omitempty"`
	// Minimum concurrent connections from a DDM node to an RDS instance.
	// Range: 0-10000000. Default: 10 (Optional)
	MinBackendConnections string `json:"min_backend_connections,omitempty"`
	// Whether the SELECT statements that do not contain any FROM clauses are pushed down.
	// Enumerated values: OFF, ON (Optional)
	NotFromPushdown string `json:"not_from_pushdown,omitempty"`
	// Threshold in seconds of the replication lag between a primary RDS instance to its read replica.
	// Range: 0-7200. Default: 30 (Optional)
	SecondsBehindMaster string `json:"seconds_behind_master,omitempty"`
	// Whether SQL audit is enabled.
	// Enumerated values: OFF, ON (Optional)
	SQLAudit string `json:"sql_audit,omitempty"`
	// Number of seconds to wait for a SQL statement to execute before it times out.
	// Range: 100-28800. Default: 28800 (Optional)
	SQLExecuteTimeout string `json:"sql_execute_timeout,omitempty"`
	// Whether a binlog hint is added to each DDL statement.
	// Enumerated values: OFF, ON (Optional)
	SupportDDLBinlogHint string `json:"support_ddl_binlog_hint,omitempty"`
	// Transactions supported by DDM.
	// Enumerated values: XA, FREE, NO_DTX (Optional)
	TransactionPolicy string `json:"transaction_policy,omitempty"`
	// Whether the SQL execution plan is optimized based on parameter values.
	// Enumerated values: OFF, ON (Optional)
	UltimateOptimize string `json:"ultimate_optimize,omitempty"`
}

// ModifyParameters is used to modify parameters of a DDM instance.
func ModifyParameters(client *golangsdk.ServiceClient, instanceId string, opts ModifyParametersOpts) (*ModifyParametersResponse, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// PUT /v3/{project_id}/instances/{instance_id}/configurations
	raw, err := client.Put(client.ServiceURL("instances", instanceId, "configurations"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res ModifyParametersResponse
	return &res, extract.Into(raw.Body, &res)
}

type ModifyParametersResponse struct {
	// DDM instance nodes
	NodeList string `json:"nodeList"`
	// Whether the instance needs to be restarted
	NeedRestart bool `json:"needRestart"`
	// Task ID
	JobID string `json:"jobId"`
	// Parameter group ID
	ConfigID string `json:"configId"`
	// Parameter group name
	ConfigName string `json:"configName"`
}
