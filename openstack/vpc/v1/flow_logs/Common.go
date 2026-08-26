package flow_logs

type FlowLog struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	TenantID     string `json:"tenant_id"`
	Description  string `json:"description"`
	ResourceType string `json:"resource_type"`
	ResourceID   string `json:"resource_id"`
	TrafficType  string `json:"traffic_type"`
	LogGroupID   string `json:"log_group_id"`
	LogTopicID   string `json:"log_topic_id"`
	IndexEnabled bool   `json:"index_enabled"`
	AdminState   bool   `json:"admin_state"`
	Status       string `json:"status"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}

type flowLogResponse struct {
	FlowLog FlowLog `json:"flow_log"`
}
