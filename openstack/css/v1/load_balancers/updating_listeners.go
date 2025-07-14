package load_balancers

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type UpdatingListenerOpts struct {
	// These parameters passed to the load_balancers.UpdatingLoadBalancingListeners function.
	// Listener object.
	Listener EsListenerRequest `json:"listener" required:"true"`
	// Type: searchTool indicates that the load balancer is modified for Elasticsearch/OpenSearch.
	// viewTool indicates that the load balancer is modified for Kibana/OpenSearch Dashboards.
	// The default value is searchTool.
	Type string `json:"type"`
}

type EsListenerRequest struct {
	// These parameters are part of an EsListenerRequest object.
	// ID of the server certificate used by the listener.
	DefaultTlsContainerRef string `json:"default_tls_container_ref" required:"true"`
	// ID of the CA certificate used by the listener.
	// This parameter is mandatory when bidirectional authentication is to be updated.
	ClientCaTlsContainerRef string `json:"client_ca_tls_container_ref"`
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
	ProtocolPort string `json:"protocol_port"`
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

// UpdatingLoadBalancingListeners will update load balancing listeners for a CSS cluster based on UpdatingListenerOpts.
func UpdatingLoadBalancingListeners(client *golangsdk.ServiceClient, clusterID string, listenerID string, opts UpdatingListenerOpts) (*EsListenerResponse, error) {
	b, err := build.RequestBody(opts, "listener")
	if err != nil {
		return nil, err
	}

	url := client.ServiceURL("clusters", clusterID, "es-listeners", listenerID)

	raw, err := client.Post(url, b, nil, &golangsdk.RequestOpts{
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
