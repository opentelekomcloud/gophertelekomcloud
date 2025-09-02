package load_balancer

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type UpdateListenerOpts struct {
	// Listener Object.
	Listener struct {
		// Server certificate ID. This parameter is mandatory when protocol is set to HTTPS.
		ServerCertId string `json:"default_tls_container_ref"`
		// CA certificate ID. This parameter is mandatory when protocol is set to HTTPS and bidirectional authentication is used.
		CaCertId string `json:"client_ca_tls_container_ref,omitempty"`
	} `json:"listener"`
	// Type: searchTool indicates that the load balancer is modified for Elasticsearch/OpenSearch.
	// viewTool indicates that the load balancer is modified for Kibana/OpenSearch Dashboards.
	// The default value is searchTool.
	Type string `json:"type,omitempty"`
}

// UpdatingLoadBalancingListeners will update load balancing listeners for a CSS cluster based on UpdatingListenerOpts.
func UpdatingLoadBalancingListeners(client *golangsdk.ServiceClient, clusterID string, listenerID string, opts UpdateListenerOpts) (*EsListenerResponse, error) {

	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	url := client.ServiceURL("clusters", clusterID, "es-listeners", listenerID)

	raw, err := client.Put(url, b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
		MoreHeaders: map[string]string{
			"Content-Type": "application/json",
		},
	})

	if err != nil {
		return nil, err
	}

	var res EsListenerResponse
	err = extract.IntoStructPtr(raw.Body, &res, "listener")
	return &res, err
}

type EsListenerResponse struct {
	// These parameters are part of an EsListenerResponse object.
	// Protocol used by the listener.
	Protocol string `json:"protocol"`
	// Listener ID.
	Id string `json:"id"`
	// Listener name.
	Name string `json:"name"`
	// Port used by the listener.
	ProtocolPort int `json:"protocol_port"`
	// Access control information of the listener object.
	Ipgroup EsIpgroupResource `json:"ipgroup"`
}

type EsIpgroupResource struct {
	// These parameters are part of an EsIpgroupResource object.
	// ID of the IP address group associated with the listener.
	IpgroupId string `json:"ipgroup_id"`
	// Status of an access control group.
	EnableIpgroup bool `json:"enable_ipgroup"`
	// Type of an access control group.
	Type string `json:"type"`
}
