package connection

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ListSessionOpts struct {
	// Specifies the node ID. The following nodes can be queried: mongos nodes in the cluster, and all nodes in the replica set and single node instances.
	NodeId string `json:"-"`
	// Specifies the index position. If offset is set to N, the resource query starts from the N+1 piece of data. The value is 0 by default, indicating that the query starts from the first piece of data. The value cannot be a negative number.
	Offset int `q:"offset"`
	// Specifies the number of records to be queried. The value range is [1, 20]. The default value is 10, indicating that 10 records are returned.
	Limit int `q:"limit"`
	// Specifies the execution plan description. If this parameter is left empty, sessions in which plan_summary is empty are queried. You can also specify an execution plan, for example, COLLSCAN IXSCAN FETCH SORT LIMIT SKIP COUNT COUNT_SCAN TEXT PROJECTION
	PlanSummary string `q:"plan_summary"`
	// Specifies the operation type. If this parameter is left empty, sessions in which type is empty are queried. You can also specify an operation type, for example, none update insert query command getmore remove killcursors.
	Type string `q:"type"`
	// Specifies the namespace. If this parameter is left blank, the sessions in which namespace is empty are queried. You can also specify the value based on the service requirements.
	NameSpace string `q:"namespace"`
	// Specifies the duration. The unit is us. If this parameter is left empty, the sessions in which cost_time is empty are queried. You can also set this parameter based on the service requirements, indicating that the sessions in which the value of cost_time exceeds the specified value are queried.
	CostTime int `q:"cost_time"`
}

func ListSessions(client *golangsdk.ServiceClient, opts ListSessionOpts) (*ListSessionsResponse, error) {
	url, err := golangsdk.NewURLBuilder().WithEndpoints("nodes", opts.NodeId, "sessions").WithQueryParams(&opts).Build()
	if err != nil {
		return nil, err
	}

	// GET https://{Endpoint}/v3/{project_id}/nodes/{node_id}/sessions
	raw, err := client.Get(client.ServiceURL(url.String()), nil, nil)
	if err != nil {
		return nil, err
	}

	var res ListSessionsResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type ListSessionsResponse struct {
	Sessions   []SessionResponse `json:"sessions"`
	TotalCount int               `json:"total_count"`
}

type SessionResponse struct {
	// Indicates the session ID.
	Id string `json:"id"`
	// Indicates that whether the current session is active. If the value is "true", the session is active. If the value is "false", the session is inactive.
	Active bool `json:"active"`
	// Indicates the operation.
	Operation string `json:"operation"`
	// Indicates the operation type.
	Type string `json:"type"`
	// Specifies the duration. The unit is us.
	CostTime string `json:"cost_time"`
	// Indicates the execution plan description.
	PlanSummary string `json:"plan_summary"`
	// Indicates the host.
	Host string `json:"engine_versions"`
	// Indicates the client address.
	Client string `json:"client"`
	// Indicates the connection description.
	Description string `json:"description"`
	// Indicates the namespace.
	Namespace string `json:"namespace"`
}
