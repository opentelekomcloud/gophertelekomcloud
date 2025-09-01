package load_balancer

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type LoadBalancingListenerOpts struct {
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

// ConfigureLoadBalancingListeners will configure load balancing listeners for a CSS cluster based on LoadBalancingListenerOpts.
func ConfigureLoadBalancingListeners(client *golangsdk.ServiceClient, clusterID string, opts LoadBalancingListenerOpts) (string, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return "", err
	}

	url := client.ServiceURL("clusters", clusterID, "es-listeners")

	raw, err := client.Post(url, b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
		MoreHeaders: map[string]string{
			"Content-Type": "application/json",
		},
	})

	if err != nil {
		return "", err
	}

	var res struct {
		ElbId string `json:"elb_id"`
	}
	err = extract.IntoStructPtr(raw.Body, &res, "")
	return res.ElbId, err
}
