package monitors

import "github.com/opentelekomcloud/gophertelekomcloud/openstack/common/structs"

type Type string

const (
	TypeTCP        Type = "TCP"
	TypeHTTP       Type = "HTTP"
	TypeHTTPS      Type = "HTTPS"
	TypeUDPConnect Type = "UDP_CONNECT"
	// TypePING is not documented in the reviewed OTC documentation.
	TypePING Type = "PING"
)

type Monitor struct {
	ID             string                `json:"id"`
	Name           string                `json:"name"`
	ProjectID      string                `json:"project_id"`
	Type           Type                  `json:"type"`
	Delay          int                   `json:"delay"`
	Timeout        int                   `json:"timeout"`
	MaxRetries     int                   `json:"max_retries"`
	MaxRetriesDown int                   `json:"max_retries_down"`
	HTTPMethod     string                `json:"http_method"`
	URLPath        string                `json:"url_path"`
	DomainName     string                `json:"domain_name"`
	ExpectedCodes  string                `json:"expected_codes"`
	AdminStateUp   bool                  `json:"admin_state_up"`
	MonitorPort    int                   `json:"monitor_port"`
	Pools          []structs.ResourceRef `json:"pools"`
	CreatedAt      string                `json:"created_at"`
	UpdatedAt      string                `json:"updated_at"`

	// These fields are documented but marked unsupported by the reviewed OTC API.
	HTTPVersion        string `json:"http_version"`
	ContentCheckEnable bool   `json:"content_check_enable"`
	SendContent        string `json:"send_content"`
	ReceiveContent     string `json:"receive_content"`
	ContentFormat      string `json:"content_format"`
}
