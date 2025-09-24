package load_balancer

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ConfigureListenerOpts struct {
	// These parameters passed to the load_balancers.ConfigureLoadBalancingListeners function.
	// Protocol type. HTTP and HTTPS are supported.
	Protocol string `json:"protocol" required:"true"`
	// Port.
	ProtocolPort int `json:"protocol_port" required:"true"`
	// Server certificate ID. This parameter is mandatory when protocol is set to HTTPS.
	ServerCertId string `json:"server_cert_id,omitempty"`
	// CA certificate ID. This parameter is mandatory when protocol is set to HTTPS and bidirectional authentication is used.
	CaCertId string `json:"ca_cert_id,omitempty"`
	// Type: searchTool indicates that the listener is configured for Elasticsearch/OpenSearch.
	// viewTool indicates that the listener is configured for Kibana/OpenSearch Dashboards.
	// The default value is searchTool.
	Type string `json:"type,omitempty"`
}

// ConfigureListener will configure load balancer listener for a CSS cluster based on ConfigureListenerOpts.
func ConfigureListener(client *golangsdk.ServiceClient, clusterID string, opts ConfigureListenerOpts) (*ConfiguredListenerResp, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	url := client.ServiceURL("clusters", clusterID, "es-listeners")

	raw, err := client.Post(url, b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
		MoreHeaders: map[string]string{
			"Content-Type": "application/json",
		},
	})

	if err != nil {
		return nil, err
	}

	var res ConfiguredListenerResp
	err = extract.IntoStructPtr(raw.Body, &res, "")
	return &res, err
}

type ConfiguredListenerResp struct {
	// ELB ID
	ElbId string `json:"elb_id"`
}
