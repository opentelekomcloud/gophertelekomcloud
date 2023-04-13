package listeners

// Protocol represents a listener protocol.
type Protocol string

// Supported attributes for create/update operations.
const (
	ProtocolTCP   Protocol = "TCP"
	ProtocolUDP   Protocol = "UDP"
	ProtocolHTTP  Protocol = "HTTP"
	ProtocolHTTPS Protocol = "HTTPS"
)

type IpGroup struct {
	IpGroupID string `json:"ipgroup_id" required:"true"`
	Enable    *bool  `json:"enable_ipgroup,omitempty"`
	Type      string `json:"type,omitempty"`
}

type InsertHeaders struct {
	ForwardedELBIP   *bool `json:"X-Forwarded-ELB-IP,omitempty"`
	ForwardedPort    *bool `json:"X-Forwarded-Port,omitempty"`
	ForwardedForPort *bool `json:"X-Forwarded-For-Port,omitempty"`
	ForwardedHost    *bool `json:"X-Forwarded-Host" required:"true"`
}

type IpGroupUpdate struct {
	IpGroupId string `json:"ipgroup_id,omitempty"`
	Type      string `json:"type,omitempty"`
}
